package security
type AppInfoDataRetentionPolicy int

const (
    DATARETAINED_APPINFODATARETENTIONPOLICY AppInfoDataRetentionPolicy = iota
    DELETEDIMMEDIATELY_APPINFODATARETENTIONPOLICY
    DELETEDWITHINTWOWEEKS_APPINFODATARETENTIONPOLICY
    DELETEDWITHINONEMONTH_APPINFODATARETENTIONPOLICY
    DELETEDWITHINTHREEMONTHS_APPINFODATARETENTIONPOLICY
    DELETEDWITHINMORETHANTHREEMONTHS_APPINFODATARETENTIONPOLICY
    UNKNOWN_APPINFODATARETENTIONPOLICY
    UNKNOWNFUTUREVALUE_APPINFODATARETENTIONPOLICY
)

func (i AppInfoDataRetentionPolicy) String() string {
    return []string{"dataRetained", "deletedImmediately", "deletedWithinTwoWeeks", "deletedWithinOneMonth", "deletedWithinThreeMonths", "deletedWithinMoreThanThreeMonths", "unknown", "unknownFutureValue"}[i]
}
func ParseAppInfoDataRetentionPolicy(v string) (any, error) {
    result := DATARETAINED_APPINFODATARETENTIONPOLICY
    switch v {
        case "dataRetained":
            result = DATARETAINED_APPINFODATARETENTIONPOLICY
        case "deletedImmediately":
            result = DELETEDIMMEDIATELY_APPINFODATARETENTIONPOLICY
        case "deletedWithinTwoWeeks":
            result = DELETEDWITHINTWOWEEKS_APPINFODATARETENTIONPOLICY
        case "deletedWithinOneMonth":
            result = DELETEDWITHINONEMONTH_APPINFODATARETENTIONPOLICY
        case "deletedWithinThreeMonths":
            result = DELETEDWITHINTHREEMONTHS_APPINFODATARETENTIONPOLICY
        case "deletedWithinMoreThanThreeMonths":
            result = DELETEDWITHINMORETHANTHREEMONTHS_APPINFODATARETENTIONPOLICY
        case "unknown":
            result = UNKNOWN_APPINFODATARETENTIONPOLICY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFODATARETENTIONPOLICY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoDataRetentionPolicy(values []AppInfoDataRetentionPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoDataRetentionPolicy) isMultiValue() bool {
    return false
}
