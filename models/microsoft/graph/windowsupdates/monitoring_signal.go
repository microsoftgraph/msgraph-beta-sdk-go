package windowsupdates
import (
    "strings"
    "errors"
)
// 
type MonitoringSignal int

const (
    ROLLBACK_MONITORINGSIGNAL MonitoringSignal = iota
    UNKNOWNFUTUREVALUE_MONITORINGSIGNAL
)

func (i MonitoringSignal) String() string {
    return []string{"ROLLBACK", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMonitoringSignal(v string) (interface{}, error) {
    result := ROLLBACK_MONITORINGSIGNAL
    switch strings.ToUpper(v) {
        case "ROLLBACK":
            result = ROLLBACK_MONITORINGSIGNAL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MONITORINGSIGNAL
        default:
            return 0, errors.New("Unknown MonitoringSignal value: " + v)
    }
    return &result, nil
}
func SerializeMonitoringSignal(values []MonitoringSignal) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
