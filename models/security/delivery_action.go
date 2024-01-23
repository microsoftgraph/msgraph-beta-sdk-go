package security
import (
    "errors"
)
// 
type DeliveryAction int

const (
    UNKNOWN_DELIVERYACTION DeliveryAction = iota
    DELIVEREDTOJUNK_DELIVERYACTION
    DELIVERED_DELIVERYACTION
    BLOCKED_DELIVERYACTION
    REPLACED_DELIVERYACTION
    UNKNOWNFUTUREVALUE_DELIVERYACTION
)

func (i DeliveryAction) String() string {
    return []string{"unknown", "deliveredToJunk", "delivered", "blocked", "replaced", "unknownFutureValue"}[i]
}
func ParseDeliveryAction(v string) (any, error) {
    result := UNKNOWN_DELIVERYACTION
    switch v {
        case "unknown":
            result = UNKNOWN_DELIVERYACTION
        case "deliveredToJunk":
            result = DELIVEREDTOJUNK_DELIVERYACTION
        case "delivered":
            result = DELIVERED_DELIVERYACTION
        case "blocked":
            result = BLOCKED_DELIVERYACTION
        case "replaced":
            result = REPLACED_DELIVERYACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DELIVERYACTION
        default:
            return 0, errors.New("Unknown DeliveryAction value: " + v)
    }
    return &result, nil
}
func SerializeDeliveryAction(values []DeliveryAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeliveryAction) isMultiValue() bool {
    return false
}
