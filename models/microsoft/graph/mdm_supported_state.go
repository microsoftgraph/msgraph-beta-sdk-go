package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MDMSUPPORTEDSTATE, nil
        case "SUPPORTED":
            return SUPPORTED_MDMSUPPORTEDSTATE, nil
        case "UNSUPPORTED":
            return UNSUPPORTED_MDMSUPPORTEDSTATE, nil
        case "DEPRECATED":
            return DEPRECATED_MDMSUPPORTEDSTATE, nil
    }
    return 0, errors.New("Unknown MdmSupportedState value: " + v)
}
func SerializeMdmSupportedState(values []MdmSupportedState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
