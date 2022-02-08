package windowsupdates
import (
    "strings"
    "errors"
)
// 
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
    result := NONE_REQUESTEDDEPLOYMENTSTATEVALUE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_REQUESTEDDEPLOYMENTSTATEVALUE
        case "PAUSED":
            result = PAUSED_REQUESTEDDEPLOYMENTSTATEVALUE
        case "ARCHIVED":
            result = ARCHIVED_REQUESTEDDEPLOYMENTSTATEVALUE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REQUESTEDDEPLOYMENTSTATEVALUE
        default:
            return 0, errors.New("Unknown RequestedDeploymentStateValue value: " + v)
    }
    return &result, nil
}
func SerializeRequestedDeploymentStateValue(values []RequestedDeploymentStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
