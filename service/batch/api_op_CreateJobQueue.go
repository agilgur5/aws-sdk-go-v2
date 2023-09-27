// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Batch job queue. When you create a job queue, you associate one or
// more compute environments to the queue and assign an order of preference for the
// compute environments. You also set a priority to the job queue that determines
// the order that the Batch scheduler places jobs onto its associated compute
// environments. For example, if a compute environment is associated with more than
// one job queue, the job queue with a higher priority is given preference for
// scheduling jobs to that compute environment.
func (c *Client) CreateJobQueue(ctx context.Context, params *CreateJobQueueInput, optFns ...func(*Options)) (*CreateJobQueueOutput, error) {
	if params == nil {
		params = &CreateJobQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJobQueue", params, optFns, c.addOperationCreateJobQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateJobQueue .
type CreateJobQueueInput struct {

	// The set of compute environments mapped to a job queue and their order relative
	// to each other. The job scheduler uses this parameter to determine which compute
	// environment runs a specific job. Compute environments must be in the VALID
	// state before you can associate them with a job queue. You can associate up to
	// three compute environments with a job queue. All of the compute environments
	// must be either EC2 ( EC2 or SPOT ) or Fargate ( FARGATE or FARGATE_SPOT ); EC2
	// and Fargate compute environments can't be mixed. All compute environments that
	// are associated with a job queue must share the same architecture. Batch doesn't
	// support mixing compute environment architecture types in a single job queue.
	//
	// This member is required.
	ComputeEnvironmentOrder []types.ComputeEnvironmentOrder

	// The name of the job queue. It can be up to 128 letters long. It can contain
	// uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	//
	// This member is required.
	JobQueueName *string

	// The priority of the job queue. Job queues with a higher priority (or a higher
	// integer value for the priority parameter) are evaluated first when associated
	// with the same compute environment. Priority is determined in descending order.
	// For example, a job queue with a priority value of 10 is given scheduling
	// preference over a job queue with a priority value of 1 . All of the compute
	// environments must be either EC2 ( EC2 or SPOT ) or Fargate ( FARGATE or
	// FARGATE_SPOT ); EC2 and Fargate compute environments can't be mixed.
	//
	// This member is required.
	Priority *int32

	// The Amazon Resource Name (ARN) of the fair share scheduling policy. If this
	// parameter is specified, the job queue uses a fair share scheduling policy. If
	// this parameter isn't specified, the job queue uses a first in, first out (FIFO)
	// scheduling policy. After a job queue is created, you can replace but can't
	// remove the fair share scheduling policy. The format is
	// aws:Partition:batch:Region:Account:scheduling-policy/Name . An example is
	// aws:aws:batch:us-west-2:123456789012:scheduling-policy/MySchedulingPolicy .
	SchedulingPolicyArn *string

	// The state of the job queue. If the job queue state is ENABLED , it is able to
	// accept jobs. If the job queue state is DISABLED , new jobs can't be added to the
	// queue, but jobs already in the queue can finish.
	State types.JQState

	// The tags that you apply to the job queue to help you categorize and organize
	// your resources. Each tag consists of a key and an optional value. For more
	// information, see Tagging your Batch resources (https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html)
	// in Batch User Guide.
	Tags map[string]string

	noSmithyDocumentSerde
}

func (*CreateJobQueueInput) operationName() string {
	return "CreateJobQueue"
}

type CreateJobQueueOutput struct {

	// The Amazon Resource Name (ARN) of the job queue.
	//
	// This member is required.
	JobQueueArn *string

	// The name of the job queue.
	//
	// This member is required.
	JobQueueName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateJobQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJobQueue{}, middleware.After)
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
	if err = addCreateJobQueueResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateJobQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJobQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateJobQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CreateJobQueue",
	}
}

type opCreateJobQueueResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateJobQueueResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateJobQueueResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "batch"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "batch"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("batch")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateJobQueueResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateJobQueueResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
