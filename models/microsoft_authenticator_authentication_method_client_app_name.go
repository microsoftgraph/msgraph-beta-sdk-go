package models
type MicrosoftAuthenticatorAuthenticationMethodClientAppName int

const (
    MICROSOFTAUTHENTICATOR_MICROSOFTAUTHENTICATORAUTHENTICATIONMETHODCLIENTAPPNAME MicrosoftAuthenticatorAuthenticationMethodClientAppName = iota
    OUTLOOKMOBILE_MICROSOFTAUTHENTICATORAUTHENTICATIONMETHODCLIENTAPPNAME
    UNKNOWNFUTUREVALUE_MICROSOFTAUTHENTICATORAUTHENTICATIONMETHODCLIENTAPPNAME
)

func (i MicrosoftAuthenticatorAuthenticationMethodClientAppName) String() string {
    return []string{"microsoftAuthenticator", "outlookMobile", "unknownFutureValue"}[i]
}
func ParseMicrosoftAuthenticatorAuthenticationMethodClientAppName(v string) (any, error) {
    result := MICROSOFTAUTHENTICATOR_MICROSOFTAUTHENTICATORAUTHENTICATIONMETHODCLIENTAPPNAME
    switch v {
        case "microsoftAuthenticator":
            result = MICROSOFTAUTHENTICATOR_MICROSOFTAUTHENTICATORAUTHENTICATIONMETHODCLIENTAPPNAME
        case "outlookMobile":
            result = OUTLOOKMOBILE_MICROSOFTAUTHENTICATORAUTHENTICATIONMETHODCLIENTAPPNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MICROSOFTAUTHENTICATORAUTHENTICATIONMETHODCLIENTAPPNAME
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMicrosoftAuthenticatorAuthenticationMethodClientAppName(values []MicrosoftAuthenticatorAuthenticationMethodClientAppName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MicrosoftAuthenticatorAuthenticationMethodClientAppName) isMultiValue() bool {
    return false
}
