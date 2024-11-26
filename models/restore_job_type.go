package models
type RestoreJobType int

const (
    STANDARD_RESTOREJOBTYPE RestoreJobType = iota
    BULK_RESTOREJOBTYPE
    UNKNOWNFUTUREVALUE_RESTOREJOBTYPE
)

func (i RestoreJobType) String() string {
    return []string{"standard", "bulk", "unknownFutureValue"}[i]
}
func ParseRestoreJobType(v string) (any, error) {
    result := STANDARD_RESTOREJOBTYPE
    switch v {
        case "standard":
            result = STANDARD_RESTOREJOBTYPE
        case "bulk":
            result = BULK_RESTOREJOBTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RESTOREJOBTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRestoreJobType(values []RestoreJobType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RestoreJobType) isMultiValue() bool {
    return false
}
