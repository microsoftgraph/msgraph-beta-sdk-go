package models
// Non-EAP methods for authentication.
type NonEapAuthenticationMethodForEapTtlsType int

const (
    // Unencrypted password (PAP).
    UNENCRYPTEDPASSWORD_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE NonEapAuthenticationMethodForEapTtlsType = iota
    // Challenge Handshake Authentication Protocol (CHAP).
    CHALLENGEHANDSHAKEAUTHENTICATIONPROTOCOL_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
    //  Microsoft CHAP (MS-CHAP).
    MICROSOFTCHAP_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
    // Microsoft CHAP Version 2 (MS-CHAP v2).
    MICROSOFTCHAPVERSIONTWO_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
)

func (i NonEapAuthenticationMethodForEapTtlsType) String() string {
    return []string{"unencryptedPassword", "challengeHandshakeAuthenticationProtocol", "microsoftChap", "microsoftChapVersionTwo"}[i]
}
func ParseNonEapAuthenticationMethodForEapTtlsType(v string) (any, error) {
    result := UNENCRYPTEDPASSWORD_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
    switch v {
        case "unencryptedPassword":
            result = UNENCRYPTEDPASSWORD_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
        case "challengeHandshakeAuthenticationProtocol":
            result = CHALLENGEHANDSHAKEAUTHENTICATIONPROTOCOL_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
        case "microsoftChap":
            result = MICROSOFTCHAP_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
        case "microsoftChapVersionTwo":
            result = MICROSOFTCHAPVERSIONTWO_NONEAPAUTHENTICATIONMETHODFOREAPTTLSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNonEapAuthenticationMethodForEapTtlsType(values []NonEapAuthenticationMethodForEapTtlsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NonEapAuthenticationMethodForEapTtlsType) isMultiValue() bool {
    return false
}
