package models
// Exchange Active Sync authentication method.
type EasAuthenticationMethod int

const (
    // Authenticate with a username and password.
    USERNAMEANDPASSWORD_EASAUTHENTICATIONMETHOD EasAuthenticationMethod = iota
    // Authenticate with a certificate.
    CERTIFICATE_EASAUTHENTICATIONMETHOD
    // Authenticate with derived credential.
    DERIVEDCREDENTIAL_EASAUTHENTICATIONMETHOD
)

func (i EasAuthenticationMethod) String() string {
    return []string{"usernameAndPassword", "certificate", "derivedCredential"}[i]
}
func ParseEasAuthenticationMethod(v string) (any, error) {
    result := USERNAMEANDPASSWORD_EASAUTHENTICATIONMETHOD
    switch v {
        case "usernameAndPassword":
            result = USERNAMEANDPASSWORD_EASAUTHENTICATIONMETHOD
        case "certificate":
            result = CERTIFICATE_EASAUTHENTICATIONMETHOD
        case "derivedCredential":
            result = DERIVEDCREDENTIAL_EASAUTHENTICATIONMETHOD
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEasAuthenticationMethod(values []EasAuthenticationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EasAuthenticationMethod) isMultiValue() bool {
    return false
}
