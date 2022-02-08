package security
import (
    "strings"
    "errors"
)
// 
type IncidentStatus int

const (
    ACTIVE_INCIDENTSTATUS IncidentStatus = iota
    RESOLVED_INCIDENTSTATUS
    REDIRECTED_INCIDENTSTATUS
    UNKNOWNFUTUREVALUE_INCIDENTSTATUS
)

func (i IncidentStatus) String() string {
    return []string{"ACTIVE", "RESOLVED", "REDIRECTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIncidentStatus(v string) (interface{}, error) {
    result := ACTIVE_INCIDENTSTATUS
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_INCIDENTSTATUS
        case "RESOLVED":
            result = RESOLVED_INCIDENTSTATUS
        case "REDIRECTED":
            result = REDIRECTED_INCIDENTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_INCIDENTSTATUS
        default:
            return 0, errors.New("Unknown IncidentStatus value: " + v)
    }
    return &result, nil
}
func SerializeIncidentStatus(values []IncidentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
