package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementDomainJoinConnectorState int

const (
    ACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE DeviceManagementDomainJoinConnectorState = iota
    ERROR_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
    INACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
)

func (i DeviceManagementDomainJoinConnectorState) String() string {
    return []string{"ACTIVE", "ERROR", "INACTIVE"}[i]
}
func ParseDeviceManagementDomainJoinConnectorState(v string) (interface{}, error) {
    result := ACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
        case "ERROR":
            result = ERROR_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
        case "INACTIVE":
            result = INACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
        default:
            return 0, errors.New("Unknown DeviceManagementDomainJoinConnectorState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementDomainJoinConnectorState(values []DeviceManagementDomainJoinConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
