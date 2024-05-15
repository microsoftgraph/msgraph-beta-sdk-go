package models
type AuthorizationSystemActionType int

const (
    DELETE_AUTHORIZATIONSYSTEMACTIONTYPE AuthorizationSystemActionType = iota
    READ_AUTHORIZATIONSYSTEMACTIONTYPE
    UNKNOWNFUTUREVALUE_AUTHORIZATIONSYSTEMACTIONTYPE
)

func (i AuthorizationSystemActionType) String() string {
    return []string{"delete", "read", "unknownFutureValue"}[i]
}
func ParseAuthorizationSystemActionType(v string) (any, error) {
    result := DELETE_AUTHORIZATIONSYSTEMACTIONTYPE
    switch v {
        case "delete":
            result = DELETE_AUTHORIZATIONSYSTEMACTIONTYPE
        case "read":
            result = READ_AUTHORIZATIONSYSTEMACTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHORIZATIONSYSTEMACTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthorizationSystemActionType(values []AuthorizationSystemActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthorizationSystemActionType) isMultiValue() bool {
    return false
}
