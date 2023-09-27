// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a dataset import job created using the CreateDatasetImportJob (https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html)
// operation. In addition to listing the parameters provided in the
// CreateDatasetImportJob request, this operation includes the following
// properties:
//   - CreationTime
//   - LastModificationTime
//   - DataSize
//   - FieldStatistics
//   - Status
//   - Message - If an error occurred, information about the error.
func (c *Client) DescribeDatasetImportJob(ctx context.Context, params *DescribeDatasetImportJobInput, optFns ...func(*Options)) (*DescribeDatasetImportJobOutput, error) {
	if params == nil {
		params = &DescribeDatasetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDatasetImportJob", params, optFns, c.addOperationDescribeDatasetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetImportJobInput struct {

	// The Amazon Resource Name (ARN) of the dataset import job.
	//
	// This member is required.
	DatasetImportJobArn *string

	noSmithyDocumentSerde
}

func (*DescribeDatasetImportJobInput) operationName() string {
	return "DescribeDatasetImportJob"
}

type DescribeDatasetImportJobOutput struct {

	// When the dataset import job was created.
	CreationTime *time.Time

	// The size of the dataset in gigabytes (GB) after the import job has finished.
	DataSize *float64

	// The location of the training data to import and an Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the data. If
	// encryption is used, DataSource includes an Key Management Service (KMS) key.
	DataSource *types.DataSource

	// The Amazon Resource Name (ARN) of the dataset that the training data was
	// imported to.
	DatasetArn *string

	// The ARN of the dataset import job.
	DatasetImportJobArn *string

	// The name of the dataset import job.
	DatasetImportJobName *string

	// The estimated time remaining in minutes for the dataset import job to complete.
	EstimatedTimeRemainingInMinutes *int64

	// Statistical information about each field in the input data.
	FieldStatistics map[string]types.Statistics

	// The format of the imported data, CSV or PARQUET.
	Format *string

	// The format of the geolocation attribute. Valid Values: "LAT_LONG" and
	// "CC_POSTALCODE" .
	GeolocationFormat *string

	// The import mode of the dataset import job, FULL or INCREMENTAL.
	ImportMode types.ImportMode

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//   - CREATE_PENDING - The CreationTime .
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//   - CREATE_STOPPING - The current timestamp.
	//   - CREATE_STOPPED - When the job stopped.
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the dataset import job. States include:
	//   - ACTIVE
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	//   - CREATE_STOPPING , CREATE_STOPPED
	Status *string

	// The single time zone applied to every item in the dataset
	TimeZone *string

	// The format of timestamps in the dataset. The format that you specify depends on
	// the DataFrequency specified when the dataset was created. The following formats
	// are supported
	//   - "yyyy-MM-dd" For the following data frequencies: Y, M, W, and D
	//   - "yyyy-MM-dd HH:mm:ss" For the following data frequencies: H, 30min, 15min,
	//   and 1min; and optionally, for: Y, M, W, and D
	TimestampFormat *string

	// Whether TimeZone is automatically derived from the geolocation attribute.
	UseGeolocationForTimeZone bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDatasetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDatasetImportJob{}, middleware.After)
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
	if err = addDescribeDatasetImportJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeDatasetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatasetImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDatasetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeDatasetImportJob",
	}
}

type opDescribeDatasetImportJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeDatasetImportJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeDatasetImportJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "forecast"
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
				signingName = "forecast"
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
				v4aScheme.SigningName = aws.String("forecast")
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

func addDescribeDatasetImportJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeDatasetImportJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
