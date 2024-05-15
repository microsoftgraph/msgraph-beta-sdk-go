package models
// Provider type for Derived Credentials.
type DerivedCredentialProviderType int

const (
    // No Derived Credential Provider Configured.
    NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE DerivedCredentialProviderType = iota
    // Entrust.
    ENTRUSTDATACARD_DERIVEDCREDENTIALPROVIDERTYPE
    // Purebred - Defense Information Systems Agency.
    PUREBRED_DERIVEDCREDENTIALPROVIDERTYPE
    // Xtec - AuthentX.
    XTEC_DERIVEDCREDENTIALPROVIDERTYPE
    // Intercede.
    INTERCEDE_DERIVEDCREDENTIALPROVIDERTYPE
)

func (i DerivedCredentialProviderType) String() string {
    return []string{"notConfigured", "entrustDataCard", "purebred", "xTec", "intercede"}[i]
}
func ParseDerivedCredentialProviderType(v string) (any, error) {
    result := NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE
        case "entrustDataCard":
            result = ENTRUSTDATACARD_DERIVEDCREDENTIALPROVIDERTYPE
        case "purebred":
            result = PUREBRED_DERIVEDCREDENTIALPROVIDERTYPE
        case "xTec":
            result = XTEC_DERIVEDCREDENTIALPROVIDERTYPE
        case "intercede":
            result = INTERCEDE_DERIVEDCREDENTIALPROVIDERTYPE
        default:
            return nil, nil
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
func (i DerivedCredentialProviderType) isMultiValue() bool {
    return false
}
