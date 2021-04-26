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
func (Config) elemValueOrNil(v interface{}) interface{} {
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

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "httpAuthorizationHeader"), DefaultConfig.HTTPAuthorizationHeader, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "grpcAuthorizationHeader"), DefaultConfig.GrpcAuthorizationHeader, "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "disableForHttp"), DefaultConfig.DisableForHTTP, "Disables auth enforcement on HTTP Endpoints.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "disableForGrpc"), DefaultConfig.DisableForGrpc, "Disables auth enforcement on Grpc Endpoints.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "httpPublicUri"), DefaultConfig.HTTPPublicURI.String(), "Optional: The publicly accessible http endpoint. This is used to build absolute URLs for endpoints that are only exposed over http (e.g. /authorize and /token for OAuth2). If not provided,  the url will be deduced based on the request url.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "secure"), DefaultConfig.Secure, "Optional: Sets whether the system is serving over SSL/TLS.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "userAuth.redirectUrl"), DefaultConfig.UserAuth.RedirectURL.String(), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "userAuth.openId.clientId"), DefaultConfig.UserAuth.OpenID.ClientID, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "userAuth.openId.clientSecretName"), DefaultConfig.UserAuth.OpenID.ClientSecretName, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "userAuth.openId.clientSecretFile"), DefaultConfig.UserAuth.OpenID.DeprecatedClientSecretFile, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "userAuth.openId.baseUrl"), DefaultConfig.UserAuth.OpenID.BaseURL.String(), "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "userAuth.openId.scopes"), []string{}, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "userAuth.cookieHashKeySecretName"), DefaultConfig.UserAuth.CookieHashKeySecretName, "OPTIONAL: Secret name to use for cookie hash key.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "userAuth.cookieBlockKeySecretName"), DefaultConfig.UserAuth.CookieBlockKeySecretName, "OPTIONAL: Secret name to use for cookie block key.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.selfAuthServer.issuer"), DefaultConfig.AppAuth.SelfAuthServer.Issuer, "Defines the issuer to use when issuing and validating tokens. The default value is https://<requestUri.HostAndPort>/")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.selfAuthServer.accessTokenLifespan"), DefaultConfig.AppAuth.SelfAuthServer.AccessTokenLifespan.String(), "Defines the lifespan of issued access tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.selfAuthServer.refreshTokenLifespan"), DefaultConfig.AppAuth.SelfAuthServer.RefreshTokenLifespan.String(), "Defines the lifespan of issued access tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.selfAuthServer.authorizationCodeLifespan"), DefaultConfig.AppAuth.SelfAuthServer.AuthorizationCodeLifespan.String(), "Defines the lifespan of issued access tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.selfAuthServer.claimSymmetricEncryptionKeySecretName"), DefaultConfig.AppAuth.SelfAuthServer.ClaimSymmetricEncryptionKeySecretName, "OPTIONAL: Secret name to use to encrypt claims in authcode token.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.selfAuthServer.tokenSigningRSAKeySecretName"), DefaultConfig.AppAuth.SelfAuthServer.TokenSigningRSAKeySecretName, "OPTIONAL: Secret name to use to retrieve RSA Signing Key.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.selfAuthServer.oldTokenSigningRSAKeySecretName"), DefaultConfig.AppAuth.SelfAuthServer.OldTokenSigningRSAKeySecretName, "OPTIONAL: Secret name to use to retrieve Old RSA Signing Key. This can be useful during key rotation to continue to accept older tokens.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.externalAuthServer.baseUrl"), DefaultConfig.AppAuth.ExternalAuthServer.BaseURL.String(), "This should be the base url of the authorization server that you are trying to hit. With Okta for instance,  it will look something like https://company.okta.com/oauth2/abcdef123456789/")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.thirdPartyConfig.flyteClient.clientId"), DefaultConfig.AppAuth.ThirdParty.FlyteClientConfig.ClientID, "public identifier for the app which handles authorization for a Flyte deployment")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "appAuth.thirdPartyConfig.flyteClient.redirectUri"), DefaultConfig.AppAuth.ThirdParty.FlyteClientConfig.RedirectURI, "This is the callback uri registered with the app which handles authorization for a Flyte deployment")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "appAuth.thirdPartyConfig.flyteClient.scopes"), []string{}, "Recommended scopes for the client to request.")
	return cmdFlags
}
