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

// Creates a version of a robot application.
func (c *Client) CreateRobotApplicationVersion(ctx context.Context, params *CreateRobotApplicationVersionInput, optFns ...func(*Options)) (*CreateRobotApplicationVersionOutput, error) {
	if params == nil {
		params = &CreateRobotApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRobotApplicationVersion", params, optFns, c.addOperationCreateRobotApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRobotApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRobotApplicationVersionInput struct {

	// The application information for the robot application.
	//
	// This member is required.
	Application *string

	// The current revision id for the robot application. If you provide a value and
	// it matches the latest revision ID, a new version will be created.
	CurrentRevisionId *string

	// A SHA256 identifier for the Docker image that you use for your robot
	// application.
	ImageDigest *string

	// The Amazon S3 identifier for the zip file bundle that you use for your robot
	// application.
	S3Etags []string

	noSmithyDocumentSerde
}

type CreateRobotApplicationVersionOutput struct {

	// The Amazon Resource Name (ARN) of the robot application.
	Arn *string

	// The object that contains the Docker image URI used to create your robot
	// application.
	Environment *types.Environment

	// The time, in milliseconds since the epoch, when the robot application was last
	// updated.
	LastUpdatedAt *time.Time

	// The name of the robot application.
	Name *string

	// The revision id of the robot application.
	RevisionId *string

	// The robot software suite (ROS distribution) used by the robot application.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The sources of the robot application.
	Sources []types.Source

	// The version of the robot application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRobotApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRobotApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRobotApplicationVersion{}, middleware.After)
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
	if err = addOpCreateRobotApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRobotApplicationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRobotApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateRobotApplicationVersion",
	}
}
