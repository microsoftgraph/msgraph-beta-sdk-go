package graph
import (
    "strings"
    "errors"
)
// 
type NdesConnectorState int

const (
    NONE_NDESCONNECTORSTATE NdesConnectorState = iota
    ACTIVE_NDESCONNECTORSTATE
    INACTIVE_NDESCONNECTORSTATE
)

func (i NdesConnectorState) String() string {
    return []string{"NONE", "ACTIVE", "INACTIVE"}[i]
}
func ParseNdesConnectorState(v string) (interface{}, error) {
    result := NONE_NDESCONNECTORSTATE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_NDESCONNECTORSTATE
        case "ACTIVE":
            result = ACTIVE_NDESCONNECTORSTATE
        case "INACTIVE":
            result = INACTIVE_NDESCONNECTORSTATE
        default:
            return 0, errors.New("Unknown NdesConnectorState value: " + v)
    }
    return &result, nil
}
func SerializeNdesConnectorState(values []NdesConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
