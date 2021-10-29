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
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE, nil
        case "ENTRUSTDATACARD":
            return ENTRUSTDATACARD_DERIVEDCREDENTIALPROVIDERTYPE, nil
        case "PUREBRED":
            return PUREBRED_DERIVEDCREDENTIALPROVIDERTYPE, nil
        case "XTEC":
            return XTEC_DERIVEDCREDENTIALPROVIDERTYPE, nil
        case "INTERCEDE":
            return INTERCEDE_DERIVEDCREDENTIALPROVIDERTYPE, nil
    }
    return 0, errors.New("Unknown DerivedCredentialProviderType value: " + v)
}
func SerializeDerivedCredentialProviderType(values []DerivedCredentialProviderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
