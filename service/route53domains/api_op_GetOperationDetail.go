// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns the current status of an operation that is not completed.
func (c *Client) GetOperationDetail(ctx context.Context, params *GetOperationDetailInput, optFns ...func(*Options)) (*GetOperationDetailOutput, error) {
	if params == nil {
		params = &GetOperationDetailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOperationDetail", params, optFns, c.addOperationGetOperationDetailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOperationDetailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// request includes the following element.
type GetOperationDetailInput struct {

	// The identifier for the operation for which you want to get the status. Route 53
	// returned the identifier in the response to the original request.
	//
	// This member is required.
	OperationId *string

	noSmithyDocumentSerde
}

// The GetOperationDetail response includes the following elements.
type GetOperationDetailOutput struct {

	// The name of a domain.
	DomainName *string

	// The date when the operation was last updated.
	LastUpdatedDate *time.Time

	// Detailed information on the status including possible errors.
	Message *string

	// The identifier for the operation.
	OperationId *string

	// The current status of the requested operation in the system.
	Status types.OperationStatus

	// Lists any outstanding operations that require customer action. Valid values
	// are:
	//   - PENDING_ACCEPTANCE : The operation is waiting for acceptance from the
	//   account that is receiving the domain.
	//   - PENDING_CUSTOMER_ACTION : The operation is waiting for customer action, for
	//   example, returning an email.
	//   - PENDING_AUTHORIZATION : The operation is waiting for the form of
	//   authorization. For more information, see ResendOperationAuthorization (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ResendOperationAuthorization.html)
	//   .
	//   - PENDING_PAYMENT_VERIFICATION : The operation is waiting for the payment
	//   method to validate.
	//   - PENDING_SUPPORT_CASE : The operation includes a support case and is waiting
	//   for its resolution.
	StatusFlag types.StatusFlag

	// The date when the request was submitted.
	SubmittedDate *time.Time

	// The type of operation that was requested.
	Type types.OperationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOperationDetailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOperationDetail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOperationDetail{}, middleware.After)
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
	if err = addOpGetOperationDetailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOperationDetail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOperationDetail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "GetOperationDetail",
	}
}
