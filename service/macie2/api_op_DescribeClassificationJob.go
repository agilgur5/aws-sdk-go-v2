// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the status and settings for a classification job.
func (c *Client) DescribeClassificationJob(ctx context.Context, params *DescribeClassificationJobInput, optFns ...func(*Options)) (*DescribeClassificationJobOutput, error) {
	if params == nil {
		params = &DescribeClassificationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClassificationJob", params, optFns, c.addOperationDescribeClassificationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClassificationJobInput struct {

	// The unique identifier for the classification job.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

func (*DescribeClassificationJobInput) operationName() string {
	return "DescribeClassificationJob"
}

type DescribeClassificationJobOutput struct {

	// An array of unique identifiers, one for each allow list that the job uses when
	// it analyzes data.
	AllowListIds []string

	// The token that was provided to ensure the idempotency of the request to create
	// the job.
	ClientToken *string

	// The date and time, in UTC and extended ISO 8601 format, when the job was
	// created.
	CreatedAt *time.Time

	// An array of unique identifiers, one for each custom data identifier that the
	// job uses when it analyzes data. This value is null if the job uses only managed
	// data identifiers to analyze data.
	CustomDataIdentifierIds []string

	// The custom description of the job.
	Description *string

	// For a recurring job, specifies whether you configured the job to analyze all
	// existing, eligible objects immediately after the job was created (true). If you
	// configured the job to analyze only those objects that were created or changed
	// after the job was created and before the job's first scheduled run, this value
	// is false. This value is also false for a one-time job.
	InitialRun bool

	// The Amazon Resource Name (ARN) of the job.
	JobArn *string

	// The unique identifier for the job.
	JobId *string

	// The current status of the job. Possible values are:
	//   - CANCELLED - You cancelled the job or, if it's a one-time job, you paused
	//   the job and didn't resume it within 30 days.
	//   - COMPLETE - For a one-time job, Amazon Macie finished processing the data
	//   specified for the job. This value doesn't apply to recurring jobs.
	//   - IDLE - For a recurring job, the previous scheduled run is complete and the
	//   next scheduled run is pending. This value doesn't apply to one-time jobs.
	//   - PAUSED - Macie started running the job but additional processing would
	//   exceed the monthly sensitive data discovery quota for your account or one or
	//   more member accounts that the job analyzes data for.
	//   - RUNNING - For a one-time job, the job is in progress. For a recurring job,
	//   a scheduled run is in progress.
	//   - USER_PAUSED - You paused the job. If you paused the job while it had a
	//   status of RUNNING and you don't resume it within 30 days of pausing it, the job
	//   or job run will expire and be cancelled, depending on the job's type. To check
	//   the expiration date, refer to the UserPausedDetails.jobExpiresAt property.
	JobStatus types.JobStatus

	// The schedule for running the job. Possible values are:
	//   - ONE_TIME - The job runs only once.
	//   - SCHEDULED - The job runs on a daily, weekly, or monthly basis. The
	//   scheduleFrequency property indicates the recurrence pattern for the job.
	JobType types.JobType

	// Specifies whether any account- or bucket-level access errors occurred when the
	// job ran. For a recurring job, this value indicates the error status of the job's
	// most recent run.
	LastRunErrorStatus *types.LastRunErrorStatus

	// The date and time, in UTC and extended ISO 8601 format, when the job started.
	// If the job is a recurring job, this value indicates when the most recent run
	// started or, if the job hasn't run yet, when the job was created.
	LastRunTime *time.Time

	// An array of unique identifiers, one for each managed data identifier that the
	// job is explicitly configured to include (use) or exclude (not use) when it
	// analyzes data. Inclusion or exclusion depends on the managed data identifier
	// selection type specified for the job (managedDataIdentifierSelector).This value
	// is null if the job's managed data identifier selection type is ALL, NONE, or
	// RECOMMENDED.
	ManagedDataIdentifierIds []string

	// The selection type that determines which managed data identifiers the job uses
	// when it analyzes data. Possible values are:
	//   - ALL - Use all managed data identifiers.
	//   - EXCLUDE - Use all managed data identifiers except the ones specified by the
	//   managedDataIdentifierIds property.
	//   - INCLUDE - Use only the managed data identifiers specified by the
	//   managedDataIdentifierIds property.
	//   - NONE - Don't use any managed data identifiers. Use only custom data
	//   identifiers (customDataIdentifierIds).
	//   - RECOMMENDED (default) - Use the recommended set of managed data identifiers.
	// If this value is null, the job uses the recommended set of managed data
	// identifiers. If the job is a recurring job and this value is ALL or EXCLUDE,
	// each job run automatically uses new managed data identifiers that are released.
	// If this value is null or RECOMMENDED for a recurring job, each job run uses all
	// the managed data identifiers that are in the recommended set when the run
	// starts. For information about individual managed data identifiers or to
	// determine which ones are in the recommended set, see Using managed data
	// identifiers (https://docs.aws.amazon.com/macie/latest/user/managed-data-identifiers.html)
	// and Recommended managed data identifiers (https://docs.aws.amazon.com/macie/latest/user/discovery-jobs-mdis-recommended.html)
	// in the Amazon Macie User Guide.
	ManagedDataIdentifierSelector types.ManagedDataIdentifierSelector

	// The custom name of the job.
	Name *string

	// The S3 buckets that contain the objects to analyze, and the scope of that
	// analysis.
	S3JobDefinition *types.S3JobDefinition

	// The sampling depth, as a percentage, that determines the percentage of eligible
	// objects that the job analyzes.
	SamplingPercentage int32

	// The recurrence pattern for running the job. This value is null if the job is
	// configured to run only once.
	ScheduleFrequency *types.JobScheduleFrequency

	// The number of times that the job has run and processing statistics for the
	// job's current run.
	Statistics *types.Statistics

	// A map of key-value pairs that specifies which tags (keys and values) are
	// associated with the classification job.
	Tags map[string]string

	// If the current status of the job is USER_PAUSED, specifies when the job was
	// paused and when the job or job run will expire and be cancelled if it isn't
	// resumed. This value is present only if the value for jobStatus is USER_PAUSED.
	UserPausedDetails *types.UserPausedDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClassificationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeClassificationJob{}, middleware.After)
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
	if err = addDescribeClassificationJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeClassificationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClassificationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeClassificationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "DescribeClassificationJob",
	}
}

type opDescribeClassificationJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeClassificationJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeClassificationJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "macie2"
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
				signingName = "macie2"
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
				v4aScheme.SigningName = aws.String("macie2")
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

func addDescribeClassificationJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeClassificationJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
