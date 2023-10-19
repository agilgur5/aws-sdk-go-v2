// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the name of the label.
func (c *Client) DescribeLabel(ctx context.Context, params *DescribeLabelInput, optFns ...func(*Options)) (*DescribeLabelOutput, error) {
	if params == nil {
		params = &DescribeLabelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLabel", params, optFns, c.addOperationDescribeLabelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLabelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLabelInput struct {

	// Returns the name of the group containing the label.
	//
	// This member is required.
	LabelGroupName *string

	// Returns the ID of the label.
	//
	// This member is required.
	LabelId *string

	noSmithyDocumentSerde
}

type DescribeLabelOutput struct {

	// The time at which the label was created.
	CreatedAt *time.Time

	// The end time of the requested label.
	EndTime *time.Time

	// Indicates that a label pertains to a particular piece of equipment.
	Equipment *string

	// Indicates the type of anomaly associated with the label. Data in this field
	// will be retained for service usage. Follow best practices for the security of
	// your data.
	FaultCode *string

	// The Amazon Resource Name (ARN) of the requested label group.
	LabelGroupArn *string

	// The name of the requested label group.
	LabelGroupName *string

	// The ID of the requested label.
	LabelId *string

	// Metadata providing additional information about the label. Data in this field
	// will be retained for service usage. Follow best practices for the security of
	// your data.
	Notes *string

	// Indicates whether a labeled event represents an anomaly.
	Rating types.LabelRating

	// The start time of the requested label.
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLabelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeLabel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeLabel{}, middleware.After)
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
	if err = addOpDescribeLabelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLabel(options.Region), middleware.Before); err != nil {
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
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "DescribeLabel",
		options:   options,
	}, "ResolveEndpointV2", middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&signRequestMiddleware{}, middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&getIdentityMiddleware{
		options: options,
	}, middleware.Before)
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeLabel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "DescribeLabel",
	}
}
