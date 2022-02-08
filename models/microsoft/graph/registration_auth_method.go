package graph
import (
    "strings"
    "errors"
)
// 
type RegistrationAuthMethod int

const (
    EMAIL_REGISTRATIONAUTHMETHOD RegistrationAuthMethod = iota
    MOBILEPHONE_REGISTRATIONAUTHMETHOD
    OFFICEPHONE_REGISTRATIONAUTHMETHOD
    SECURITYQUESTION_REGISTRATIONAUTHMETHOD
    APPNOTIFICATION_REGISTRATIONAUTHMETHOD
    APPCODE_REGISTRATIONAUTHMETHOD
    ALTERNATEMOBILEPHONE_REGISTRATIONAUTHMETHOD
    FIDO_REGISTRATIONAUTHMETHOD
    APPPASSWORD_REGISTRATIONAUTHMETHOD
    UNKNOWNFUTUREVALUE_REGISTRATIONAUTHMETHOD
)

func (i RegistrationAuthMethod) String() string {
    return []string{"EMAIL", "MOBILEPHONE", "OFFICEPHONE", "SECURITYQUESTION", "APPNOTIFICATION", "APPCODE", "ALTERNATEMOBILEPHONE", "FIDO", "APPPASSWORD", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRegistrationAuthMethod(v string) (interface{}, error) {
    result := EMAIL_REGISTRATIONAUTHMETHOD
    switch strings.ToUpper(v) {
        case "EMAIL":
            result = EMAIL_REGISTRATIONAUTHMETHOD
        case "MOBILEPHONE":
            result = MOBILEPHONE_REGISTRATIONAUTHMETHOD
        case "OFFICEPHONE":
            result = OFFICEPHONE_REGISTRATIONAUTHMETHOD
        case "SECURITYQUESTION":
            result = SECURITYQUESTION_REGISTRATIONAUTHMETHOD
        case "APPNOTIFICATION":
            result = APPNOTIFICATION_REGISTRATIONAUTHMETHOD
        case "APPCODE":
            result = APPCODE_REGISTRATIONAUTHMETHOD
        case "ALTERNATEMOBILEPHONE":
            result = ALTERNATEMOBILEPHONE_REGISTRATIONAUTHMETHOD
        case "FIDO":
            result = FIDO_REGISTRATIONAUTHMETHOD
        case "APPPASSWORD":
            result = APPPASSWORD_REGISTRATIONAUTHMETHOD
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REGISTRATIONAUTHMETHOD
        default:
            return 0, errors.New("Unknown RegistrationAuthMethod value: " + v)
    }
    return &result, nil
}
func SerializeRegistrationAuthMethod(values []RegistrationAuthMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
