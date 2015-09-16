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
