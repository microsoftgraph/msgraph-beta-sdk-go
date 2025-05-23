// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement
type AlertStatusType int

const (
    ACTIVE_ALERTSTATUSTYPE AlertStatusType = iota
    RESOLVED_ALERTSTATUSTYPE
    UNKNOWNFUTUREVALUE_ALERTSTATUSTYPE
)

func (i AlertStatusType) String() string {
    return []string{"active", "resolved", "unknownFutureValue"}[i]
}
func ParseAlertStatusType(v string) (any, error) {
    result := ACTIVE_ALERTSTATUSTYPE
    switch v {
        case "active":
            result = ACTIVE_ALERTSTATUSTYPE
        case "resolved":
            result = RESOLVED_ALERTSTATUSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTSTATUSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlertStatusType(values []AlertStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AlertStatusType) isMultiValue() bool {
    return false
}
