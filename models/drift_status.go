// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type DriftStatus int

const (
    ACTIVE_DRIFTSTATUS DriftStatus = iota
    FIXED_DRIFTSTATUS
    // A marker value for members added after the release of this API.
    UNKNOWNFUTUREVALUE_DRIFTSTATUS
)

func (i DriftStatus) String() string {
    return []string{"active", "fixed", "unknownFutureValue"}[i]
}
func ParseDriftStatus(v string) (any, error) {
    result := ACTIVE_DRIFTSTATUS
    switch v {
        case "active":
            result = ACTIVE_DRIFTSTATUS
        case "fixed":
            result = FIXED_DRIFTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DRIFTSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDriftStatus(values []DriftStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DriftStatus) isMultiValue() bool {
    return false
}
