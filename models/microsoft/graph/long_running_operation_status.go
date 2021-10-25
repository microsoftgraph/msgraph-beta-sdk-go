package graph
import (
    "strings"
    "errors"
)
type LongRunningOperationStatus int

const (
    NOTSTARTED_LONGRUNNINGOPERATIONSTATUS LongRunningOperationStatus = iota
    RUNNING_LONGRUNNINGOPERATIONSTATUS
    SUCCEEDED_LONGRUNNINGOPERATIONSTATUS
    FAILED_LONGRUNNINGOPERATIONSTATUS
)

func (i LongRunningOperationStatus) String() string {
    return []string{"NOTSTARTED", "RUNNING", "SUCCEEDED", "FAILED"}[i]
}
func ParseLongRunningOperationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_LONGRUNNINGOPERATIONSTATUS, nil
        case "RUNNING":
            return RUNNING_LONGRUNNINGOPERATIONSTATUS, nil
        case "SUCCEEDED":
            return SUCCEEDED_LONGRUNNINGOPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_LONGRUNNINGOPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown LongRunningOperationStatus value: " + v)
}
func SerializeLongRunningOperationStatus(values []LongRunningOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
