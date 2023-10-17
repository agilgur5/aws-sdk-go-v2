// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a robot. This API will no longer be supported as of May 2, 2022. Use
// it to remove resources that were created for Deployment Service.
//
// Deprecated: Support for the AWS RoboMaker application deployment feature has
// ended. For additional information, see
// https://docs.aws.amazon.com/robomaker/latest/dg/fleets.html.
func (c *Client) DescribeRobot(ctx context.Context, params *DescribeRobotInput, optFns ...func(*Options)) (*DescribeRobotOutput, error) {
	if params == nil {
		params = &DescribeRobotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRobot", params, optFns, c.addOperationDescribeRobotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRobotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRobotInput struct {

	// The Amazon Resource Name (ARN) of the robot to be described.
	//
	// This member is required.
	Robot *string

	noSmithyDocumentSerde
}

type DescribeRobotOutput struct {

	// The target architecture of the robot application.
	Architecture types.Architecture

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the fleet.
	FleetArn *string

	// The Greengrass group id.
	GreengrassGroupId *string

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string

	// The time of the last deployment job.
	LastDeploymentTime *time.Time

	// The name of the robot.
	Name *string

	// The status of the fleet.
	Status types.RobotStatus

	// The list of all tags added to the specified robot.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRobotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRobot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRobot{}, middleware.After)
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
	if err = addOpDescribeRobotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRobot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRobot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeRobot",
	}
}
