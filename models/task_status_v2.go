package models
import (
    "errors"
)
// Provides operations to manage the compliance singleton.
type TaskStatus_v2 int

const (
    NOTSTARTED_TASKSTATUS_V2 TaskStatus_v2 = iota
    INPROGRESS_TASKSTATUS_V2
    COMPLETED_TASKSTATUS_V2
    UNKNOWNFUTUREVALUE_TASKSTATUS_V2
)

func (i TaskStatus_v2) String() string {
    return []string{"notStarted", "inProgress", "completed", "unknownFutureValue"}[i]
}
func ParseTaskStatus_v2(v string) (interface{}, error) {
    result := NOTSTARTED_TASKSTATUS_V2
    switch v {
        case "notStarted":
            result = NOTSTARTED_TASKSTATUS_V2
        case "inProgress":
            result = INPROGRESS_TASKSTATUS_V2
        case "completed":
            result = COMPLETED_TASKSTATUS_V2
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TASKSTATUS_V2
        default:
            return 0, errors.New("Unknown TaskStatus_v2 value: " + v)
    }
    return &result, nil
}
func SerializeTaskStatus_v2(values []TaskStatus_v2) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
