package models
type StatusDetail int

const (
    SUBMITTED_STATUSDETAIL StatusDetail = iota
    APPROVED_STATUSDETAIL
    COMPLETED_STATUSDETAIL
    CANCELED_STATUSDETAIL
    REJECTED_STATUSDETAIL
    UNKNOWNFUTUREVALUE_STATUSDETAIL
)

func (i StatusDetail) String() string {
    return []string{"submitted", "approved", "completed", "canceled", "rejected", "unknownFutureValue"}[i]
}
func ParseStatusDetail(v string) (any, error) {
    result := SUBMITTED_STATUSDETAIL
    switch v {
        case "submitted":
            result = SUBMITTED_STATUSDETAIL
        case "approved":
            result = APPROVED_STATUSDETAIL
        case "completed":
            result = COMPLETED_STATUSDETAIL
        case "canceled":
            result = CANCELED_STATUSDETAIL
        case "rejected":
            result = REJECTED_STATUSDETAIL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_STATUSDETAIL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStatusDetail(values []StatusDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StatusDetail) isMultiValue() bool {
    return false
}
