package windowsupdates
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SCHEDULED":
            return SCHEDULED_DEPLOYMENTSTATEVALUE, nil
        case "OFFERING":
            return OFFERING_DEPLOYMENTSTATEVALUE, nil
        case "PAUSED":
            return PAUSED_DEPLOYMENTSTATEVALUE, nil
        case "FAULTED":
            return FAULTED_DEPLOYMENTSTATEVALUE, nil
        case "ARCHIVED":
            return ARCHIVED_DEPLOYMENTSTATEVALUE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEVALUE, nil
    }
    return 0, errors.New("Unknown DeploymentStateValue value: " + v)
}
func SerializeDeploymentStateValue(values []DeploymentStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
