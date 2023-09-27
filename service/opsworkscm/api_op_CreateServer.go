// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and immedately starts a new server. The server is ready to use when it
// is in the HEALTHY state. By default, you can create a maximum of 10 servers.
// This operation is asynchronous. A LimitExceededException is thrown when you
// have created the maximum number of servers (10). A
// ResourceAlreadyExistsException is thrown when a server with the same name
// already exists in the account. A ResourceNotFoundException is thrown when you
// specify a backup ID that is not valid or is for a backup that does not exist. A
// ValidationException is thrown when parameters of the request are not valid. If
// you do not specify a security group by adding the SecurityGroupIds parameter,
// AWS OpsWorks creates a new security group. Chef Automate: The default security
// group opens the Chef server to the world on TCP port 443. If a KeyName is
// present, AWS OpsWorks enables SSH access. SSH is also open to the world on TCP
// port 22. Puppet Enterprise: The default security group opens TCP ports 22, 443,
// 4433, 8140, 8142, 8143, and 8170. If a KeyName is present, AWS OpsWorks enables
// SSH access. SSH is also open to the world on TCP port 22. By default, your
// server is accessible from any IP address. We recommend that you update your
// security group rules to allow access from known IP addresses and address ranges
// only. To edit security group rules, open Security Groups in the navigation pane
// of the EC2 management console. To specify your own domain for a server, and
// provide your own self-signed or CA-signed certificate and private key, specify
// values for CustomDomain , CustomCertificate , and CustomPrivateKey .
func (c *Client) CreateServer(ctx context.Context, params *CreateServerInput, optFns ...func(*Options)) (*CreateServerOutput, error) {
	if params == nil {
		params = &CreateServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServer", params, optFns, c.addOperationCreateServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServerInput struct {

	// The configuration management engine to use. Valid values include ChefAutomate
	// and Puppet .
	//
	// This member is required.
	Engine *string

	// The ARN of the instance profile that your Amazon EC2 instances use. Although
	// the AWS OpsWorks console typically creates the instance profile for you, if you
	// are using API commands instead, run the service-role-creation.yaml AWS
	// CloudFormation template, located at
	// https://s3.amazonaws.com/opsworks-cm-us-east-1-prod-default-assets/misc/opsworks-cm-roles.yaml.
	// This template creates a CloudFormation stack that includes the instance profile
	// you need.
	//
	// This member is required.
	InstanceProfileArn *string

	// The Amazon EC2 instance type to use. For example, m5.large .
	//
	// This member is required.
	InstanceType *string

	// The name of the server. The server name must be unique within your AWS account,
	// within each region. Server names must start with a letter; then letters,
	// numbers, or hyphens (-) are allowed, up to a maximum of 40 characters.
	//
	// This member is required.
	ServerName *string

	// The service role that the AWS OpsWorks CM service backend uses to work with
	// your account. Although the AWS OpsWorks management console typically creates the
	// service role for you, if you are using the AWS CLI or API commands, run the
	// service-role-creation.yaml AWS CloudFormation template, located at
	// https://s3.amazonaws.com/opsworks-cm-us-east-1-prod-default-assets/misc/opsworks-cm-roles.yaml.
	// This template creates a CloudFormation stack that includes the service role and
	// instance profile that you need.
	//
	// This member is required.
	ServiceRoleArn *string

	// Associate a public IP address with a server that you are launching. Valid
	// values are true or false . The default value is true .
	AssociatePublicIpAddress *bool

	// If you specify this field, AWS OpsWorks CM creates the server by using the
	// backup represented by BackupId.
	BackupId *string

	// The number of automated backups that you want to keep. Whenever a new backup is
	// created, AWS OpsWorks CM deletes the oldest backups if this number is exceeded.
	// The default value is 1 .
	BackupRetentionCount *int32

	// A PEM-formatted HTTPS certificate. The value can be be a single, self-signed
	// certificate, or a certificate chain. If you specify a custom certificate, you
	// must also specify values for CustomDomain and CustomPrivateKey . The following
	// are requirements for the CustomCertificate value:
	//   - You can provide either a self-signed, custom certificate, or the full
	//   certificate chain.
	//   - The certificate must be a valid X509 certificate, or a certificate chain in
	//   PEM format.
	//   - The certificate must be valid at the time of upload. A certificate can't be
	//   used before its validity period begins (the certificate's NotBefore date), or
	//   after it expires (the certificate's NotAfter date).
	//   - The certificate’s common name or subject alternative names (SANs), if
	//   present, must match the value of CustomDomain .
	//   - The certificate must match the value of CustomPrivateKey .
	CustomCertificate *string

	// An optional public endpoint of a server, such as https://aws.my-company.com . To
	// access the server, create a CNAME DNS record in your preferred DNS service that
	// points the custom domain to the endpoint that is generated when the server is
	// created (the value of the CreateServer Endpoint attribute). You cannot access
	// the server by using the generated Endpoint value if the server is using a
	// custom domain. If you specify a custom domain, you must also specify values for
	// CustomCertificate and CustomPrivateKey .
	CustomDomain *string

	// A private key in PEM format for connecting to the server by using HTTPS. The
	// private key must not be encrypted; it cannot be protected by a password or
	// passphrase. If you specify a custom private key, you must also specify values
	// for CustomDomain and CustomCertificate .
	CustomPrivateKey *string

	// Enable or disable scheduled backups. Valid values are true or false . The
	// default value is true .
	DisableAutomatedBackup *bool

	// Optional engine attributes on a specified server. Attributes accepted in a Chef
	// createServer request:
	//   - CHEF_AUTOMATE_PIVOTAL_KEY : A base64-encoded RSA public key. The
	//   corresponding private key is required to access the Chef API. When no
	//   CHEF_AUTOMATE_PIVOTAL_KEY is set, a private key is generated and returned in the
	//   response.
	//   - CHEF_AUTOMATE_ADMIN_PASSWORD : The password for the administrative user in
	//   the Chef Automate web-based dashboard. The password length is a minimum of eight
	//   characters, and a maximum of 32. The password can contain letters, numbers, and
	//   special characters (!/@#$%^&+=_). The password must contain at least one lower
	//   case letter, one upper case letter, one number, and one special character. When
	//   no CHEF_AUTOMATE_ADMIN_PASSWORD is set, one is generated and returned in the
	//   response.
	// Attributes accepted in a Puppet createServer request:
	//   - PUPPET_ADMIN_PASSWORD : To work with the Puppet Enterprise console, a
	//   password must use ASCII characters.
	//   - PUPPET_R10K_REMOTE : The r10k remote is the URL of your control repository
	//   (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying an
	//   r10k remote opens TCP port 8170.
	//   - PUPPET_R10K_PRIVATE_KEY : If you are using a private Git repository, add
	//   PUPPET_R10K_PRIVATE_KEY to specify a PEM-encoded private SSH key.
	EngineAttributes []types.EngineAttribute

	// The engine model of the server. Valid values in this release include Monolithic
	// for Puppet and Single for Chef.
	EngineModel *string

	// The major release version of the engine that you want to use. For a Chef
	// server, the valid value for EngineVersion is currently 2 . For a Puppet server,
	// valid values are 2019 or 2017 .
	EngineVersion *string

	// The Amazon EC2 key pair to set for the instance. This parameter is optional; if
	// desired, you may specify this parameter to connect to your instances by using
	// SSH.
	KeyPair *string

	// The start time for a one-hour period during which AWS OpsWorks CM backs up
	// application-level data on your server if automated backups are enabled. Valid
	// values must be specified in one of the following formats:
	//   - HH:MM for daily backups
	//   - DDD:HH:MM for weekly backups
	// MM must be specified as 00 . The specified time is in coordinated universal time
	// (UTC). The default value is a random, daily start time. Example: 08:00 , which
	// represents a daily start time of 08:00 UTC. Example: Mon:08:00 , which
	// represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredBackupWindow *string

	// The start time for a one-hour period each week during which AWS OpsWorks CM
	// performs maintenance on the instance. Valid values must be specified in the
	// following format: DDD:HH:MM . MM must be specified as 00 . The specified time is
	// in coordinated universal time (UTC). The default value is a random one-hour
	// period on Tuesday, Wednesday, or Friday. See TimeWindowDefinition for more
	// information. Example: Mon:08:00 , which represents a start time of every Monday
	// at 08:00 UTC. (8:00 a.m.)
	PreferredMaintenanceWindow *string

	// A list of security group IDs to attach to the Amazon EC2 instance. If you add
	// this parameter, the specified security groups must be within the VPC that is
	// specified by SubnetIds . If you do not specify this parameter, AWS OpsWorks CM
	// creates one new security group that uses TCP ports 22 and 443, open to 0.0.0.0/0
	// (everyone).
	SecurityGroupIds []string

	// The IDs of subnets in which to launch the server EC2 instance. Amazon
	// EC2-Classic customers: This field is required. All servers must run within a
	// VPC. The VPC must have "Auto Assign Public IP" enabled. EC2-VPC customers: This
	// field is optional. If you do not specify subnet IDs, your EC2 instances are
	// created in a default subnet that is selected by Amazon EC2. If you specify
	// subnet IDs, the VPC must have "Auto Assign Public IP" enabled. For more
	// information about supported Amazon EC2 platforms, see Supported Platforms (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html)
	// .
	SubnetIds []string

	// A map that contains tag keys and tag values to attach to an AWS OpsWorks for
	// Chef Automate or AWS OpsWorks for Puppet Enterprise server.
	//   - The key cannot be empty.
	//   - The key can be a maximum of 127 characters, and can contain only Unicode
	//   letters, numbers, or separators, or the following special characters: + - = .
	//   _ : / @
	//   - The value can be a maximum 255 characters, and contain only Unicode
	//   letters, numbers, or separators, or the following special characters: + - = .
	//   _ : / @
	//   - Leading and trailing white spaces are trimmed from both the key and value.
	//   - A maximum of 50 user-applied tags is allowed for any AWS OpsWorks-CM
	//   server.
	Tags []types.Tag

	noSmithyDocumentSerde
}

func (*CreateServerInput) operationName() string {
	return "CreateServer"
}

type CreateServerOutput struct {

	// The server that is created by the request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateServer{}, middleware.After)
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
	if err = addCreateServerResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "CreateServer",
	}
}

type opCreateServerResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateServerResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateServerResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "opsworks-cm"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "opsworks-cm"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("opsworks-cm")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateServerResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateServerResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
