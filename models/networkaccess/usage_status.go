package networkaccess
type UsageStatus int

const (
    FREQUENTLYUSED_USAGESTATUS UsageStatus = iota
    RARELYUSED_USAGESTATUS
    UNKNOWNFUTUREVALUE_USAGESTATUS
)

func (i UsageStatus) String() string {
    return []string{"frequentlyUsed", "rarelyUsed", "unknownFutureValue"}[i]
}
func ParseUsageStatus(v string) (any, error) {
    result := FREQUENTLYUSED_USAGESTATUS
    switch v {
        case "frequentlyUsed":
            result = FREQUENTLYUSED_USAGESTATUS
        case "rarelyUsed":
            result = RARELYUSED_USAGESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USAGESTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUsageStatus(values []UsageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UsageStatus) isMultiValue() bool {
    return false
}
