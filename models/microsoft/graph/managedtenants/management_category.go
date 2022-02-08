package managedtenants
import (
    "strings"
    "errors"
)
// 
type ManagementCategory int

const (
    CUSTOM_MANAGEMENTCATEGORY ManagementCategory = iota
    DEVICES_MANAGEMENTCATEGORY
    IDENTITY_MANAGEMENTCATEGORY
    DATA_MANAGEMENTCATEGORY
    UNKNOWNFUTUREVALUE_MANAGEMENTCATEGORY
)

func (i ManagementCategory) String() string {
    return []string{"CUSTOM", "DEVICES", "IDENTITY", "DATA", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseManagementCategory(v string) (interface{}, error) {
    result := CUSTOM_MANAGEMENTCATEGORY
    switch strings.ToUpper(v) {
        case "CUSTOM":
            result = CUSTOM_MANAGEMENTCATEGORY
        case "DEVICES":
            result = DEVICES_MANAGEMENTCATEGORY
        case "IDENTITY":
            result = IDENTITY_MANAGEMENTCATEGORY
        case "DATA":
            result = DATA_MANAGEMENTCATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTCATEGORY
        default:
            return 0, errors.New("Unknown ManagementCategory value: " + v)
    }
    return &result, nil
}
func SerializeManagementCategory(values []ManagementCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
