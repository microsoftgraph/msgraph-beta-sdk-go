package graph
import (
    "strings"
    "errors"
)
// 
type DerivedCredentialProviderType int

const (
    NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE DerivedCredentialProviderType = iota
    ENTRUSTDATACARD_DERIVEDCREDENTIALPROVIDERTYPE
    PUREBRED_DERIVEDCREDENTIALPROVIDERTYPE
    XTEC_DERIVEDCREDENTIALPROVIDERTYPE
    INTERCEDE_DERIVEDCREDENTIALPROVIDERTYPE
)

func (i DerivedCredentialProviderType) String() string {
    return []string{"NOTCONFIGURED", "ENTRUSTDATACARD", "PUREBRED", "XTEC", "INTERCEDE"}[i]
}
func ParseDerivedCredentialProviderType(v string) (interface{}, error) {
    result := NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE
        case "ENTRUSTDATACARD":
            result = ENTRUSTDATACARD_DERIVEDCREDENTIALPROVIDERTYPE
        case "PUREBRED":
            result = PUREBRED_DERIVEDCREDENTIALPROVIDERTYPE
        case "XTEC":
            result = XTEC_DERIVEDCREDENTIALPROVIDERTYPE
        case "INTERCEDE":
            result = INTERCEDE_DERIVEDCREDENTIALPROVIDERTYPE
        default:
            return 0, errors.New("Unknown DerivedCredentialProviderType value: " + v)
    }
    return &result, nil
}
func SerializeDerivedCredentialProviderType(values []DerivedCredentialProviderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
