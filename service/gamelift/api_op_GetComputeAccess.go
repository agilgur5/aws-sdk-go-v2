// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests authorization to remotely connect to a compute resource in an Amazon
// GameLift fleet. Call this action to connect to an instance in a managed EC2
// fleet if the fleet's game build uses Amazon GameLift server SDK 5.x or later. To
// connect to instances with game builds that use server SDK 4.x or earlier, call
// GetInstanceAccess . To request access to a compute, identify the specific EC2
// instance and the fleet it belongs to. You can retrieve instances for a managed
// EC2 fleet by calling ListCompute . If successful, this operation returns a set
// of temporary Amazon Web Services credentials, including a two-part access key
// and a session token. Use these credentials with Amazon EC2 Systems Manager (SSM)
// to start a session with the compute. For more details, see Starting a session
// (CLI) (https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html#sessions-start-cli)
// in the Amazon EC2 Systems Manager User Guide. Learn more Remotely connect to
// fleet instances (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html)
// Debug fleet issues (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html)
func (c *Client) GetComputeAccess(ctx context.Context, params *GetComputeAccessInput, optFns ...func(*Options)) (*GetComputeAccessOutput, error) {
	if params == nil {
		params = &GetComputeAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComputeAccess", params, optFns, c.addOperationGetComputeAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComputeAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComputeAccessInput struct {

	// A unique identifier for the compute resource that you want to connect to. You
	// can use either a registered compute name or an instance ID.
	//
	// This member is required.
	ComputeName *string

	// A unique identifier for the fleet that contains the compute resource you want
	// to connect to. You can use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

type GetComputeAccessOutput struct {

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to an Amazon GameLift compute resource and uniquely
	// identifies it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::compute/compute-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	ComputeArn *string

	// The identifier of the compute resource to be accessed. This value might be
	// either a compute name or an instance ID.
	ComputeName *string

	// A set of temporary Amazon Web Services credentials for use when connecting to
	// the compute resource with Amazon EC2 Systems Manager (SSM).
	Credentials *types.AwsCredentials

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to a Amazon GameLift fleet resource and uniquely identifies
	// it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	FleetArn *string

	// The ID of the fleet that contains the compute resource to be accessed.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComputeAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComputeAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComputeAccess{}, middleware.After)
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
	if err = addOpGetComputeAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComputeAccess(options.Region), middleware.Before); err != nil {
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
		operation: "GetComputeAccess",
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

func newServiceMetadataMiddleware_opGetComputeAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "GetComputeAccess",
	}
}
