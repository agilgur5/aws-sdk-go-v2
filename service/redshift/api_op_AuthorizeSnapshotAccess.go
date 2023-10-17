// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Authorizes the specified Amazon Web Services account to restore the specified
// snapshot. For more information about working with snapshots, go to Amazon
// Redshift Snapshots (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) AuthorizeSnapshotAccess(ctx context.Context, params *AuthorizeSnapshotAccessInput, optFns ...func(*Options)) (*AuthorizeSnapshotAccessOutput, error) {
	if params == nil {
		params = &AuthorizeSnapshotAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeSnapshotAccess", params, optFns, c.addOperationAuthorizeSnapshotAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeSnapshotAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeSnapshotAccessInput struct {

	// The identifier of the Amazon Web Services account authorized to restore the
	// specified snapshot. To share a snapshot with Amazon Web Services Support,
	// specify amazon-redshift-support.
	//
	// This member is required.
	AccountWithRestoreAccess *string

	// The Amazon Resource Name (ARN) of the snapshot to authorize access to.
	SnapshotArn *string

	// The identifier of the cluster the snapshot was created from. This parameter is
	// required if your IAM user has a policy containing a snapshot resource element
	// that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string

	// The identifier of the snapshot the account is authorized to restore.
	SnapshotIdentifier *string

	noSmithyDocumentSerde
}

type AuthorizeSnapshotAccessOutput struct {

	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAuthorizeSnapshotAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAuthorizeSnapshotAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAuthorizeSnapshotAccess{}, middleware.After)
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
	if err = addOpAuthorizeSnapshotAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeSnapshotAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeSnapshotAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "AuthorizeSnapshotAccess",
	}
}
