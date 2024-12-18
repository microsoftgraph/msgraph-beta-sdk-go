package security
type AppInfoFedRampLevel int

const (
    HIGH_APPINFOFEDRAMPLEVEL AppInfoFedRampLevel = iota
    MODERATE_APPINFOFEDRAMPLEVEL
    LOW_APPINFOFEDRAMPLEVEL
    LISAAS_APPINFOFEDRAMPLEVEL
    UNKNOWN_APPINFOFEDRAMPLEVEL
    UNKNOWNFUTUREVALUE_APPINFOFEDRAMPLEVEL
    NOTSUPPORTED_APPINFOFEDRAMPLEVEL
)

func (i AppInfoFedRampLevel) String() string {
    return []string{"high", "moderate", "low", "liSaaS", "unknown", "unknownFutureValue", "notSupported"}[i]
}
func ParseAppInfoFedRampLevel(v string) (any, error) {
    result := HIGH_APPINFOFEDRAMPLEVEL
    switch v {
        case "high":
            result = HIGH_APPINFOFEDRAMPLEVEL
        case "moderate":
            result = MODERATE_APPINFOFEDRAMPLEVEL
        case "low":
            result = LOW_APPINFOFEDRAMPLEVEL
        case "liSaaS":
            result = LISAAS_APPINFOFEDRAMPLEVEL
        case "unknown":
            result = UNKNOWN_APPINFOFEDRAMPLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFOFEDRAMPLEVEL
        case "notSupported":
            result = NOTSUPPORTED_APPINFOFEDRAMPLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoFedRampLevel(values []AppInfoFedRampLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoFedRampLevel) isMultiValue() bool {
    return false
}
