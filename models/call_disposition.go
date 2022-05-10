package models
import (
    "errors"
)
// Provides operations to call the redirect method.
type CallDisposition int

const (
    DEFAULT_ESCAPED_CALLDISPOSITION CallDisposition = iota
    SIMULTANEOUSRING_CALLDISPOSITION
    FORWARD_CALLDISPOSITION
)

func (i CallDisposition) String() string {
    return []string{"default", "simultaneousRing", "forward"}[i]
}
func ParseCallDisposition(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_CALLDISPOSITION
    switch v {
        case "default":
            result = DEFAULT_ESCAPED_CALLDISPOSITION
        case "simultaneousRing":
            result = SIMULTANEOUSRING_CALLDISPOSITION
        case "forward":
            result = FORWARD_CALLDISPOSITION
        default:
            return 0, errors.New("Unknown CallDisposition value: " + v)
    }
    return &result, nil
}
func SerializeCallDisposition(values []CallDisposition) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
