// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a filter resource using specified filter criteria. When the filter
// action is set to SUPPRESS this action creates a suppression rule.
func (c *Client) CreateFilter(ctx context.Context, params *CreateFilterInput, optFns ...func(*Options)) (*CreateFilterOutput, error) {
	if params == nil {
		params = &CreateFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFilter", params, optFns, c.addOperationCreateFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFilterInput struct {

	// Defines the action that is to be applied to the findings that match the filter.
	//
	// This member is required.
	Action types.FilterAction

	// Defines the criteria to be used in the filter for querying findings.
	//
	// This member is required.
	FilterCriteria *types.FilterCriteria

	// The name of the filter. Minimum length of 3. Maximum length of 64. Valid
	// characters include alphanumeric characters, dot (.), underscore (_), and dash
	// (-). Spaces are not allowed.
	//
	// This member is required.
	Name *string

	// A description of the filter.
	Description *string

	// The reason for creating the filter.
	Reason *string

	// A list of tags for the filter.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateFilterOutput struct {

	// The Amazon Resource Number (ARN) of the successfully created filter.
	//
	// This member is required.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFilter{}, middleware.After)
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
	if err = addOpCreateFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "CreateFilter",
	}
}
