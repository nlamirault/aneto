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
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/mitchellh/cli"

	"github.com/nlamirault/aneto/logging"
	"github.com/nlamirault/aneto/providers/glacier"
)

type VaultCommand struct {
	UI cli.Ui
}

func (c *VaultCommand) Help() string {
	helpText := `
Usage: aneto vault [options]
	Manage vault from Amazon Glacier
Options:
	--debug                       Debug mode enabled
	--name=name                   Vault name
        --region=name                 Region name
        --action=action               Action to perform

Action :
        list                        List vaults
        desc                        Describe vault
        display                     Display the vault inventory
        create                      Create a new vault
        delete                      Delete an existing vault
`
	return strings.TrimSpace(helpText)
}

func (c *VaultCommand) Synopsis() string {
	return "Manage vault from Amazon Glacier"
}

func (c *VaultCommand) Run(args []string) int {
	var debug bool
	var region, name, action string
	f := flag.NewFlagSet("vault", flag.ContinueOnError)
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
	log.Printf("[DEBUG] Args : %s %s", args, action)
	switch action {
	case "list":
		c.doListGlacierVaults(region)
	case "get":
		c.doDescribeGlacierVault(region, name)
	case "display":
		c.doDisplayGlacierVault(region, name)
	case "create":
		c.doCreateGlacierVault(region, name)
	case "delete":
		c.doDeleteGlacierVault(region, name)
	}
	return 0
}

func (c *VaultCommand) doListGlacierVaults(region string) {
	c.UI.Info("List Vaults")
	result, err := glacier.ListVaults(
		glacier.GetGlacierClient(getAWSConfig(region)))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info(awsutil.Prettify(result))
}

func (c *VaultCommand) doDescribeGlacierVault(region string, name string) {
	c.UI.Info(fmt.Sprintf("Describe vault : %s\n", name))
	result, err := glacier.DescribeVault(
		glacier.GetGlacierClient(getAWSConfig(region)), name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info(awsutil.Prettify(result))
}

func (c *VaultCommand) doDisplayGlacierVault(region string, name string) {
	c.UI.Info(fmt.Sprintf("Display vault inventory : %s\n", name))
	result, err := glacier.DisplayVault(
		glacier.GetGlacierClient(getAWSConfig(region)), name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info(awsutil.Prettify(result))
}

func (c *VaultCommand) doCreateGlacierVault(region string, name string) {
	c.UI.Info(fmt.Sprintf("Create vault : %s\n", name))
	result, err := glacier.CreateVault(
		glacier.GetGlacierClient(getAWSConfig(region)), name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info(awsutil.Prettify(result))
}

func (c *VaultCommand) doDeleteGlacierVault(region string, name string) {
	c.UI.Info(fmt.Sprintf("Delete vault : %s\n", name))
	result, err := glacier.DeleteVault(
		glacier.GetGlacierClient(getAWSConfig(region)), name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info(awsutil.Prettify(result))
}
