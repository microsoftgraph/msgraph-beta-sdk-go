package models
type SamlNameIDFormat int

const (
    DEFAULT_SAMLNAMEIDFORMAT SamlNameIDFormat = iota
    UNSPECIFIED_SAMLNAMEIDFORMAT
    EMAILADDRESS_SAMLNAMEIDFORMAT
    WINDOWSDOMAINQUALIFIEDNAME_SAMLNAMEIDFORMAT
    PERSISTENT_SAMLNAMEIDFORMAT
    UNKNOWNFUTUREVALUE_SAMLNAMEIDFORMAT
)

func (i SamlNameIDFormat) String() string {
    return []string{"default", "unspecified", "emailAddress", "windowsDomainQualifiedName", "persistent", "unknownFutureValue"}[i]
}
func ParseSamlNameIDFormat(v string) (any, error) {
    result := DEFAULT_SAMLNAMEIDFORMAT
    switch v {
        case "default":
            result = DEFAULT_SAMLNAMEIDFORMAT
        case "unspecified":
            result = UNSPECIFIED_SAMLNAMEIDFORMAT
        case "emailAddress":
            result = EMAILADDRESS_SAMLNAMEIDFORMAT
        case "windowsDomainQualifiedName":
            result = WINDOWSDOMAINQUALIFIEDNAME_SAMLNAMEIDFORMAT
        case "persistent":
            result = PERSISTENT_SAMLNAMEIDFORMAT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SAMLNAMEIDFORMAT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSamlNameIDFormat(values []SamlNameIDFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SamlNameIDFormat) isMultiValue() bool {
    return false
}
