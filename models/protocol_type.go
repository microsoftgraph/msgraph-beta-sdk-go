package models
import (
    "errors"
    "strings"
)
// 
type ProtocolType int

const (
    NONE_PROTOCOLTYPE ProtocolType = iota
    OAUTH2_PROTOCOLTYPE
    ROPC_PROTOCOLTYPE
    WSFEDERATION_PROTOCOLTYPE
    SAML20_PROTOCOLTYPE
    DEVICECODE_PROTOCOLTYPE
    UNKNOWNFUTUREVALUE_PROTOCOLTYPE
    AUTHENTICATIONTRANSFER_PROTOCOLTYPE
)

func (i ProtocolType) String() string {
    var values []string
    for p := ProtocolType(1); p <= AUTHENTICATIONTRANSFER_PROTOCOLTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "oAuth2", "ropc", "wsFederation", "saml20", "deviceCode", "unknownFutureValue", "authenticationTransfer"}[p])
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
            default:
                return 0, errors.New("Unknown ProtocolType value: " + v)
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
