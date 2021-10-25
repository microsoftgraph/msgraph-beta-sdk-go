package graph
import (
    "strings"
    "errors"
)
type DeviceManagementExchangeAccessLevel int

const (
    NONE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL DeviceManagementExchangeAccessLevel = iota
    ALLOW_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
    BLOCK_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
    QUARANTINE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
)

func (i DeviceManagementExchangeAccessLevel) String() string {
    return []string{"NONE", "ALLOW", "BLOCK", "QUARANTINE"}[i]
}
func ParseDeviceManagementExchangeAccessLevel(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL, nil
        case "ALLOW":
            return ALLOW_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL, nil
        case "BLOCK":
            return BLOCK_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL, nil
        case "QUARANTINE":
            return QUARANTINE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL, nil
    }
    return 0, errors.New("Unknown DeviceManagementExchangeAccessLevel value: " + v)
}
func SerializeDeviceManagementExchangeAccessLevel(values []DeviceManagementExchangeAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
