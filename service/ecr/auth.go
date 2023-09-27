// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

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

type operationNamer interface {
	operationName() string
}

func bindAuthResolverParams(input interface{}, options Options) *AuthResolverParameters {
	params := &AuthResolverParameters{
		Operation: input.(operationNamer).operationName(),
	}

	bindAuthParamsRegion(params, input, options)

	return params
}
