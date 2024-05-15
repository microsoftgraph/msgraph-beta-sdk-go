package models
type GcpRoleType int

const (
    SYSTEM_GCPROLETYPE GcpRoleType = iota
    CUSTOM_GCPROLETYPE
    UNKNOWNFUTUREVALUE_GCPROLETYPE
)

func (i GcpRoleType) String() string {
    return []string{"system", "custom", "unknownFutureValue"}[i]
}
func ParseGcpRoleType(v string) (any, error) {
    result := SYSTEM_GCPROLETYPE
    switch v {
        case "system":
            result = SYSTEM_GCPROLETYPE
        case "custom":
            result = CUSTOM_GCPROLETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_GCPROLETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGcpRoleType(values []GcpRoleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GcpRoleType) isMultiValue() bool {
    return false
}
