// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatchlogs provides a client for Amazon CloudWatch Logs.
package cloudwatchlogs

import (
	"sync"

	"github.com/awslabs/aws-sdk-go/aws"
)

var oprw sync.Mutex

// CreateLogGroupRequest generates a request for the CreateLogGroup operation.
func (c *CloudWatchLogs) CreateLogGroupRequest(input *CreateLogGroupInput) (req *aws.Request, output *CreateLogGroupOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateLogGroup == nil {
		opCreateLogGroup = &aws.Operation{
			Name:       "CreateLogGroup",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateLogGroupInput{}
	}

	req = c.newRequest(opCreateLogGroup, input, output)
	output = &CreateLogGroupOutput{}
	req.Data = output
	return
}

// Creates a new log group with the specified name. The name of the log group
// must be unique within a region for an AWS account. You can create up to 500
// log groups per account.
//
//  You must use the following guidelines when naming a log group:  Log group
// names can be between 1 and 512 characters long. Allowed characters are a-z,
// A-Z, 0-9, '_' (underscore), '-' (hyphen), '/' (forward slash), and '.' (period).
func (c *CloudWatchLogs) CreateLogGroup(input *CreateLogGroupInput) (*CreateLogGroupOutput, error) {
	req, out := c.CreateLogGroupRequest(input)
	err := req.Send()
	return out, err
}

var opCreateLogGroup *aws.Operation

// CreateLogStreamRequest generates a request for the CreateLogStream operation.
func (c *CloudWatchLogs) CreateLogStreamRequest(input *CreateLogStreamInput) (req *aws.Request, output *CreateLogStreamOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateLogStream == nil {
		opCreateLogStream = &aws.Operation{
			Name:       "CreateLogStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateLogStreamInput{}
	}

	req = c.newRequest(opCreateLogStream, input, output)
	output = &CreateLogStreamOutput{}
	req.Data = output
	return
}

// Creates a new log stream in the specified log group. The name of the log
// stream must be unique within the log group. There is no limit on the number
// of log streams that can exist in a log group.
//
//  You must use the following guidelines when naming a log stream:  Log stream
// names can be between 1 and 512 characters long. The ':' colon character is
// not allowed.
func (c *CloudWatchLogs) CreateLogStream(input *CreateLogStreamInput) (*CreateLogStreamOutput, error) {
	req, out := c.CreateLogStreamRequest(input)
	err := req.Send()
	return out, err
}

var opCreateLogStream *aws.Operation

// DeleteLogGroupRequest generates a request for the DeleteLogGroup operation.
func (c *CloudWatchLogs) DeleteLogGroupRequest(input *DeleteLogGroupInput) (req *aws.Request, output *DeleteLogGroupOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteLogGroup == nil {
		opDeleteLogGroup = &aws.Operation{
			Name:       "DeleteLogGroup",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteLogGroupInput{}
	}

	req = c.newRequest(opDeleteLogGroup, input, output)
	output = &DeleteLogGroupOutput{}
	req.Data = output
	return
}

// Deletes the log group with the specified name and permanently deletes all
// the archived log events associated with it.
func (c *CloudWatchLogs) DeleteLogGroup(input *DeleteLogGroupInput) (*DeleteLogGroupOutput, error) {
	req, out := c.DeleteLogGroupRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteLogGroup *aws.Operation

// DeleteLogStreamRequest generates a request for the DeleteLogStream operation.
func (c *CloudWatchLogs) DeleteLogStreamRequest(input *DeleteLogStreamInput) (req *aws.Request, output *DeleteLogStreamOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteLogStream == nil {
		opDeleteLogStream = &aws.Operation{
			Name:       "DeleteLogStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteLogStreamInput{}
	}

	req = c.newRequest(opDeleteLogStream, input, output)
	output = &DeleteLogStreamOutput{}
	req.Data = output
	return
}

// Deletes a log stream and permanently deletes all the archived log events
// associated with it.
func (c *CloudWatchLogs) DeleteLogStream(input *DeleteLogStreamInput) (*DeleteLogStreamOutput, error) {
	req, out := c.DeleteLogStreamRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteLogStream *aws.Operation

// DeleteMetricFilterRequest generates a request for the DeleteMetricFilter operation.
func (c *CloudWatchLogs) DeleteMetricFilterRequest(input *DeleteMetricFilterInput) (req *aws.Request, output *DeleteMetricFilterOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteMetricFilter == nil {
		opDeleteMetricFilter = &aws.Operation{
			Name:       "DeleteMetricFilter",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteMetricFilterInput{}
	}

	req = c.newRequest(opDeleteMetricFilter, input, output)
	output = &DeleteMetricFilterOutput{}
	req.Data = output
	return
}

// Deletes a metric filter associated with the specified log group.
func (c *CloudWatchLogs) DeleteMetricFilter(input *DeleteMetricFilterInput) (*DeleteMetricFilterOutput, error) {
	req, out := c.DeleteMetricFilterRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteMetricFilter *aws.Operation

// DeleteRetentionPolicyRequest generates a request for the DeleteRetentionPolicy operation.
func (c *CloudWatchLogs) DeleteRetentionPolicyRequest(input *DeleteRetentionPolicyInput) (req *aws.Request, output *DeleteRetentionPolicyOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteRetentionPolicy == nil {
		opDeleteRetentionPolicy = &aws.Operation{
			Name:       "DeleteRetentionPolicy",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteRetentionPolicyInput{}
	}

	req = c.newRequest(opDeleteRetentionPolicy, input, output)
	output = &DeleteRetentionPolicyOutput{}
	req.Data = output
	return
}

// Deletes the retention policy of the specified log group. Log events would
// not expire if they belong to log groups without a retention policy.
func (c *CloudWatchLogs) DeleteRetentionPolicy(input *DeleteRetentionPolicyInput) (*DeleteRetentionPolicyOutput, error) {
	req, out := c.DeleteRetentionPolicyRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteRetentionPolicy *aws.Operation

// DescribeLogGroupsRequest generates a request for the DescribeLogGroups operation.
func (c *CloudWatchLogs) DescribeLogGroupsRequest(input *DescribeLogGroupsInput) (req *aws.Request, output *DescribeLogGroupsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeLogGroups == nil {
		opDescribeLogGroups = &aws.Operation{
			Name:       "DescribeLogGroups",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeLogGroupsInput{}
	}

	req = c.newRequest(opDescribeLogGroups, input, output)
	output = &DescribeLogGroupsOutput{}
	req.Data = output
	return
}

// Returns all the log groups that are associated with the AWS account making
// the request. The list returned in the response is ASCII-sorted by log group
// name.
//
//  By default, this operation returns up to 50 log groups. If there are more
// log groups to list, the response would contain a nextToken value in the response
// body. You can also limit the number of log groups returned in the response
// by specifying the limit parameter in the request.
func (c *CloudWatchLogs) DescribeLogGroups(input *DescribeLogGroupsInput) (*DescribeLogGroupsOutput, error) {
	req, out := c.DescribeLogGroupsRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeLogGroups *aws.Operation

// DescribeLogStreamsRequest generates a request for the DescribeLogStreams operation.
func (c *CloudWatchLogs) DescribeLogStreamsRequest(input *DescribeLogStreamsInput) (req *aws.Request, output *DescribeLogStreamsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeLogStreams == nil {
		opDescribeLogStreams = &aws.Operation{
			Name:       "DescribeLogStreams",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeLogStreamsInput{}
	}

	req = c.newRequest(opDescribeLogStreams, input, output)
	output = &DescribeLogStreamsOutput{}
	req.Data = output
	return
}

// Returns all the log streams that are associated with the specified log group.
// The list returned in the response is ASCII-sorted by log stream name.
//
//  By default, this operation returns up to 50 log streams. If there are more
// log streams to list, the response would contain a nextToken value in the
// response body. You can also limit the number of log streams returned in the
// response by specifying the limit parameter in the request. This operation
// has a limit of five transactions per second, after which transactions are
// throttled.
func (c *CloudWatchLogs) DescribeLogStreams(input *DescribeLogStreamsInput) (*DescribeLogStreamsOutput, error) {
	req, out := c.DescribeLogStreamsRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeLogStreams *aws.Operation

// DescribeMetricFiltersRequest generates a request for the DescribeMetricFilters operation.
func (c *CloudWatchLogs) DescribeMetricFiltersRequest(input *DescribeMetricFiltersInput) (req *aws.Request, output *DescribeMetricFiltersOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeMetricFilters == nil {
		opDescribeMetricFilters = &aws.Operation{
			Name:       "DescribeMetricFilters",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeMetricFiltersInput{}
	}

	req = c.newRequest(opDescribeMetricFilters, input, output)
	output = &DescribeMetricFiltersOutput{}
	req.Data = output
	return
}

// Returns all the metrics filters associated with the specified log group.
// The list returned in the response is ASCII-sorted by filter name.
//
//  By default, this operation returns up to 50 metric filters. If there are
// more metric filters to list, the response would contain a nextToken value
// in the response body. You can also limit the number of metric filters returned
// in the response by specifying the limit parameter in the request.
func (c *CloudWatchLogs) DescribeMetricFilters(input *DescribeMetricFiltersInput) (*DescribeMetricFiltersOutput, error) {
	req, out := c.DescribeMetricFiltersRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeMetricFilters *aws.Operation

// GetLogEventsRequest generates a request for the GetLogEvents operation.
func (c *CloudWatchLogs) GetLogEventsRequest(input *GetLogEventsInput) (req *aws.Request, output *GetLogEventsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetLogEvents == nil {
		opGetLogEvents = &aws.Operation{
			Name:       "GetLogEvents",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetLogEventsInput{}
	}

	req = c.newRequest(opGetLogEvents, input, output)
	output = &GetLogEventsOutput{}
	req.Data = output
	return
}

// Retrieves log events from the specified log stream. You can provide an optional
// time range to filter the results on the event timestamp.
//
//  By default, this operation returns as much log events as can fit in a response
// size of 1MB, up to 10,000 log events. The response will always include a
// nextForwardToken and a nextBackwardToken in the response body. You can use
// any of these tokens in subsequent GetLogEvents requests to paginate through
// events in either forward or backward direction. You can also limit the number
// of log events returned in the response by specifying the limit parameter
// in the request.
func (c *CloudWatchLogs) GetLogEvents(input *GetLogEventsInput) (*GetLogEventsOutput, error) {
	req, out := c.GetLogEventsRequest(input)
	err := req.Send()
	return out, err
}

var opGetLogEvents *aws.Operation

// PutLogEventsRequest generates a request for the PutLogEvents operation.
func (c *CloudWatchLogs) PutLogEventsRequest(input *PutLogEventsInput) (req *aws.Request, output *PutLogEventsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opPutLogEvents == nil {
		opPutLogEvents = &aws.Operation{
			Name:       "PutLogEvents",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &PutLogEventsInput{}
	}

	req = c.newRequest(opPutLogEvents, input, output)
	output = &PutLogEventsOutput{}
	req.Data = output
	return
}

// Uploads a batch of log events to the specified log stream.
//
//  Every PutLogEvents request must include the sequenceToken obtained from
// the response of the previous request. An upload in a newly created log stream
// does not require a sequenceToken.
//
//  The batch of events must satisfy the following constraints:  The maximum
// batch size is 1,048,576 bytes, and this size is calculated as the sum of
// all event messages in UTF-8, plus 26 bytes for each log event. None of the
// log events in the batch can be more than 2 hours in the future. None of the
// log events in the batch can be older than 14 days or the retention period
// of the log group. The log events in the batch must be in chronological ordered
// by their timestamp. The maximum number of log events in a batch is 10,000.
func (c *CloudWatchLogs) PutLogEvents(input *PutLogEventsInput) (*PutLogEventsOutput, error) {
	req, out := c.PutLogEventsRequest(input)
	err := req.Send()
	return out, err
}

var opPutLogEvents *aws.Operation

// PutMetricFilterRequest generates a request for the PutMetricFilter operation.
func (c *CloudWatchLogs) PutMetricFilterRequest(input *PutMetricFilterInput) (req *aws.Request, output *PutMetricFilterOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opPutMetricFilter == nil {
		opPutMetricFilter = &aws.Operation{
			Name:       "PutMetricFilter",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &PutMetricFilterInput{}
	}

	req = c.newRequest(opPutMetricFilter, input, output)
	output = &PutMetricFilterOutput{}
	req.Data = output
	return
}

// Creates or updates a metric filter and associates it with the specified log
// group. Metric filters allow you to configure rules to extract metric data
// from log events ingested through PutLogEvents requests.
func (c *CloudWatchLogs) PutMetricFilter(input *PutMetricFilterInput) (*PutMetricFilterOutput, error) {
	req, out := c.PutMetricFilterRequest(input)
	err := req.Send()
	return out, err
}

var opPutMetricFilter *aws.Operation

// PutRetentionPolicyRequest generates a request for the PutRetentionPolicy operation.
func (c *CloudWatchLogs) PutRetentionPolicyRequest(input *PutRetentionPolicyInput) (req *aws.Request, output *PutRetentionPolicyOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opPutRetentionPolicy == nil {
		opPutRetentionPolicy = &aws.Operation{
			Name:       "PutRetentionPolicy",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &PutRetentionPolicyInput{}
	}

	req = c.newRequest(opPutRetentionPolicy, input, output)
	output = &PutRetentionPolicyOutput{}
	req.Data = output
	return
}

// Sets the retention of the specified log group. A retention policy allows
// you to configure the number of days you want to retain log events in the
// specified log group.
func (c *CloudWatchLogs) PutRetentionPolicy(input *PutRetentionPolicyInput) (*PutRetentionPolicyOutput, error) {
	req, out := c.PutRetentionPolicyRequest(input)
	err := req.Send()
	return out, err
}

var opPutRetentionPolicy *aws.Operation

// TestMetricFilterRequest generates a request for the TestMetricFilter operation.
func (c *CloudWatchLogs) TestMetricFilterRequest(input *TestMetricFilterInput) (req *aws.Request, output *TestMetricFilterOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opTestMetricFilter == nil {
		opTestMetricFilter = &aws.Operation{
			Name:       "TestMetricFilter",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &TestMetricFilterInput{}
	}

	req = c.newRequest(opTestMetricFilter, input, output)
	output = &TestMetricFilterOutput{}
	req.Data = output
	return
}

// Tests the filter pattern of a metric filter against a sample of log event
// messages. You can use this operation to validate the correctness of a metric
// filter pattern.
func (c *CloudWatchLogs) TestMetricFilter(input *TestMetricFilterInput) (*TestMetricFilterOutput, error) {
	req, out := c.TestMetricFilterRequest(input)
	err := req.Send()
	return out, err
}

var opTestMetricFilter *aws.Operation

type CreateLogGroupInput struct {
	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	metadataCreateLogGroupInput `json:"-" xml:"-"`
}

type metadataCreateLogGroupInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateLogGroupOutput struct {
	metadataCreateLogGroupOutput `json:"-" xml:"-"`
}

type metadataCreateLogGroupOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateLogStreamInput struct {
	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	LogStreamName *string `locationName:"logStreamName" type:"string" required:"true"`

	metadataCreateLogStreamInput `json:"-" xml:"-"`
}

type metadataCreateLogStreamInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateLogStreamOutput struct {
	metadataCreateLogStreamOutput `json:"-" xml:"-"`
}

type metadataCreateLogStreamOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteLogGroupInput struct {
	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	metadataDeleteLogGroupInput `json:"-" xml:"-"`
}

type metadataDeleteLogGroupInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteLogGroupOutput struct {
	metadataDeleteLogGroupOutput `json:"-" xml:"-"`
}

type metadataDeleteLogGroupOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteLogStreamInput struct {
	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	LogStreamName *string `locationName:"logStreamName" type:"string" required:"true"`

	metadataDeleteLogStreamInput `json:"-" xml:"-"`
}

type metadataDeleteLogStreamInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteLogStreamOutput struct {
	metadataDeleteLogStreamOutput `json:"-" xml:"-"`
}

type metadataDeleteLogStreamOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMetricFilterInput struct {
	// The name of the metric filter.
	FilterName *string `locationName:"filterName" type:"string" required:"true"`

	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	metadataDeleteMetricFilterInput `json:"-" xml:"-"`
}

type metadataDeleteMetricFilterInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMetricFilterOutput struct {
	metadataDeleteMetricFilterOutput `json:"-" xml:"-"`
}

type metadataDeleteMetricFilterOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteRetentionPolicyInput struct {
	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	metadataDeleteRetentionPolicyInput `json:"-" xml:"-"`
}

type metadataDeleteRetentionPolicyInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteRetentionPolicyOutput struct {
	metadataDeleteRetentionPolicyOutput `json:"-" xml:"-"`
}

type metadataDeleteRetentionPolicyOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLogGroupsInput struct {
	// The maximum number of items returned in the response. If you don't specify
	// a value, the request would return up to 50 items.
	Limit *int64 `locationName:"limit" type:"integer"`

	LogGroupNamePrefix *string `locationName:"logGroupNamePrefix" type:"string"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous DescribeLogGroups
	// request.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataDescribeLogGroupsInput `json:"-" xml:"-"`
}

type metadataDescribeLogGroupsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLogGroupsOutput struct {
	// A list of log groups.
	LogGroups []*LogGroup `locationName:"logGroups" type:"list"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous request. The
	// token expires after 24 hours.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataDescribeLogGroupsOutput `json:"-" xml:"-"`
}

type metadataDescribeLogGroupsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLogStreamsInput struct {
	// If set to true, results are returned in descending order. If you don't specify
	// a value or set it to false, results are returned in ascending order.
	Descending *bool `locationName:"descending" type:"boolean"`

	// The maximum number of items returned in the response. If you don't specify
	// a value, the request would return up to 50 items.
	Limit *int64 `locationName:"limit" type:"integer"`

	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	// Will only return log streams that match the provided logStreamNamePrefix.
	// If you don't specify a value, no prefix filter is applied.
	LogStreamNamePrefix *string `locationName:"logStreamNamePrefix" type:"string"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous DescribeLogStreams
	// request.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Specifies what to order the returned log streams by. Valid arguments are
	// 'LogStreamName' or 'LastEventTime'. If you don't specify a value, results
	// are ordered by LogStreamName. If 'LastEventTime' is chosen, the request cannot
	// also contain a logStreamNamePrefix.
	OrderBy *string `locationName:"orderBy" type:"string"`

	metadataDescribeLogStreamsInput `json:"-" xml:"-"`
}

type metadataDescribeLogStreamsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLogStreamsOutput struct {
	// A list of log streams.
	LogStreams []*LogStream `locationName:"logStreams" type:"list"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous request. The
	// token expires after 24 hours.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataDescribeLogStreamsOutput `json:"-" xml:"-"`
}

type metadataDescribeLogStreamsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeMetricFiltersInput struct {
	// The name of the metric filter.
	FilterNamePrefix *string `locationName:"filterNamePrefix" type:"string"`

	// The maximum number of items returned in the response. If you don't specify
	// a value, the request would return up to 50 items.
	Limit *int64 `locationName:"limit" type:"integer"`

	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous DescribeMetricFilters
	// request.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataDescribeMetricFiltersInput `json:"-" xml:"-"`
}

type metadataDescribeMetricFiltersInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeMetricFiltersOutput struct {
	MetricFilters []*MetricFilter `locationName:"metricFilters" type:"list"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous request. The
	// token expires after 24 hours.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataDescribeMetricFiltersOutput `json:"-" xml:"-"`
}

type metadataDescribeMetricFiltersOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetLogEventsInput struct {
	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	EndTime *int64 `locationName:"endTime" type:"long"`

	// The maximum number of log events returned in the response. If you don't specify
	// a value, the request would return as much log events as can fit in a response
	// size of 1MB, up to 10,000 log events.
	Limit *int64 `locationName:"limit" type:"integer"`

	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	LogStreamName *string `locationName:"logStreamName" type:"string" required:"true"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the nextForwardToken or nextBackwardToken
	// fields in the response of the previous GetLogEvents request.
	NextToken *string `locationName:"nextToken" type:"string"`

	// If set to true, the earliest log events would be returned first. The default
	// is false (the latest log events are returned first).
	StartFromHead *bool `locationName:"startFromHead" type:"boolean"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	StartTime *int64 `locationName:"startTime" type:"long"`

	metadataGetLogEventsInput `json:"-" xml:"-"`
}

type metadataGetLogEventsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetLogEventsOutput struct {
	Events []*OutputLogEvent `locationName:"events" type:"list"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous request. The
	// token expires after 24 hours.
	NextBackwardToken *string `locationName:"nextBackwardToken" type:"string"`

	// A string token used for pagination that points to the next page of results.
	// It must be a value obtained from the response of the previous request. The
	// token expires after 24 hours.
	NextForwardToken *string `locationName:"nextForwardToken" type:"string"`

	metadataGetLogEventsOutput `json:"-" xml:"-"`
}

type metadataGetLogEventsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A log event is a record of some activity that was recorded by the application
// or resource being monitored. The log event record that Amazon CloudWatch
// Logs understands contains two properties: the timestamp of when the event
// occurred, and the raw event message.
type InputLogEvent struct {
	Message *string `locationName:"message" type:"string" required:"true"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	Timestamp *int64 `locationName:"timestamp" type:"long" required:"true"`

	metadataInputLogEvent `json:"-" xml:"-"`
}

type metadataInputLogEvent struct {
	SDKShapeTraits bool `type:"structure"`
}

type LogGroup struct {
	ARN *string `locationName:"arn" type:"string"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	LogGroupName *string `locationName:"logGroupName" type:"string"`

	// The number of metric filters associated with the log group.
	MetricFilterCount *int64 `locationName:"metricFilterCount" type:"integer"`

	// Specifies the number of days you want to retain log events in the specified
	// log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180,
	// 365, 400, 545, 731, 1827, 3653.
	RetentionInDays *int64 `locationName:"retentionInDays" type:"integer"`

	StoredBytes *int64 `locationName:"storedBytes" type:"long"`

	metadataLogGroup `json:"-" xml:"-"`
}

type metadataLogGroup struct {
	SDKShapeTraits bool `type:"structure"`
}

// A log stream is sequence of log events that share the same emitter.
type LogStream struct {
	ARN *string `locationName:"arn" type:"string"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	FirstEventTimestamp *int64 `locationName:"firstEventTimestamp" type:"long"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	LastEventTimestamp *int64 `locationName:"lastEventTimestamp" type:"long"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	LastIngestionTime *int64 `locationName:"lastIngestionTime" type:"long"`

	LogStreamName *string `locationName:"logStreamName" type:"string"`

	StoredBytes *int64 `locationName:"storedBytes" type:"long"`

	// A string token used for making PutLogEvents requests. A sequenceToken can
	// only be used once, and PutLogEvents requests must include the sequenceToken
	// obtained from the response of the previous request.
	UploadSequenceToken *string `locationName:"uploadSequenceToken" type:"string"`

	metadataLogStream `json:"-" xml:"-"`
}

type metadataLogStream struct {
	SDKShapeTraits bool `type:"structure"`
}

// Metric filters can be used to express how Amazon CloudWatch Logs would extract
// metric observations from ingested log events and transform them to metric
// data in a CloudWatch metric.
type MetricFilter struct {
	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	// The name of the metric filter.
	FilterName *string `locationName:"filterName" type:"string"`

	// A symbolic description of how Amazon CloudWatch Logs should interpret the
	// data in each log entry. For example, a log entry may contain timestamps,
	// IP addresses, strings, and so on. You use the pattern to specify what to
	// look for in the log stream.
	FilterPattern *string `locationName:"filterPattern" type:"string"`

	MetricTransformations []*MetricTransformation `locationName:"metricTransformations" type:"list"`

	metadataMetricFilter `json:"-" xml:"-"`
}

type metadataMetricFilter struct {
	SDKShapeTraits bool `type:"structure"`
}

type MetricFilterMatchRecord struct {
	EventMessage *string `locationName:"eventMessage" type:"string"`

	EventNumber *int64 `locationName:"eventNumber" type:"long"`

	ExtractedValues *map[string]*string `locationName:"extractedValues" type:"map"`

	metadataMetricFilterMatchRecord `json:"-" xml:"-"`
}

type metadataMetricFilterMatchRecord struct {
	SDKShapeTraits bool `type:"structure"`
}

type MetricTransformation struct {
	// The name of the CloudWatch metric to which the monitored log information
	// should be published. For example, you may publish to a metric called ErrorCount.
	MetricName *string `locationName:"metricName" type:"string" required:"true"`

	// The destination namespace of the new CloudWatch metric.
	MetricNamespace *string `locationName:"metricNamespace" type:"string" required:"true"`

	// What to publish to the metric. For example, if you're counting the occurrences
	// of a particular term like "Error", the value will be "1" for each occurrence.
	// If you're counting the bytes transferred the published value will be the
	// value in the log event.
	MetricValue *string `locationName:"metricValue" type:"string" required:"true"`

	metadataMetricTransformation `json:"-" xml:"-"`
}

type metadataMetricTransformation struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputLogEvent struct {
	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	IngestionTime *int64 `locationName:"ingestionTime" type:"long"`

	Message *string `locationName:"message" type:"string"`

	// A point in time expressed as the number milliseconds since Jan 1, 1970 00:00:00
	// UTC.
	Timestamp *int64 `locationName:"timestamp" type:"long"`

	metadataOutputLogEvent `json:"-" xml:"-"`
}

type metadataOutputLogEvent struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutLogEventsInput struct {
	// A list of events belonging to a log stream.
	LogEvents []*InputLogEvent `locationName:"logEvents" type:"list" required:"true"`

	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	LogStreamName *string `locationName:"logStreamName" type:"string" required:"true"`

	// A string token that must be obtained from the response of the previous PutLogEvents
	// request.
	SequenceToken *string `locationName:"sequenceToken" type:"string"`

	metadataPutLogEventsInput `json:"-" xml:"-"`
}

type metadataPutLogEventsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutLogEventsOutput struct {
	// A string token used for making PutLogEvents requests. A sequenceToken can
	// only be used once, and PutLogEvents requests must include the sequenceToken
	// obtained from the response of the previous request.
	NextSequenceToken *string `locationName:"nextSequenceToken" type:"string"`

	RejectedLogEventsInfo *RejectedLogEventsInfo `locationName:"rejectedLogEventsInfo" type:"structure"`

	metadataPutLogEventsOutput `json:"-" xml:"-"`
}

type metadataPutLogEventsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutMetricFilterInput struct {
	// The name of the metric filter.
	FilterName *string `locationName:"filterName" type:"string" required:"true"`

	// A symbolic description of how Amazon CloudWatch Logs should interpret the
	// data in each log entry. For example, a log entry may contain timestamps,
	// IP addresses, strings, and so on. You use the pattern to specify what to
	// look for in the log stream.
	FilterPattern *string `locationName:"filterPattern" type:"string" required:"true"`

	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	MetricTransformations []*MetricTransformation `locationName:"metricTransformations" type:"list" required:"true"`

	metadataPutMetricFilterInput `json:"-" xml:"-"`
}

type metadataPutMetricFilterInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutMetricFilterOutput struct {
	metadataPutMetricFilterOutput `json:"-" xml:"-"`
}

type metadataPutMetricFilterOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutRetentionPolicyInput struct {
	LogGroupName *string `locationName:"logGroupName" type:"string" required:"true"`

	// Specifies the number of days you want to retain log events in the specified
	// log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180,
	// 365, 400, 545, 731, 1827, 3653.
	RetentionInDays *int64 `locationName:"retentionInDays" type:"integer" required:"true"`

	metadataPutRetentionPolicyInput `json:"-" xml:"-"`
}

type metadataPutRetentionPolicyInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutRetentionPolicyOutput struct {
	metadataPutRetentionPolicyOutput `json:"-" xml:"-"`
}

type metadataPutRetentionPolicyOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RejectedLogEventsInfo struct {
	ExpiredLogEventEndIndex *int64 `locationName:"expiredLogEventEndIndex" type:"integer"`

	TooNewLogEventStartIndex *int64 `locationName:"tooNewLogEventStartIndex" type:"integer"`

	TooOldLogEventEndIndex *int64 `locationName:"tooOldLogEventEndIndex" type:"integer"`

	metadataRejectedLogEventsInfo `json:"-" xml:"-"`
}

type metadataRejectedLogEventsInfo struct {
	SDKShapeTraits bool `type:"structure"`
}

type TestMetricFilterInput struct {
	// A symbolic description of how Amazon CloudWatch Logs should interpret the
	// data in each log entry. For example, a log entry may contain timestamps,
	// IP addresses, strings, and so on. You use the pattern to specify what to
	// look for in the log stream.
	FilterPattern *string `locationName:"filterPattern" type:"string" required:"true"`

	LogEventMessages []*string `locationName:"logEventMessages" type:"list" required:"true"`

	metadataTestMetricFilterInput `json:"-" xml:"-"`
}

type metadataTestMetricFilterInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type TestMetricFilterOutput struct {
	Matches []*MetricFilterMatchRecord `locationName:"matches" type:"list"`

	metadataTestMetricFilterOutput `json:"-" xml:"-"`
}

type metadataTestMetricFilterOutput struct {
	SDKShapeTraits bool `type:"structure"`
}