// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Returns configuration for an Object Lambda Access Point. The following actions
// are related to GetAccessPointConfigurationForObjectLambda :
//   - PutAccessPointConfigurationForObjectLambda (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointConfigurationForObjectLambda.html)
func (c *Client) GetAccessPointConfigurationForObjectLambda(ctx context.Context, params *GetAccessPointConfigurationForObjectLambdaInput, optFns ...func(*Options)) (*GetAccessPointConfigurationForObjectLambdaOutput, error) {
	if params == nil {
		params = &GetAccessPointConfigurationForObjectLambdaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessPointConfigurationForObjectLambda", params, optFns, c.addOperationGetAccessPointConfigurationForObjectLambdaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessPointConfigurationForObjectLambdaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointConfigurationForObjectLambdaInput struct {

	// The account ID for the account that owns the specified Object Lambda Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// The name of the Object Lambda Access Point you want to return the configuration
	// for.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

func (*GetAccessPointConfigurationForObjectLambdaInput) operationName() string {
	return "GetAccessPointConfigurationForObjectLambda"
}

func (in *GetAccessPointConfigurationForObjectLambdaInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type GetAccessPointConfigurationForObjectLambdaOutput struct {

	// Object Lambda Access Point configuration document.
	Configuration *types.ObjectLambdaConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessPointConfigurationForObjectLambdaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPointConfigurationForObjectLambda{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPointConfigurationForObjectLambda{}, middleware.After)
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetAccessPointConfigurationForObjectLambdaMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetAccessPointConfigurationForObjectLambdaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPointConfigurationForObjectLambda(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetAccessPointConfigurationForObjectLambdaUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetAccessPointConfigurationForObjectLambdaMiddleware struct {
}

func (*endpointPrefix_opGetAccessPointConfigurationForObjectLambdaMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessPointConfigurationForObjectLambdaMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetAccessPointConfigurationForObjectLambdaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetAccessPointConfigurationForObjectLambdaMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetAccessPointConfigurationForObjectLambdaMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessPointConfigurationForObjectLambda(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetAccessPointConfigurationForObjectLambda",
	}
}

func copyGetAccessPointConfigurationForObjectLambdaInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetAccessPointConfigurationForObjectLambdaInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetAccessPointConfigurationForObjectLambdaInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillGetAccessPointConfigurationForObjectLambdaAccountID(input interface{}, v string) error {
	in := input.(*GetAccessPointConfigurationForObjectLambdaInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetAccessPointConfigurationForObjectLambdaUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetAccessPointConfigurationForObjectLambdaInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
