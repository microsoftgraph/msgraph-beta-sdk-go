package graph
import (
    "strings"
    "errors"
)
// 
type CallDisposition int

const (
    DEFAULT_CALLDISPOSITION CallDisposition = iota
    SIMULTANEOUSRING_CALLDISPOSITION
    FORWARD_CALLDISPOSITION
)

func (i CallDisposition) String() string {
    return []string{"DEFAULT", "SIMULTANEOUSRING", "FORWARD"}[i]
}
func ParseCallDisposition(v string) (interface{}, error) {
    result := DEFAULT_CALLDISPOSITION
    switch strings.ToUpper(v) {
        case "DEFAULT":
            result = DEFAULT_CALLDISPOSITION
        case "SIMULTANEOUSRING":
            result = SIMULTANEOUSRING_CALLDISPOSITION
        case "FORWARD":
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
