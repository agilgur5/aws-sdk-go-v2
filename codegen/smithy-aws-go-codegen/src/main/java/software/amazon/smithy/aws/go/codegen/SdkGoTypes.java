/*
 * Copyright 2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package software.amazon.smithy.aws.go.codegen;

import software.amazon.smithy.aws.go.codegen.customization.AwsCustomGoDependency;
import software.amazon.smithy.codegen.core.Symbol;

/**
 * Collection of Symbol constants for types in the aws-sdk-go-v2 runtime.
 */
public final class SdkGoTypes {
    private SdkGoTypes() { }

    public static final class Aws {
        public static final Symbol String = AwsGoDependency.AWS_CORE.valueSymbol("String");
        public static final Symbol Bool = AwsGoDependency.AWS_CORE.valueSymbol("Bool");

        public static final Symbol FIPSEndpointStateEnabled = AwsGoDependency.AWS_CORE.valueSymbol("FIPSEndpointStateEnabled");
        public static final Symbol DualStackEndpointStateEnabled = AwsGoDependency.AWS_CORE.valueSymbol("DualStackEndpointStateEnabled");
    }

    public static final class Internal {
        public static final class Auth {
            public static final Symbol HTTPAuthScheme = AwsGoDependency.INTERNAL_AUTH.pointableSymbol("HTTPAuthScheme");
            public static final Symbol NewHTTPAuthScheme = AwsGoDependency.INTERNAL_AUTH.valueSymbol("NewHTTPAuthScheme");

            public static final class Smithy {
                public static final Symbol CredentialsAdapter = AwsGoDependency.INTERNAL_AUTH_SMITHY.pointableSymbol("CredentialsAdapter");
                public static final Symbol CredentialsProviderAdapter = AwsGoDependency.INTERNAL_AUTH_SMITHY.pointableSymbol("CredentialsProviderAdapter");
                public static final Symbol V4SignerAdapter = AwsGoDependency.INTERNAL_AUTH_SMITHY.pointableSymbol("V4SignerAdapter");
                public static final Symbol BearerTokenAdapter = AwsGoDependency.INTERNAL_AUTH_SMITHY.pointableSymbol("BearerTokenAdapter");
                public static final Symbol BearerTokenProviderAdapter = AwsGoDependency.INTERNAL_AUTH_SMITHY.pointableSymbol("BearerTokenProviderAdapter");
                public static final Symbol BearerTokenSignerAdapter = AwsGoDependency.INTERNAL_AUTH_SMITHY.pointableSymbol("BearerTokenSignerAdapter");
            }
        }
        public static final class Endpoints {
            public static final Symbol MapFIPSRegion = AwsGoDependency.INTERNAL_ENDPOINTS.valueSymbol("MapFIPSRegion");
        }
    }

    public static final class ServiceCustomizations {
        public static final class S3Control {
            public static final Symbol AddDisableHostPrefixMiddleware = AwsCustomGoDependency.S3CONTROL_CUSTOMIZATION.valueSymbol("AddDisableHostPrefixMiddleware");
        }
    }
}
