// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	smithyauth "github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func bindAuthParamsRegion(params *AuthResolverParameters, _ interface{}, options Options) {
	params.Region = options.Region
}

// AuthResolverParameters contains the set of inputs necessary for auth scheme
// resolution.
type AuthResolverParameters struct {
	// The name of the operation being invoked.
	Operation string

	// The region in which the operation is being invoked.
	Region string
}

func bindAuthResolverParams(operation string, input interface{}, options Options) *AuthResolverParameters {
	params := &AuthResolverParameters{
		Operation: operation,
	}

	bindAuthParamsRegion(params, input, options)

	return params
}

// AuthSchemeResolver returns a set of possible authentication options for an
// operation.
type AuthSchemeResolver interface {
	ResolveAuthSchemes(context.Context, *AuthResolverParameters) ([]*smithyauth.Option, error)
}

type defaultAuthSchemeResolver struct{}

var _ AuthSchemeResolver = (*defaultAuthSchemeResolver)(nil)

func (*defaultAuthSchemeResolver) ResolveAuthSchemes(ctx context.Context, params *AuthResolverParameters) ([]*smithyauth.Option, error) {
	if overrides, ok := operationAuthOptions[params.Operation]; ok {
		return overrides(params), nil
	}
	return serviceAuthOptions(params), nil
}

var operationAuthOptions = map[string]func(*AuthResolverParameters) []*smithyauth.Option{}

func serviceAuthOptions(params *AuthResolverParameters) []*smithyauth.Option {
	return []*smithyauth.Option{
		smithyhttp.NewSigV4Option(func(props *smithyhttp.SigV4Properties) {
			props.SigningName = "mq"
			props.SigningRegion = params.Region
		}),
	}
}

type resolveAuthSchemeMiddleware struct {
	operation string
	options   Options
}

func (*resolveAuthSchemeMiddleware) ID() string {
	return "ResolveAuthScheme"
}

func (m *resolveAuthSchemeMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	params := bindAuthResolverParams(m.operation, in.Parameters, m.options)
	options, err := m.options.AuthSchemeResolver.ResolveAuthSchemes(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("resolve auth scheme: %v", err)
	}

	scheme, ok := m.selectScheme(options)
	if !ok {
		return out, metadata, fmt.Errorf("could not select an auth scheme")
	}

	ctx = setResolvedAuthScheme(ctx, scheme)
	return next.HandleSerialize(ctx, in)
}

func (m *resolveAuthSchemeMiddleware) selectScheme(options []*smithyauth.Option) (*resolvedAuthScheme, bool) {
	for _, option := range options {
		if option.SchemeID == smithyhttp.SchemeIDAnonymous {
			return newResolvedAuthScheme(smithyhttp.NewAnonymousScheme(), option), true
		}

		for _, scheme := range m.options.AuthSchemes {
			if scheme.SchemeID() != option.SchemeID {
				continue
			}

			if scheme.IdentityResolver(m.options) != nil {
				return newResolvedAuthScheme(scheme, option), true
			}
		}
	}

	return nil, false
}

type resolvedAuthSchemeKey struct{}

type resolvedAuthScheme struct {
	Scheme             smithyhttp.AuthScheme
	IdentityProperties smithy.Properties
	SignerProperties   smithy.Properties
}

func newResolvedAuthScheme(scheme smithyhttp.AuthScheme, option *smithyauth.Option) *resolvedAuthScheme {
	return &resolvedAuthScheme{
		Scheme:             scheme,
		IdentityProperties: option.IdentityProperties,
		SignerProperties:   option.SignerProperties,
	}
}

func setResolvedAuthScheme(ctx context.Context, scheme *resolvedAuthScheme) context.Context {
	return middleware.WithStackValue(ctx, resolvedAuthSchemeKey{}, scheme)
}

func getResolvedAuthScheme(ctx context.Context) *resolvedAuthScheme {
	v, _ := middleware.GetStackValue(ctx, resolvedAuthSchemeKey{}).(*resolvedAuthScheme)
	return v
}

type getIdentityMiddleware struct {
	options Options
}

func (*getIdentityMiddleware) ID() string {
	return "GetIdentity"
}

func (m *getIdentityMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	rscheme := getResolvedAuthScheme(ctx)
	if rscheme == nil {
		return out, metadata, fmt.Errorf("no resolved auth scheme")
	}

	resolver := rscheme.Scheme.IdentityResolver(m.options)
	if resolver == nil {
		return out, metadata, fmt.Errorf("no identity resolver")
	}

	identity, err := resolver.GetIdentity(ctx, rscheme.IdentityProperties)
	if err != nil {
		return out, metadata, fmt.Errorf("get identity: %v", err)
	}

	ctx = setIdentity(ctx, identity)
	return next.HandleFinalize(ctx, in)
}

type identityKey struct{}

func setIdentity(ctx context.Context, identity smithyauth.Identity) context.Context {
	return middleware.WithStackValue(ctx, identityKey{}, identity)
}

func getIdentity(ctx context.Context) smithyauth.Identity {
	v, _ := middleware.GetStackValue(ctx, identityKey{}).(smithyauth.Identity)
	return v
}

type signRequestMiddleware struct {
}

func (*signRequestMiddleware) ID() string {
	return "Signing"
}

func (m *signRequestMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	rscheme := getResolvedAuthScheme(ctx)
	if rscheme == nil {
		return out, metadata, fmt.Errorf("no resolved auth scheme")
	}

	identity := getIdentity(ctx)
	if identity == nil {
		return out, metadata, fmt.Errorf("no identity")
	}

	signer := rscheme.Scheme.Signer()
	if signer == nil {
		return out, metadata, fmt.Errorf("no signer")
	}

	if err := signer.SignRequest(ctx, req, identity, rscheme.SignerProperties); err != nil {
		return out, metadata, fmt.Errorf("sign request: %v", err)
	}

	return next.HandleFinalize(ctx, in)
}
