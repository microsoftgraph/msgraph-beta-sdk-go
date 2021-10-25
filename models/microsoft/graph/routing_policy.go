package graph
import (
    "strings"
    "errors"
)
type RoutingPolicy int

const (
    NONE_ROUTINGPOLICY RoutingPolicy = iota
    NOMISSEDCALL_ROUTINGPOLICY
    DISABLEFORWARDINGEXCEPTPHONE_ROUTINGPOLICY
    DISABLEFORWARDING_ROUTINGPOLICY
    PREFERSKYPEFORBUSINESS_ROUTINGPOLICY
    UNKNOWNFUTUREVALUE_ROUTINGPOLICY
)

func (i RoutingPolicy) String() string {
    return []string{"NONE", "NOMISSEDCALL", "DISABLEFORWARDINGEXCEPTPHONE", "DISABLEFORWARDING", "PREFERSKYPEFORBUSINESS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRoutingPolicy(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ROUTINGPOLICY, nil
        case "NOMISSEDCALL":
            return NOMISSEDCALL_ROUTINGPOLICY, nil
        case "DISABLEFORWARDINGEXCEPTPHONE":
            return DISABLEFORWARDINGEXCEPTPHONE_ROUTINGPOLICY, nil
        case "DISABLEFORWARDING":
            return DISABLEFORWARDING_ROUTINGPOLICY, nil
        case "PREFERSKYPEFORBUSINESS":
            return PREFERSKYPEFORBUSINESS_ROUTINGPOLICY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ROUTINGPOLICY, nil
    }
    return 0, errors.New("Unknown RoutingPolicy value: " + v)
}
func SerializeRoutingPolicy(values []RoutingPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
