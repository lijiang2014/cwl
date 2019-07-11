package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/lijiang2014/tugboat/docker"
  "github.com/lijiang2014/tugboat/localos"
  
  "github.com/lijiang2014/cwl"
  "github.com/lijiang2014/cwl/process"
  localfs "github.com/lijiang2014/cwl/process/fs/local"
  "path/filepath"
  //gsfs "github.com/lijiang2014/cwl/process/fs/gs"
  
  tug "github.com/lijiang2014/tugboat"
  "github.com/lijiang2014/tugboat/storage/local"
  //gsstore "github.com/buchanae/tugboat/storage/gs"
  
  "github.com/rs/xid"
  "github.com/spf13/cobra"
)

func init() {
  outdir := "cwl-output"
  debug := false

  cmd := &cobra.Command{
    Use: "run <doc.cwl> <inputs.json>",
    Short: "run",
    Args: cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
      return run(args[0], args[1], outdir, debug)
    },
  }
  root.AddCommand(cmd)
  f := cmd.Flags()

  f.StringVar(&outdir, "outdir", outdir, "")
  f.BoolVar(&debug, "debug", debug, "")
}

func run(path, inputsPath, outdir string, debug bool) error {
  fmt.Println("local cwl run.")
  vals, err := cwl.LoadValuesFile(inputsPath)
  if err != nil {
    return err
  }
  inputsDir := filepath.Dir(inputsPath)

  doc, err := cwl.Load(path)
  if err != nil {
    return err
  }

  r := runner{inputsDir, outdir, debug}

  outvals, err := r.runDoc(doc, vals)
  if err != nil {
    return err
  }

  b, err := json.MarshalIndent(outvals, "", "  ")
  if err != nil {
    return err
  }
  fmt.Println(string(b))

  return err
}

type runner struct {
  inputsDir string
  outdir string
  debug bool
}

func (r *runner) runDoc(doc cwl.Document, vals cwl.Values) (cwl.Values, error) {
  switch z := doc.(type) {
  case *cwl.Tool:
    return r.runTool(z, vals)
  case *cwl.Workflow:
    return r.runWorkflow(z, vals)
  default:
    return nil, fmt.Errorf(`running doc: unknown doc type "%s"`, doc.Doctype())
  }
}

func (r *runner) runWorkflow(wf *cwl.Workflow, vals cwl.Values) (cwl.Values, error) {
  process.DebugWorkflow(wf, vals)
  return nil, nil
}

func (r *runner) runTool(tool *cwl.Tool, vals cwl.Values) (cwl.Values, error) {
  // TODO hack. need to think carefully about how resource requirement and runtime
  //      actually get scheduled.
  var resources *cwl.ResourceRequirement
	reqs := append([]cwl.Requirement{}, tool.Requirements...)
	reqs = append(reqs, tool.Hints...)
  for _, req := range reqs {
    if r, ok := req.(cwl.ResourceRequirement); ok {
      resources = &r
    }
  }

  rt := process.Runtime{
    Outdir: "/cwl",
  }
  // TODO related to the resource requirement search above. basically a hack
  //      for the conformance tests, for now.
  if resources != nil {
    rt.Cores = string(resources.CoresMin)
  }

  fs := localfs.NewLocal(r.inputsDir)
  fs.CalcChecksum = true
  //fs, err := gsfs.NewGS("buchanae-funnel")
  //if err != nil {
    //return nil, err
  //}

  proc, err := process.NewProcess(tool, vals, rt, fs)
  if err != nil {
    return nil, err
  }

  cmd, err := proc.Command()
  if err != nil {
    return nil, err
  }

  //fmt.Fprintln(os.Stderr, cmd)

  workdir := "/cwl"
  // TODO necessary for cwl conformance tests
  image := "python:2"

  if d, ok := tool.RequiresDocker(); ok {
    image = d.Pull
    if d.OutputDirectory != "" {
      workdir = d.OutputDirectory
    }
  }

  task := &tug.Task{
    ID: "cwl-test1-" + xid.New().String(),
    ContainerImage: image,
    Command: cmd,
    Workdir: workdir,
    Volumes: []string{workdir, "/tmp"},
    Env: proc.Env(),

    /* TODO need process.OutputBindings() */
    Outputs: []tug.File{
      {
        URL: r.outdir,
        Path: workdir,
      },
    },
  }
  task.Env["HOME"] = workdir
  task.Env["TMPDIR"] = "/tmp"

  stdout := proc.Stdout()
  stderr := proc.Stderr()
  if stdout != "" {
    task.Stdout = workdir + "/" + stdout
  }
  if stderr != "" {
    task.Stderr = workdir + "/" + stderr
  }

  files := []cwl.File{}
  for _, in := range proc.InputBindings() {
    if f, ok := in.Value.(cwl.File); ok {
      files = append(files, flattenFiles(f)...)
    }
  }
  for _, f := range files {
    task.Inputs = append(task.Inputs, tug.File{
      URL: f.Location,
      // TODO
      Path: f.Path,
    })
  }

  ctx := context.Background()
  store, _ := local.NewLocal()
  //store, _ := gsstore.NewGS("buchanae-funnel")
  var log tug.Logger
  if r.debug {
    log = tug.StderrLogger{}
  } else {
    log = tug.EmptyLogger{}
  }
  
  var exec tug.Executor
  if _, got :=tool.RequiresDocker() ; got {
    exec = &docker.Docker{
      Logger: log,
      NoPull: true,
    }
  } else {
    exec = &localos.LocalOS{  Logger:log, EnvAppend:true}
  }

	stage, err := tug.NewStage("cwl-workdir", 0755)
  if err != nil {
    panic(err)
  }
  stage.LeaveDir = r.debug
  defer stage.RemoveAll()

  err = tug.Run(ctx, task, stage, log, store, exec)
  if err != nil {
    if e, ok := err.(*tug.ExecError); ok {
      for _, code := range tool.SuccessCodes {
        if e.ExitCode == code {
          err = nil
        }
      }
    }
  }
  if err != nil {
    return nil, err
  }

  //fmt.Println(strings.Join(cmd, " "))

  outfs := localfs.NewLocal(r.outdir)
  outfs.CalcChecksum = true
  //outfs, err := gsfs.NewGS("buchanae-cwl-output")
  if err != nil {
    return nil, err
  }

  return proc.Outputs(outfs)
}

func flattenFiles(file cwl.File) []cwl.File {
  files := []cwl.File{file}
  for _, fd := range file.SecondaryFiles {
    // TODO fix the mismatch between cwl.File and *cwl.File
    if f, ok := fd.(*cwl.File); ok {
      files = append(files, flattenFiles(*f)...)
    }
  }
  return files
}


