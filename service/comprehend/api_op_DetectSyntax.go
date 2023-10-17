// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Inspects text for syntax and the part of speech of words in the document. For
// more information, see Syntax (https://docs.aws.amazon.com/comprehend/latest/dg/how-syntax.html)
// in the Comprehend Developer Guide.
func (c *Client) DetectSyntax(ctx context.Context, params *DetectSyntaxInput, optFns ...func(*Options)) (*DetectSyntaxOutput, error) {
	if params == nil {
		params = &DetectSyntaxInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectSyntax", params, optFns, c.addOperationDetectSyntaxMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectSyntaxOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectSyntaxInput struct {

	// The language code of the input documents. You can specify any of the following
	// languages supported by Amazon Comprehend: German ("de"), English ("en"), Spanish
	// ("es"), French ("fr"), Italian ("it"), or Portuguese ("pt").
	//
	// This member is required.
	LanguageCode types.SyntaxLanguageCode

	// A UTF-8 string. The maximum string size is 5 KB.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type DetectSyntaxOutput struct {

	// A collection of syntax tokens describing the text. For each token, the response
	// provides the text, the token type, where the text begins and ends, and the level
	// of confidence that Amazon Comprehend has that the token is correct. For a list
	// of token types, see Syntax (https://docs.aws.amazon.com/comprehend/latest/dg/how-syntax.html)
	// in the Comprehend Developer Guide.
	SyntaxTokens []types.SyntaxToken

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectSyntaxMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectSyntax{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectSyntax{}, middleware.After)
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
	if err = addOpDetectSyntaxValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectSyntax(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectSyntax(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectSyntax",
	}
}
