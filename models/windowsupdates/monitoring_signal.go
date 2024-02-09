package windowsupdates
import (
    "errors"
)
type MonitoringSignal int

const (
    ROLLBACK_MONITORINGSIGNAL MonitoringSignal = iota
    INELIGIBLE_MONITORINGSIGNAL
    UNKNOWNFUTUREVALUE_MONITORINGSIGNAL
)

func (i MonitoringSignal) String() string {
    return []string{"rollback", "ineligible", "unknownFutureValue"}[i]
}
func ParseMonitoringSignal(v string) (any, error) {
    result := ROLLBACK_MONITORINGSIGNAL
    switch v {
        case "rollback":
            result = ROLLBACK_MONITORINGSIGNAL
        case "ineligible":
            result = INELIGIBLE_MONITORINGSIGNAL
        case "unknownFutureValue":
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
func (i MonitoringSignal) isMultiValue() bool {
    return false
}
