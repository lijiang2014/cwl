package simple

import (
  "github.com/lijiang2014/cwl"
  "github.com/lijiang2014/cwl/process"
  "github.com/lijiang2014/cwl/process/fs/local"
)

type SimpleEnv struct {
  fs *local.Local
}

func NewSimpleEnv() *SimpleEnv {
	/*
	  id, err := uuid.NewRandom()
	  if err != nil {
	    return errf("error generating unique file location: %s", err)
	  }
	*/
  return &SimpleEnv{fs: local.NewLocal(".")}
}

func (s *SimpleEnv) Runtime() process.Runtime {
  return process.Runtime{}
}

func (s *SimpleEnv) Filesystem() process.Filesystem {
  return s.fs
}

func (s *SimpleEnv) CheckResources(req cwl.ResourceRequirement) error {
  return nil
}
