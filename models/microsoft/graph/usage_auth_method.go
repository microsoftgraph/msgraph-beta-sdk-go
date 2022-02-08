package graph
import (
    "strings"
    "errors"
)
// 
type UsageAuthMethod int

const (
    EMAIL_USAGEAUTHMETHOD UsageAuthMethod = iota
    MOBILESMS_USAGEAUTHMETHOD
    MOBILECALL_USAGEAUTHMETHOD
    OFFICEPHONE_USAGEAUTHMETHOD
    SECURITYQUESTION_USAGEAUTHMETHOD
    APPNOTIFICATION_USAGEAUTHMETHOD
    APPCODE_USAGEAUTHMETHOD
    ALTERNATEMOBILECALL_USAGEAUTHMETHOD
    FIDO_USAGEAUTHMETHOD
    APPPASSWORD_USAGEAUTHMETHOD
    UNKNOWNFUTUREVALUE_USAGEAUTHMETHOD
)

func (i UsageAuthMethod) String() string {
    return []string{"EMAIL", "MOBILESMS", "MOBILECALL", "OFFICEPHONE", "SECURITYQUESTION", "APPNOTIFICATION", "APPCODE", "ALTERNATEMOBILECALL", "FIDO", "APPPASSWORD", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUsageAuthMethod(v string) (interface{}, error) {
    result := EMAIL_USAGEAUTHMETHOD
    switch strings.ToUpper(v) {
        case "EMAIL":
            result = EMAIL_USAGEAUTHMETHOD
        case "MOBILESMS":
            result = MOBILESMS_USAGEAUTHMETHOD
        case "MOBILECALL":
            result = MOBILECALL_USAGEAUTHMETHOD
        case "OFFICEPHONE":
            result = OFFICEPHONE_USAGEAUTHMETHOD
        case "SECURITYQUESTION":
            result = SECURITYQUESTION_USAGEAUTHMETHOD
        case "APPNOTIFICATION":
            result = APPNOTIFICATION_USAGEAUTHMETHOD
        case "APPCODE":
            result = APPCODE_USAGEAUTHMETHOD
        case "ALTERNATEMOBILECALL":
            result = ALTERNATEMOBILECALL_USAGEAUTHMETHOD
        case "FIDO":
            result = FIDO_USAGEAUTHMETHOD
        case "APPPASSWORD":
            result = APPPASSWORD_USAGEAUTHMETHOD
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_USAGEAUTHMETHOD
        default:
            return 0, errors.New("Unknown UsageAuthMethod value: " + v)
    }
    return &result, nil
}
func SerializeUsageAuthMethod(values []UsageAuthMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
