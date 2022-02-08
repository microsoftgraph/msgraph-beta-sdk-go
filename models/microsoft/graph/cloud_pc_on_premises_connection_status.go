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
    result := PENDING_CLOUDPCONPREMISESCONNECTIONSTATUS
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "RUNNING":
            result = RUNNING_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "PASSED":
            result = PASSED_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "FAILED":
            result = FAILED_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "WARNING":
            result = WARNING_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONSTATUS
        default:
            return 0, errors.New("Unknown CloudPcOnPremisesConnectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcOnPremisesConnectionStatus(values []CloudPcOnPremisesConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
