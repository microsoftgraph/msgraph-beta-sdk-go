package models
import (
    "errors"
)
type AppKeyCredentialRestrictionType int

const (
    ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE AppKeyCredentialRestrictionType = iota
    TRUSTEDCERTIFICATEAUTHORITY_APPKEYCREDENTIALRESTRICTIONTYPE
    UNKNOWNFUTUREVALUE_APPKEYCREDENTIALRESTRICTIONTYPE
)

func (i AppKeyCredentialRestrictionType) String() string {
    return []string{"asymmetricKeyLifetime", "trustedCertificateAuthority", "unknownFutureValue"}[i]
}
func ParseAppKeyCredentialRestrictionType(v string) (any, error) {
    result := ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE
    switch v {
        case "asymmetricKeyLifetime":
            result = ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE
        case "trustedCertificateAuthority":
            result = TRUSTEDCERTIFICATEAUTHORITY_APPKEYCREDENTIALRESTRICTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPKEYCREDENTIALRESTRICTIONTYPE
        default:
            return 0, errors.New("Unknown AppKeyCredentialRestrictionType value: " + v)
    }
    return &result, nil
}
func SerializeAppKeyCredentialRestrictionType(values []AppKeyCredentialRestrictionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppKeyCredentialRestrictionType) isMultiValue() bool {
    return false
}
