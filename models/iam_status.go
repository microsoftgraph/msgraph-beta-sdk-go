package models
type IamStatus int

const (
    ACTIVE_IAMSTATUS IamStatus = iota
    INACTIVE_IAMSTATUS
    DISABLED_IAMSTATUS
    UNKNOWNFUTUREVALUE_IAMSTATUS
)

func (i IamStatus) String() string {
    return []string{"active", "inactive", "disabled", "unknownFutureValue"}[i]
}
func ParseIamStatus(v string) (any, error) {
    result := ACTIVE_IAMSTATUS
    switch v {
        case "active":
            result = ACTIVE_IAMSTATUS
        case "inactive":
            result = INACTIVE_IAMSTATUS
        case "disabled":
            result = DISABLED_IAMSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IAMSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIamStatus(values []IamStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IamStatus) isMultiValue() bool {
    return false
}
