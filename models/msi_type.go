package models
type MsiType int

const (
    NONE_MSITYPE MsiType = iota
    USERASSIGNED_MSITYPE
    SYSTEMASSIGNED_MSITYPE
    UNKNOWNFUTUREVALUE_MSITYPE
)

func (i MsiType) String() string {
    return []string{"none", "userAssigned", "systemAssigned", "unknownFutureValue"}[i]
}
func ParseMsiType(v string) (any, error) {
    result := NONE_MSITYPE
    switch v {
        case "none":
            result = NONE_MSITYPE
        case "userAssigned":
            result = USERASSIGNED_MSITYPE
        case "systemAssigned":
            result = SYSTEMASSIGNED_MSITYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MSITYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMsiType(values []MsiType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MsiType) isMultiValue() bool {
    return false
}
