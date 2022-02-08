package graph
import (
    "strings"
    "errors"
)
// 
type PayloadDeliveryPlatform int

const (
    UNKNOWN_PAYLOADDELIVERYPLATFORM PayloadDeliveryPlatform = iota
    SMS_PAYLOADDELIVERYPLATFORM
    EMAIL_PAYLOADDELIVERYPLATFORM
    TEAMS_PAYLOADDELIVERYPLATFORM
    UNKNOWNFUTUREVALUE_PAYLOADDELIVERYPLATFORM
)

func (i PayloadDeliveryPlatform) String() string {
    return []string{"UNKNOWN", "SMS", "EMAIL", "TEAMS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePayloadDeliveryPlatform(v string) (interface{}, error) {
    result := UNKNOWN_PAYLOADDELIVERYPLATFORM
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_PAYLOADDELIVERYPLATFORM
        case "SMS":
            result = SMS_PAYLOADDELIVERYPLATFORM
        case "EMAIL":
            result = EMAIL_PAYLOADDELIVERYPLATFORM
        case "TEAMS":
            result = TEAMS_PAYLOADDELIVERYPLATFORM
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PAYLOADDELIVERYPLATFORM
        default:
            return 0, errors.New("Unknown PayloadDeliveryPlatform value: " + v)
    }
    return &result, nil
}
func SerializePayloadDeliveryPlatform(values []PayloadDeliveryPlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
