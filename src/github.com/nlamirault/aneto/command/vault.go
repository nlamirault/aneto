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
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/mitchellh/cli"

	"github.com/nlamirault/aneto/providers/glacier"
)

type VaultCommand struct {
	UI cli.Ui
}

func (c *VaultCommand) Help() string {
	helpText := `
Usage: aneto vault [options] --action=action
	Manage vault from Amazon Glacier
Options:
	--debug                       Debug mode enabled
	--name=name                   Vault name
        --region=name                 Region name
        --action=action               Action to perform

Action :
        list                        List vaults
        get                         Describe vault
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
	config := getAWSConfig(region, debug)
	switch action {
	case "list":
		c.doListGlacierVaults(config)
	case "get":
		c.doDescribeGlacierVault(config, name)
	case "display":
		c.doDisplayGlacierVault(config, name)
	case "create":
		c.doCreateGlacierVault(config, name)
	case "delete":
		c.doDeleteGlacierVault(config, name)
	}
	return 0
}

func (c *VaultCommand) doListGlacierVaults(config *aws.Config) {
	c.UI.Info("List Vaults :")
	result, err := glacier.ListVaults(glacier.GetGlacierClient(config))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(result))
	for _, vault := range result.VaultList {
		c.UI.Output(fmt.Sprintf("- %s %s [%d]",
			*vault.VaultName,
			*vault.CreationDate,
			*vault.NumberOfArchives))
	}
}

func (c *VaultCommand) doDescribeGlacierVault(config *aws.Config, name string) {
	c.UI.Info(fmt.Sprintf("Describe vault %s :", name))
	vault, err := glacier.DescribeVault(
		glacier.GetGlacierClient(config), name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(vault))
	c.UI.Output(fmt.Sprintf("- %s %s [%d]",
		*vault.VaultName,
		*vault.CreationDate,
		*vault.NumberOfArchives))
}

func (c *VaultCommand) doDisplayGlacierVault(config *aws.Config, name string) {
	c.UI.Info(fmt.Sprintf("Display vault inventory : %s", name))
	client := glacier.GetGlacierClient(config)
	job, err := glacier.DisplayVault(client, name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(job))
	statusCode := "InProgress"
	for statusCode == "InProgress" {
		desc, err := glacier.DescribeJob(client, name, *job.JobId)
		if err != nil {
			c.UI.Error(err.Error())
			return
		}
		log.Printf("[DEBUG] %s", awsutil.Prettify(desc))
		c.UI.Output(fmt.Sprintf("- %s ", *desc.StatusCode))
		time.Sleep(2000 * time.Millisecond)
	}
}

func (c *VaultCommand) doCreateGlacierVault(config *aws.Config, name string) {
	c.UI.Info(fmt.Sprintf("Create vault : %s", name))
	result, err := glacier.CreateVault(
		glacier.GetGlacierClient(config), name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(result))
}

func (c *VaultCommand) doDeleteGlacierVault(config *aws.Config, name string) {
	c.UI.Info(fmt.Sprintf("Delete vault : %s", name))
	result, err := glacier.DeleteVault(
		glacier.GetGlacierClient(config), name)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(result))
}
