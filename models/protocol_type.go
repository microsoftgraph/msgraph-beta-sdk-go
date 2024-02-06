package models
import (
    "errors"
    "math"
    "strings"
)
// 
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
)

func (i ProtocolType) String() string {
    var values []string
    options := []string{"none", "oAuth2", "ropc", "wsFederation", "saml20", "deviceCode", "unknownFutureValue", "authenticationTransfer"}
    for p := 0; p < 8; p++ {
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
