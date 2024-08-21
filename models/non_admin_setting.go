package models
type NonAdminSetting int

const (
    FALSE_NONADMINSETTING NonAdminSetting = iota
    TRUE_NONADMINSETTING
    UNKNOWNFUTUREVALUE_NONADMINSETTING
)

func (i NonAdminSetting) String() string {
    return []string{"false", "true", "unknownFutureValue"}[i]
}
func ParseNonAdminSetting(v string) (any, error) {
    result := FALSE_NONADMINSETTING
    switch v {
        case "false":
            result = FALSE_NONADMINSETTING
        case "true":
            result = TRUE_NONADMINSETTING
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NONADMINSETTING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNonAdminSetting(values []NonAdminSetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NonAdminSetting) isMultiValue() bool {
    return false
}
