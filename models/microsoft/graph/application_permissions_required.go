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
    result := UNKNOWN_APPLICATIONPERMISSIONSREQUIRED
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_APPLICATIONPERMISSIONSREQUIRED
        case "ANONYMOUS":
            result = ANONYMOUS_APPLICATIONPERMISSIONSREQUIRED
        case "GUEST":
            result = GUEST_APPLICATIONPERMISSIONSREQUIRED
        case "USER":
            result = USER_APPLICATIONPERMISSIONSREQUIRED
        case "ADMINISTRATOR":
            result = ADMINISTRATOR_APPLICATIONPERMISSIONSREQUIRED
        case "SYSTEM":
            result = SYSTEM_APPLICATIONPERMISSIONSREQUIRED
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_APPLICATIONPERMISSIONSREQUIRED
        default:
            return 0, errors.New("Unknown ApplicationPermissionsRequired value: " + v)
    }
    return &result, nil
}
func SerializeApplicationPermissionsRequired(values []ApplicationPermissionsRequired) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
