package industrydata
type ApiFormat int

const (
    ONEROSTER_APIFORMAT ApiFormat = iota
    UNKNOWNFUTUREVALUE_APIFORMAT
)

func (i ApiFormat) String() string {
    return []string{"oneRoster", "unknownFutureValue"}[i]
}
func ParseApiFormat(v string) (any, error) {
    result := ONEROSTER_APIFORMAT
    switch v {
        case "oneRoster":
            result = ONEROSTER_APIFORMAT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APIFORMAT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApiFormat(values []ApiFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApiFormat) isMultiValue() bool {
    return false
}
