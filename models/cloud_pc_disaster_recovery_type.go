package models
type CloudPcDisasterRecoveryType int

const (
    NOTCONFIGURED_CLOUDPCDISASTERRECOVERYTYPE CloudPcDisasterRecoveryType = iota
    CROSSREGION_CLOUDPCDISASTERRECOVERYTYPE
    PREMIUM_CLOUDPCDISASTERRECOVERYTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYTYPE
)

func (i CloudPcDisasterRecoveryType) String() string {
    return []string{"notConfigured", "crossRegion", "premium", "unknownFutureValue"}[i]
}
func ParseCloudPcDisasterRecoveryType(v string) (any, error) {
    result := NOTCONFIGURED_CLOUDPCDISASTERRECOVERYTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_CLOUDPCDISASTERRECOVERYTYPE
        case "crossRegion":
            result = CROSSREGION_CLOUDPCDISASTERRECOVERYTYPE
        case "premium":
            result = PREMIUM_CLOUDPCDISASTERRECOVERYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcDisasterRecoveryType(values []CloudPcDisasterRecoveryType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcDisasterRecoveryType) isMultiValue() bool {
    return false
}
