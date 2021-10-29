package graph
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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_INCIDENTSTATUS, nil
        case "RESOLVED":
            return RESOLVED_INCIDENTSTATUS, nil
        case "REDIRECTED":
            return REDIRECTED_INCIDENTSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_INCIDENTSTATUS, nil
    }
    return 0, errors.New("Unknown IncidentStatus value: " + v)
}
func SerializeIncidentStatus(values []IncidentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
