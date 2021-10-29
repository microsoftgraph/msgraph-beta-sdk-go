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
    UNKNOWNFUTUREVALUE_MANAGEMENTCATEGORY
)

func (i ManagementCategory) String() string {
    return []string{"CUSTOM", "DEVICES", "IDENTITY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseManagementCategory(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CUSTOM":
            return CUSTOM_MANAGEMENTCATEGORY, nil
        case "DEVICES":
            return DEVICES_MANAGEMENTCATEGORY, nil
        case "IDENTITY":
            return IDENTITY_MANAGEMENTCATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MANAGEMENTCATEGORY, nil
    }
    return 0, errors.New("Unknown ManagementCategory value: " + v)
}
func SerializeManagementCategory(values []ManagementCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
