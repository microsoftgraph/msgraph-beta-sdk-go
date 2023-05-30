package models
import (
    "errors"
)
// 
type UserDefaultAuthenticationMethodType int

const (
    PUSH_USERDEFAULTAUTHENTICATIONMETHODTYPE UserDefaultAuthenticationMethodType = iota
    OATH_USERDEFAULTAUTHENTICATIONMETHODTYPE
    VOICEMOBILE_USERDEFAULTAUTHENTICATIONMETHODTYPE
    VOICEALTERNATEMOBILE_USERDEFAULTAUTHENTICATIONMETHODTYPE
    VOICEOFFICE_USERDEFAULTAUTHENTICATIONMETHODTYPE
    SMS_USERDEFAULTAUTHENTICATIONMETHODTYPE
    UNKNOWNFUTUREVALUE_USERDEFAULTAUTHENTICATIONMETHODTYPE
)

func (i UserDefaultAuthenticationMethodType) String() string {
    return []string{"push", "oath", "voiceMobile", "voiceAlternateMobile", "voiceOffice", "sms", "unknownFutureValue"}[i]
}
func ParseUserDefaultAuthenticationMethodType(v string) (any, error) {
    result := PUSH_USERDEFAULTAUTHENTICATIONMETHODTYPE
    switch v {
        case "push":
            result = PUSH_USERDEFAULTAUTHENTICATIONMETHODTYPE
        case "oath":
            result = OATH_USERDEFAULTAUTHENTICATIONMETHODTYPE
        case "voiceMobile":
            result = VOICEMOBILE_USERDEFAULTAUTHENTICATIONMETHODTYPE
        case "voiceAlternateMobile":
            result = VOICEALTERNATEMOBILE_USERDEFAULTAUTHENTICATIONMETHODTYPE
        case "voiceOffice":
            result = VOICEOFFICE_USERDEFAULTAUTHENTICATIONMETHODTYPE
        case "sms":
            result = SMS_USERDEFAULTAUTHENTICATIONMETHODTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERDEFAULTAUTHENTICATIONMETHODTYPE
        default:
            return 0, errors.New("Unknown UserDefaultAuthenticationMethodType value: " + v)
    }
    return &result, nil
}
func SerializeUserDefaultAuthenticationMethodType(values []UserDefaultAuthenticationMethodType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
