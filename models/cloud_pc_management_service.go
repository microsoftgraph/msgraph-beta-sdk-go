package models
import (
    "errors"
    "strings"
)
// 
type CloudPcManagementService int

const (
    WINDOWS365_CLOUDPCMANAGEMENTSERVICE CloudPcManagementService = iota
    DEVBOX_CLOUDPCMANAGEMENTSERVICE
    UNKNOWNFUTUREVALUE_CLOUDPCMANAGEMENTSERVICE
    RPABOX_CLOUDPCMANAGEMENTSERVICE
)

func (i CloudPcManagementService) String() string {
    var values []string
    for p := CloudPcManagementService(1); p <= RPABOX_CLOUDPCMANAGEMENTSERVICE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"windows365", "devBox", "unknownFutureValue", "rpaBox"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseCloudPcManagementService(v string) (any, error) {
    var result CloudPcManagementService
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "windows365":
                result |= WINDOWS365_CLOUDPCMANAGEMENTSERVICE
            case "devBox":
                result |= DEVBOX_CLOUDPCMANAGEMENTSERVICE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CLOUDPCMANAGEMENTSERVICE
            case "rpaBox":
                result |= RPABOX_CLOUDPCMANAGEMENTSERVICE
            default:
                return 0, errors.New("Unknown CloudPcManagementService value: " + v)
        }
    }
    return &result, nil
}
func SerializeCloudPcManagementService(values []CloudPcManagementService) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcManagementService) isMultiValue() bool {
    return true
}
