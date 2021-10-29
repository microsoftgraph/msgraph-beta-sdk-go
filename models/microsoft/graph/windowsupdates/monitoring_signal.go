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
    switch strings.ToUpper(v) {
        case "ROLLBACK":
            return ROLLBACK_MONITORINGSIGNAL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MONITORINGSIGNAL, nil
    }
    return 0, errors.New("Unknown MonitoringSignal value: " + v)
}
func SerializeMonitoringSignal(values []MonitoringSignal) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
