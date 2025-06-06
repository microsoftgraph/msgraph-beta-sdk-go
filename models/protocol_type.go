// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
import (
    "math"
    "strings"
)
type ProtocolType int

const (
    NONE_PROTOCOLTYPE = 1
    OAUTH2_PROTOCOLTYPE = 2
    ROPC_PROTOCOLTYPE = 4
    WSFEDERATION_PROTOCOLTYPE = 8
    SAML20_PROTOCOLTYPE = 16
    DEVICECODE_PROTOCOLTYPE = 32
    UNKNOWNFUTUREVALUE_PROTOCOLTYPE = 64
    AUTHENTICATIONTRANSFER_PROTOCOLTYPE = 128
    NATIVEAUTH_PROTOCOLTYPE = 256
    IMPLICITACCESSTOKENANDGETRESPONSEMODE_PROTOCOLTYPE = 512
    IMPLICITIDTOKENANDGETRESPONSEMODE_PROTOCOLTYPE = 1024
    IMPLICITACCESSTOKENANDPOSTRESPONSEMODE_PROTOCOLTYPE = 2048
    IMPLICITIDTOKENANDPOSTRESPONSEMODE_PROTOCOLTYPE = 4096
    AUTHORIZATIONCODEWITHOUTPKCE_PROTOCOLTYPE = 8192
    AUTHORIZATIONCODEWITHPKCE_PROTOCOLTYPE = 16384
    CLIENTCREDENTIALS_PROTOCOLTYPE = 32768
    REFRESHTOKENGRANT_PROTOCOLTYPE = 65536
    ENCRYPTEDAUTHORIZERESPONSE_PROTOCOLTYPE = 131072
    DIRECTUSERGRANT_PROTOCOLTYPE = 262144
    KERBEROS_PROTOCOLTYPE = 524288
    PRTGRANT_PROTOCOLTYPE = 1048576
    SEAMLESSSSO_PROTOCOLTYPE = 2097152
    PRTBROKERBASED_PROTOCOLTYPE = 4194304
    PRTNONBROKERBASED_PROTOCOLTYPE = 8388608
    ONBEHALFOF_PROTOCOLTYPE = 16777216
    SAMLONBEHALFOF_PROTOCOLTYPE = 33554432
)

func (i ProtocolType) String() string {
    var values []string
    options := []string{"none", "oAuth2", "ropc", "wsFederation", "saml20", "deviceCode", "unknownFutureValue", "authenticationTransfer", "nativeAuth", "implicitAccessTokenAndGetResponseMode", "implicitIdTokenAndGetResponseMode", "implicitAccessTokenAndPostResponseMode", "implicitIdTokenAndPostResponseMode", "authorizationCodeWithoutPkce", "authorizationCodeWithPkce", "clientCredentials", "refreshTokenGrant", "encryptedAuthorizeResponse", "directUserGrant", "kerberos", "prtGrant", "seamlessSso", "prtBrokerBased", "prtNonBrokerBased", "onBehalfOf", "samlOnBehalfOf"}
    for p := 0; p < 26; p++ {
        mantis := ProtocolType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseProtocolType(v string) (any, error) {
    var result ProtocolType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_PROTOCOLTYPE
            case "oAuth2":
                result |= OAUTH2_PROTOCOLTYPE
            case "ropc":
                result |= ROPC_PROTOCOLTYPE
            case "wsFederation":
                result |= WSFEDERATION_PROTOCOLTYPE
            case "saml20":
                result |= SAML20_PROTOCOLTYPE
            case "deviceCode":
                result |= DEVICECODE_PROTOCOLTYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_PROTOCOLTYPE
            case "authenticationTransfer":
                result |= AUTHENTICATIONTRANSFER_PROTOCOLTYPE
            case "nativeAuth":
                result |= NATIVEAUTH_PROTOCOLTYPE
            case "implicitAccessTokenAndGetResponseMode":
                result |= IMPLICITACCESSTOKENANDGETRESPONSEMODE_PROTOCOLTYPE
            case "implicitIdTokenAndGetResponseMode":
                result |= IMPLICITIDTOKENANDGETRESPONSEMODE_PROTOCOLTYPE
            case "implicitAccessTokenAndPostResponseMode":
                result |= IMPLICITACCESSTOKENANDPOSTRESPONSEMODE_PROTOCOLTYPE
            case "implicitIdTokenAndPostResponseMode":
                result |= IMPLICITIDTOKENANDPOSTRESPONSEMODE_PROTOCOLTYPE
            case "authorizationCodeWithoutPkce":
                result |= AUTHORIZATIONCODEWITHOUTPKCE_PROTOCOLTYPE
            case "authorizationCodeWithPkce":
                result |= AUTHORIZATIONCODEWITHPKCE_PROTOCOLTYPE
            case "clientCredentials":
                result |= CLIENTCREDENTIALS_PROTOCOLTYPE
            case "refreshTokenGrant":
                result |= REFRESHTOKENGRANT_PROTOCOLTYPE
            case "encryptedAuthorizeResponse":
                result |= ENCRYPTEDAUTHORIZERESPONSE_PROTOCOLTYPE
            case "directUserGrant":
                result |= DIRECTUSERGRANT_PROTOCOLTYPE
            case "kerberos":
                result |= KERBEROS_PROTOCOLTYPE
            case "prtGrant":
                result |= PRTGRANT_PROTOCOLTYPE
            case "seamlessSso":
                result |= SEAMLESSSSO_PROTOCOLTYPE
            case "prtBrokerBased":
                result |= PRTBROKERBASED_PROTOCOLTYPE
            case "prtNonBrokerBased":
                result |= PRTNONBROKERBASED_PROTOCOLTYPE
            case "onBehalfOf":
                result |= ONBEHALFOF_PROTOCOLTYPE
            case "samlOnBehalfOf":
                result |= SAMLONBEHALFOF_PROTOCOLTYPE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeProtocolType(values []ProtocolType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProtocolType) isMultiValue() bool {
    return true
}
