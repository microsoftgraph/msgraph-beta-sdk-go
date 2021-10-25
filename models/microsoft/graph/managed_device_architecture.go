package graph
import (
    "strings"
    "errors"
)
type ManagedDeviceArchitecture int

const (
    UNKNOWN_MANAGEDDEVICEARCHITECTURE ManagedDeviceArchitecture = iota
    X86_MANAGEDDEVICEARCHITECTURE
    X64_MANAGEDDEVICEARCHITECTURE
    ARM_MANAGEDDEVICEARCHITECTURE
    ARM64_MANAGEDDEVICEARCHITECTURE
)

func (i ManagedDeviceArchitecture) String() string {
    return []string{"UNKNOWN", "X86", "X64", "ARM", "ARM64"}[i]
}
func ParseManagedDeviceArchitecture(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MANAGEDDEVICEARCHITECTURE, nil
        case "X86":
            return X86_MANAGEDDEVICEARCHITECTURE, nil
        case "X64":
            return X64_MANAGEDDEVICEARCHITECTURE, nil
        case "ARM":
            return ARM_MANAGEDDEVICEARCHITECTURE, nil
        case "ARM64":
            return ARM64_MANAGEDDEVICEARCHITECTURE, nil
    }
    return 0, errors.New("Unknown ManagedDeviceArchitecture value: " + v)
}
func SerializeManagedDeviceArchitecture(values []ManagedDeviceArchitecture) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
