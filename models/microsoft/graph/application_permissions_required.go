package graph
import (
    "strings"
    "errors"
)
// 
type ApplicationPermissionsRequired int

const (
    UNKNOWN_APPLICATIONPERMISSIONSREQUIRED ApplicationPermissionsRequired = iota
    ANONYMOUS_APPLICATIONPERMISSIONSREQUIRED
    GUEST_APPLICATIONPERMISSIONSREQUIRED
    USER_APPLICATIONPERMISSIONSREQUIRED
    ADMINISTRATOR_APPLICATIONPERMISSIONSREQUIRED
    SYSTEM_APPLICATIONPERMISSIONSREQUIRED
    UNKNOWNFUTUREVALUE_APPLICATIONPERMISSIONSREQUIRED
)

func (i ApplicationPermissionsRequired) String() string {
    return []string{"UNKNOWN", "ANONYMOUS", "GUEST", "USER", "ADMINISTRATOR", "SYSTEM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseApplicationPermissionsRequired(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_APPLICATIONPERMISSIONSREQUIRED, nil
        case "ANONYMOUS":
            return ANONYMOUS_APPLICATIONPERMISSIONSREQUIRED, nil
        case "GUEST":
            return GUEST_APPLICATIONPERMISSIONSREQUIRED, nil
        case "USER":
            return USER_APPLICATIONPERMISSIONSREQUIRED, nil
        case "ADMINISTRATOR":
            return ADMINISTRATOR_APPLICATIONPERMISSIONSREQUIRED, nil
        case "SYSTEM":
            return SYSTEM_APPLICATIONPERMISSIONSREQUIRED, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_APPLICATIONPERMISSIONSREQUIRED, nil
    }
    return 0, errors.New("Unknown ApplicationPermissionsRequired value: " + v)
}
func SerializeApplicationPermissionsRequired(values []ApplicationPermissionsRequired) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
