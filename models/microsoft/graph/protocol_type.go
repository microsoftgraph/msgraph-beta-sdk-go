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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_PROTOCOLTYPE, nil
        case "OAUTH2":
            return OAUTH2_PROTOCOLTYPE, nil
        case "ROPC":
            return ROPC_PROTOCOLTYPE, nil
        case "WSFEDERATION":
            return WSFEDERATION_PROTOCOLTYPE, nil
        case "SAML20":
            return SAML20_PROTOCOLTYPE, nil
        case "DEVICECODE":
            return DEVICECODE_PROTOCOLTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROTOCOLTYPE, nil
    }
    return 0, errors.New("Unknown ProtocolType value: " + v)
}
func SerializeProtocolType(values []ProtocolType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
