package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcOnPremisesConnectionStatus int

const (
    PENDING_CLOUDPCONPREMISESCONNECTIONSTATUS CloudPcOnPremisesConnectionStatus = iota
    RUNNING_CLOUDPCONPREMISESCONNECTIONSTATUS
    PASSED_CLOUDPCONPREMISESCONNECTIONSTATUS
    FAILED_CLOUDPCONPREMISESCONNECTIONSTATUS
    WARNING_CLOUDPCONPREMISESCONNECTIONSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONSTATUS
)

func (i CloudPcOnPremisesConnectionStatus) String() string {
    return []string{"PENDING", "RUNNING", "PASSED", "FAILED", "WARNING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcOnPremisesConnectionStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_CLOUDPCONPREMISESCONNECTIONSTATUS, nil
        case "RUNNING":
            return RUNNING_CLOUDPCONPREMISESCONNECTIONSTATUS, nil
        case "PASSED":
            return PASSED_CLOUDPCONPREMISESCONNECTIONSTATUS, nil
        case "FAILED":
            return FAILED_CLOUDPCONPREMISESCONNECTIONSTATUS, nil
        case "WARNING":
            return WARNING_CLOUDPCONPREMISESCONNECTIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONSTATUS, nil
    }
    return 0, errors.New("Unknown CloudPcOnPremisesConnectionStatus value: " + v)
}
func SerializeCloudPcOnPremisesConnectionStatus(values []CloudPcOnPremisesConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
