package process

import (
	"fmt"
	"github.com/lijiang2014/cwl"
	"github.com/kr/pretty"
	"os"
	"strings"
)

// errf makes fmt.Errorf shorter
func errf(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}

func wrap(err error, msg string, args ...interface{}) error {
	return errf("%s: %s", fmt.Sprintf(msg, args...), err)
}

// getPos is a helper for accessing the Position field
// of a possibly nil CommandLineBinding
func getPos(in *cwl.CommandLineBinding) int {
	if in == nil {
		return 0
	}
	return in.Position
}

func debug(args ...interface{}) {
	var fmts []string
	var formatters []interface{}
	for _, arg := range args {
		fmts = append(fmts, "%# v")
		formatters = append(formatters, pretty.Formatter(arg))
	}
	fmt.Fprintf(os.Stderr, strings.Join(fmts, " ")+"\n", formatters...)
}
