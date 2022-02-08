package graph
import (
    "strings"
    "errors"
)
// 
type EducationSynchronizationStatus int

const (
    PAUSED_EDUCATIONSYNCHRONIZATIONSTATUS EducationSynchronizationStatus = iota
    INPROGRESS_EDUCATIONSYNCHRONIZATIONSTATUS
    SUCCESS_EDUCATIONSYNCHRONIZATIONSTATUS
    ERROR_EDUCATIONSYNCHRONIZATIONSTATUS
    VALIDATIONERROR_EDUCATIONSYNCHRONIZATIONSTATUS
    QUARANTINED_EDUCATIONSYNCHRONIZATIONSTATUS
    UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONSTATUS
    EXTRACTING_EDUCATIONSYNCHRONIZATIONSTATUS
    VALIDATING_EDUCATIONSYNCHRONIZATIONSTATUS
)

func (i EducationSynchronizationStatus) String() string {
    return []string{"PAUSED", "INPROGRESS", "SUCCESS", "ERROR", "VALIDATIONERROR", "QUARANTINED", "UNKNOWNFUTUREVALUE", "EXTRACTING", "VALIDATING"}[i]
}
func ParseEducationSynchronizationStatus(v string) (interface{}, error) {
    result := PAUSED_EDUCATIONSYNCHRONIZATIONSTATUS
    switch strings.ToUpper(v) {
        case "PAUSED":
            result = PAUSED_EDUCATIONSYNCHRONIZATIONSTATUS
        case "INPROGRESS":
            result = INPROGRESS_EDUCATIONSYNCHRONIZATIONSTATUS
        case "SUCCESS":
            result = SUCCESS_EDUCATIONSYNCHRONIZATIONSTATUS
        case "ERROR":
            result = ERROR_EDUCATIONSYNCHRONIZATIONSTATUS
        case "VALIDATIONERROR":
            result = VALIDATIONERROR_EDUCATIONSYNCHRONIZATIONSTATUS
        case "QUARANTINED":
            result = QUARANTINED_EDUCATIONSYNCHRONIZATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONSTATUS
        case "EXTRACTING":
            result = EXTRACTING_EDUCATIONSYNCHRONIZATIONSTATUS
        case "VALIDATING":
            result = VALIDATING_EDUCATIONSYNCHRONIZATIONSTATUS
        default:
            return 0, errors.New("Unknown EducationSynchronizationStatus value: " + v)
    }
    return &result, nil
}
func SerializeEducationSynchronizationStatus(values []EducationSynchronizationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
