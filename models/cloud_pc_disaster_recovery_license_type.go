package models
type CloudPcDisasterRecoveryLicenseType int

const (
    NONE_CLOUDPCDISASTERRECOVERYLICENSETYPE CloudPcDisasterRecoveryLicenseType = iota
    STANDARD_CLOUDPCDISASTERRECOVERYLICENSETYPE
    UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYLICENSETYPE
)

func (i CloudPcDisasterRecoveryLicenseType) String() string {
    return []string{"none", "standard", "unknownFutureValue"}[i]
}
func ParseCloudPcDisasterRecoveryLicenseType(v string) (any, error) {
    result := NONE_CLOUDPCDISASTERRECOVERYLICENSETYPE
    switch v {
        case "none":
            result = NONE_CLOUDPCDISASTERRECOVERYLICENSETYPE
        case "standard":
            result = STANDARD_CLOUDPCDISASTERRECOVERYLICENSETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYLICENSETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcDisasterRecoveryLicenseType(values []CloudPcDisasterRecoveryLicenseType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcDisasterRecoveryLicenseType) isMultiValue() bool {
    return false
}
