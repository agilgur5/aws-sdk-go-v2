// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Inserts or deletes IPSetDescriptor objects in an IPSet . For
// each IPSetDescriptor object, you specify the following values:
//   - Whether to insert or delete the object from the array. If you want to
//     change an IPSetDescriptor object, you delete the existing object and add a new
//     one.
//   - The IP address version, IPv4 or IPv6 .
//   - The IP address in CIDR notation, for example, 192.0.2.0/24 (for the range of
//     IP addresses from 192.0.2.0 to 192.0.2.255 ) or 192.0.2.44/32 (for the
//     individual IP address 192.0.2.44 ).
//
// AWS WAF supports IPv4 address ranges: /8 and any range between /16 through /32.
// AWS WAF supports IPv6 address ranges: /24, /32, /48, /56, /64, and /128. For
// more information about CIDR notation, see the Wikipedia entry Classless
// Inter-Domain Routing (https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
// . IPv6 addresses can be represented using any of the following formats:
//   - 1111:0000:0000:0000:0000:0000:0000:0111/128
//   - 1111:0:0:0:0:0:0:0111/128
//   - 1111::0111/128
//   - 1111::111/128
//
// You use an IPSet to specify which web requests you want to allow or block based
// on the IP addresses that the requests originated from. For example, if you're
// receiving a lot of requests from one or a small number of IP addresses and you
// want to block the requests, you can create an IPSet that specifies those IP
// addresses, and then configure AWS WAF to block the requests. To create and
// configure an IPSet , perform the following steps:
//   - Submit a CreateIPSet request.
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateIPSet request.
//   - Submit an UpdateIPSet request to specify the IP addresses that you want AWS
//     WAF to watch for.
//
// When you update an IPSet , you specify the IP addresses that you want to add
// and/or the IP addresses that you want to delete. If you want to change an IP
// address, you delete the existing IP address and add the new one. You can insert
// a maximum of 1000 addresses in a single request. For more information about how
// to use the AWS WAF API to allow or block HTTP requests, see the AWS WAF
// Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/) .
func (c *Client) UpdateIPSet(ctx context.Context, params *UpdateIPSetInput, optFns ...func(*Options)) (*UpdateIPSetOutput, error) {
	if params == nil {
		params = &UpdateIPSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIPSet", params, optFns, c.addOperationUpdateIPSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIPSetInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// The IPSetId of the IPSet that you want to update. IPSetId is returned by
	// CreateIPSet and by ListIPSets .
	//
	// This member is required.
	IPSetId *string

	// An array of IPSetUpdate objects that you want to insert into or delete from an
	// IPSet . For more information, see the applicable data types:
	//   - IPSetUpdate : Contains Action and IPSetDescriptor
	//   - IPSetDescriptor : Contains Type and Value
	// You can insert a maximum of 1000 addresses in a single request.
	//
	// This member is required.
	Updates []types.IPSetUpdate

	noSmithyDocumentSerde
}

type UpdateIPSetOutput struct {

	// The ChangeToken that you used to submit the UpdateIPSet request. You can also
	// use this value to query the status of the request. For more information, see
	// GetChangeTokenStatus .
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIPSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateIPSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateIPSet{}, middleware.After)
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
	if err = addOpUpdateIPSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIPSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIPSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateIPSet",
	}
}
