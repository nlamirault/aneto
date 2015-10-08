// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	// "flag"
	"fmt"
	// "log"
	"os"

	"github.com/mitchellh/cli"

	"github.com/nlamirault/aneto/version"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	cli := &cli.CLI{
		Args:       os.Args[1:],
		Commands:   Commands,
		HelpFunc:   cli.BasicHelpFunc("aneto"),
		HelpWriter: os.Stdout,
		Version:    version.Version,
	}

	exitCode, err := cli.Run()
	if err != nil {
		Ui.Error(fmt.Sprintf("Error executing CLI: %s", err.Error()))
		return 1
	}

	return exitCode
}
