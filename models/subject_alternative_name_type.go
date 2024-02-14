package models
import (
    "errors"
    "math"
    "strings"
)
// Subject Alternative Name Options.
type SubjectAlternativeNameType int

const (
    // No subject alternative name.
    NONE_SUBJECTALTERNATIVENAMETYPE = 1
    // Email address.
    EMAILADDRESS_SUBJECTALTERNATIVENAMETYPE = 2
    // User Principal Name (UPN).
    USERPRINCIPALNAME_SUBJECTALTERNATIVENAMETYPE = 4
    // Custom Azure AD Attribute.
    CUSTOMAZUREADATTRIBUTE_SUBJECTALTERNATIVENAMETYPE = 8
    // Domain Name Service (DNS).
    DOMAINNAMESERVICE_SUBJECTALTERNATIVENAMETYPE = 16
    // Universal Resource Identifier (URI).
    UNIVERSALRESOURCEIDENTIFIER_SUBJECTALTERNATIVENAMETYPE = 32
)

func (i SubjectAlternativeNameType) String() string {
    var values []string
    options := []string{"none", "emailAddress", "userPrincipalName", "customAzureADAttribute", "domainNameService", "universalResourceIdentifier"}
    for p := 0; p < 6; p++ {
        mantis := SubjectAlternativeNameType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
