package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "EMAIL":
            return EMAIL_REGISTRATIONAUTHMETHOD, nil
        case "MOBILEPHONE":
            return MOBILEPHONE_REGISTRATIONAUTHMETHOD, nil
        case "OFFICEPHONE":
            return OFFICEPHONE_REGISTRATIONAUTHMETHOD, nil
        case "SECURITYQUESTION":
            return SECURITYQUESTION_REGISTRATIONAUTHMETHOD, nil
        case "APPNOTIFICATION":
            return APPNOTIFICATION_REGISTRATIONAUTHMETHOD, nil
        case "APPCODE":
            return APPCODE_REGISTRATIONAUTHMETHOD, nil
        case "ALTERNATEMOBILEPHONE":
            return ALTERNATEMOBILEPHONE_REGISTRATIONAUTHMETHOD, nil
        case "FIDO":
            return FIDO_REGISTRATIONAUTHMETHOD, nil
        case "APPPASSWORD":
            return APPPASSWORD_REGISTRATIONAUTHMETHOD, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REGISTRATIONAUTHMETHOD, nil
    }
    return 0, errors.New("Unknown RegistrationAuthMethod value: " + v)
}
func SerializeRegistrationAuthMethod(values []RegistrationAuthMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
