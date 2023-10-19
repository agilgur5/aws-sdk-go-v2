// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a specific runtime environment.
func (c *Client) GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) {
	if params == nil {
		params = &GetEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnvironment", params, optFns, c.addOperationGetEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentInput struct {

	// The unique identifier of the runtime environment.
	//
	// This member is required.
	EnvironmentId *string

	noSmithyDocumentSerde
}

type GetEnvironmentOutput struct {

	// The timestamp when the runtime environment was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The target platform for the runtime environment.
	//
	// This member is required.
	EngineType types.EngineType

	// The version of the runtime engine.
	//
	// This member is required.
	EngineVersion *string

	// The Amazon Resource Name (ARN) of the runtime environment.
	//
	// This member is required.
	EnvironmentArn *string

	// The unique identifier of the runtime environment.
	//
	// This member is required.
	EnvironmentId *string

	// The type of instance underlying the runtime environment.
	//
	// This member is required.
	InstanceType *string

	// The name of the runtime environment. Must be unique within the account.
	//
	// This member is required.
	Name *string

	// The unique identifiers of the security groups assigned to this runtime
	// environment.
	//
	// This member is required.
	SecurityGroupIds []string

	// The status of the runtime environment.
	//
	// This member is required.
	Status types.EnvironmentLifecycle

	// The unique identifiers of the subnets assigned to this runtime environment.
	//
	// This member is required.
	SubnetIds []string

	// The unique identifier for the VPC used with this runtime environment.
	//
	// This member is required.
	VpcId *string

	// The number of instances included in the runtime environment. A standalone
	// runtime environment has a maxiumum of one instance. Currently, a high
	// availability runtime environment has a maximum of two instances.
	ActualCapacity *int32

	// The description of the runtime environment.
	Description *string

	// The desired capacity of the high availability configuration for the runtime
	// environment.
	HighAvailabilityConfig *types.HighAvailabilityConfig

	// The identifier of a customer managed key.
	KmsKeyId *string

	// The Amazon Resource Name (ARN) for the load balancer used with the runtime
	// environment.
	LoadBalancerArn *string

	// Indicates the pending maintenance scheduled on this environment.
	PendingMaintenance *types.PendingMaintenance

	// Configures the maintenance window you want for the runtime environment. If you
	// do not provide a value, a random system-generated value will be assigned.
	PreferredMaintenanceWindow *string

	// Whether applications running in this runtime environment are publicly
	// accessible.
	PubliclyAccessible bool

	// The reason for the reported status.
	StatusReason *string

	// The storage configurations defined for the runtime environment.
	StorageConfigurations []types.StorageConfiguration

	// The tags defined for this runtime environment.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEnvironment{}, middleware.After)
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
	if err = addOpGetEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnvironment(options.Region), middleware.Before); err != nil {
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
		operation: "GetEnvironment",
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

func newServiceMetadataMiddleware_opGetEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "GetEnvironment",
	}
}
