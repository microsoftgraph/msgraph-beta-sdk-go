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
    switch strings.ToUpper(v) {
        case "DISABLED":
            return DISABLED_MANAGEDINSTALLERSTATUS, nil
        case "ENABLED":
            return ENABLED_MANAGEDINSTALLERSTATUS, nil
    }
    return 0, errors.New("Unknown ManagedInstallerStatus value: " + v)
}
func SerializeManagedInstallerStatus(values []ManagedInstallerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
