package ediscovery
import (
    "strings"
    "errors"
)
// 
type CaseOperationStatus int

const (
    NOTSTARTED_CASEOPERATIONSTATUS CaseOperationStatus = iota
    SUBMISSIONFAILED_CASEOPERATIONSTATUS
    RUNNING_CASEOPERATIONSTATUS
    SUCCEEDED_CASEOPERATIONSTATUS
    PARTIALLYSUCCEEDED_CASEOPERATIONSTATUS
    FAILED_CASEOPERATIONSTATUS
)

func (i CaseOperationStatus) String() string {
    return []string{"NOTSTARTED", "SUBMISSIONFAILED", "RUNNING", "SUCCEEDED", "PARTIALLYSUCCEEDED", "FAILED"}[i]
}
func ParseCaseOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_CASEOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_CASEOPERATIONSTATUS
        case "SUBMISSIONFAILED":
            result = SUBMISSIONFAILED_CASEOPERATIONSTATUS
        case "RUNNING":
            result = RUNNING_CASEOPERATIONSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_CASEOPERATIONSTATUS
        case "PARTIALLYSUCCEEDED":
            result = PARTIALLYSUCCEEDED_CASEOPERATIONSTATUS
        case "FAILED":
            result = FAILED_CASEOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown CaseOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeCaseOperationStatus(values []CaseOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
