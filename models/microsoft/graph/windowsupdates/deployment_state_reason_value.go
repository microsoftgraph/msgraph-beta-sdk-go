package windowsupdates
import (
    "strings"
    "errors"
)
// 
type DeploymentStateReasonValue int

const (
    SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE DeploymentStateReasonValue = iota
    OFFERINGBYREQUEST_DEPLOYMENTSTATEREASONVALUE
    PAUSEDBYREQUEST_DEPLOYMENTSTATEREASONVALUE
    PAUSEDBYMONITORING_DEPLOYMENTSTATEREASONVALUE
    UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEREASONVALUE
    FAULTEDBYCONTENTOUTDATED_DEPLOYMENTSTATEREASONVALUE
)

func (i DeploymentStateReasonValue) String() string {
    return []string{"SCHEDULEDBYOFFERWINDOW", "OFFERINGBYREQUEST", "PAUSEDBYREQUEST", "PAUSEDBYMONITORING", "UNKNOWNFUTUREVALUE", "FAULTEDBYCONTENTOUTDATED"}[i]
}
func ParseDeploymentStateReasonValue(v string) (interface{}, error) {
    result := SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE
    switch strings.ToUpper(v) {
        case "SCHEDULEDBYOFFERWINDOW":
            result = SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE
        case "OFFERINGBYREQUEST":
            result = OFFERINGBYREQUEST_DEPLOYMENTSTATEREASONVALUE
        case "PAUSEDBYREQUEST":
            result = PAUSEDBYREQUEST_DEPLOYMENTSTATEREASONVALUE
        case "PAUSEDBYMONITORING":
            result = PAUSEDBYMONITORING_DEPLOYMENTSTATEREASONVALUE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEREASONVALUE
        case "FAULTEDBYCONTENTOUTDATED":
            result = FAULTEDBYCONTENTOUTDATED_DEPLOYMENTSTATEREASONVALUE
        default:
            return 0, errors.New("Unknown DeploymentStateReasonValue value: " + v)
    }
    return &result, nil
}
func SerializeDeploymentStateReasonValue(values []DeploymentStateReasonValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
