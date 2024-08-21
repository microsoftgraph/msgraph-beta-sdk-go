package models
type ApprovalItemType int

const (
    BASIC_APPROVALITEMTYPE ApprovalItemType = iota
    BASICAWAITALL_APPROVALITEMTYPE
    CUSTOM_APPROVALITEMTYPE
    CUSTOMAWAITALL_APPROVALITEMTYPE
    UNKNOWNFUTUREVALUE_APPROVALITEMTYPE
)

func (i ApprovalItemType) String() string {
    return []string{"basic", "basicAwaitAll", "custom", "customAwaitAll", "unknownFutureValue"}[i]
}
func ParseApprovalItemType(v string) (any, error) {
    result := BASIC_APPROVALITEMTYPE
    switch v {
        case "basic":
            result = BASIC_APPROVALITEMTYPE
        case "basicAwaitAll":
            result = BASICAWAITALL_APPROVALITEMTYPE
        case "custom":
            result = CUSTOM_APPROVALITEMTYPE
        case "customAwaitAll":
            result = CUSTOMAWAITALL_APPROVALITEMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPROVALITEMTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApprovalItemType(values []ApprovalItemType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApprovalItemType) isMultiValue() bool {
    return false
}
