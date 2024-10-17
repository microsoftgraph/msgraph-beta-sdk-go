package models
// Enum type of possible certificate authority types. This feature supports a two-tier certification authority model with a root certification authority and one or more child issuing (intermediate) certification authorities.
type CloudCertificationAuthorityType int

const (
    // Default. Unknown or invalid certification authority type.
    UNKNOWN_CLOUDCERTIFICATIONAUTHORITYTYPE CloudCertificationAuthorityType = iota
    // Indicates root certification authority. Can be used as the parent of an issuing certification authority. Root Certification Authority cannot issue leaf certificates.
    ROOTCERTIFICATIONAUTHORITY_CLOUDCERTIFICATIONAUTHORITYTYPE
    // Indicates issuing (subordinate) certification authority. Must have a parent root certification authority. Issuing Certification Authority can issue leaf certificates.
    ISSUINGCERTIFICATIONAUTHORITY_CLOUDCERTIFICATIONAUTHORITYTYPE
    // Indicates issuing (subordinate) certification authority that has an external root certification authority. Issuing Certification Authority with external root can issue leaf certificates.
    ISSUINGCERTIFICATIONAUTHORITYWITHEXTERNALROOT_CLOUDCERTIFICATIONAUTHORITYTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYTYPE
)

func (i CloudCertificationAuthorityType) String() string {
    return []string{"unknown", "rootCertificationAuthority", "issuingCertificationAuthority", "issuingCertificationAuthorityWithExternalRoot", "unknownFutureValue"}[i]
}
func ParseCloudCertificationAuthorityType(v string) (any, error) {
    result := UNKNOWN_CLOUDCERTIFICATIONAUTHORITYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDCERTIFICATIONAUTHORITYTYPE
        case "rootCertificationAuthority":
            result = ROOTCERTIFICATIONAUTHORITY_CLOUDCERTIFICATIONAUTHORITYTYPE
        case "issuingCertificationAuthority":
            result = ISSUINGCERTIFICATIONAUTHORITY_CLOUDCERTIFICATIONAUTHORITYTYPE
        case "issuingCertificationAuthorityWithExternalRoot":
            result = ISSUINGCERTIFICATIONAUTHORITYWITHEXTERNALROOT_CLOUDCERTIFICATIONAUTHORITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudCertificationAuthorityType(values []CloudCertificationAuthorityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudCertificationAuthorityType) isMultiValue() bool {
    return false
}
