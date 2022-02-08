package graph
import (
    "strings"
    "errors"
)
// 
type NetworkType int

const (
    INTRANET_NETWORKTYPE NetworkType = iota
    EXTRANET_NETWORKTYPE
    NAMEDNETWORK_NETWORKTYPE
    TRUSTED_NETWORKTYPE
    TRUSTEDNAMEDLOCATION_NETWORKTYPE
    UNKNOWNFUTUREVALUE_NETWORKTYPE
)

func (i NetworkType) String() string {
    return []string{"INTRANET", "EXTRANET", "NAMEDNETWORK", "TRUSTED", "TRUSTEDNAMEDLOCATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseNetworkType(v string) (interface{}, error) {
    result := INTRANET_NETWORKTYPE
    switch strings.ToUpper(v) {
        case "INTRANET":
            result = INTRANET_NETWORKTYPE
        case "EXTRANET":
            result = EXTRANET_NETWORKTYPE
        case "NAMEDNETWORK":
            result = NAMEDNETWORK_NETWORKTYPE
        case "TRUSTED":
            result = TRUSTED_NETWORKTYPE
        case "TRUSTEDNAMEDLOCATION":
            result = TRUSTEDNAMEDLOCATION_NETWORKTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_NETWORKTYPE
        default:
            return 0, errors.New("Unknown NetworkType value: " + v)
    }
    return &result, nil
}
func SerializeNetworkType(values []NetworkType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
