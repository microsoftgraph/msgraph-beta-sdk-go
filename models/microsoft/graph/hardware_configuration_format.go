package graph
import (
    "strings"
    "errors"
)
// 
type HardwareConfigurationFormat int

const (
    DELL_HARDWARECONFIGURATIONFORMAT HardwareConfigurationFormat = iota
    SURFACE_HARDWARECONFIGURATIONFORMAT
    SURFACEDOCK_HARDWARECONFIGURATIONFORMAT
)

func (i HardwareConfigurationFormat) String() string {
    return []string{"DELL", "SURFACE", "SURFACEDOCK"}[i]
}
func ParseHardwareConfigurationFormat(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DELL":
            return DELL_HARDWARECONFIGURATIONFORMAT, nil
        case "SURFACE":
            return SURFACE_HARDWARECONFIGURATIONFORMAT, nil
        case "SURFACEDOCK":
            return SURFACEDOCK_HARDWARECONFIGURATIONFORMAT, nil
    }
    return 0, errors.New("Unknown HardwareConfigurationFormat value: " + v)
}
func SerializeHardwareConfigurationFormat(values []HardwareConfigurationFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
