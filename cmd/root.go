package cmd

import (
	"fmt"
	"os"

	"github.com/Morizz00/codebase-analyzer/analyzer"
)

func Execute() {
	args := os.Args
	targ := "."
	if len(args) > 1 {
		targ = args[1]
	}
	stats := analyzer.Scan(targ)
	for _, s := range stats {
		fmt.Printf("%s (%s)--%d lines,%d bytes\n", s.Path, s.Language, s.Lines, s.Size)
	}
}
