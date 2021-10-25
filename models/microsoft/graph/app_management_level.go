package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNSPECIFIED":
            return UNSPECIFIED_APPMANAGEMENTLEVEL, nil
        case "UNMANAGED":
            return UNMANAGED_APPMANAGEMENTLEVEL, nil
        case "MDM":
            return MDM_APPMANAGEMENTLEVEL, nil
        case "ANDROIDENTERPRISE":
            return ANDROIDENTERPRISE_APPMANAGEMENTLEVEL, nil
    }
    return 0, errors.New("Unknown AppManagementLevel value: " + v)
}
func SerializeAppManagementLevel(values []AppManagementLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
