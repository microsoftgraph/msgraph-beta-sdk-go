package models
type CloudPcBulkActionStatus int

const (
    PENDING_CLOUDPCBULKACTIONSTATUS CloudPcBulkActionStatus = iota
    SUCCEEDED_CLOUDPCBULKACTIONSTATUS
    FAILED_CLOUDPCBULKACTIONSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCBULKACTIONSTATUS
)

func (i CloudPcBulkActionStatus) String() string {
    return []string{"pending", "succeeded", "failed", "unknownFutureValue"}[i]
}
func ParseCloudPcBulkActionStatus(v string) (any, error) {
    result := PENDING_CLOUDPCBULKACTIONSTATUS
    switch v {
        case "pending":
            result = PENDING_CLOUDPCBULKACTIONSTATUS
        case "succeeded":
            result = SUCCEEDED_CLOUDPCBULKACTIONSTATUS
        case "failed":
            result = FAILED_CLOUDPCBULKACTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCBULKACTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcBulkActionStatus(values []CloudPcBulkActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcBulkActionStatus) isMultiValue() bool {
    return false
}
