package graph
import (
    "strings"
    "errors"
)
// 
type TaskStatus_v2 int

const (
    NOTSTARTED_TASKSTATUS_V2 TaskStatus_v2 = iota
    INPROGRESS_TASKSTATUS_V2
    COMPLETED_TASKSTATUS_V2
    UNKNOWNFUTUREVALUE_TASKSTATUS_V2
)

func (i TaskStatus_v2) String() string {
    return []string{"NOTSTARTED", "INPROGRESS", "COMPLETED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTaskStatus_v2(v string) (interface{}, error) {
    result := NOTSTARTED_TASKSTATUS_V2
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_TASKSTATUS_V2
        case "INPROGRESS":
            result = INPROGRESS_TASKSTATUS_V2
        case "COMPLETED":
            result = COMPLETED_TASKSTATUS_V2
        case "UNKNOWNFUTUREVALUE":
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
