package graph
import (
    "strings"
    "errors"
)
// 
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
    result := UNKNOWN_MANAGEDDEVICEARCHITECTURE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MANAGEDDEVICEARCHITECTURE
        case "X86":
            result = X86_MANAGEDDEVICEARCHITECTURE
        case "X64":
            result = X64_MANAGEDDEVICEARCHITECTURE
        case "ARM":
            result = ARM_MANAGEDDEVICEARCHITECTURE
        case "ARM64":
            result = ARM64_MANAGEDDEVICEARCHITECTURE
        default:
            return 0, errors.New("Unknown ManagedDeviceArchitecture value: " + v)
    }
    return &result, nil
}
func SerializeManagedDeviceArchitecture(values []ManagedDeviceArchitecture) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
