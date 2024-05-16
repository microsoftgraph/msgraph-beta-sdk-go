package windowsupdates
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
    return []string{"scheduled", "offering", "paused", "faulted", "archived", "unknownFutureValue"}[i]
}
func ParseDeploymentStateValue(v string) (any, error) {
    result := SCHEDULED_DEPLOYMENTSTATEVALUE
    switch v {
        case "scheduled":
            result = SCHEDULED_DEPLOYMENTSTATEVALUE
        case "offering":
            result = OFFERING_DEPLOYMENTSTATEVALUE
        case "paused":
            result = PAUSED_DEPLOYMENTSTATEVALUE
        case "faulted":
            result = FAULTED_DEPLOYMENTSTATEVALUE
        case "archived":
            result = ARCHIVED_DEPLOYMENTSTATEVALUE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEVALUE
        default:
            return nil, nil
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
func (i DeploymentStateValue) isMultiValue() bool {
    return false
}
