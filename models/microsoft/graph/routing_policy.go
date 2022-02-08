package graph
import (
    "strings"
    "errors"
)
// 
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
    result := NONE_ROUTINGPOLICY
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ROUTINGPOLICY
        case "NOMISSEDCALL":
            result = NOMISSEDCALL_ROUTINGPOLICY
        case "DISABLEFORWARDINGEXCEPTPHONE":
            result = DISABLEFORWARDINGEXCEPTPHONE_ROUTINGPOLICY
        case "DISABLEFORWARDING":
            result = DISABLEFORWARDING_ROUTINGPOLICY
        case "PREFERSKYPEFORBUSINESS":
            result = PREFERSKYPEFORBUSINESS_ROUTINGPOLICY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ROUTINGPOLICY
        default:
            return 0, errors.New("Unknown RoutingPolicy value: " + v)
    }
    return &result, nil
}
func SerializeRoutingPolicy(values []RoutingPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
