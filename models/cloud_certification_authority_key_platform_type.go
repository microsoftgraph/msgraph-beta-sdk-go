package models
// Enum type of possible key platforms used by the certification authority.
type CloudCertificationAuthorityKeyPlatformType int

const (
    // Default. The key platform type is unknown or invalid.
    UNKNOWN_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE CloudCertificationAuthorityKeyPlatformType = iota
    // The certification authority keys are stored in software.
    SOFTWARE_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
    // The certification authority keys are stored in a hardware security module.
    HARDWARESECURITYMODULE_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
)

func (i CloudCertificationAuthorityKeyPlatformType) String() string {
    return []string{"unknown", "software", "hardwareSecurityModule", "unknownFutureValue"}[i]
}
func ParseCloudCertificationAuthorityKeyPlatformType(v string) (any, error) {
    result := UNKNOWN_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
        case "software":
            result = SOFTWARE_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
        case "hardwareSecurityModule":
            result = HARDWARESECURITYMODULE_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYKEYPLATFORMTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudCertificationAuthorityKeyPlatformType(values []CloudCertificationAuthorityKeyPlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudCertificationAuthorityKeyPlatformType) isMultiValue() bool {
    return false
}
