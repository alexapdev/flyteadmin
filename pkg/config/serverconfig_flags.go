// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (ServerConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (ServerConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in ServerConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg ServerConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("ServerConfig", pflag.ExitOnError)
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "httpPort"), defaultServerConfig.HTTPPort, "On which http port to serve admin")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "grpcPort"), defaultServerConfig.GrpcPort, "On which grpc port to serve admin")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "grpcServerReflection"), defaultServerConfig.GrpcServerReflection, "Enable GRPC Server Reflection")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "kube-config"), defaultServerConfig.KubeConfig, "Path to kubernetes client config file.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "master"), defaultServerConfig.Master, "The address of the Kubernetes API server.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.secure"), defaultServerConfig.Security.Secure, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.ssl.certificateFile"), defaultServerConfig.Security.Ssl.CertificateFile, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.ssl.keyFile"), defaultServerConfig.Security.Ssl.KeyFile, "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.useAuth"), defaultServerConfig.Security.UseAuth, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.httpAuthorizationHeader"), defaultServerConfig.Security.Auth.HTTPAuthorizationHeader, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.grpcAuthorizationHeader"), defaultServerConfig.Security.Auth.GrpcAuthorizationHeader, "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.auth.disableForHttp"), defaultServerConfig.Security.Auth.DisableForHTTP, "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.auth.disableForGrpc"), defaultServerConfig.Security.Auth.DisableForGrpc, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.httpPublicUri"), defaultServerConfig.Security.Auth.HTTPPublicUri.String(), "The publicly accessible http endpoint. This is used to build absolute URLs for endpoints that are only exposed over http (e.g. /authorize and /token for OAuth2).")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.redirectUrl"), defaultServerConfig.Security.Auth.UserAuth.RedirectURL.String(), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.openId.clientId"), defaultServerConfig.Security.Auth.UserAuth.OpenID.ClientID, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.openId.clientSecretName"), defaultServerConfig.Security.Auth.UserAuth.OpenID.ClientSecretName, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.openId.clientSecretFile"), defaultServerConfig.Security.Auth.UserAuth.OpenID.DeprecatedClientSecretFile, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.openId.baseUrl"), defaultServerConfig.Security.Auth.UserAuth.OpenID.BaseURL.String(), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.openId.callbackUrl"), defaultServerConfig.Security.Auth.UserAuth.OpenID.CallbackURL.String(), "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.openId.scopes"), []string{}, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.cookie_hash_key_secret_name"), defaultServerConfig.Security.Auth.UserAuth.CookieHashKeySecretName, "OPTIONAL: Secret name to use for cookie hash key.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.userAuth.cookie_block_key_secret_name"), defaultServerConfig.Security.Auth.UserAuth.CookieBlockKeySecretName, "OPTIONAL: Secret name to use for cookie block key.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.selfAuthServer.issuer"), defaultServerConfig.Security.Auth.AppAuth.SelfAuthServer.Issuer, "Defines the issuer to use when issuing and validating tokens. The default value is https://<requestUri.HostAndPort>/")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.selfAuthServer.accessTokenLifespan"), defaultServerConfig.Security.Auth.AppAuth.SelfAuthServer.AccessTokenLifespan.String(), "Defines the lifespan of issued access tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.selfAuthServer.refreshTokenLifespan"), defaultServerConfig.Security.Auth.AppAuth.SelfAuthServer.RefreshTokenLifespan.String(), "Defines the lifespan of issued access tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.selfAuthServer.authorizationCodeLifespan"), defaultServerConfig.Security.Auth.AppAuth.SelfAuthServer.AuthorizationCodeLifespan.String(), "Defines the lifespan of issued access tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.selfAuthServer.claimSymmetricEncryptionKeySecretName"), defaultServerConfig.Security.Auth.AppAuth.SelfAuthServer.ClaimSymmetricEncryptionKeySecretName, "OPTIONAL")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.selfAuthServer.tokenSigningRSAKeySecretName"), defaultServerConfig.Security.Auth.AppAuth.SelfAuthServer.TokenSigningRSAKeySecretName, "OPTIONAL: Secret name to use to retrieve RSA Signing Key.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.selfAuthServer.oldTokenSigningRSAKeySecretName"), defaultServerConfig.Security.Auth.AppAuth.SelfAuthServer.OldTokenSigningRSAKeySecretName, "OPTIONAL: Secret name to use to retrieve Old RSA Signing Key. This can be useful during key rotation to continue to accept older tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.externalAuthServer.baseUrl"), defaultServerConfig.Security.Auth.AppAuth.ExternalAuthServer.BaseURL.String(), "This should be the base url of the authorization server that you are trying to hit. With Okta for instance,  it will look something like https://company.okta.com/oauth2/abcdef123456789/")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.thirdPartyConfig.flyteClient.clientId"), defaultServerConfig.Security.Auth.AppAuth.ThirdParty.FlyteClientConfig.ClientID, "public identifier for the app which handles authorization for a Flyte deployment")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.thirdPartyConfig.flyteClient.redirectUri"), defaultServerConfig.Security.Auth.AppAuth.ThirdParty.FlyteClientConfig.RedirectURI, "This is the callback uri registered with the app which handles authorization for a Flyte deployment")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "security.auth.appAuth.thirdPartyConfig.flyteClient.scopes"), []string{}, "Recommended scopes for the client to request.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.auditAccess"), defaultServerConfig.Security.AuditAccess, "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.allowCors"), defaultServerConfig.Security.AllowCors, "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "security.allowedOrigins"), []string{}, "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "security.allowedHeaders"), []string{}, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "thirdPartyConfig.flyteClient.clientId"), defaultServerConfig.DeprecatedThirdPartyConfig.FlyteClientConfig.ClientID, "public identifier for the app which handles authorization for a Flyte deployment")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "thirdPartyConfig.flyteClient.redirectUri"), defaultServerConfig.DeprecatedThirdPartyConfig.FlyteClientConfig.RedirectURI, "This is the callback uri registered with the app which handles authorization for a Flyte deployment")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "thirdPartyConfig.flyteClient.scopes"), []string{}, "Recommended scopes for the client to request.")
	return cmdFlags
}
