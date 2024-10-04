package healthmonitoring
type AlertType int

const (
    UNKNOWN_ALERTTYPE AlertType = iota
    MFASIGNINFAILURE_ALERTTYPE
    MANAGEDDEVICESIGNINFAILURE_ALERTTYPE
    COMPLIANTDEVICESIGNINFAILURE_ALERTTYPE
    UNKNOWNFUTUREVALUE_ALERTTYPE
)

func (i AlertType) String() string {
    return []string{"unknown", "mfaSignInFailure", "managedDeviceSignInFailure", "compliantDeviceSignInFailure", "unknownFutureValue"}[i]
}
func ParseAlertType(v string) (any, error) {
    result := UNKNOWN_ALERTTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_ALERTTYPE
        case "mfaSignInFailure":
            result = MFASIGNINFAILURE_ALERTTYPE
        case "managedDeviceSignInFailure":
            result = MANAGEDDEVICESIGNINFAILURE_ALERTTYPE
        case "compliantDeviceSignInFailure":
            result = COMPLIANTDEVICESIGNINFAILURE_ALERTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlertType(values []AlertType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AlertType) isMultiValue() bool {
    return false
}
