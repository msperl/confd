package support

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// Support is a client for AWS Support.
type Support struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new Support client.
func New(config *aws.Config) *Support {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "support",
		APIVersion:   "2013-04-15",
		JSONVersion:  "1.1",
		TargetPrefix: "AWSSupport_20130415",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &Support{service}
}

// newRequest creates a new request for a Support operation and runs any
// custom request initialization.
func (c *Support) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
