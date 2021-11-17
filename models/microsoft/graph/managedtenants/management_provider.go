package managedtenants
import (
    "strings"
    "errors"
)
// 
type ManagementProvider int

const (
    MICROSOFT_MANAGEMENTPROVIDER ManagementProvider = iota
    COMMUNITY_MANAGEMENTPROVIDER
    INDIRECTPROVIDER_MANAGEMENTPROVIDER
    SELF_MANAGEMENTPROVIDER
    UNKNOWNFUTUREVALUE_MANAGEMENTPROVIDER
)

func (i ManagementProvider) String() string {
    return []string{"MICROSOFT", "COMMUNITY", "INDIRECTPROVIDER", "SELF", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseManagementProvider(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "MICROSOFT":
            return MICROSOFT_MANAGEMENTPROVIDER, nil
        case "COMMUNITY":
            return COMMUNITY_MANAGEMENTPROVIDER, nil
        case "INDIRECTPROVIDER":
            return INDIRECTPROVIDER_MANAGEMENTPROVIDER, nil
        case "SELF":
            return SELF_MANAGEMENTPROVIDER, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MANAGEMENTPROVIDER, nil
    }
    return 0, errors.New("Unknown ManagementProvider value: " + v)
}
func SerializeManagementProvider(values []ManagementProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
