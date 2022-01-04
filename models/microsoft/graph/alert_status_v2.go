package graph
import (
    "strings"
    "errors"
)
// 
type AlertStatus_v2 int

const (
    UNKNOWN_ALERTSTATUS_V2 AlertStatus_v2 = iota
    NEW_ALERTSTATUS_V2
    INPROGRESS_ALERTSTATUS_V2
    RESOLVED_ALERTSTATUS_V2
    UNKNOWNFUTUREVALUE_ALERTSTATUS_V2
)

func (i AlertStatus_v2) String() string {
    return []string{"UNKNOWN", "NEW", "INPROGRESS", "RESOLVED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAlertStatus_v2(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ALERTSTATUS_V2, nil
        case "NEW":
            return NEW_ALERTSTATUS_V2, nil
        case "INPROGRESS":
            return INPROGRESS_ALERTSTATUS_V2, nil
        case "RESOLVED":
            return RESOLVED_ALERTSTATUS_V2, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALERTSTATUS_V2, nil
    }
    return 0, errors.New("Unknown AlertStatus_v2 value: " + v)
}
func SerializeAlertStatus_v2(values []AlertStatus_v2) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
