package main

import (
  "fmt"
  "encoding/json"
  "github.com/buchanae/cwl"
  "os"
  "os/exec"
  "github.com/spf13/cobra"
)

var root = cobra.Command{
  Use: "cwl",
	SilenceUsage:  true,
}

func init() {
  cmd := &cobra.Command{
    Use: "dump <doc.cwl>",
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
      return dump(args[0])
    },
  }
  root.AddCommand(cmd)
}

func main() {
  if err := root.Execute(); err != nil {
    os.Exit(1)
  }
}

func dump(path string) error {
  doc, err := cwl.LoadFile(path)
  if err != nil {
    return err
  }

  b, err := json.MarshalIndent(doc, "", "  ")
  if err != nil {
    return err
  }

  fmt.Println(string(b))
  return nil
}

func init() {
  cmd := &cobra.Command{
    Use: "build <doc.cwl> <inputs.json>",
    Args: cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
      return build(args[0], args[1])
    },
  }
  root.AddCommand(cmd)
}

func build(path, inputsPath string) error {
  vals, err := cwl.LoadInputValuesFile(inputsPath)
  if err != nil {
    return err
  }

  doc, err := cwl.LoadFile(path)
  if err != nil {
    return err
  }

  clt, ok := doc.(*cwl.CommandLineTool)
  if !ok {
    return fmt.Errorf("can only build command line tools")
  }

  args, err := cwl.BuildCommand(clt, vals)
  if err != nil {
    return err
  }

  fmt.Println(args)
  return nil
}

func init() {
  cmd := &cobra.Command{
    Use: "run <doc.cwl> <inputs.json>",
    Args: cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
      return run(args[0], args[1])
    },
  }
  root.AddCommand(cmd)
}

func run(path, inputsPath string) error {
  vals, err := cwl.LoadInputValuesFile(inputsPath)
  if err != nil {
    return err
  }

  doc, err := cwl.LoadFile(path)
  if err != nil {
    return err
  }

  clt, ok := doc.(*cwl.CommandLineTool)
  if !ok {
    return fmt.Errorf("can only build command line tools")
  }

  args, err := cwl.BuildCommand(clt, vals)
  if err != nil {
    return err
  }

  cmd := exec.Command(args[0], args[1:]...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}
