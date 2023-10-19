// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get information about a resource that's been registered for zonal shifts with
// Amazon Route 53 Application Recovery Controller in this AWS Region. Resources
// that are registered for zonal shifts are managed resources in Route 53 ARC. At
// this time, you can only start a zonal shift for Network Load Balancers and
// Application Load Balancers with cross-zone load balancing turned off.
func (c *Client) GetManagedResource(ctx context.Context, params *GetManagedResourceInput, optFns ...func(*Options)) (*GetManagedResourceOutput, error) {
	if params == nil {
		params = &GetManagedResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetManagedResource", params, optFns, c.addOperationGetManagedResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetManagedResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetManagedResourceInput struct {

	// The identifier for the resource to include in a zonal shift. The identifier is
	// the Amazon Resource Name (ARN) for the resource. At this time, you can only
	// start a zonal shift for Network Load Balancers and Application Load Balancers
	// with cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type GetManagedResourceOutput struct {

	// A collection of key-value pairs that indicate whether resources are active in
	// Availability Zones or not. The key name is the Availability Zone where the
	// resource is deployed. The value is 1 or 0.
	//
	// This member is required.
	AppliedWeights map[string]float32

	// The zonal shifts that are currently active for a resource.
	//
	// This member is required.
	ZonalShifts []types.ZonalShiftInResource

	// The Amazon Resource Name (ARN) for the resource.
	Arn *string

	// The name of the resource.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetManagedResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetManagedResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetManagedResource{}, middleware.After)
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
	if err = addOpGetManagedResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetManagedResource(options.Region), middleware.Before); err != nil {
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
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "GetManagedResource",
		options:   options,
	}, "ResolveEndpointV2", middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&signRequestMiddleware{}, middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&getIdentityMiddleware{
		options: options,
	}, middleware.Before)
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetManagedResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "arc-zonal-shift",
		OperationName: "GetManagedResource",
	}
}
