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

package glacier

import (
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glacier"
)

// func getGlacierClient return Glacier service client
func GetGlacierClient(cfg *aws.Config) *glacier.Glacier {
	c := glacier.New(cfg)
	return c
}

func ListVaults(glacierClient *glacier.Glacier) (*glacier.ListVaultsOutput, error) {
	log.Printf("[DEBUG] List Glacier vaults")
	return glacierClient.ListVaults(&glacier.ListVaultsInput{})
}

func DescribeVault(glacierClient *glacier.Glacier, name string) (*glacier.DescribeVaultOutput, error) {
	log.Printf("[DEBUG] Call describe Glacier vault : %s", name)
	return glacierClient.DescribeVault(&glacier.DescribeVaultInput{
		VaultName: aws.String(name),
	})
}

func DisplayVault(glacierClient *glacier.Glacier, name string) (*glacier.InitiateJobOutput, error) {
	log.Printf("[DEBUG] Call Display Glacier vault : %s", name)
	params := &glacier.InitiateJobInput{
		VaultName: aws.String(name),
		JobParameters: &glacier.JobParameters{
			InventoryRetrievalParameters: &glacier.InventoryRetrievalJobInput{},
			Type: aws.String("inventory-retrieval"),
		},
	}
	return glacierClient.InitiateJob(params)
}

func CreateVault(glacierClient *glacier.Glacier, name string) (*glacier.CreateVaultOutput, error) {
	log.Printf("[DEBUG] Call create Glacier vault : %s", name)
	return glacierClient.CreateVault(&glacier.CreateVaultInput{
		VaultName: aws.String(name),
	})
}

func DeleteVault(glacierClient *glacier.Glacier, name string) (*glacier.DeleteVaultOutput, error) {
	log.Printf("[DEBUG] Call Delete Glacier vault : %s", name)
	return glacierClient.DeleteVault(&glacier.DeleteVaultInput{
		VaultName: aws.String(name),
	})
}

func UploadArchive(glacierClient *glacier.Glacier, name string, archive []byte,
	description string) (*glacier.ArchiveCreationOutput, error) {
	log.Printf("[DEBUG] Call Upload archive to Glacier vault : %s", name)
	return glacierClient.UploadArchive(&glacier.UploadArchiveInput{
		VaultName:          aws.String(name),
		ArchiveDescription: aws.String(description),
		Body:               bytes.NewReader(archive),
	})
}

func DownloadArchive(glacierClient *glacier.Glacier, name string) {
	log.Printf("[DEBUG] Call Download archive from Glacier vault : %s", name)
}

func DeleteArchive(glacierClient *glacier.Glacier, name string,
	archiveID string) (*glacier.DeleteArchiveOutput, error) {
	log.Printf("[DEBUG] Call Delete archive from Glacier vault : %s %s",
		name, archiveID)
	return glacierClient.DeleteArchive(&glacier.DeleteArchiveInput{
		VaultName: aws.String(name),
		ArchiveId: aws.String(archiveID),
	})
}
