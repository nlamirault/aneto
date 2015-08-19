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
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
)

var (
	region        string
	vault         string
	printVersion  bool
	doCreateVault bool
	doDeleteVault bool
)

func init() {
	flag.StringVar(&region, "region", "eu-west-1", "aws region")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")
	flag.BoolVar(&doCreateVault, "create", false, "create vault")
	flag.BoolVar(&doDeleteVault, "delete", false, "delete vault")
	flag.StringVar(&vault, "vault", "", "vault name")

}

func main() {
	flag.Parse()
	if printVersion {
		log.Println("Version", Version)
		os.Exit(0)
	}
	if doCreateVault {
		checkArgument(vault, "Glacier vault name")
		createGlacierVault()
	}
	if doDeleteVault {
		checkArgument(vault, "Glacier vault name")
		// deleteVault()
	}

}

func checkArgument(key string, value string) {
	if key == "" {
		log.Printf("Please specify %s. Exiting.\n", value)
		os.Exit(1)
	}
}

func getVaultName() string {
	return os.Getenv("ANETO_VAULT_NAME")
}

func getAWSConfig(region string) *aws.Config {
	return &aws.Config{Region: aws.String(region)}

}

func createGlacierVault() {
	log.Printf("Create vault : %s\n", vault)
	result, err := createVault(getGlacierClient(getAWSConfig(region)), vault)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}
