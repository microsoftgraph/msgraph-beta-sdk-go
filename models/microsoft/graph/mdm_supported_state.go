package graph
import (
    "strings"
    "errors"
)
// 
type MdmSupportedState int

const (
    UNKNOWN_MDMSUPPORTEDSTATE MdmSupportedState = iota
    SUPPORTED_MDMSUPPORTEDSTATE
    UNSUPPORTED_MDMSUPPORTEDSTATE
    DEPRECATED_MDMSUPPORTEDSTATE
)

func (i MdmSupportedState) String() string {
    return []string{"UNKNOWN", "SUPPORTED", "UNSUPPORTED", "DEPRECATED"}[i]
}
func ParseMdmSupportedState(v string) (interface{}, error) {
    result := UNKNOWN_MDMSUPPORTEDSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MDMSUPPORTEDSTATE
        case "SUPPORTED":
            result = SUPPORTED_MDMSUPPORTEDSTATE
        case "UNSUPPORTED":
            result = UNSUPPORTED_MDMSUPPORTEDSTATE
        case "DEPRECATED":
            result = DEPRECATED_MDMSUPPORTEDSTATE
        default:
            return 0, errors.New("Unknown MdmSupportedState value: " + v)
    }
    return &result, nil
}
func SerializeMdmSupportedState(values []MdmSupportedState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
