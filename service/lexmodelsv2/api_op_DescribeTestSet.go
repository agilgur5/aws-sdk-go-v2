// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets metadata information about the test set.
func (c *Client) DescribeTestSet(ctx context.Context, params *DescribeTestSetInput, optFns ...func(*Options)) (*DescribeTestSetOutput, error) {
	if params == nil {
		params = &DescribeTestSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTestSet", params, optFns, c.addOperationDescribeTestSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTestSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTestSetInput struct {

	// The test set Id for the test set request.
	//
	// This member is required.
	TestSetId *string

	noSmithyDocumentSerde
}

type DescribeTestSetOutput struct {

	// The creation date and time for the test set data.
	CreationDateTime *time.Time

	// The description of the test set.
	Description *string

	// The date and time for the last update of the test set data.
	LastUpdatedDateTime *time.Time

	// Indicates whether the test set is audio or text data.
	Modality types.TestSetModality

	// The total number of agent and user turn in the test set.
	NumTurns *int32

	// The roleARN used for any operation in the test set to access resources in the
	// Amazon Web Services account.
	RoleArn *string

	// The status of the test set.
	Status types.TestSetStatus

	// The Amazon S3 storage location for the test set data.
	StorageLocation *types.TestSetStorageLocation

	// The test set Id for the test set response.
	TestSetId *string

	// The test set name of the test set.
	TestSetName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTestSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTestSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTestSet{}, middleware.After)
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
	if err = addOpDescribeTestSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTestSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTestSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeTestSet",
	}
}
