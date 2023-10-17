// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about a given change set.
func (c *Client) DescribeChangeSet(ctx context.Context, params *DescribeChangeSetInput, optFns ...func(*Options)) (*DescribeChangeSetOutput, error) {
	if params == nil {
		params = &DescribeChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChangeSet", params, optFns, c.addOperationDescribeChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChangeSetInput struct {

	// Required. The catalog related to the request. Fixed value: AWSMarketplace
	//
	// This member is required.
	Catalog *string

	// Required. The unique identifier for the StartChangeSet request that you want to
	// describe the details for.
	//
	// This member is required.
	ChangeSetId *string

	noSmithyDocumentSerde
}

type DescribeChangeSetOutput struct {

	// An array of ChangeSummary objects.
	ChangeSet []types.ChangeSummary

	// The ARN associated with the unique identifier for the change set referenced in
	// this request.
	ChangeSetArn *string

	// Required. The unique identifier for the change set referenced in this request.
	ChangeSetId *string

	// The optional name provided in the StartChangeSet request. If you do not provide
	// a name, one is set by default.
	ChangeSetName *string

	// The date and time, in ISO 8601 format (2018-02-27T13:45:22Z), the request
	// transitioned to a terminal state. The change cannot transition to a different
	// state. Null if the request is not in a terminal state.
	EndTime *string

	// Returned if the change set is in FAILED status. Can be either CLIENT_ERROR ,
	// which means that there are issues with the request (see the ErrorDetailList ),
	// or SERVER_FAULT , which means that there is a problem in the system, and you
	// should retry your request.
	FailureCode types.FailureCode

	// Returned if there is a failure on the change set, but that failure is not
	// related to any of the changes in the request.
	FailureDescription *string

	// The date and time, in ISO 8601 format (2018-02-27T13:45:22Z), the request
	// started.
	StartTime *string

	// The status of the change request.
	Status types.ChangeStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChangeSet(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "DescribeChangeSet",
	}
}
