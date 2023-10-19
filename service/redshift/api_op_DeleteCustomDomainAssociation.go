// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Contains information about deleting a custom domain association for a cluster.
func (c *Client) DeleteCustomDomainAssociation(ctx context.Context, params *DeleteCustomDomainAssociationInput, optFns ...func(*Options)) (*DeleteCustomDomainAssociationOutput, error) {
	if params == nil {
		params = &DeleteCustomDomainAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomDomainAssociation", params, optFns, c.addOperationDeleteCustomDomainAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomDomainAssociationInput struct {

	// The identifier of the cluster to delete a custom domain association for.
	//
	// This member is required.
	ClusterIdentifier *string

	noSmithyDocumentSerde
}

type DeleteCustomDomainAssociationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCustomDomainAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCustomDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCustomDomainAssociation{}, middleware.After)
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
	if err = addOpDeleteCustomDomainAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomDomainAssociation(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteCustomDomainAssociation",
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

func newServiceMetadataMiddleware_opDeleteCustomDomainAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeleteCustomDomainAssociation",
	}
}
