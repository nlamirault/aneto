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
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/codegangsta/cli"
)

// Commands is the CLI commands
var Commands = []cli.Command{
	{
		Name:  "vault",
		Usage: "Manage Glacier vaults",
		Subcommands: []cli.Command{
			commandListVault,
			commandDescribeVault,
			commandDisplayVault,
			commandCreateVault,
			commandDeleteVault,
		},
	},
	{
		Name:  "archive",
		Usage: "Manage archives from backup",
		Subcommands: []cli.Command{
			commandGetArchive,
			commandPutArchive,
			commandDeleteArchive,
		},
	},
}

var commandListVault = cli.Command{
	Name:        "list",
	Usage:       "List Glacier vaults",
	Description: ``,
	Action:      doListGlacierVaults,
}
var commandDescribeVault = cli.Command{
	Name:        "desc",
	Usage:       "Describe a Glacier vault",
	Description: ``,
	Action:      doDescribeGlacierVault,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "Vault name",
		},
	},
}

var commandDisplayVault = cli.Command{
	Name:        "display",
	Usage:       "Display the Glacier vault inventory",
	Description: ``,
	Action:      doDisplayGlacierVault,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "Vault name",
		},
	},
}

var commandCreateVault = cli.Command{
	Name:        "create",
	Usage:       "Create a Glacier vault",
	Description: ``,
	Action:      doCreateGlacierVault,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "Vault name",
		},
	},
}

var commandDeleteVault = cli.Command{
	Name:        "delete",
	Usage:       "Delete a Glacier vault",
	Description: ``,
	Action:      doDeleteGlacierVault,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "Vault name",
		},
	},
}

var commandGetArchive = cli.Command{
	Name:        "get",
	Usage:       "Retrieve an archive from backup",
	Description: ``,
	Action:      doGetArchive,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "vault name",
		},
		cli.StringFlag{
			Name:  "key",
			Usage: "archive key name",
		},
	},
}

var commandPutArchive = cli.Command{
	Name:        "put",
	Usage:       "Store an archive into backup",
	Description: ``,
	Action:      doPutArchive,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "vault name",
		},
		cli.StringFlag{
			Name:  "key",
			Usage: "archive key name",
		},
		cli.StringFlag{
			Name:  "file",
			Usage: "archive filename",
		},
	},
}

var commandDeleteArchive = cli.Command{
	Name:        "delete",
	Usage:       "Delete an archive from backup",
	Description: ``,
	Action:      doDeleteArchive,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "vault name",
		},
		cli.StringFlag{
			Name:  "archiveID",
			Usage: "ID of the archive",
		},
	},
}

func doListGlacierVaults(c *cli.Context) {
	log.Println("Vaults")
	result, err := listVaults(getGlacierClient(getAWSConfig(region)))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}

func doDescribeGlacierVault(c *cli.Context) {
	log.Printf("Describe vault : %s\n", c.String("name"))
	result, err := describeVault(
		getGlacierClient(getAWSConfig(region)), c.String("name"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}

func doDisplayGlacierVault(c *cli.Context) {
	log.Printf("Display vault inventory : %s\n", c.String("name"))
	result, err := displayVault(
		getGlacierClient(getAWSConfig(region)), c.String("name"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}

func doCreateGlacierVault(c *cli.Context) {
	log.Printf("Create vault : %s\n", c.String("name"))
	result, err := createVault(
		getGlacierClient(getAWSConfig(region)), c.String("name"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}

func doDeleteGlacierVault(c *cli.Context) {
	log.Printf("Delete vault : %s\n", c.String("name"))
	result, err := deleteVault(
		getGlacierClient(getAWSConfig(region)), c.String("name"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}

func doGetArchive(c *cli.Context) {
	log.Printf("Retrieve archive %s\n", c.String("key"))
	// result, err := uploadArchive(
	// 	getAWSConfig(region), c.String("description"))
}

func doPutArchive(c *cli.Context) {
	log.Printf("Upload archive %s : %s\n", c.String("key"), c.String("file"))
	data, err := ioutil.ReadFile(c.String("file"))
	if err != nil {
		log.Printf("Can't read archive file : %v\n", err)
		return
	}
	result, err := uploadArchive(
		getGlacierClient(getAWSConfig(region)),
		c.String("name"), data, c.String("key"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}

func doDeleteArchive(c *cli.Context) {
	log.Printf("Delete archive %s\n", c.String("archiveID"))
	result, err := deleteArchive(getGlacierClient(getAWSConfig(region)),
		c.String("name"), c.String("archiveID"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}
