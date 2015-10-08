// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sts

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/internal/protocol/query"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// The AWS Security Token Service (STS) is a web service that enables you to
// request temporary, limited-privilege credentials for AWS Identity and Access
// Management (IAM) users or for users that you authenticate (federated users).
// This guide provides descriptions of the STS API. For more detailed information
// about using this service, go to Using Temporary Security Credentials (http://docs.aws.amazon.com/STS/latest/UsingSTS/Welcome.html"
// target="_blank).
//
//  As an alternative to using the API, you can use one of the AWS SDKs, which
// consist of libraries and sample code for various programming languages and
// platforms (Java, Ruby, .NET, iOS, Android, etc.). The SDKs provide a convenient
// way to create programmatic access to STS. For example, the SDKs take care
// of cryptographically signing requests, managing errors, and retrying requests
// automatically. For information about the AWS SDKs, including how to download
// and install them, see the Tools for Amazon Web Services page (http://aws.amazon.com/tools/).
//  For information about setting up signatures and authorization through the
// API, go to Signing AWS API Requests (http://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html"
// target="_blank) in the AWS General Reference. For general information about
// the Query API, go to Making Query Requests (http://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html"
// target="_blank) in Using IAM. For information about using security tokens
// with other AWS products, go to Using Temporary Security Credentials to Access
// AWS (http://docs.aws.amazon.com/STS/latest/UsingSTS/UsingTokens.html) in
// Using Temporary Security Credentials.
//
// If you're new to AWS and need additional technical information about a specific
// AWS product, you can find the product's technical documentation at http://aws.amazon.com/documentation/
// (http://aws.amazon.com/documentation/" target="_blank).
//
//  Endpoints
//
// The AWS Security Token Service (STS) has a default endpoint of https://sts.amazonaws.com
// that maps to the US East (N. Virginia) region. Additional regions are available,
// but must first be activated in the AWS Management Console before you can
// use a different region's endpoint. For more information about activating
// a region for STS see Activating STS in a New Region (http://docs.aws.amazon.com/STS/latest/UsingSTS/sts-enableregions.html)
// in the Using Temporary Security Credentials guide.
//
// For information about STS endpoints, see Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#sts_region)
// in the AWS General Reference.
//
//  Recording API requests
//
// STS supports AWS CloudTrail, which is a service that records AWS calls for
// your AWS account and delivers log files to an Amazon S3 bucket. By using
// information collected by CloudTrail, you can determine what requests were
// successfully made to STS, who made the request, when it was made, and so
// on. To learn more about CloudTrail, including how to turn it on and find
// your log files, see the AWS CloudTrail User Guide (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/what_is_cloud_trail_top_level.html).
type STS struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new STS client.
func New(config *aws.Config) *STS {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "sts",
			APIVersion:  "2011-06-15",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &STS{service}
}

// newRequest creates a new request for a STS operation and runs any
// custom request initialization.
func (c *STS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
