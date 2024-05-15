package models
// Indicates the status of the device in the correlation group. Eg: Device status can be anomalous, affected, at risk.
type UserExperienceAnalyticsDeviceStatus int

const (
    // Indicates the the device is part of the anomaly.
    ANOMALOUS_USEREXPERIENCEANALYTICSDEVICESTATUS UserExperienceAnalyticsDeviceStatus = iota
    // Indicates the device is affected by the anomaly and is part of the correlation group.
    AFFECTED_USEREXPERIENCEANALYTICSDEVICESTATUS
    // Indicates the device is not part of the anomaly but is part of the correlation group.
    ATRISK_USEREXPERIENCEANALYTICSDEVICESTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSDEVICESTATUS
)

func (i UserExperienceAnalyticsDeviceStatus) String() string {
    return []string{"anomalous", "affected", "atRisk", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsDeviceStatus(v string) (any, error) {
    result := ANOMALOUS_USEREXPERIENCEANALYTICSDEVICESTATUS
    switch v {
        case "anomalous":
            result = ANOMALOUS_USEREXPERIENCEANALYTICSDEVICESTATUS
        case "affected":
            result = AFFECTED_USEREXPERIENCEANALYTICSDEVICESTATUS
        case "atRisk":
            result = ATRISK_USEREXPERIENCEANALYTICSDEVICESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSDEVICESTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsDeviceStatus(values []UserExperienceAnalyticsDeviceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserExperienceAnalyticsDeviceStatus) isMultiValue() bool {
    return false
}
