package graph
import (
    "strings"
    "errors"
)
// 
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
    result := NONE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        case "ALLOW":
            result = ALLOW_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        case "BLOCK":
            result = BLOCK_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        case "QUARANTINE":
            result = QUARANTINE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeAccessLevel value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeAccessLevel(values []DeviceManagementExchangeAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
