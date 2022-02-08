package graph
import (
    "strings"
    "errors"
)
// 
type ManagedInstallerStatus int

const (
    DISABLED_MANAGEDINSTALLERSTATUS ManagedInstallerStatus = iota
    ENABLED_MANAGEDINSTALLERSTATUS
)

func (i ManagedInstallerStatus) String() string {
    return []string{"DISABLED", "ENABLED"}[i]
}
func ParseManagedInstallerStatus(v string) (interface{}, error) {
    result := DISABLED_MANAGEDINSTALLERSTATUS
    switch strings.ToUpper(v) {
        case "DISABLED":
            result = DISABLED_MANAGEDINSTALLERSTATUS
        case "ENABLED":
            result = ENABLED_MANAGEDINSTALLERSTATUS
        default:
            return 0, errors.New("Unknown ManagedInstallerStatus value: " + v)
    }
    return &result, nil
}
func SerializeManagedInstallerStatus(values []ManagedInstallerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
