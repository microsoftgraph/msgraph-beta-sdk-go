package models
type UsageRightState int

const (
    ACTIVE_USAGERIGHTSTATE UsageRightState = iota
    INACTIVE_USAGERIGHTSTATE
    WARNING_USAGERIGHTSTATE
    SUSPENDED_USAGERIGHTSTATE
    UNKNOWNFUTUREVALUE_USAGERIGHTSTATE
)

func (i UsageRightState) String() string {
    return []string{"active", "inactive", "warning", "suspended", "unknownFutureValue"}[i]
}
func ParseUsageRightState(v string) (any, error) {
    result := ACTIVE_USAGERIGHTSTATE
    switch v {
        case "active":
            result = ACTIVE_USAGERIGHTSTATE
        case "inactive":
            result = INACTIVE_USAGERIGHTSTATE
        case "warning":
            result = WARNING_USAGERIGHTSTATE
        case "suspended":
            result = SUSPENDED_USAGERIGHTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USAGERIGHTSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUsageRightState(values []UsageRightState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UsageRightState) isMultiValue() bool {
    return false
}
