package security
type PolicyStatus int

const (
    PENDING_POLICYSTATUS PolicyStatus = iota
    ERROR_POLICYSTATUS
    SUCCESS_POLICYSTATUS
    UNKNOWNFUTUREVALUE_POLICYSTATUS
)

func (i PolicyStatus) String() string {
    return []string{"pending", "error", "success", "unknownFutureValue"}[i]
}
func ParsePolicyStatus(v string) (any, error) {
    result := PENDING_POLICYSTATUS
    switch v {
        case "pending":
            result = PENDING_POLICYSTATUS
        case "error":
            result = ERROR_POLICYSTATUS
        case "success":
            result = SUCCESS_POLICYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_POLICYSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePolicyStatus(values []PolicyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PolicyStatus) isMultiValue() bool {
    return false
}
