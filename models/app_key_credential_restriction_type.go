package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type AppKeyCredentialRestrictionType int

const (
    ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE AppKeyCredentialRestrictionType = iota
    UNKNOWNFUTUREVALUE_APPKEYCREDENTIALRESTRICTIONTYPE
)

func (i AppKeyCredentialRestrictionType) String() string {
    return []string{"asymmetricKeyLifetime", "unknownFutureValue"}[i]
}
func ParseAppKeyCredentialRestrictionType(v string) (interface{}, error) {
    result := ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE
    switch v {
        case "asymmetricKeyLifetime":
            result = ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE
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
