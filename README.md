[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Godoc](https://img.shields.io/badge/godoc-ref-blue.svg)](http://godoc.org/github.com/buchanae/cwl)

This is a Go library and command line tool for working with the [common workflow language (CWL)](http://commonwl.org). 

The core library includes the data type
present in the CWL spec (CommandLineTool, Workflow, etc.), with a few extra utilities for loading and working with CWL documents. 

For details, see the [reference docs](https://godoc.org/github.com/buchanae/cwl).

## Subpackages

The [process](./process) library contains experimental, unfinished code for processing CWL documents in order to execute commands and workflows.

The [expr](./expr) library contains utilities for parsing CWL expressions out of strings. This parser is not yet robust (see the known issues below).

## Alpha quality

At the time of this writing, this library is only a couple weeks old. I feel that the core CWL document loading library is fairly stable, but I can't promise that there aren't plenty of bugs lurking. 

The command line tool is far from stable, and needs work before becoming useful.

The [process](./process) library is highly experimental. Processing CWL tools and workflows in a robust manner is not a trivial task.

## CLI

The `cwl` command line tool includes a few commands for loading, inspecting, and experimental support for running commands. See the [releases](https://github.com/buchanae/cwl/releases) page to download the command line tool binary.

## Usage (CLI)

The command line tool is still young and therefore fairly useless. Still:
```
cwl dump https://raw.githubusercontent.com/buchanae/cwl/master/examples/array-inputs.cwl
// ...outputs the normalized document in JSON.
```

`cwl run` exists and is experimental. This command will run a CWL document, similar `cwltool`.

## Usage (library)

```go
package main

import (
  "log"
  "github.com/buchanae/cwl"
)

func main() {
  // Load a CWL CommandLineTool document.
  path := "https://raw.githubusercontent.com/buchanae/cwl/master/examples/array-inputs.cwl"
  doc, err := cwl.Load(path)
  if err != nil {
    log.Fatal(err)
  }
  
  tool := doc.(*cwl.Tool)
  
  // Print the ID of all the input fields
  for _, input := range tool.Inputs {
    log.Println(input.ID)
  }
}
```

```
go run main.go
2018/03/12 18:16:22 filesA
2018/03/12 18:16:22 filesB
2018/03/12 18:16:22 filesC
```

## Normalization

The CWL spec allows multiple different types for some fields, e.g. `CommandLineTool.inputs` may be a string, a list of strings, a map of string to string, a map of string to object, and so on. This is rather difficult to program against, especially in a statically typed language without generics or union types (i.e. Go).

This library normalizes all fields to a single type, often prefering a list where a string and map might be allowed. In the example above, [`CommandLineTool.inputs`](https://godoc.org/github.com/buchanae/cwl#Tool) is a list.

Similarly, many fields might be a [CWL expression](http://www.commonwl.org/v1.0/CommandLineTool.html#Expressions). In this library, any field which *might* be an expression has the type [`Expression`](https://godoc.org/github.com/buchanae/cwl#Expression).

## Notable changes and known issues

I've taken some liberties with the CWL spec:

- `CommandLineTool` is named `Tool` instead, for brevity.
- [Schema Salad](http://www.commonwl.org/v1.0/SchemaSalad.html) is not implemented and likely won't be implemented.
- `$include` and `$import` statements are not yet implemented, but will be.
- The CWL expression parser is not robust and will not correctly parse complex expressions, especially those containing `$()` and escaping.
- documentation and examples are still sparse, more on the way soon.

## Tests and Features 

* 0 PASS  General test of command line generation
* 1 PASS Test nested prefixes with arrays
* 2 FAIL: SchemaDefRequirement is not supported (yet)
* 3 PASS  Test command line with optional input (missing)
* 4 PASS  Test command line with optional input (provided)
* 5 FAIL: InitialWorkDirRequirement is not supported (yet)
* 6 PASS Test command execution in Docker with stdout redirection
* 7 PASS  Test command execution in Docker with simplified syntax stdout redirection
* 8 PASS  Test command execution in Docker with stdout redirection
* 9 PASS Test command line with stderr redirection
* 10 PASS  Test command line with stderr redirection, brief syntax
* 11 P Test command line with stderr redirection, named brief syntax
* 12 FAIL 未实现 STDIN 的功能
* 13 - 18 FAIL: Error: running doc: unknown doc type "ExpressionTool" ; Test default usage of Any in expressions.
* 19 FAIL: 未实现 ANY 的功能 ;   Testing Any type compatibility in outputSource

共 126 个测试，后面的进行省略

* 采用实际 App 进行测试 ：

* hpc-batch lammps
* hpc-session dcv
* k8s-batch fit-10
* k8s-session tesorflow

** 添加了 ClusterRequirement 用来进行作业的路由描述

和 endpoint 的关联关系

* job 状态为 running 后建立 endpoint ,并 建立 job 和 endpoint 的关联
* 从R队列中删除 job 时 ， 关联删除 endpoint 记录

