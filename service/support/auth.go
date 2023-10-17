// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	smithyauth "github.com/aws/smithy-go/auth"
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

func bindAuthResolverParams(input interface{}, options Options) *AuthResolverParameters {
	params := &AuthResolverParameters{
		Operation: "",
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
			props.SigningName = "support"
			props.SigningRegion = params.Region
		}),
	}
}
