package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_NDESCONNECTORSTATE, nil
        case "ACTIVE":
            return ACTIVE_NDESCONNECTORSTATE, nil
        case "INACTIVE":
            return INACTIVE_NDESCONNECTORSTATE, nil
    }
    return 0, errors.New("Unknown NdesConnectorState value: " + v)
}
func SerializeNdesConnectorState(values []NdesConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
