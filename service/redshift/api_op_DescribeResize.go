// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the last resize operation for the specified cluster.
// If no resize operation has ever been initiated for the specified cluster, a
// HTTP 404 error is returned. If a resize operation was initiated and completed,
// the status of the resize remains as SUCCEEDED until the next resize. A resize
// operation can be requested using ModifyCluster and specifying a different
// number or type of nodes for the cluster.
func (c *Client) DescribeResize(ctx context.Context, params *DescribeResizeInput, optFns ...func(*Options)) (*DescribeResizeOutput, error) {
	if params == nil {
		params = &DescribeResizeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeResize", params, optFns, c.addOperationDescribeResizeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeResizeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeResizeInput struct {

	// The unique identifier of a cluster whose resize progress you are requesting.
	// This parameter is case-sensitive. By default, resize operations for all clusters
	// defined for an Amazon Web Services account are returned.
	//
	// This member is required.
	ClusterIdentifier *string

	noSmithyDocumentSerde
}

func (*DescribeResizeInput) operationName() string {
	return "DescribeResize"
}

// Describes the result of a cluster resize operation.
type DescribeResizeOutput struct {

	// The average rate of the resize operation over the last few minutes, measured in
	// megabytes per second. After the resize operation completes, this value shows the
	// average rate of the entire resize operation.
	AvgResizeRateInMegaBytesPerSecond *float64

	// The percent of data transferred from source cluster to target cluster.
	DataTransferProgressPercent *float64

	// The amount of seconds that have elapsed since the resize operation began. After
	// the resize operation completes, this value shows the total actual time, in
	// seconds, for the resize operation.
	ElapsedTimeInSeconds *int64

	// The estimated time remaining, in seconds, until the resize operation is
	// complete. This value is calculated based on the average resize rate and the
	// estimated amount of data remaining to be processed. Once the resize operation is
	// complete, this value will be 0.
	EstimatedTimeToCompletionInSeconds *int64

	// The names of tables that have been completely imported . Valid Values: List of
	// table names.
	ImportTablesCompleted []string

	// The names of tables that are being currently imported. Valid Values: List of
	// table names.
	ImportTablesInProgress []string

	// The names of tables that have not been yet imported. Valid Values: List of
	// table names
	ImportTablesNotStarted []string

	// An optional string to provide additional details about the resize action.
	Message *string

	// While the resize operation is in progress, this value shows the current amount
	// of data, in megabytes, that has been processed so far. When the resize operation
	// is complete, this value shows the total amount of data, in megabytes, on the
	// cluster, which may be more or less than TotalResizeDataInMegaBytes (the
	// estimated total amount of data before resize).
	ProgressInMegaBytes *int64

	// An enum with possible values of ClassicResize and ElasticResize . These values
	// describe the type of resize operation being performed.
	ResizeType *string

	// The status of the resize operation. Valid Values: NONE | IN_PROGRESS | FAILED |
	// SUCCEEDED | CANCELLING
	Status *string

	// The cluster type after the resize operation is complete. Valid Values:
	// multi-node | single-node
	TargetClusterType *string

	// The type of encryption for the cluster after the resize is complete. Possible
	// values are KMS and None .
	TargetEncryptionType *string

	// The node type that the cluster will have after the resize operation is complete.
	TargetNodeType *string

	// The number of nodes that the cluster will have after the resize operation is
	// complete.
	TargetNumberOfNodes *int32

	// The estimated total amount of data, in megabytes, on the cluster before the
	// resize operation began.
	TotalResizeDataInMegaBytes *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeResizeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeResize{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeResize{}, middleware.After)
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
	if err = addDescribeResizeResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeResizeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeResize(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeResize(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeResize",
	}
}

type opDescribeResizeResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeResizeResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeResizeResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "redshift"
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
				signingName = "redshift"
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
				v4aScheme.SigningName = aws.String("redshift")
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

func addDescribeResizeResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeResizeResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
