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
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glacier"
)

// func getGlacierClient return Glacier service client
func getGlacierClient(cfg *aws.Config) *glacier.Glacier {
	c := glacier.New(cfg)
	return c
}

func listVaults(glacierClient *glacier.Glacier) (*glacier.ListVaultsOutput, error) {
	return glacierClient.ListVaults(&glacier.ListVaultsInput{})
}

func describeVault(glacierClient *glacier.Glacier, name string) (*glacier.DescribeVaultOutput, error) {
	return glacierClient.DescribeVault(&glacier.DescribeVaultInput{
		VaultName: aws.String(name),
	})
}

func displayVault(glacierClient *glacier.Glacier, name string) (*glacier.InitiateJobOutput, error) {

	params := &glacier.InitiateJobInput{
		VaultName: aws.String(name),
		JobParameters: &glacier.JobParameters{
			InventoryRetrievalParameters: &glacier.InventoryRetrievalJobInput{},
			Type: aws.String("inventory-retrieval"),
		},
	}

	return glacierClient.InitiateJob(params)
}

func createVault(glacierClient *glacier.Glacier, name string) (*glacier.CreateVaultOutput, error) {
	return glacierClient.CreateVault(&glacier.CreateVaultInput{
		VaultName: aws.String(name),
	})
}

func deleteVault(glacierClient *glacier.Glacier, name string) (*glacier.DeleteVaultOutput, error) {
	return glacierClient.DeleteVault(&glacier.DeleteVaultInput{
		VaultName: aws.String(name),
	})
}

func uploadArchive(glacierClient *glacier.Glacier, name string, archive []byte,
	description string) (*glacier.ArchiveCreationOutput, error) {
	return glacierClient.UploadArchive(&glacier.UploadArchiveInput{
		VaultName:          aws.String(name),
		ArchiveDescription: aws.String(description),
		Body:               bytes.NewReader(archive),
	})
}

func downloadArchive() {
}

func deleteArchive(glacierClient *glacier.Glacier, name string,
	archiveID string) (*glacier.DeleteArchiveOutput, error) {
	return glacierClient.DeleteArchive(&glacier.DeleteArchiveInput{
		VaultName: aws.String(name),
		ArchiveId: aws.String(archiveID),
	})
}
