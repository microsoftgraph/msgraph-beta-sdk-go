package graph
import (
    "strings"
    "errors"
)
// 
type AppManagementLevel int

const (
    UNSPECIFIED_APPMANAGEMENTLEVEL AppManagementLevel = iota
    UNMANAGED_APPMANAGEMENTLEVEL
    MDM_APPMANAGEMENTLEVEL
    ANDROIDENTERPRISE_APPMANAGEMENTLEVEL
)

func (i AppManagementLevel) String() string {
    return []string{"UNSPECIFIED", "UNMANAGED", "MDM", "ANDROIDENTERPRISE"}[i]
}
func ParseAppManagementLevel(v string) (interface{}, error) {
    result := UNSPECIFIED_APPMANAGEMENTLEVEL
    switch strings.ToUpper(v) {
        case "UNSPECIFIED":
            result = UNSPECIFIED_APPMANAGEMENTLEVEL
        case "UNMANAGED":
            result = UNMANAGED_APPMANAGEMENTLEVEL
        case "MDM":
            result = MDM_APPMANAGEMENTLEVEL
        case "ANDROIDENTERPRISE":
            result = ANDROIDENTERPRISE_APPMANAGEMENTLEVEL
        default:
            return 0, errors.New("Unknown AppManagementLevel value: " + v)
    }
    return &result, nil
}
func SerializeAppManagementLevel(values []AppManagementLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
