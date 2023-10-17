// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an existing Amazon Kendra thesaurus.
func (c *Client) DescribeThesaurus(ctx context.Context, params *DescribeThesaurusInput, optFns ...func(*Options)) (*DescribeThesaurusOutput, error) {
	if params == nil {
		params = &DescribeThesaurusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeThesaurus", params, optFns, c.addOperationDescribeThesaurusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeThesaurusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeThesaurusInput struct {

	// The identifier of the thesaurus you want to get information on.
	//
	// This member is required.
	Id *string

	// The identifier of the index for the thesaurus.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeThesaurusOutput struct {

	// The Unix timestamp when the thesaurus was created.
	CreatedAt *time.Time

	// The thesaurus description.
	Description *string

	// When the Status field value is FAILED , the ErrorMessage field provides more
	// information.
	ErrorMessage *string

	// The size of the thesaurus file in bytes.
	FileSizeBytes *int64

	// The identifier of the thesaurus.
	Id *string

	// The identifier of the index for the thesaurus.
	IndexId *string

	// The thesaurus name.
	Name *string

	// An IAM role that gives Amazon Kendra permissions to access thesaurus file
	// specified in SourceS3Path .
	RoleArn *string

	// Information required to find a specific file in an Amazon S3 bucket.
	SourceS3Path *types.S3Path

	// The current status of the thesaurus. When the value is ACTIVE , queries are able
	// to use the thesaurus. If the Status field value is FAILED , the ErrorMessage
	// field provides more information. If the status is ACTIVE_BUT_UPDATE_FAILED , it
	// means that Amazon Kendra could not ingest the new thesaurus file. The old
	// thesaurus file is still active.
	Status types.ThesaurusStatus

	// The number of synonym rules in the thesaurus file.
	SynonymRuleCount *int64

	// The number of unique terms in the thesaurus file. For example, the synonyms
	// a,b,c and a=>d , the term count would be 4.
	TermCount *int64

	// The Unix timestamp when the thesaurus was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeThesaurusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeThesaurus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeThesaurus{}, middleware.After)
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
	if err = addOpDescribeThesaurusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThesaurus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeThesaurus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeThesaurus",
	}
}
