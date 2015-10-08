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

package command

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/mitchellh/cli"

	"github.com/nlamirault/aneto/logging"
	"github.com/nlamirault/aneto/providers/glacier"
)

type ArchiveCommand struct {
	UI cli.Ui
}

func (c *ArchiveCommand) Help() string {
	helpText := `
Usage: aneto archive [action]
	Manage archive from Amazon Glacier
Options:
	--debug                       Debug mode enabled
	--name=name                   Archive name
        --region=name                 Region name
        --action=action               Action to perform

Action :
        get                         Get an archive
        put                         Put an archive
        delete                      Delete an archive
`
	return strings.TrimSpace(helpText)
}

func (c *ArchiveCommand) Synopsis() string {
	return "Manage archive from Amazon Glacier"
}

func (c *ArchiveCommand) Run(args []string) int {
	var debug bool
	var name, region, action string
	f := flag.NewFlagSet("archive", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&name, "name", "", "Glacier vault's name")
	f.StringVar(&region, "region", "", "AWS region name")
	f.StringVar(&action, "action", "", "Action to perform")

	if err := f.Parse(args); err != nil {
		return 1
	}
	if debug {
		c.UI.Info("Debug mode enabled.")
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	log.Printf("[DEBUG] Args : %s", args)
	return 0
}

func (c *ArchiveCommand) doGetArchive(key string) {
	c.UI.Info(fmt.Sprintf("Retrieve archive %s\n", key))
	// result, err := glacier.UploadArchive(
	// 	getAWSConfig(region), c.String("description"))
}

func (c *ArchiveCommand) doPutArchive(region string, name string, key string, file string) {
	c.UI.Info(fmt.Sprintf("Upload archive %s : %s\n", key, file))
	data, err := ioutil.ReadFile(file)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	result, err := glacier.UploadArchive(
		glacier.GetGlacierClient(getAWSConfig(region)),
		name,
		data,
		key)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info(awsutil.Prettify(result))
}

func (c *ArchiveCommand) doDeleteArchive(region string, name string, archiveID string) {
	c.UI.Info(fmt.Sprintf("Delete archive %s %s\n", name, archiveID))
	result, err := glacier.DeleteArchive(
		glacier.GetGlacierClient(getAWSConfig(region)),
		name,
		archiveID)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info(awsutil.Prettify(result))
}
