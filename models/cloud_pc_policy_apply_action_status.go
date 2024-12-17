package models
type CloudPcPolicyApplyActionStatus int

const (
    PROCESSING_CLOUDPCPOLICYAPPLYACTIONSTATUS CloudPcPolicyApplyActionStatus = iota
    SUCCEEDED_CLOUDPCPOLICYAPPLYACTIONSTATUS
    FAILED_CLOUDPCPOLICYAPPLYACTIONSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCPOLICYAPPLYACTIONSTATUS
)

func (i CloudPcPolicyApplyActionStatus) String() string {
    return []string{"processing", "succeeded", "failed", "unknownFutureValue"}[i]
}
func ParseCloudPcPolicyApplyActionStatus(v string) (any, error) {
    result := PROCESSING_CLOUDPCPOLICYAPPLYACTIONSTATUS
    switch v {
        case "processing":
            result = PROCESSING_CLOUDPCPOLICYAPPLYACTIONSTATUS
        case "succeeded":
            result = SUCCEEDED_CLOUDPCPOLICYAPPLYACTIONSTATUS
        case "failed":
            result = FAILED_CLOUDPCPOLICYAPPLYACTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPOLICYAPPLYACTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcPolicyApplyActionStatus(values []CloudPcPolicyApplyActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcPolicyApplyActionStatus) isMultiValue() bool {
    return false
}
