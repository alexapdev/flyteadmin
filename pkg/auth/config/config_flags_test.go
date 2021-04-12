// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeSlice_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_httpAuthorizationHeader", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("httpAuthorizationHeader"); err == nil {
				assert.Equal(t, string(defaultConfig.HTTPAuthorizationHeader), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("httpAuthorizationHeader", testValue)
			if vString, err := cmdFlags.GetString("httpAuthorizationHeader"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.HTTPAuthorizationHeader)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_grpcAuthorizationHeader", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("grpcAuthorizationHeader"); err == nil {
				assert.Equal(t, string(defaultConfig.GrpcAuthorizationHeader), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("grpcAuthorizationHeader", testValue)
			if vString, err := cmdFlags.GetString("grpcAuthorizationHeader"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.GrpcAuthorizationHeader)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_disableForHttp", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("disableForHttp"); err == nil {
				assert.Equal(t, bool(defaultConfig.DisableForHTTP), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("disableForHttp", testValue)
			if vBool, err := cmdFlags.GetBool("disableForHttp"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.DisableForHTTP)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_disableForGrpc", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("disableForGrpc"); err == nil {
				assert.Equal(t, bool(defaultConfig.DisableForGrpc), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("disableForGrpc", testValue)
			if vBool, err := cmdFlags.GetBool("disableForGrpc"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.DisableForGrpc)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_httpPublicUri", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("httpPublicUri"); err == nil {
				assert.Equal(t, string(defaultConfig.HTTPPublicUri.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.HTTPPublicUri.String()

			cmdFlags.Set("httpPublicUri", testValue)
			if vString, err := cmdFlags.GetString("httpPublicUri"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.HTTPPublicUri)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.redirectUrl", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.redirectUrl"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.RedirectURL.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.UserAuth.RedirectURL.String()

			cmdFlags.Set("userAuth.redirectUrl", testValue)
			if vString, err := cmdFlags.GetString("userAuth.redirectUrl"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.RedirectURL)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.openId.clientId", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.openId.clientId"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.OpenID.ClientID), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("userAuth.openId.clientId", testValue)
			if vString, err := cmdFlags.GetString("userAuth.openId.clientId"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.OpenID.ClientID)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.openId.clientSecretName", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.openId.clientSecretName"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.OpenID.ClientSecretName), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("userAuth.openId.clientSecretName", testValue)
			if vString, err := cmdFlags.GetString("userAuth.openId.clientSecretName"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.OpenID.ClientSecretName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.openId.clientSecretFile", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.openId.clientSecretFile"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.OpenID.DeprecatedClientSecretFile), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("userAuth.openId.clientSecretFile", testValue)
			if vString, err := cmdFlags.GetString("userAuth.openId.clientSecretFile"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.OpenID.DeprecatedClientSecretFile)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.openId.baseUrl", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.openId.baseUrl"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.OpenID.BaseURL.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.UserAuth.OpenID.BaseURL.String()

			cmdFlags.Set("userAuth.openId.baseUrl", testValue)
			if vString, err := cmdFlags.GetString("userAuth.openId.baseUrl"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.OpenID.BaseURL)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.openId.callbackUrl", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.openId.callbackUrl"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.OpenID.CallbackURL.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.UserAuth.OpenID.CallbackURL.String()

			cmdFlags.Set("userAuth.openId.callbackUrl", testValue)
			if vString, err := cmdFlags.GetString("userAuth.openId.callbackUrl"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.OpenID.CallbackURL)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.openId.scopes", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vStringSlice, err := cmdFlags.GetStringSlice("userAuth.openId.scopes"); err == nil {
				assert.Equal(t, []string([]string{}), vStringSlice)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := join_Config("1,1", ",")

			cmdFlags.Set("userAuth.openId.scopes", testValue)
			if vStringSlice, err := cmdFlags.GetStringSlice("userAuth.openId.scopes"); err == nil {
				testDecodeSlice_Config(t, join_Config(vStringSlice, ","), &actual.UserAuth.OpenID.Scopes)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.cookie_hash_key_secret_name", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.cookie_hash_key_secret_name"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.CookieHashKeySecretName), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("userAuth.cookie_hash_key_secret_name", testValue)
			if vString, err := cmdFlags.GetString("userAuth.cookie_hash_key_secret_name"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.CookieHashKeySecretName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_userAuth.cookie_block_key_secret_name", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("userAuth.cookie_block_key_secret_name"); err == nil {
				assert.Equal(t, string(defaultConfig.UserAuth.CookieBlockKeySecretName), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("userAuth.cookie_block_key_secret_name", testValue)
			if vString, err := cmdFlags.GetString("userAuth.cookie_block_key_secret_name"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.UserAuth.CookieBlockKeySecretName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.selfAuthServer.issuer", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.issuer"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.SelfAuthServer.Issuer), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("appAuth.selfAuthServer.issuer", testValue)
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.issuer"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.SelfAuthServer.Issuer)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.selfAuthServer.accessTokenLifespan", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.accessTokenLifespan"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.SelfAuthServer.AccessTokenLifespan.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.AppAuth.SelfAuthServer.AccessTokenLifespan.String()

			cmdFlags.Set("appAuth.selfAuthServer.accessTokenLifespan", testValue)
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.accessTokenLifespan"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.SelfAuthServer.AccessTokenLifespan)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.selfAuthServer.refreshTokenLifespan", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.refreshTokenLifespan"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.SelfAuthServer.RefreshTokenLifespan.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.AppAuth.SelfAuthServer.RefreshTokenLifespan.String()

			cmdFlags.Set("appAuth.selfAuthServer.refreshTokenLifespan", testValue)
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.refreshTokenLifespan"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.SelfAuthServer.RefreshTokenLifespan)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.selfAuthServer.authorizationCodeLifespan", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.authorizationCodeLifespan"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.SelfAuthServer.AuthorizationCodeLifespan.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.AppAuth.SelfAuthServer.AuthorizationCodeLifespan.String()

			cmdFlags.Set("appAuth.selfAuthServer.authorizationCodeLifespan", testValue)
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.authorizationCodeLifespan"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.SelfAuthServer.AuthorizationCodeLifespan)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.selfAuthServer.claimSymmetricEncryptionKeySecretName", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.claimSymmetricEncryptionKeySecretName"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.SelfAuthServer.ClaimSymmetricEncryptionKeySecretName), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("appAuth.selfAuthServer.claimSymmetricEncryptionKeySecretName", testValue)
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.claimSymmetricEncryptionKeySecretName"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.SelfAuthServer.ClaimSymmetricEncryptionKeySecretName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.selfAuthServer.tokenSigningRSAKeySecretName", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.tokenSigningRSAKeySecretName"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.SelfAuthServer.TokenSigningRSAKeySecretName), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("appAuth.selfAuthServer.tokenSigningRSAKeySecretName", testValue)
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.tokenSigningRSAKeySecretName"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.SelfAuthServer.TokenSigningRSAKeySecretName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.selfAuthServer.oldTokenSigningRSAKeySecretName", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.oldTokenSigningRSAKeySecretName"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.SelfAuthServer.OldTokenSigningRSAKeySecretName), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("appAuth.selfAuthServer.oldTokenSigningRSAKeySecretName", testValue)
			if vString, err := cmdFlags.GetString("appAuth.selfAuthServer.oldTokenSigningRSAKeySecretName"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.SelfAuthServer.OldTokenSigningRSAKeySecretName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.externalAuthServer.baseUrl", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.externalAuthServer.baseUrl"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.ExternalAuthServer.BaseURL.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.AppAuth.ExternalAuthServer.BaseURL.String()

			cmdFlags.Set("appAuth.externalAuthServer.baseUrl", testValue)
			if vString, err := cmdFlags.GetString("appAuth.externalAuthServer.baseUrl"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.ExternalAuthServer.BaseURL)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.thirdPartyConfig.flyteClient.clientId", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.thirdPartyConfig.flyteClient.clientId"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.ThirdParty.FlyteClientConfig.ClientID), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("appAuth.thirdPartyConfig.flyteClient.clientId", testValue)
			if vString, err := cmdFlags.GetString("appAuth.thirdPartyConfig.flyteClient.clientId"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.ThirdParty.FlyteClientConfig.ClientID)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.thirdPartyConfig.flyteClient.redirectUri", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("appAuth.thirdPartyConfig.flyteClient.redirectUri"); err == nil {
				assert.Equal(t, string(defaultConfig.AppAuth.ThirdParty.FlyteClientConfig.RedirectURI), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("appAuth.thirdPartyConfig.flyteClient.redirectUri", testValue)
			if vString, err := cmdFlags.GetString("appAuth.thirdPartyConfig.flyteClient.redirectUri"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AppAuth.ThirdParty.FlyteClientConfig.RedirectURI)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_appAuth.thirdPartyConfig.flyteClient.scopes", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vStringSlice, err := cmdFlags.GetStringSlice("appAuth.thirdPartyConfig.flyteClient.scopes"); err == nil {
				assert.Equal(t, []string([]string{}), vStringSlice)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := join_Config("1,1", ",")

			cmdFlags.Set("appAuth.thirdPartyConfig.flyteClient.scopes", testValue)
			if vStringSlice, err := cmdFlags.GetStringSlice("appAuth.thirdPartyConfig.flyteClient.scopes"); err == nil {
				testDecodeSlice_Config(t, join_Config(vStringSlice, ","), &actual.AppAuth.ThirdParty.FlyteClientConfig.Scopes)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
