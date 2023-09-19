package models
import (
    "errors"
    "strings"
)
// Subject Alternative Name Options.
type SubjectAlternativeNameType int

const (
    // No subject alternative name.
    NONE_SUBJECTALTERNATIVENAMETYPE SubjectAlternativeNameType = iota
    // Email address.
    EMAILADDRESS_SUBJECTALTERNATIVENAMETYPE
    // User Principal Name (UPN).
    USERPRINCIPALNAME_SUBJECTALTERNATIVENAMETYPE
    // Custom Azure AD Attribute.
    CUSTOMAZUREADATTRIBUTE_SUBJECTALTERNATIVENAMETYPE
    // Domain Name Service (DNS).
    DOMAINNAMESERVICE_SUBJECTALTERNATIVENAMETYPE
    // Universal Resource Identifier (URI).
    UNIVERSALRESOURCEIDENTIFIER_SUBJECTALTERNATIVENAMETYPE
)

func (i SubjectAlternativeNameType) String() string {
    var values []string
    for p := SubjectAlternativeNameType(1); p <= UNIVERSALRESOURCEIDENTIFIER_SUBJECTALTERNATIVENAMETYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "emailAddress", "userPrincipalName", "customAzureADAttribute", "domainNameService", "universalResourceIdentifier"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSubjectAlternativeNameType(v string) (any, error) {
    var result SubjectAlternativeNameType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_SUBJECTALTERNATIVENAMETYPE
            case "emailAddress":
                result |= EMAILADDRESS_SUBJECTALTERNATIVENAMETYPE
            case "userPrincipalName":
                result |= USERPRINCIPALNAME_SUBJECTALTERNATIVENAMETYPE
            case "customAzureADAttribute":
                result |= CUSTOMAZUREADATTRIBUTE_SUBJECTALTERNATIVENAMETYPE
            case "domainNameService":
                result |= DOMAINNAMESERVICE_SUBJECTALTERNATIVENAMETYPE
            case "universalResourceIdentifier":
                result |= UNIVERSALRESOURCEIDENTIFIER_SUBJECTALTERNATIVENAMETYPE
            default:
                return 0, errors.New("Unknown SubjectAlternativeNameType value: " + v)
        }
    }
    return &result, nil
}
func SerializeSubjectAlternativeNameType(values []SubjectAlternativeNameType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SubjectAlternativeNameType) isMultiValue() bool {
    return true
}
