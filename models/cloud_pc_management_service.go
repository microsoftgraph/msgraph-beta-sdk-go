package models
import (
    "math"
    "strings"
)
type CloudPcManagementService int

const (
    WINDOWS365_CLOUDPCMANAGEMENTSERVICE = 1
    DEVBOX_CLOUDPCMANAGEMENTSERVICE = 2
    UNKNOWNFUTUREVALUE_CLOUDPCMANAGEMENTSERVICE = 4
    RPABOX_CLOUDPCMANAGEMENTSERVICE = 8
)

func (i CloudPcManagementService) String() string {
    var values []string
    options := []string{"windows365", "devBox", "unknownFutureValue", "rpaBox"}
    for p := 0; p < 4; p++ {
        mantis := CloudPcManagementService(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
