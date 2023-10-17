// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the progress of a pre-update dry run analysis on an Amazon OpenSearch
// Service domain. For more information, see Determining whether a change will
// cause a blue/green deployment (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-configuration-changes#dryrun)
// .
func (c *Client) DescribeDryRunProgress(ctx context.Context, params *DescribeDryRunProgressInput, optFns ...func(*Options)) (*DescribeDryRunProgressOutput, error) {
	if params == nil {
		params = &DescribeDryRunProgressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDryRunProgress", params, optFns, c.addOperationDescribeDryRunProgressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDryRunProgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDryRunProgressInput struct {

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// The unique identifier of the dry run.
	DryRunId *string

	// Whether to include the configuration of the dry run in the response. The
	// configuration specifies the updates that you're planning to make on the domain.
	LoadDryRunConfig *bool

	noSmithyDocumentSerde
}

type DescribeDryRunProgressOutput struct {

	// Details about the changes you're planning to make on the domain.
	DryRunConfig *types.DomainStatus

	// The current status of the dry run, including any validation errors.
	DryRunProgressStatus *types.DryRunProgressStatus

	// The results of the dry run.
	DryRunResults *types.DryRunResults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDryRunProgressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDryRunProgress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDryRunProgress{}, middleware.After)
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
	if err = addOpDescribeDryRunProgressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDryRunProgress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDryRunProgress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeDryRunProgress",
	}
}
