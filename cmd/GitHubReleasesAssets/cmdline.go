package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var CmdLine struct {
	Repository string
}

func init() {
	var re = regexp.MustCompile("^[A-Za-z0-9-_]+/[A-Za-z0-9-_]+$")

	if len(os.Args) >= 2 && re.MatchString(os.Args[1]) {
		CmdLine.Repository = os.Args[1]
	} else {
		var BaseName string = filepath.Base(os.Args[0])

		fmt.Printf("USAGE:  %s <owner>/<repo>\n", BaseName)
		os.Exit(1)
	}
}
