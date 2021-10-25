package windowsupdates
import (
    "strings"
    "errors"
)
type RequestedDeploymentStateValue int

const (
    NONE_REQUESTEDDEPLOYMENTSTATEVALUE RequestedDeploymentStateValue = iota
    PAUSED_REQUESTEDDEPLOYMENTSTATEVALUE
    ARCHIVED_REQUESTEDDEPLOYMENTSTATEVALUE
    UNKNOWNFUTUREVALUE_REQUESTEDDEPLOYMENTSTATEVALUE
)

func (i RequestedDeploymentStateValue) String() string {
    return []string{"NONE", "PAUSED", "ARCHIVED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRequestedDeploymentStateValue(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_REQUESTEDDEPLOYMENTSTATEVALUE, nil
        case "PAUSED":
            return PAUSED_REQUESTEDDEPLOYMENTSTATEVALUE, nil
        case "ARCHIVED":
            return ARCHIVED_REQUESTEDDEPLOYMENTSTATEVALUE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REQUESTEDDEPLOYMENTSTATEVALUE, nil
    }
    return 0, errors.New("Unknown RequestedDeploymentStateValue value: " + v)
}
func SerializeRequestedDeploymentStateValue(values []RequestedDeploymentStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
