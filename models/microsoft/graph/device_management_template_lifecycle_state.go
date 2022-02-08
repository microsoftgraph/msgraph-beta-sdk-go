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
    result := INVALID_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    switch strings.ToUpper(v) {
        case "INVALID":
            result = INVALID_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "DRAFT":
            result = DRAFT_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "ACTIVE":
            result = ACTIVE_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "SUPERSEDED":
            result = SUPERSEDED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "DEPRECATED":
            result = DEPRECATED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "RETIRED":
            result = RETIRED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        default:
            return 0, errors.New("Unknown DeviceManagementTemplateLifecycleState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementTemplateLifecycleState(values []DeviceManagementTemplateLifecycleState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
