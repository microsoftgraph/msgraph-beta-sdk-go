package security
import (
    "errors"
)
// 
type ScopeType int

const (
    DEVICEGROUP_SCOPETYPE ScopeType = iota
    UNKNOWNFUTUREVALUE_SCOPETYPE
)

func (i ScopeType) String() string {
    return []string{"deviceGroup", "unknownFutureValue"}[i]
}
func ParseScopeType(v string) (any, error) {
    result := DEVICEGROUP_SCOPETYPE
    switch v {
        case "deviceGroup":
            result = DEVICEGROUP_SCOPETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SCOPETYPE
        default:
            return 0, errors.New("Unknown ScopeType value: " + v)
    }
    return &result, nil
}
func SerializeScopeType(values []ScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ScopeType) isMultiValue() bool {
    return false
}
