// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The SetVisibleToAllUsers parameter is no longer supported. Your cluster may be
// visible to all users in your account. To restrict cluster access using an IAM
// policy, see Identity and Access Management for Amazon EMR (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-access-IAM.html)
// . Sets the Cluster$VisibleToAllUsers value for an Amazon EMR cluster. When true
// , IAM principals in the Amazon Web Services account can perform Amazon EMR
// cluster actions that their IAM policies allow. When false , only the IAM
// principal that created the cluster and the Amazon Web Services account root user
// can perform Amazon EMR actions on the cluster, regardless of IAM permissions
// policies attached to other IAM principals. This action works on running
// clusters. When you create a cluster, use the RunJobFlowInput$VisibleToAllUsers
// parameter. For more information, see Understanding the Amazon EMR Cluster
// VisibleToAllUsers Setting (https://docs.aws.amazon.com/emr/latest/ManagementGuide/security_IAM_emr-with-IAM.html#security_set_visible_to_all_users)
// in the Amazon EMR Management Guide.
func (c *Client) SetVisibleToAllUsers(ctx context.Context, params *SetVisibleToAllUsersInput, optFns ...func(*Options)) (*SetVisibleToAllUsersOutput, error) {
	if params == nil {
		params = &SetVisibleToAllUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetVisibleToAllUsers", params, optFns, c.addOperationSetVisibleToAllUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetVisibleToAllUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the SetVisibleToAllUsers action.
type SetVisibleToAllUsersInput struct {

	// The unique identifier of the job flow (cluster).
	//
	// This member is required.
	JobFlowIds []string

	// A value of true indicates that an IAM principal in the Amazon Web Services
	// account can perform Amazon EMR actions on the cluster that the IAM policies
	// attached to the principal allow. A value of false indicates that only the IAM
	// principal that created the cluster and the Amazon Web Services root user can
	// perform Amazon EMR actions on the cluster.
	//
	// This member is required.
	VisibleToAllUsers bool

	noSmithyDocumentSerde
}

type SetVisibleToAllUsersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetVisibleToAllUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetVisibleToAllUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetVisibleToAllUsers{}, middleware.After)
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
	if err = addOpSetVisibleToAllUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetVisibleToAllUsers(options.Region), middleware.Before); err != nil {
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
		operation: "SetVisibleToAllUsers",
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

func newServiceMetadataMiddleware_opSetVisibleToAllUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "SetVisibleToAllUsers",
	}
}
