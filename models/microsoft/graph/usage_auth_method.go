package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "EMAIL":
            return EMAIL_USAGEAUTHMETHOD, nil
        case "MOBILESMS":
            return MOBILESMS_USAGEAUTHMETHOD, nil
        case "MOBILECALL":
            return MOBILECALL_USAGEAUTHMETHOD, nil
        case "OFFICEPHONE":
            return OFFICEPHONE_USAGEAUTHMETHOD, nil
        case "SECURITYQUESTION":
            return SECURITYQUESTION_USAGEAUTHMETHOD, nil
        case "APPNOTIFICATION":
            return APPNOTIFICATION_USAGEAUTHMETHOD, nil
        case "APPCODE":
            return APPCODE_USAGEAUTHMETHOD, nil
        case "ALTERNATEMOBILECALL":
            return ALTERNATEMOBILECALL_USAGEAUTHMETHOD, nil
        case "FIDO":
            return FIDO_USAGEAUTHMETHOD, nil
        case "APPPASSWORD":
            return APPPASSWORD_USAGEAUTHMETHOD, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_USAGEAUTHMETHOD, nil
    }
    return 0, errors.New("Unknown UsageAuthMethod value: " + v)
}
func SerializeUsageAuthMethod(values []UsageAuthMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
