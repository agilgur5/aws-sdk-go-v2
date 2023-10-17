// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the settings to use for a cluster.
func (c *Client) UpdateClusterSettings(ctx context.Context, params *UpdateClusterSettingsInput, optFns ...func(*Options)) (*UpdateClusterSettingsOutput, error) {
	if params == nil {
		params = &UpdateClusterSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClusterSettings", params, optFns, c.addOperationUpdateClusterSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterSettingsInput struct {

	// The name of the cluster to modify the settings for.
	//
	// This member is required.
	Cluster *string

	// The setting to use by default for a cluster. This parameter is used to turn on
	// CloudWatch Container Insights for a cluster. If this value is specified, it
	// overrides the containerInsights value set with PutAccountSetting or
	// PutAccountSettingDefault . Currently, if you delete an existing cluster that
	// does not have Container Insights turned on, and then create a new cluster with
	// the same name with Container Insights tuned on, Container Insights will not
	// actually be turned on. If you want to preserve the same name for your existing
	// cluster and turn on Container Insights, you must wait 7 days before you can
	// re-create it.
	//
	// This member is required.
	Settings []types.ClusterSetting

	noSmithyDocumentSerde
}

type UpdateClusterSettingsOutput struct {

	// Details about the cluster
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateClusterSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateClusterSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateClusterSettings{}, middleware.After)
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
	if err = addOpUpdateClusterSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClusterSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateClusterSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateClusterSettings",
	}
}
