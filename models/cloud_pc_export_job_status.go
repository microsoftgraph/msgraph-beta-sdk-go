package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type CloudPcExportJobStatus int

const (
    NOTSTARTED_CLOUDPCEXPORTJOBSTATUS CloudPcExportJobStatus = iota
    INPROGRESS_CLOUDPCEXPORTJOBSTATUS
    COMPLETED_CLOUDPCEXPORTJOBSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCEXPORTJOBSTATUS
)

func (i CloudPcExportJobStatus) String() string {
    return []string{"notStarted", "inProgress", "completed", "unknownFutureValue"}[i]
}
func ParseCloudPcExportJobStatus(v string) (interface{}, error) {
    result := NOTSTARTED_CLOUDPCEXPORTJOBSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_CLOUDPCEXPORTJOBSTATUS
        case "inProgress":
            result = INPROGRESS_CLOUDPCEXPORTJOBSTATUS
        case "completed":
            result = COMPLETED_CLOUDPCEXPORTJOBSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCEXPORTJOBSTATUS
        default:
            return 0, errors.New("Unknown CloudPcExportJobStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcExportJobStatus(values []CloudPcExportJobStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
