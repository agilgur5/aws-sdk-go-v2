// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the access that is assigned to the specific file transfer
// protocol-enabled server, as identified by its ServerId property and its
// ExternalId . The response from this call returns the properties of the access
// that is associated with the ServerId value that was specified.
func (c *Client) DescribeAccess(ctx context.Context, params *DescribeAccessInput, optFns ...func(*Options)) (*DescribeAccessOutput, error) {
	if params == nil {
		params = &DescribeAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccess", params, optFns, c.addOperationDescribeAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccessInput struct {

	// A unique identifier that is required to identify specific groups within your
	// directory. The users of the group that you associate have access to your Amazon
	// S3 or Amazon EFS resources over the enabled protocols using Transfer Family. If
	// you know the group name, you can view the SID values by running the following
	// command using Windows PowerShell. Get-ADGroup -Filter {samAccountName -like
	// "YourGroupName*"} -Properties * | Select SamAccountName,ObjectSid In that
	// command, replace YourGroupName with the name of your Active Directory group. The
	// regular expression used to validate this parameter is a string of characters
	// consisting of uppercase and lowercase alphanumeric characters with no spaces.
	// You can also include underscores or any of the following characters: =,.@:/-
	//
	// This member is required.
	ExternalId *string

	// A system-assigned unique identifier for a server that has this access assigned.
	//
	// This member is required.
	ServerId *string

	noSmithyDocumentSerde
}

type DescribeAccessOutput struct {

	// The external identifier of the server that the access is attached to.
	//
	// This member is required.
	Access *types.DescribedAccess

	// A system-assigned unique identifier for a server that has this access assigned.
	//
	// This member is required.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAccess{}, middleware.After)
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
	if err = addOpDescribeAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccess(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeAccess",
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

func newServiceMetadataMiddleware_opDescribeAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DescribeAccess",
	}
}
