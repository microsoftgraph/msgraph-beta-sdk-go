package models
type AuthenticationFailureReasonCode int

const (
    INCOMPLETE_AUTHENTICATIONFAILUREREASONCODE AuthenticationFailureReasonCode = iota
    DENIED_AUTHENTICATIONFAILUREREASONCODE
    SYSTEMFAILURE_AUTHENTICATIONFAILUREREASONCODE
    BADREQUEST_AUTHENTICATIONFAILUREREASONCODE
    OTHER_AUTHENTICATIONFAILUREREASONCODE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONFAILUREREASONCODE
    USERERROR_AUTHENTICATIONFAILUREREASONCODE
    CONFIGERROR_AUTHENTICATIONFAILUREREASONCODE
)

func (i AuthenticationFailureReasonCode) String() string {
    return []string{"incomplete", "denied", "systemFailure", "badRequest", "other", "unknownFutureValue", "userError", "configError"}[i]
}
func ParseAuthenticationFailureReasonCode(v string) (any, error) {
    result := INCOMPLETE_AUTHENTICATIONFAILUREREASONCODE
    switch v {
        case "incomplete":
            result = INCOMPLETE_AUTHENTICATIONFAILUREREASONCODE
        case "denied":
            result = DENIED_AUTHENTICATIONFAILUREREASONCODE
        case "systemFailure":
            result = SYSTEMFAILURE_AUTHENTICATIONFAILUREREASONCODE
        case "badRequest":
            result = BADREQUEST_AUTHENTICATIONFAILUREREASONCODE
        case "other":
            result = OTHER_AUTHENTICATIONFAILUREREASONCODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONFAILUREREASONCODE
        case "userError":
            result = USERERROR_AUTHENTICATIONFAILUREREASONCODE
        case "configError":
            result = CONFIGERROR_AUTHENTICATIONFAILUREREASONCODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthenticationFailureReasonCode(values []AuthenticationFailureReasonCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationFailureReasonCode) isMultiValue() bool {
    return false
}
