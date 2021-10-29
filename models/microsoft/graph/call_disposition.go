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
    switch strings.ToUpper(v) {
        case "DEFAULT":
            return DEFAULT_CALLDISPOSITION, nil
        case "SIMULTANEOUSRING":
            return SIMULTANEOUSRING_CALLDISPOSITION, nil
        case "FORWARD":
            return FORWARD_CALLDISPOSITION, nil
    }
    return 0, errors.New("Unknown CallDisposition value: " + v)
}
func SerializeCallDisposition(values []CallDisposition) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
