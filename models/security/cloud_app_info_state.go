package security
type CloudAppInfoState int

const (
    TRUE_CLOUDAPPINFOSTATE CloudAppInfoState = iota
    FALSE_CLOUDAPPINFOSTATE
    UNKNOWN_CLOUDAPPINFOSTATE
    UNKNOWNFUTUREVALUE_CLOUDAPPINFOSTATE
)

func (i CloudAppInfoState) String() string {
    return []string{"true", "false", "unknown", "unknownFutureValue"}[i]
}
func ParseCloudAppInfoState(v string) (any, error) {
    result := TRUE_CLOUDAPPINFOSTATE
    switch v {
        case "true":
            result = TRUE_CLOUDAPPINFOSTATE
        case "false":
            result = FALSE_CLOUDAPPINFOSTATE
        case "unknown":
            result = UNKNOWN_CLOUDAPPINFOSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDAPPINFOSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudAppInfoState(values []CloudAppInfoState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudAppInfoState) isMultiValue() bool {
    return false
}
