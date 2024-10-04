package healthmonitoring
type AlertState int

const (
    ACTIVE_ALERTSTATE AlertState = iota
    RESOLVED_ALERTSTATE
    UNKNOWNFUTUREVALUE_ALERTSTATE
)

func (i AlertState) String() string {
    return []string{"active", "resolved", "unknownFutureValue"}[i]
}
func ParseAlertState(v string) (any, error) {
    result := ACTIVE_ALERTSTATE
    switch v {
        case "active":
            result = ACTIVE_ALERTSTATE
        case "resolved":
            result = RESOLVED_ALERTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlertState(values []AlertState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AlertState) isMultiValue() bool {
    return false
}
