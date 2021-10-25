package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SUCCEEDED":
            return SUCCEEDED_SYNCHRONIZATIONTASKEXECUTIONRESULT, nil
        case "FAILED":
            return FAILED_SYNCHRONIZATIONTASKEXECUTIONRESULT, nil
        case "ENTRYLEVELERRORS":
            return ENTRYLEVELERRORS_SYNCHRONIZATIONTASKEXECUTIONRESULT, nil
    }
    return 0, errors.New("Unknown SynchronizationTaskExecutionResult value: " + v)
}
func SerializeSynchronizationTaskExecutionResult(values []SynchronizationTaskExecutionResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
