// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Simulate how a set of IAM policies and optionally a resource-based policy works
// with a list of API operations and Amazon Web Services resources to determine the
// policies' effective permissions. The policies are provided as strings. The
// simulation does not perform the API operations; it only checks the authorization
// to determine if the simulated policies allow or deny the operations. You can
// simulate resources that don't exist in your account. If you want to simulate
// existing policies that are attached to an IAM user, group, or role, use
// SimulatePrincipalPolicy instead. Context keys are variables that are maintained
// by Amazon Web Services and its services and which provide details about the
// context of an API query request. You can use the Condition element of an IAM
// policy to evaluate context keys. To get the list of context keys that the
// policies require for correct simulation, use GetContextKeysForCustomPolicy . If
// the output is long, you can use MaxItems and Marker parameters to paginate the
// results. The IAM policy simulator evaluates statements in the identity-based
// policy and the inputs that you provide during simulation. The policy simulator
// results can differ from your live Amazon Web Services environment. We recommend
// that you check your policies against your live Amazon Web Services environment
// after testing using the policy simulator to confirm that you have the desired
// results. For more information about using the policy simulator, see Testing IAM
// policies with the IAM policy simulator  (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html)
// in the IAM User Guide.
func (c *Client) SimulateCustomPolicy(ctx context.Context, params *SimulateCustomPolicyInput, optFns ...func(*Options)) (*SimulateCustomPolicyOutput, error) {
	if params == nil {
		params = &SimulateCustomPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SimulateCustomPolicy", params, optFns, c.addOperationSimulateCustomPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SimulateCustomPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimulateCustomPolicyInput struct {

	// A list of names of API operations to evaluate in the simulation. Each operation
	// is evaluated against each resource. Each operation must include the service
	// identifier, such as iam:CreateUser . This operation does not support using
	// wildcards (*) in an action name.
	//
	// This member is required.
	ActionNames []string

	// A list of policy documents to include in the simulation. Each document is
	// specified as a string containing the complete, valid JSON text of an IAM policy.
	// Do not include any resource-based policies in this parameter. Any resource-based
	// policy must be submitted with the ResourcePolicy parameter. The policies cannot
	// be "scope-down" policies, such as you could include in a call to
	// GetFederationToken (https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetFederationToken.html)
	// or one of the AssumeRole (https://docs.aws.amazon.com/IAM/latest/APIReference/API_AssumeRole.html)
	// API operations. In other words, do not use policies designed to restrict what a
	// user can do while using the temporary credentials. The maximum length of the
	// policy document that you can pass in this operation, including whitespace, is
	// listed below. To view the maximum character counts of a managed policy with no
	// whitespaces, see IAM and STS character quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length)
	// . The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// This member is required.
	PolicyInputList []string

	// The ARN of the IAM user that you want to use as the simulated caller of the API
	// operations. CallerArn is required if you include a ResourcePolicy so that the
	// policy's Principal element has a value to use in evaluating the policy. You can
	// specify only the ARN of an IAM user. You cannot specify the ARN of an assumed
	// role, federated user, or a service principal.
	CallerArn *string

	// A list of context keys and corresponding values for the simulation to use.
	// Whenever a context key is evaluated in one of the simulated IAM permissions
	// policies, the corresponding value is supplied.
	ContextEntries []types.ContextEntry

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	// The IAM permissions boundary policy to simulate. The permissions boundary sets
	// the maximum permissions that an IAM entity can have. You can input only one
	// permissions boundary when you pass a policy to this operation. For more
	// information about permissions boundaries, see Permissions boundaries for IAM
	// entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide. The policy input is specified as a string that contains
	// the complete, valid JSON text of a permissions boundary policy. The maximum
	// length of the policy document that you can pass in this operation, including
	// whitespace, is listed below. To view the maximum character counts of a managed
	// policy with no whitespaces, see IAM and STS character quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length)
	// . The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	PermissionsBoundaryPolicyInputList []string

	// A list of ARNs of Amazon Web Services resources to include in the simulation.
	// If this parameter is not provided, then the value defaults to * (all
	// resources). Each API in the ActionNames parameter is evaluated for each
	// resource in this list. The simulation determines the access result (allowed or
	// denied) of each combination and reports it in the response. You can simulate
	// resources that don't exist in your account. The simulation does not
	// automatically retrieve policies for the specified resources. If you want to
	// include a resource policy in the simulation, then you must include the policy as
	// a string in the ResourcePolicy parameter. If you include a ResourcePolicy , then
	// it must be applicable to all of the resources included in the simulation or you
	// receive an invalid input error. For more information about ARNs, see Amazon
	// Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference. Simulation of resource-based
	// policies isn't supported for IAM roles.
	ResourceArns []string

	// Specifies the type of simulation to run. Different API operations that support
	// resource-based policies require different combinations of resources. By
	// specifying the type of simulation to run, you enable the policy simulator to
	// enforce the presence of the required resources to ensure reliable simulation
	// results. If your simulation does not match one of the following scenarios, then
	// you can omit this parameter. The following list shows each of the supported
	// scenario values and the resources that you must define to run the simulation.
	// Each of the EC2 scenarios requires that you specify instance, image, and
	// security group resources. If your scenario includes an EBS volume, then you must
	// specify that volume as a resource. If the EC2 scenario includes VPC, then you
	// must supply the network interface resource. If it includes an IP subnet, then
	// you must specify the subnet resource. For more information on the EC2 scenario
	// options, see Supported platforms (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html)
	// in the Amazon EC2 User Guide.
	//   - EC2-VPC-InstanceStore instance, image, security group, network interface
	//   - EC2-VPC-InstanceStore-Subnet instance, image, security group, network
	//   interface, subnet
	//   - EC2-VPC-EBS instance, image, security group, network interface, volume
	//   - EC2-VPC-EBS-Subnet instance, image, security group, network interface,
	//   subnet, volume
	ResourceHandlingOption *string

	// An ARN representing the Amazon Web Services account ID that specifies the owner
	// of any simulated resource that does not identify its owner in the resource ARN.
	// Examples of resource ARNs include an S3 bucket or object. If ResourceOwner is
	// specified, it is also used as the account owner of any ResourcePolicy included
	// in the simulation. If the ResourceOwner parameter is not specified, then the
	// owner of the resources and the resource policy defaults to the account of the
	// identity provided in CallerArn . This parameter is required only if you specify
	// a resource-based policy and account that owns the resource is different from the
	// account that owns the simulated calling user CallerArn . The ARN for an account
	// uses the following syntax: arn:aws:iam::AWS-account-ID:root . For example, to
	// represent the account with the 112233445566 ID, use the following ARN:
	// arn:aws:iam::112233445566-ID:root .
	ResourceOwner *string

	// A resource-based policy to include in the simulation provided as a string. Each
	// resource in the simulation is treated as if it had this policy attached. You can
	// include only one resource-based policy in a simulation. The maximum length of
	// the policy document that you can pass in this operation, including whitespace,
	// is listed below. To view the maximum character counts of a managed policy with
	// no whitespaces, see IAM and STS character quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length)
	// . The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	// Simulation of resource-based policies isn't supported for IAM roles.
	ResourcePolicy *string

	noSmithyDocumentSerde
}

// Contains the response to a successful SimulatePrincipalPolicy or
// SimulateCustomPolicy request.
type SimulateCustomPolicyOutput struct {

	// The results of the simulation.
	EvaluationResults []types.EvaluationResult

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSimulateCustomPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSimulateCustomPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSimulateCustomPolicy{}, middleware.After)
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
	if err = addOpSimulateCustomPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSimulateCustomPolicy(options.Region), middleware.Before); err != nil {
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

// SimulateCustomPolicyAPIClient is a client that implements the
// SimulateCustomPolicy operation.
type SimulateCustomPolicyAPIClient interface {
	SimulateCustomPolicy(context.Context, *SimulateCustomPolicyInput, ...func(*Options)) (*SimulateCustomPolicyOutput, error)
}

var _ SimulateCustomPolicyAPIClient = (*Client)(nil)

// SimulateCustomPolicyPaginatorOptions is the paginator options for
// SimulateCustomPolicy
type SimulateCustomPolicyPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SimulateCustomPolicyPaginator is a paginator for SimulateCustomPolicy
type SimulateCustomPolicyPaginator struct {
	options   SimulateCustomPolicyPaginatorOptions
	client    SimulateCustomPolicyAPIClient
	params    *SimulateCustomPolicyInput
	nextToken *string
	firstPage bool
}

// NewSimulateCustomPolicyPaginator returns a new SimulateCustomPolicyPaginator
func NewSimulateCustomPolicyPaginator(client SimulateCustomPolicyAPIClient, params *SimulateCustomPolicyInput, optFns ...func(*SimulateCustomPolicyPaginatorOptions)) *SimulateCustomPolicyPaginator {
	if params == nil {
		params = &SimulateCustomPolicyInput{}
	}

	options := SimulateCustomPolicyPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SimulateCustomPolicyPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SimulateCustomPolicyPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SimulateCustomPolicy page.
func (p *SimulateCustomPolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SimulateCustomPolicyOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.SimulateCustomPolicy(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opSimulateCustomPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "SimulateCustomPolicy",
	}
}
