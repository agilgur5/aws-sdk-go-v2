// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of names for one or more on-premises instances. Unless otherwise
// specified, both registered and deregistered on-premises instance names are
// listed. To list only registered or deregistered on-premises instance names, use
// the registration status parameter.
func (c *Client) ListOnPremisesInstances(ctx context.Context, params *ListOnPremisesInstancesInput, optFns ...func(*Options)) (*ListOnPremisesInstancesOutput, error) {
	if params == nil {
		params = &ListOnPremisesInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOnPremisesInstances", params, optFns, c.addOperationListOnPremisesInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOnPremisesInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListOnPremisesInstances operation.
type ListOnPremisesInstancesInput struct {

	// An identifier returned from the previous list on-premises instances call. It
	// can be used to return the next set of on-premises instances in the list.
	NextToken *string

	// The registration status of the on-premises instances:
	//   - Deregistered : Include deregistered on-premises instances in the resulting
	//   list.
	//   - Registered : Include registered on-premises instances in the resulting list.
	RegistrationStatus types.RegistrationStatus

	// The on-premises instance tags that are used to restrict the on-premises
	// instance names returned.
	TagFilters []types.TagFilter

	noSmithyDocumentSerde
}

// Represents the output of the list on-premises instances operation.
type ListOnPremisesInstancesOutput struct {

	// The list of matching on-premises instance names.
	InstanceNames []string

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list on-premises instances call to return the
	// next set of on-premises instances in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOnPremisesInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOnPremisesInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOnPremisesInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOnPremisesInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOnPremisesInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListOnPremisesInstances",
	}
}
