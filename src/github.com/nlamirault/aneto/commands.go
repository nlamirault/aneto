// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
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
			commandCreateVault,
			commandDeleteVault,
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

func doUploadArchive(c *cli.Context) {
	log.Printf("Upload archive %s\n", c.String("file"))
	// result, err := uploadArchive(
	// 	getAWSConfig(region), c.String("description"))
}
