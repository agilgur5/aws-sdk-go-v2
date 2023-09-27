// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a transform job. A transform job uses a trained model to get inferences
// on a dataset and saves these results to an Amazon S3 location that you specify.
// To perform batch transformations, you create a transform job and use the data
// that you have readily available. In the request body, you provide the following:
//
//   - TransformJobName - Identifies the transform job. The name must be unique
//     within an Amazon Web Services Region in an Amazon Web Services account.
//   - ModelName - Identifies the model to use. ModelName must be the name of an
//     existing Amazon SageMaker model in the same Amazon Web Services Region and
//     Amazon Web Services account. For information on creating a model, see
//     CreateModel (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html)
//     .
//   - TransformInput - Describes the dataset to be transformed and the Amazon S3
//     location where it is stored.
//   - TransformOutput - Identifies the Amazon S3 location where you want Amazon
//     SageMaker to save the results from the transform job.
//   - TransformResources - Identifies the ML compute instances for the transform
//     job.
//
// For more information about how batch transformation works, see Batch Transform (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform.html)
// .
func (c *Client) CreateTransformJob(ctx context.Context, params *CreateTransformJobInput, optFns ...func(*Options)) (*CreateTransformJobOutput, error) {
	if params == nil {
		params = &CreateTransformJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransformJob", params, optFns, c.addOperationCreateTransformJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransformJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransformJobInput struct {

	// The name of the model that you want to use for the transform job. ModelName
	// must be the name of an existing Amazon SageMaker model within an Amazon Web
	// Services Region in an Amazon Web Services account.
	//
	// This member is required.
	ModelName *string

	// Describes the input source and the way the transform job consumes it.
	//
	// This member is required.
	TransformInput *types.TransformInput

	// The name of the transform job. The name must be unique within an Amazon Web
	// Services Region in an Amazon Web Services account.
	//
	// This member is required.
	TransformJobName *string

	// Describes the results of the transform job.
	//
	// This member is required.
	TransformOutput *types.TransformOutput

	// Describes the resources, including ML instance types and ML instance count, to
	// use for the transform job.
	//
	// This member is required.
	TransformResources *types.TransformResources

	// Specifies the number of records to include in a mini-batch for an HTTP
	// inference request. A record is a single unit of input data that inference can be
	// made on. For example, a single line in a CSV file is a record. To enable the
	// batch strategy, you must set the SplitType property to Line , RecordIO , or
	// TFRecord . To use only one record when making an HTTP invocation request to a
	// container, set BatchStrategy to SingleRecord and SplitType to Line . To fit as
	// many records in a mini-batch as can fit within the MaxPayloadInMB limit, set
	// BatchStrategy to MultiRecord and SplitType to Line .
	BatchStrategy types.BatchStrategy

	// Configuration to control how SageMaker captures inference data.
	DataCaptureConfig *types.BatchDataCaptureConfig

	// The data structure used to specify the data to be used for inference in a batch
	// transform job and to associate the data that is relevant to the prediction
	// results in the output. The input filter provided allows you to exclude input
	// data that is not needed for inference in a batch transform job. The output
	// filter provided allows you to include input data relevant to interpreting the
	// predictions in the output from the job. For more information, see Associate
	// Prediction Results with their Corresponding Input Records (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html)
	// .
	DataProcessing *types.DataProcessing

	// The environment variables to set in the Docker container. We support up to 16
	// key and values entries in the map.
	Environment map[string]string

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//   - CreateProcessingJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html)
	//   - CreateTrainingJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html)
	//   - CreateTransformJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html)
	ExperimentConfig *types.ExperimentConfig

	// The maximum number of parallel requests that can be sent to each instance in a
	// transform job. If MaxConcurrentTransforms is set to 0 or left unset, Amazon
	// SageMaker checks the optional execution-parameters to determine the settings for
	// your chosen algorithm. If the execution-parameters endpoint is not enabled, the
	// default value is 1 . For more information on execution-parameters, see How
	// Containers Serve Requests (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-batch-code.html#your-algorithms-batch-code-how-containe-serves-requests)
	// . For built-in algorithms, you don't need to set a value for
	// MaxConcurrentTransforms .
	MaxConcurrentTransforms *int32

	// The maximum allowed size of the payload, in MB. A payload is the data portion
	// of a record (without metadata). The value in MaxPayloadInMB must be greater
	// than, or equal to, the size of a single record. To estimate the size of a record
	// in MB, divide the size of your dataset by the number of records. To ensure that
	// the records fit within the maximum payload size, we recommend using a slightly
	// larger value. The default value is 6 MB. The value of MaxPayloadInMB cannot be
	// greater than 100 MB. If you specify the MaxConcurrentTransforms parameter, the
	// value of (MaxConcurrentTransforms * MaxPayloadInMB) also cannot exceed 100 MB.
	// For cases where the payload might be arbitrarily large and is transmitted using
	// HTTP chunked encoding, set the value to 0 . This feature works only in supported
	// algorithms. Currently, Amazon SageMaker built-in algorithms do not support HTTP
	// chunked encoding.
	MaxPayloadInMB *int32

	// Configures the timeout and maximum number of retries for processing a transform
	// job invocation.
	ModelClientConfig *types.ModelClientConfig

	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the Amazon Web Services Billing and Cost Management User Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

func (*CreateTransformJobInput) operationName() string {
	return "CreateTransformJob"
}

type CreateTransformJobOutput struct {

	// The Amazon Resource Name (ARN) of the transform job.
	//
	// This member is required.
	TransformJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTransformJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTransformJob{}, middleware.After)
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
	if err = addCreateTransformJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTransformJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransformJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransformJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateTransformJob",
	}
}

type opCreateTransformJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateTransformJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateTransformJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "sagemaker"
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
				signingName = "sagemaker"
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
				v4aScheme.SigningName = aws.String("sagemaker")
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

func addCreateTransformJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateTransformJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
