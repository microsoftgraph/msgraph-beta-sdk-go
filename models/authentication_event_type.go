package models
type AuthenticationEventType int

const (
    TOKENISSUANCESTART_AUTHENTICATIONEVENTTYPE AuthenticationEventType = iota
    PAGERENDERSTART_AUTHENTICATIONEVENTTYPE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONEVENTTYPE
    ATTRIBUTECOLLECTIONSTART_AUTHENTICATIONEVENTTYPE
    ATTRIBUTECOLLECTIONSUBMIT_AUTHENTICATIONEVENTTYPE
    EMAILOTPSEND_AUTHENTICATIONEVENTTYPE
)

func (i AuthenticationEventType) String() string {
    return []string{"tokenIssuanceStart", "pageRenderStart", "unknownFutureValue", "attributeCollectionStart", "attributeCollectionSubmit", "emailOtpSend"}[i]
}
func ParseAuthenticationEventType(v string) (any, error) {
    result := TOKENISSUANCESTART_AUTHENTICATIONEVENTTYPE
    switch v {
        case "tokenIssuanceStart":
            result = TOKENISSUANCESTART_AUTHENTICATIONEVENTTYPE
        case "pageRenderStart":
            result = PAGERENDERSTART_AUTHENTICATIONEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONEVENTTYPE
        case "attributeCollectionStart":
            result = ATTRIBUTECOLLECTIONSTART_AUTHENTICATIONEVENTTYPE
        case "attributeCollectionSubmit":
            result = ATTRIBUTECOLLECTIONSUBMIT_AUTHENTICATIONEVENTTYPE
        case "emailOtpSend":
            result = EMAILOTPSEND_AUTHENTICATIONEVENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthenticationEventType(values []AuthenticationEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationEventType) isMultiValue() bool {
    return false
}
