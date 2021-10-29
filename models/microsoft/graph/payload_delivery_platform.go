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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PAYLOADDELIVERYPLATFORM, nil
        case "SMS":
            return SMS_PAYLOADDELIVERYPLATFORM, nil
        case "EMAIL":
            return EMAIL_PAYLOADDELIVERYPLATFORM, nil
        case "TEAMS":
            return TEAMS_PAYLOADDELIVERYPLATFORM, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PAYLOADDELIVERYPLATFORM, nil
    }
    return 0, errors.New("Unknown PayloadDeliveryPlatform value: " + v)
}
func SerializePayloadDeliveryPlatform(values []PayloadDeliveryPlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
