package windowsupdates
import (
    "strings"
    "errors"
)
// 
type MonitoringAction int

const (
    ALERTERROR_MONITORINGACTION MonitoringAction = iota
    PAUSEDEPLOYMENT_MONITORINGACTION
    UNKNOWNFUTUREVALUE_MONITORINGACTION
)

func (i MonitoringAction) String() string {
    return []string{"ALERTERROR", "PAUSEDEPLOYMENT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMonitoringAction(v string) (interface{}, error) {
    result := ALERTERROR_MONITORINGACTION
    switch strings.ToUpper(v) {
        case "ALERTERROR":
            result = ALERTERROR_MONITORINGACTION
        case "PAUSEDEPLOYMENT":
            result = PAUSEDEPLOYMENT_MONITORINGACTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MONITORINGACTION
        default:
            return 0, errors.New("Unknown MonitoringAction value: " + v)
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
