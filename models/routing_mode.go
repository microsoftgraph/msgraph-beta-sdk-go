package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the commsApplication singleton.
type RoutingMode int

const (
    ONETOONE_ROUTINGMODE RoutingMode = iota
    MULTICAST_ROUTINGMODE
)

func (i RoutingMode) String() string {
    return []string{"ONETOONE", "MULTICAST"}[i]
}
func ParseRoutingMode(v string) (interface{}, error) {
    result := ONETOONE_ROUTINGMODE
    switch strings.ToUpper(v) {
        case "ONETOONE":
            result = ONETOONE_ROUTINGMODE
        case "MULTICAST":
            result = MULTICAST_ROUTINGMODE
        default:
            return 0, errors.New("Unknown RoutingMode value: " + v)
    }
    return &result, nil
}
func SerializeRoutingMode(values []RoutingMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}