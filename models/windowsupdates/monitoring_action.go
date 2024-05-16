package windowsupdates
type MonitoringAction int

const (
    ALERTERROR_MONITORINGACTION MonitoringAction = iota
    OFFERFALLBACK_MONITORINGACTION
    PAUSEDEPLOYMENT_MONITORINGACTION
    UNKNOWNFUTUREVALUE_MONITORINGACTION
)

func (i MonitoringAction) String() string {
    return []string{"alertError", "offerFallback", "pauseDeployment", "unknownFutureValue"}[i]
}
func ParseMonitoringAction(v string) (any, error) {
    result := ALERTERROR_MONITORINGACTION
    switch v {
        case "alertError":
            result = ALERTERROR_MONITORINGACTION
        case "offerFallback":
            result = OFFERFALLBACK_MONITORINGACTION
        case "pauseDeployment":
            result = PAUSEDEPLOYMENT_MONITORINGACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MONITORINGACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMonitoringAction(values []MonitoringAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MonitoringAction) isMultiValue() bool {
    return false
}
