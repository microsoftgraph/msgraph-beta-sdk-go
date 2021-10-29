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
    switch strings.ToUpper(v) {
        case "INTRANET":
            return INTRANET_NETWORKTYPE, nil
        case "EXTRANET":
            return EXTRANET_NETWORKTYPE, nil
        case "NAMEDNETWORK":
            return NAMEDNETWORK_NETWORKTYPE, nil
        case "TRUSTED":
            return TRUSTED_NETWORKTYPE, nil
        case "TRUSTEDNAMEDLOCATION":
            return TRUSTEDNAMEDLOCATION_NETWORKTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_NETWORKTYPE, nil
    }
    return 0, errors.New("Unknown NetworkType value: " + v)
}
func SerializeNetworkType(values []NetworkType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
