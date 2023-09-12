package networkaccess
import (
    "errors"
)
// 
type TrafficForwardingType int

const (
    M365_TRAFFICFORWARDINGTYPE TrafficForwardingType = iota
    INTERNET_TRAFFICFORWARDINGTYPE
    PRIVATE_TRAFFICFORWARDINGTYPE
    UNKNOWNFUTUREVALUE_TRAFFICFORWARDINGTYPE
)

func (i TrafficForwardingType) String() string {
    return []string{"m365", "internet", "private", "unknownFutureValue"}[i]
}
func ParseTrafficForwardingType(v string) (any, error) {
    result := M365_TRAFFICFORWARDINGTYPE
    switch v {
        case "m365":
            result = M365_TRAFFICFORWARDINGTYPE
        case "internet":
            result = INTERNET_TRAFFICFORWARDINGTYPE
        case "private":
            result = PRIVATE_TRAFFICFORWARDINGTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TRAFFICFORWARDINGTYPE
        default:
            return 0, errors.New("Unknown TrafficForwardingType value: " + v)
    }
    return &result, nil
}
func SerializeTrafficForwardingType(values []TrafficForwardingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TrafficForwardingType) isMultiValue() bool {
    return false
}
