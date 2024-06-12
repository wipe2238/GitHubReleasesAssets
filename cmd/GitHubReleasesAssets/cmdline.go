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
	var re2 = regexp.MustCompile("^[A-Za-z0-9-_]+/[A-Za-z0-9-_]+$")
	var re3 = regexp.MustCompile("^[A-Za-z0-9-_]+$")

	if len(os.Args) == 2 &&
		re2.MatchString(os.Args[1]) {
		CmdLine.Repository = os.Args[1]
	} else if len(os.Args) == 3 &&
		re3.MatchString(os.Args[1]) &&
		re3.MatchString(os.Args[2]) {
		CmdLine.Repository = fmt.Sprintf("%s/%s", os.Args[1], os.Args[2])
	} else {
		var BaseName string = filepath.Base(os.Args[0])

		fmt.Printf("USAGE:  %s <owner>/<repo>\n        %s <owner> <repo>", BaseName, BaseName)
		os.Exit(1)
	}
}
