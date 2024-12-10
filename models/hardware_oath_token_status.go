package models
type HardwareOathTokenStatus int

const (
    AVAILABLE_HARDWAREOATHTOKENSTATUS HardwareOathTokenStatus = iota
    ASSIGNED_HARDWAREOATHTOKENSTATUS
    ACTIVATED_HARDWAREOATHTOKENSTATUS
    FAILEDACTIVATION_HARDWAREOATHTOKENSTATUS
    UNKNOWNFUTUREVALUE_HARDWAREOATHTOKENSTATUS
)

func (i HardwareOathTokenStatus) String() string {
    return []string{"available", "assigned", "activated", "failedActivation", "unknownFutureValue"}[i]
}
func ParseHardwareOathTokenStatus(v string) (any, error) {
    result := AVAILABLE_HARDWAREOATHTOKENSTATUS
    switch v {
        case "available":
            result = AVAILABLE_HARDWAREOATHTOKENSTATUS
        case "assigned":
            result = ASSIGNED_HARDWAREOATHTOKENSTATUS
        case "activated":
            result = ACTIVATED_HARDWAREOATHTOKENSTATUS
        case "failedActivation":
            result = FAILEDACTIVATION_HARDWAREOATHTOKENSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_HARDWAREOATHTOKENSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHardwareOathTokenStatus(values []HardwareOathTokenStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HardwareOathTokenStatus) isMultiValue() bool {
    return false
}
