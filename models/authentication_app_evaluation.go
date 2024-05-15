package models
type AuthenticationAppEvaluation int

const (
    SUCCESS_AUTHENTICATIONAPPEVALUATION AuthenticationAppEvaluation = iota
    FAILURE_AUTHENTICATIONAPPEVALUATION
    UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPEVALUATION
)

func (i AuthenticationAppEvaluation) String() string {
    return []string{"success", "failure", "unknownFutureValue"}[i]
}
func ParseAuthenticationAppEvaluation(v string) (any, error) {
    result := SUCCESS_AUTHENTICATIONAPPEVALUATION
    switch v {
        case "success":
            result = SUCCESS_AUTHENTICATIONAPPEVALUATION
        case "failure":
            result = FAILURE_AUTHENTICATIONAPPEVALUATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPEVALUATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthenticationAppEvaluation(values []AuthenticationAppEvaluation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationAppEvaluation) isMultiValue() bool {
    return false
}
