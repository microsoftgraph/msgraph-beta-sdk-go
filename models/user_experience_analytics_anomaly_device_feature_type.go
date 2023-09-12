package models
import (
    "errors"
)
// Indicates the device's feature type. Possible values are: manufacturer, model, osVersion, application or driver.
type UserExperienceAnalyticsAnomalyDeviceFeatureType int

const (
    // Indicates the manufacturer name as device feature type.
    MANUFACTURER_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE UserExperienceAnalyticsAnomalyDeviceFeatureType = iota
    // Indicates the model as a device feature type.
    MODEL_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
    // Indicates the OS as a device feature type.
    OSVERSION_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
    // Indicates the application as a device feature type.
    APPLICATION_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
    // Indicates the driver as a device feature type.
    DRIVER_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
)

func (i UserExperienceAnalyticsAnomalyDeviceFeatureType) String() string {
    return []string{"manufacturer", "model", "osVersion", "application", "driver", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsAnomalyDeviceFeatureType(v string) (any, error) {
    result := MANUFACTURER_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
    switch v {
        case "manufacturer":
            result = MANUFACTURER_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
        case "model":
            result = MODEL_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
        case "osVersion":
            result = OSVERSION_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
        case "application":
            result = APPLICATION_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
        case "driver":
            result = DRIVER_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYDEVICEFEATURETYPE
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsAnomalyDeviceFeatureType value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsAnomalyDeviceFeatureType(values []UserExperienceAnalyticsAnomalyDeviceFeatureType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserExperienceAnalyticsAnomalyDeviceFeatureType) isMultiValue() bool {
    return false
}
