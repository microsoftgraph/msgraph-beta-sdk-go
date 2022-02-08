package ediscovery
import (
    "strings"
    "errors"
)
// 
type CaseStatus int

const (
    UNKNOWN_CASESTATUS CaseStatus = iota
    ACTIVE_CASESTATUS
    PENDINGDELETE_CASESTATUS
    CLOSING_CASESTATUS
    CLOSED_CASESTATUS
    CLOSEDWITHERROR_CASESTATUS
)

func (i CaseStatus) String() string {
    return []string{"UNKNOWN", "ACTIVE", "PENDINGDELETE", "CLOSING", "CLOSED", "CLOSEDWITHERROR"}[i]
}
func ParseCaseStatus(v string) (interface{}, error) {
    result := UNKNOWN_CASESTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_CASESTATUS
        case "ACTIVE":
            result = ACTIVE_CASESTATUS
        case "PENDINGDELETE":
            result = PENDINGDELETE_CASESTATUS
        case "CLOSING":
            result = CLOSING_CASESTATUS
        case "CLOSED":
            result = CLOSED_CASESTATUS
        case "CLOSEDWITHERROR":
            result = CLOSEDWITHERROR_CASESTATUS
        default:
            return 0, errors.New("Unknown CaseStatus value: " + v)
    }
    return &result, nil
}
func SerializeCaseStatus(values []CaseStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
