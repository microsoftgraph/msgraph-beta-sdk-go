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
    switch strings.ToUpper(v) {
        case "ALERTERROR":
            return ALERTERROR_MONITORINGACTION, nil
        case "PAUSEDEPLOYMENT":
            return PAUSEDEPLOYMENT_MONITORINGACTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MONITORINGACTION, nil
    }
    return 0, errors.New("Unknown MonitoringAction value: " + v)
}
func SerializeMonitoringAction(values []MonitoringAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
