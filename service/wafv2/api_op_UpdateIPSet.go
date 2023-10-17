// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified IPSet . This operation completely replaces the mutable
// specifications that you already have for the IP set with the ones that you
// provide to this call. To modify an IP set, do the following:
//   - Retrieve it by calling GetIPSet
//   - Update its settings as needed
//   - Provide the complete IP set specification to this call
//
// When you make changes to web ACLs or web ACL components, like rules and rule
// groups, WAF propagates the changes everywhere that the web ACL and its
// components are stored and used. Your changes are applied within seconds, but
// there might be a brief period of inconsistency when the changes have arrived in
// some places and not in others. So, for example, if you change a rule action
// setting, the action might be the old action in one area and the new action in
// another area. Or if you add an IP address to an IP set used in a blocking rule,
// the new address might briefly be blocked in one area while still allowed in
// another. This temporary inconsistency can occur when you first associate a web
// ACL with an Amazon Web Services resource and when you change a web ACL that is
// already associated with a resource. Generally, any inconsistencies of this type
// last only a few seconds.
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

	// Contains an array of strings that specifies zero or more IP addresses or blocks
	// of IP addresses that you want WAF to inspect for in incoming requests. All
	// addresses must be specified using Classless Inter-Domain Routing (CIDR)
	// notation. WAF supports all IPv4 and IPv6 CIDR ranges except for /0 . Example
	// address strings:
	//   - For requests that originated from the IP address 192.0.2.44, specify
	//   192.0.2.44/32 .
	//   - For requests that originated from IP addresses from 192.0.2.0 to
	//   192.0.2.255, specify 192.0.2.0/24 .
	//   - For requests that originated from the IP address
	//   1111:0000:0000:0000:0000:0000:0000:0111, specify
	//   1111:0000:0000:0000:0000:0000:0000:0111/128 .
	//   - For requests that originated from IP addresses
	//   1111:0000:0000:0000:0000:0000:0000:0000 to
	//   1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify
	//   1111:0000:0000:0000:0000:0000:0000:0000/64 .
	// For more information about CIDR notation, see the Wikipedia entry Classless
	// Inter-Domain Routing (https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// . Example JSON Addresses specifications:
	//   - Empty array: "Addresses": []
	//   - Array with one address: "Addresses": ["192.0.2.44/32"]
	//   - Array with three addresses: "Addresses": ["192.0.2.44/32", "192.0.2.0/24",
	//   "192.0.0.0/16"]
	//   - INVALID specification: "Addresses": [""] INVALID
	//
	// This member is required.
	Addresses []string

	// A unique identifier for the set. This ID is returned in the responses to create
	// and list commands. You provide it to operations like update and delete.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. WAF returns a token to your get and list
	// requests, to mark the state of the entity at the time of the request. To make
	// changes to the entity associated with the token, you provide the token to
	// operations like update and delete . WAF uses the token to ensure that no changes
	// have been made to the entity since you last retrieved it. If a change has been
	// made, the update fails with a WAFOptimisticLockException . If this happens,
	// perform another get , and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the IP set. You cannot change the name of an IPSet after you create
	// it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, an App Runner service, or an Amazon Web Services Verified Access
	// instance. To work with CloudFront, you must also specify the Region US East (N.
	// Virginia) as follows:
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// A description of the IP set that helps with identification.
	Description *string

	noSmithyDocumentSerde
}

type UpdateIPSetOutput struct {

	// A token used for optimistic locking. WAF returns this token to your update
	// requests. You use NextLockToken in the same manner as you use LockToken .
	NextLockToken *string

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
		SigningName:   "wafv2",
		OperationName: "UpdateIPSet",
	}
}
