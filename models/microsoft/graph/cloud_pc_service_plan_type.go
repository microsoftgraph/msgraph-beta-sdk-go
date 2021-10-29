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
    switch strings.ToUpper(v) {
        case "ENTERPRISE":
            return ENTERPRISE_CLOUDPCSERVICEPLANTYPE, nil
        case "BUSINESS":
            return BUSINESS_CLOUDPCSERVICEPLANTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCSERVICEPLANTYPE, nil
    }
    return 0, errors.New("Unknown CloudPcServicePlanType value: " + v)
}
func SerializeCloudPcServicePlanType(values []CloudPcServicePlanType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
