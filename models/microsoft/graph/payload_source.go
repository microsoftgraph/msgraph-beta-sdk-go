package graph
import (
    "strings"
    "errors"
)
// 
type PayloadSource int

const (
    UNKNOWN_PAYLOADSOURCE PayloadSource = iota
    GLOBAL_PAYLOADSOURCE
    TENANT_PAYLOADSOURCE
    UNKNOWNFUTUREVALUE_PAYLOADSOURCE
)

func (i PayloadSource) String() string {
    return []string{"UNKNOWN", "GLOBAL", "TENANT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePayloadSource(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PAYLOADSOURCE, nil
        case "GLOBAL":
            return GLOBAL_PAYLOADSOURCE, nil
        case "TENANT":
            return TENANT_PAYLOADSOURCE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PAYLOADSOURCE, nil
    }
    return 0, errors.New("Unknown PayloadSource value: " + v)
}
func SerializePayloadSource(values []PayloadSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
