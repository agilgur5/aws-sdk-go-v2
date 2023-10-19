// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a custom terminology.
func (c *Client) GetTerminology(ctx context.Context, params *GetTerminologyInput, optFns ...func(*Options)) (*GetTerminologyOutput, error) {
	if params == nil {
		params = &GetTerminologyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTerminology", params, optFns, c.addOperationGetTerminologyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTerminologyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTerminologyInput struct {

	// The name of the custom terminology being retrieved.
	//
	// This member is required.
	Name *string

	// The data format of the custom terminology being retrieved. If you don't specify
	// this parameter, Amazon Translate returns a file with the same format as the file
	// that was imported to create the terminology. If you specify this parameter when
	// you retrieve a multi-directional terminology resource, you must specify the same
	// format as the input file that was imported to create it. Otherwise, Amazon
	// Translate throws an error.
	TerminologyDataFormat types.TerminologyDataFormat

	noSmithyDocumentSerde
}

type GetTerminologyOutput struct {

	// The Amazon S3 location of a file that provides any errors or warnings that were
	// produced by your input file. This file was created when Amazon Translate
	// attempted to create a terminology resource. The location is returned as a
	// presigned URL to that has a 30-minute expiration.
	AuxiliaryDataLocation *types.TerminologyDataLocation

	// The Amazon S3 location of the most recent custom terminology input file that
	// was successfully imported into Amazon Translate. The location is returned as a
	// presigned URL that has a 30-minute expiration. Amazon Translate doesn't scan all
	// input files for the risk of CSV injection attacks. CSV injection occurs when a
	// .csv or .tsv file is altered so that a record contains malicious code. The
	// record begins with a special character, such as =, +, -, or @. When the file is
	// opened in a spreadsheet program, the program might interpret the record as a
	// formula and run the code within it. Before you download an input file from
	// Amazon S3, ensure that you recognize the file and trust its creator.
	TerminologyDataLocation *types.TerminologyDataLocation

	// The properties of the custom terminology being retrieved.
	TerminologyProperties *types.TerminologyProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTerminologyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTerminology{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTerminology{}, middleware.After)
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
	if err = addOpGetTerminologyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTerminology(options.Region), middleware.Before); err != nil {
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
		operation: "GetTerminology",
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

func newServiceMetadataMiddleware_opGetTerminology(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "GetTerminology",
	}
}
