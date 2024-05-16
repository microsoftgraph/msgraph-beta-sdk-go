package models
// Indicates the supported oems of hardware configuration
type HardwareConfigurationFormat int

const (
    // Dell
    DELL_HARDWARECONFIGURATIONFORMAT HardwareConfigurationFormat = iota
    // Surface
    SURFACE_HARDWARECONFIGURATIONFORMAT
    // Surface dock
    SURFACEDOCK_HARDWARECONFIGURATIONFORMAT
)

func (i HardwareConfigurationFormat) String() string {
    return []string{"dell", "surface", "surfaceDock"}[i]
}
func ParseHardwareConfigurationFormat(v string) (any, error) {
    result := DELL_HARDWARECONFIGURATIONFORMAT
    switch v {
        case "dell":
            result = DELL_HARDWARECONFIGURATIONFORMAT
        case "surface":
            result = SURFACE_HARDWARECONFIGURATIONFORMAT
        case "surfaceDock":
            result = SURFACEDOCK_HARDWARECONFIGURATIONFORMAT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHardwareConfigurationFormat(values []HardwareConfigurationFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HardwareConfigurationFormat) isMultiValue() bool {
    return false
}
