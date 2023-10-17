// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to either upgrade your Amazon OpenSearch Service domain or perform
// an upgrade eligibility check to a compatible version of OpenSearch or
// Elasticsearch.
func (c *Client) UpgradeDomain(ctx context.Context, params *UpgradeDomainInput, optFns ...func(*Options)) (*UpgradeDomainOutput, error) {
	if params == nil {
		params = &UpgradeDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpgradeDomain", params, optFns, c.addOperationUpgradeDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpgradeDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the UpgradeDomain operation.
type UpgradeDomainInput struct {

	// Name of the OpenSearch Service domain that you want to upgrade.
	//
	// This member is required.
	DomainName *string

	// OpenSearch or Elasticsearch version to which you want to upgrade, in the format
	// Opensearch_X.Y or Elasticsearch_X.Y.
	//
	// This member is required.
	TargetVersion *string

	// Only supports the override_main_response_version parameter and not other
	// advanced options. You can only include this option when upgrading to an
	// OpenSearch version. Specifies whether the domain reports its version as 7.10 so
	// that it continues to work with Elasticsearch OSS clients and plugins.
	AdvancedOptions map[string]string

	// When true, indicates that an upgrade eligibility check needs to be performed.
	// Does not actually perform the upgrade.
	PerformCheckOnly *bool

	noSmithyDocumentSerde
}

// Container for the response returned by UpgradeDomain operation.
type UpgradeDomainOutput struct {

	// The advanced options configuration for the domain.
	AdvancedOptions map[string]string

	// Container for information about a configuration change happening on a domain.
	ChangeProgressDetails *types.ChangeProgressDetails

	// The name of the domain that was upgraded.
	DomainName *string

	// When true, indicates that an upgrade eligibility check was performed.
	PerformCheckOnly *bool

	// OpenSearch or Elasticsearch version that the domain was upgraded to.
	TargetVersion *string

	// The unique identifier of the domain upgrade.
	UpgradeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpgradeDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpgradeDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpgradeDomain{}, middleware.After)
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
	if err = addOpUpgradeDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpgradeDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpgradeDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "UpgradeDomain",
	}
}
