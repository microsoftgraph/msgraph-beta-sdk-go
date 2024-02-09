package industrydata
import (
    "errors"
)
type InboundDomain int

const (
    EDUCATIONROSTERING_INBOUNDDOMAIN InboundDomain = iota
    UNKNOWNFUTUREVALUE_INBOUNDDOMAIN
)

func (i InboundDomain) String() string {
    return []string{"educationRostering", "unknownFutureValue"}[i]
}
func ParseInboundDomain(v string) (any, error) {
    result := EDUCATIONROSTERING_INBOUNDDOMAIN
    switch v {
        case "educationRostering":
            result = EDUCATIONROSTERING_INBOUNDDOMAIN
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INBOUNDDOMAIN
        default:
            return 0, errors.New("Unknown InboundDomain value: " + v)
    }
    return &result, nil
}
func SerializeInboundDomain(values []InboundDomain) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InboundDomain) isMultiValue() bool {
    return false
}
