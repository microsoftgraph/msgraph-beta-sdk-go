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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE, nil
        case "ERROR":
            return ERROR_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE, nil
        case "INACTIVE":
            return INACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE, nil
    }
    return 0, errors.New("Unknown DeviceManagementDomainJoinConnectorState value: " + v)
}
func SerializeDeviceManagementDomainJoinConnectorState(values []DeviceManagementDomainJoinConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
