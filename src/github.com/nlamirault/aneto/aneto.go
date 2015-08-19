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
	// "flag"
	"fmt"
	// "log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/codegangsta/cli"
)

// region represents the default AWS region
const region string = "eu-west-1"

func getAWSConfig(region string) *aws.Config {
	return &aws.Config{Region: aws.String(region)}
}

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "aneto"
	app.Version = Version
	app.Usage = "A backup tool"
	app.Author = "Nicolas Lamirault"
	app.Email = "nicolas.lamirault@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "region",
			Value: "eu-west-1",
			Usage: fmt.Sprintf("AWS region"),
		},
	}
	app.Commands = Commands
	return app
}

func main() {
	app := makeApp()
	app.Run(os.Args)
}
