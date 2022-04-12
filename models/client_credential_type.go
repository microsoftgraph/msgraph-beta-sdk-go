package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type ClientCredentialType int

const (
    NONE_CLIENTCREDENTIALTYPE ClientCredentialType = iota
    CLIENTSECRET_CLIENTCREDENTIALTYPE
    CLIENTASSERTION_CLIENTCREDENTIALTYPE
    FEDERATEDIDENTITYCREDENTIAL_CLIENTCREDENTIALTYPE
    MANAGEDIDENTITY_CLIENTCREDENTIALTYPE
    CERTIFICATE_CLIENTCREDENTIALTYPE
    UNKNOWNFUTUREVALUE_CLIENTCREDENTIALTYPE
)

func (i ClientCredentialType) String() string {
    return []string{"NONE", "CLIENTSECRET", "CLIENTASSERTION", "FEDERATEDIDENTITYCREDENTIAL", "MANAGEDIDENTITY", "CERTIFICATE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseClientCredentialType(v string) (interface{}, error) {
    result := NONE_CLIENTCREDENTIALTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CLIENTCREDENTIALTYPE
        case "CLIENTSECRET":
            result = CLIENTSECRET_CLIENTCREDENTIALTYPE
        case "CLIENTASSERTION":
            result = CLIENTASSERTION_CLIENTCREDENTIALTYPE
        case "FEDERATEDIDENTITYCREDENTIAL":
            result = FEDERATEDIDENTITYCREDENTIAL_CLIENTCREDENTIALTYPE
        case "MANAGEDIDENTITY":
            result = MANAGEDIDENTITY_CLIENTCREDENTIALTYPE
        case "CERTIFICATE":
            result = CERTIFICATE_CLIENTCREDENTIALTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLIENTCREDENTIALTYPE
        default:
            return 0, errors.New("Unknown ClientCredentialType value: " + v)
    }
    return &result, nil
}
func SerializeClientCredentialType(values []ClientCredentialType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
