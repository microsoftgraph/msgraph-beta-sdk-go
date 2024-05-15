package models
// Indicates the category of the anomaly. Eg: anomaly type can be device, application, stop error, driver or other.
type UserExperienceAnalyticsAnomalyType int

const (
    // Indicates the detected anomaly is due to certain devices.
    DEVICE_USEREXPERIENCEANALYTICSANOMALYTYPE UserExperienceAnalyticsAnomalyType = iota
    // Indicates the detected anomaly is due to a specific application.
    APPLICATION_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Indicates the detected anomaly is due to a specific stop error.
    STOPERROR_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Indicates the detected anomaly is due to a specific driver.
    DRIVER_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Indicates the category of detected anomaly is undefined.
    OTHER_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYTYPE
)

func (i UserExperienceAnalyticsAnomalyType) String() string {
    return []string{"device", "application", "stopError", "driver", "other", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsAnomalyType(v string) (any, error) {
    result := DEVICE_USEREXPERIENCEANALYTICSANOMALYTYPE
    switch v {
        case "device":
            result = DEVICE_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "application":
            result = APPLICATION_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "stopError":
            result = STOPERROR_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "driver":
            result = DRIVER_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "other":
            result = OTHER_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsAnomalyType(values []UserExperienceAnalyticsAnomalyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserExperienceAnalyticsAnomalyType) isMultiValue() bool {
    return false
}
