// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified principal ARN with the specified portfolio. If you
// share the portfolio with principal name sharing enabled, the PrincipalARN
// association is included in the share. The PortfolioID , PrincipalARN , and
// PrincipalType parameters are required. You can associate a maximum of 10
// Principals with a portfolio using PrincipalType as IAM_PATTERN . When you
// associate a principal with portfolio, a potential privilege escalation path may
// occur when that portfolio is then shared with other accounts. For a user in a
// recipient account who is not an Service Catalog Admin, but still has the ability
// to create Principals (Users/Groups/Roles), that user could create a role that
// matches a principal name association for the portfolio. Although this user may
// not know which principal names are associated through Service Catalog, they may
// be able to guess the user. If this potential escalation path is a concern, then
// Service Catalog recommends using PrincipalType as IAM . With this configuration,
// the PrincipalARN must already exist in the recipient account before it can be
// associated.
func (c *Client) AssociatePrincipalWithPortfolio(ctx context.Context, params *AssociatePrincipalWithPortfolioInput, optFns ...func(*Options)) (*AssociatePrincipalWithPortfolioOutput, error) {
	if params == nil {
		params = &AssociatePrincipalWithPortfolioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociatePrincipalWithPortfolio", params, optFns, c.addOperationAssociatePrincipalWithPortfolioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociatePrincipalWithPortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociatePrincipalWithPortfolioInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The ARN of the principal (user, role, or group). If the PrincipalType is IAM ,
	// the supported value is a fully defined IAM Amazon Resource Name (ARN) (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns)
	// . If the PrincipalType is IAM_PATTERN , the supported value is an IAM ARN
	// without an AccountID in the following format:
	// arn:partition:iam:::resource-type/resource-id The ARN resource-id can be either:
	//
	//   - A fully formed resource-id. For example, arn:aws:iam:::role/resource-name
	//   or arn:aws:iam:::role/resource-path/resource-name
	//   - A wildcard ARN. The wildcard ARN accepts IAM_PATTERN values with a "*" or
	//   "?" in the resource-id segment of the ARN. For example
	//   arn:partition:service:::resource-type/resource-path/resource-name. The new
	//   symbols are exclusive to the resource-path and resource-name and cannot replace
	//   the resource-type or other ARN values. The ARN path and principal name allow
	//   unlimited wildcard characters.
	// Examples of an acceptable wildcard ARN:
	//   - arn:aws:iam:::role/ResourceName_*
	//   - arn:aws:iam:::role/*/ResourceName_?
	// Examples of an unacceptable wildcard ARN:
	//   - arn:aws:iam:::*/ResourceName
	// You can associate multiple IAM_PATTERN s even if the account has no principal
	// with that name. The "?" wildcard character matches zero or one of any character.
	// This is similar to ".?" in regular regex context. The "*" wildcard character
	// matches any number of any characters. This is similar to ".*" in regular regex
	// context. In the IAM Principal ARN format
	// (arn:partition:iam:::resource-type/resource-path/resource-name), valid
	// resource-type values include user/, group/, or role/. The "?" and "*" characters
	// are allowed only after the resource-type in the resource-id segment. You can use
	// special characters anywhere within the resource-id. The "*" character also
	// matches the "/" character, allowing paths to be formed within the resource-id.
	// For example, arn:aws:iam:::role/*/ResourceName_? matches both
	// arn:aws:iam:::role/pathA/pathB/ResourceName_1 and
	// arn:aws:iam:::role/pathA/ResourceName_1.
	//
	// This member is required.
	PrincipalARN *string

	// The principal type. The supported value is IAM if you use a fully defined
	// Amazon Resource Name (ARN), or IAM_PATTERN if you use an ARN with no accountID ,
	// with or without wildcard characters.
	//
	// This member is required.
	PrincipalType types.PrincipalType

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	noSmithyDocumentSerde
}

type AssociatePrincipalWithPortfolioOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociatePrincipalWithPortfolioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociatePrincipalWithPortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociatePrincipalWithPortfolio{}, middleware.After)
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
	if err = addOpAssociatePrincipalWithPortfolioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePrincipalWithPortfolio(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociatePrincipalWithPortfolio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "AssociatePrincipalWithPortfolio",
	}
}
