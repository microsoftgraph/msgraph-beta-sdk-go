package models
import (
    "errors"
)
// 
type FrontlineCloudPcAccessState int

const (
    UNASSIGNED_FRONTLINECLOUDPCACCESSSTATE FrontlineCloudPcAccessState = iota
    NOLICENSESAVAILABLE_FRONTLINECLOUDPCACCESSSTATE
    ACTIVATIONFAILED_FRONTLINECLOUDPCACCESSSTATE
    ACTIVE_FRONTLINECLOUDPCACCESSSTATE
    ACTIVATING_FRONTLINECLOUDPCACCESSSTATE
    STANDBYMODE_FRONTLINECLOUDPCACCESSSTATE
    UNKNOWNFUTUREVALUE_FRONTLINECLOUDPCACCESSSTATE
)

func (i FrontlineCloudPcAccessState) String() string {
    return []string{"unassigned", "noLicensesAvailable", "activationFailed", "active", "activating", "standbyMode", "unknownFutureValue"}[i]
}
func ParseFrontlineCloudPcAccessState(v string) (any, error) {
    result := UNASSIGNED_FRONTLINECLOUDPCACCESSSTATE
    switch v {
        case "unassigned":
            result = UNASSIGNED_FRONTLINECLOUDPCACCESSSTATE
        case "noLicensesAvailable":
            result = NOLICENSESAVAILABLE_FRONTLINECLOUDPCACCESSSTATE
        case "activationFailed":
            result = ACTIVATIONFAILED_FRONTLINECLOUDPCACCESSSTATE
        case "active":
            result = ACTIVE_FRONTLINECLOUDPCACCESSSTATE
        case "activating":
            result = ACTIVATING_FRONTLINECLOUDPCACCESSSTATE
        case "standbyMode":
            result = STANDBYMODE_FRONTLINECLOUDPCACCESSSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FRONTLINECLOUDPCACCESSSTATE
        default:
            return 0, errors.New("Unknown FrontlineCloudPcAccessState value: " + v)
    }
    return &result, nil
}
func SerializeFrontlineCloudPcAccessState(values []FrontlineCloudPcAccessState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FrontlineCloudPcAccessState) isMultiValue() bool {
    return false
}
