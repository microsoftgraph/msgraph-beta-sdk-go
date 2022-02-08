package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcServicePlanType int

const (
    ENTERPRISE_CLOUDPCSERVICEPLANTYPE CloudPcServicePlanType = iota
    BUSINESS_CLOUDPCSERVICEPLANTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCSERVICEPLANTYPE
)

func (i CloudPcServicePlanType) String() string {
    return []string{"ENTERPRISE", "BUSINESS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcServicePlanType(v string) (interface{}, error) {
    result := ENTERPRISE_CLOUDPCSERVICEPLANTYPE
    switch strings.ToUpper(v) {
        case "ENTERPRISE":
            result = ENTERPRISE_CLOUDPCSERVICEPLANTYPE
        case "BUSINESS":
            result = BUSINESS_CLOUDPCSERVICEPLANTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSERVICEPLANTYPE
        default:
            return 0, errors.New("Unknown CloudPcServicePlanType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcServicePlanType(values []CloudPcServicePlanType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
