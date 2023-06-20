package networkaccess
import (
    "errors"
)
// 
type TrafficType int

const (
    INTERNET_TRAFFICTYPE TrafficType = iota
    PRIVATE_TRAFFICTYPE
    MICROSOFT365_TRAFFICTYPE
    ALL_TRAFFICTYPE
    UNKNOWNFUTUREVALUE_TRAFFICTYPE
)

func (i TrafficType) String() string {
    return []string{"internet", "private", "microsoft365", "all", "unknownFutureValue"}[i]
}
func ParseTrafficType(v string) (any, error) {
    result := INTERNET_TRAFFICTYPE
    switch v {
        case "internet":
            result = INTERNET_TRAFFICTYPE
        case "private":
            result = PRIVATE_TRAFFICTYPE
        case "microsoft365":
            result = MICROSOFT365_TRAFFICTYPE
        case "all":
            result = ALL_TRAFFICTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TRAFFICTYPE
        default:
            return 0, errors.New("Unknown TrafficType value: " + v)
    }
    return &result, nil
}
func SerializeTrafficType(values []TrafficType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
