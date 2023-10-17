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

// Creates a simulation application.
func (c *Client) CreateSimulationApplication(ctx context.Context, params *CreateSimulationApplicationInput, optFns ...func(*Options)) (*CreateSimulationApplicationOutput, error) {
	if params == nil {
		params = &CreateSimulationApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSimulationApplication", params, optFns, c.addOperationCreateSimulationApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSimulationApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSimulationApplicationInput struct {

	// The name of the simulation application.
	//
	// This member is required.
	Name *string

	// The robot software suite (ROS distribution) used by the simulation application.
	//
	// This member is required.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The simulation software suite used by the simulation application.
	//
	// This member is required.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite

	// The object that contains the Docker image URI used to create your simulation
	// application.
	Environment *types.Environment

	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine

	// The sources of the simulation application.
	Sources []types.SourceConfig

	// A map that contains tag keys and tag values that are attached to the simulation
	// application.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSimulationApplicationOutput struct {

	// The Amazon Resource Name (ARN) of the simulation application.
	Arn *string

	// The object that contains the Docker image URI that you used to create your
	// simulation application.
	Environment *types.Environment

	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time

	// The name of the simulation application.
	Name *string

	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine

	// The revision id of the simulation application.
	RevisionId *string

	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite

	// The sources of the simulation application.
	Sources []types.Source

	// The list of all tags added to the simulation application.
	Tags map[string]string

	// The version of the simulation application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSimulationApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSimulationApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSimulationApplication{}, middleware.After)
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
	if err = addOpCreateSimulationApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSimulationApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSimulationApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateSimulationApplication",
	}
}
