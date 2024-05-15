package industrydata
type ReadinessStatus int

const (
    NOTREADY_READINESSSTATUS ReadinessStatus = iota
    READY_READINESSSTATUS
    FAILED_READINESSSTATUS
    DISABLED_READINESSSTATUS
    EXPIRED_READINESSSTATUS
    UNKNOWNFUTUREVALUE_READINESSSTATUS
)

func (i ReadinessStatus) String() string {
    return []string{"notReady", "ready", "failed", "disabled", "expired", "unknownFutureValue"}[i]
}
func ParseReadinessStatus(v string) (any, error) {
    result := NOTREADY_READINESSSTATUS
    switch v {
        case "notReady":
            result = NOTREADY_READINESSSTATUS
        case "ready":
            result = READY_READINESSSTATUS
        case "failed":
            result = FAILED_READINESSSTATUS
        case "disabled":
            result = DISABLED_READINESSSTATUS
        case "expired":
            result = EXPIRED_READINESSSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_READINESSSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeReadinessStatus(values []ReadinessStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReadinessStatus) isMultiValue() bool {
    return false
}
