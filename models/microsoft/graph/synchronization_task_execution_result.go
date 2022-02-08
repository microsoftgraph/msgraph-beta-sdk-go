package graph
import (
    "strings"
    "errors"
)
// 
type SynchronizationTaskExecutionResult int

const (
    SUCCEEDED_SYNCHRONIZATIONTASKEXECUTIONRESULT SynchronizationTaskExecutionResult = iota
    FAILED_SYNCHRONIZATIONTASKEXECUTIONRESULT
    ENTRYLEVELERRORS_SYNCHRONIZATIONTASKEXECUTIONRESULT
)

func (i SynchronizationTaskExecutionResult) String() string {
    return []string{"SUCCEEDED", "FAILED", "ENTRYLEVELERRORS"}[i]
}
func ParseSynchronizationTaskExecutionResult(v string) (interface{}, error) {
    result := SUCCEEDED_SYNCHRONIZATIONTASKEXECUTIONRESULT
    switch strings.ToUpper(v) {
        case "SUCCEEDED":
            result = SUCCEEDED_SYNCHRONIZATIONTASKEXECUTIONRESULT
        case "FAILED":
            result = FAILED_SYNCHRONIZATIONTASKEXECUTIONRESULT
        case "ENTRYLEVELERRORS":
            result = ENTRYLEVELERRORS_SYNCHRONIZATIONTASKEXECUTIONRESULT
        default:
            return 0, errors.New("Unknown SynchronizationTaskExecutionResult value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationTaskExecutionResult(values []SynchronizationTaskExecutionResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
