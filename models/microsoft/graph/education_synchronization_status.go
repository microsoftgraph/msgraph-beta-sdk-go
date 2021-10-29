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
    switch strings.ToUpper(v) {
        case "PAUSED":
            return PAUSED_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "SUCCESS":
            return SUCCESS_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "ERROR":
            return ERROR_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "VALIDATIONERROR":
            return VALIDATIONERROR_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "QUARANTINED":
            return QUARANTINED_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "EXTRACTING":
            return EXTRACTING_EDUCATIONSYNCHRONIZATIONSTATUS, nil
        case "VALIDATING":
            return VALIDATING_EDUCATIONSYNCHRONIZATIONSTATUS, nil
    }
    return 0, errors.New("Unknown EducationSynchronizationStatus value: " + v)
}
func SerializeEducationSynchronizationStatus(values []EducationSynchronizationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
