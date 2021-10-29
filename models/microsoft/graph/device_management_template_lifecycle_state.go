package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementTemplateLifecycleState int

const (
    INVALID_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE DeviceManagementTemplateLifecycleState = iota
    DRAFT_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    ACTIVE_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    SUPERSEDED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    DEPRECATED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    RETIRED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
)

func (i DeviceManagementTemplateLifecycleState) String() string {
    return []string{"INVALID", "DRAFT", "ACTIVE", "SUPERSEDED", "DEPRECATED", "RETIRED"}[i]
}
func ParseDeviceManagementTemplateLifecycleState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INVALID":
            return INVALID_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE, nil
        case "DRAFT":
            return DRAFT_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE, nil
        case "ACTIVE":
            return ACTIVE_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE, nil
        case "SUPERSEDED":
            return SUPERSEDED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE, nil
        case "DEPRECATED":
            return DEPRECATED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE, nil
        case "RETIRED":
            return RETIRED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE, nil
    }
    return 0, errors.New("Unknown DeviceManagementTemplateLifecycleState value: " + v)
}
func SerializeDeviceManagementTemplateLifecycleState(values []DeviceManagementTemplateLifecycleState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
