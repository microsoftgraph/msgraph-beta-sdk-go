package models
// Indicates the type of device script.
type DeviceHealthScriptType int

const (
    // Indicates this is a device health script.
    DEVICEHEALTHSCRIPT_DEVICEHEALTHSCRIPTTYPE DeviceHealthScriptType = iota
    // Indicates this is a managed installer script.
    MANAGEDINSTALLERSCRIPT_DEVICEHEALTHSCRIPTTYPE
)

func (i DeviceHealthScriptType) String() string {
    return []string{"deviceHealthScript", "managedInstallerScript"}[i]
}
func ParseDeviceHealthScriptType(v string) (any, error) {
    result := DEVICEHEALTHSCRIPT_DEVICEHEALTHSCRIPTTYPE
    switch v {
        case "deviceHealthScript":
            result = DEVICEHEALTHSCRIPT_DEVICEHEALTHSCRIPTTYPE
        case "managedInstallerScript":
            result = MANAGEDINSTALLERSCRIPT_DEVICEHEALTHSCRIPTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceHealthScriptType(values []DeviceHealthScriptType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceHealthScriptType) isMultiValue() bool {
    return false
}
