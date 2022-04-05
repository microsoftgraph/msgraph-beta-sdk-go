package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type CloudPcManagementService int

const (
    WINDOWS365_CLOUDPCMANAGEMENTSERVICE CloudPcManagementService = iota
    DEVBOX_CLOUDPCMANAGEMENTSERVICE
    UNKNOWNFUTUREVALUE_CLOUDPCMANAGEMENTSERVICE
)

func (i CloudPcManagementService) String() string {
    return []string{"WINDOWS365", "DEVBOX", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcManagementService(v string) (interface{}, error) {
    result := WINDOWS365_CLOUDPCMANAGEMENTSERVICE
    switch strings.ToUpper(v) {
        case "WINDOWS365":
            result = WINDOWS365_CLOUDPCMANAGEMENTSERVICE
        case "DEVBOX":
            result = DEVBOX_CLOUDPCMANAGEMENTSERVICE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCMANAGEMENTSERVICE
        default:
            return 0, errors.New("Unknown CloudPcManagementService value: " + v)
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
