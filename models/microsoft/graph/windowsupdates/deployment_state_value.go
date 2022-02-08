package windowsupdates
import (
    "strings"
    "errors"
)
// 
type DeploymentStateValue int

const (
    SCHEDULED_DEPLOYMENTSTATEVALUE DeploymentStateValue = iota
    OFFERING_DEPLOYMENTSTATEVALUE
    PAUSED_DEPLOYMENTSTATEVALUE
    FAULTED_DEPLOYMENTSTATEVALUE
    ARCHIVED_DEPLOYMENTSTATEVALUE
    UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEVALUE
)

func (i DeploymentStateValue) String() string {
    return []string{"SCHEDULED", "OFFERING", "PAUSED", "FAULTED", "ARCHIVED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeploymentStateValue(v string) (interface{}, error) {
    result := SCHEDULED_DEPLOYMENTSTATEVALUE
    switch strings.ToUpper(v) {
        case "SCHEDULED":
            result = SCHEDULED_DEPLOYMENTSTATEVALUE
        case "OFFERING":
            result = OFFERING_DEPLOYMENTSTATEVALUE
        case "PAUSED":
            result = PAUSED_DEPLOYMENTSTATEVALUE
        case "FAULTED":
            result = FAULTED_DEPLOYMENTSTATEVALUE
        case "ARCHIVED":
            result = ARCHIVED_DEPLOYMENTSTATEVALUE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEVALUE
        default:
            return 0, errors.New("Unknown DeploymentStateValue value: " + v)
    }
    return &result, nil
}
func SerializeDeploymentStateValue(values []DeploymentStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
