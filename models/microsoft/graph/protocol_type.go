package graph
import (
    "strings"
    "errors"
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
)

func (i ProtocolType) String() string {
    return []string{"NONE", "OAUTH2", "ROPC", "WSFEDERATION", "SAML20", "DEVICECODE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseProtocolType(v string) (interface{}, error) {
    result := NONE_PROTOCOLTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_PROTOCOLTYPE
        case "OAUTH2":
            result = OAUTH2_PROTOCOLTYPE
        case "ROPC":
            result = ROPC_PROTOCOLTYPE
        case "WSFEDERATION":
            result = WSFEDERATION_PROTOCOLTYPE
        case "SAML20":
            result = SAML20_PROTOCOLTYPE
        case "DEVICECODE":
            result = DEVICECODE_PROTOCOLTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROTOCOLTYPE
        default:
            return 0, errors.New("Unknown ProtocolType value: " + v)
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
