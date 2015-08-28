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
