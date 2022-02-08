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
    result := MICROSOFT_MANAGEMENTPROVIDER
    switch strings.ToUpper(v) {
        case "MICROSOFT":
            result = MICROSOFT_MANAGEMENTPROVIDER
        case "COMMUNITY":
            result = COMMUNITY_MANAGEMENTPROVIDER
        case "INDIRECTPROVIDER":
            result = INDIRECTPROVIDER_MANAGEMENTPROVIDER
        case "SELF":
            result = SELF_MANAGEMENTPROVIDER
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTPROVIDER
        default:
            return 0, errors.New("Unknown ManagementProvider value: " + v)
    }
    return &result, nil
}
func SerializeManagementProvider(values []ManagementProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
