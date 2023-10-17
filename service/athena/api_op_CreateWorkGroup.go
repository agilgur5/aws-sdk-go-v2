// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a workgroup with the specified name. A workgroup can be an Apache Spark
// enabled workgroup or an Athena SQL workgroup.
func (c *Client) CreateWorkGroup(ctx context.Context, params *CreateWorkGroupInput, optFns ...func(*Options)) (*CreateWorkGroupOutput, error) {
	if params == nil {
		params = &CreateWorkGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkGroup", params, optFns, c.addOperationCreateWorkGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkGroupInput struct {

	// The workgroup name.
	//
	// This member is required.
	Name *string

	// Contains configuration information for creating an Athena SQL workgroup or
	// Spark enabled Athena workgroup. Athena SQL workgroup configuration includes the
	// location in Amazon S3 where query and calculation results are stored, the
	// encryption configuration, if any, used for encrypting query results, whether the
	// Amazon CloudWatch Metrics are enabled for the workgroup, the limit for the
	// amount of bytes scanned (cutoff) per query, if it is specified, and whether
	// workgroup's settings (specified with EnforceWorkGroupConfiguration ) in the
	// WorkGroupConfiguration override client-side settings. See
	// WorkGroupConfiguration$EnforceWorkGroupConfiguration .
	Configuration *types.WorkGroupConfiguration

	// The workgroup description.
	Description *string

	// A list of comma separated tags to add to the workgroup that is created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateWorkGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkGroup{}, middleware.After)
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
	if err = addOpCreateWorkGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "CreateWorkGroup",
	}
}
