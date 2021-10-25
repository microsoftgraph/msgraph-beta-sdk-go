package windowsupdates
import (
    "strings"
    "errors"
)
type DeploymentStateReasonValue int

const (
    SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE DeploymentStateReasonValue = iota
    OFFERINGBYREQUEST_DEPLOYMENTSTATEREASONVALUE
    PAUSEDBYREQUEST_DEPLOYMENTSTATEREASONVALUE
    PAUSEDBYMONITORING_DEPLOYMENTSTATEREASONVALUE
    UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEREASONVALUE
)

func (i DeploymentStateReasonValue) String() string {
    return []string{"SCHEDULEDBYOFFERWINDOW", "OFFERINGBYREQUEST", "PAUSEDBYREQUEST", "PAUSEDBYMONITORING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeploymentStateReasonValue(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SCHEDULEDBYOFFERWINDOW":
            return SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE, nil
        case "OFFERINGBYREQUEST":
            return OFFERINGBYREQUEST_DEPLOYMENTSTATEREASONVALUE, nil
        case "PAUSEDBYREQUEST":
            return PAUSEDBYREQUEST_DEPLOYMENTSTATEREASONVALUE, nil
        case "PAUSEDBYMONITORING":
            return PAUSEDBYMONITORING_DEPLOYMENTSTATEREASONVALUE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEREASONVALUE, nil
    }
    return 0, errors.New("Unknown DeploymentStateReasonValue value: " + v)
}
func SerializeDeploymentStateReasonValue(values []DeploymentStateReasonValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
