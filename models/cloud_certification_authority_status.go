package models
// Enum type of possible certification authority statuses. These statuses indicate whether a certification authority is currently able to issue certificates or temporarily paused or permanently revoked.
type CloudCertificationAuthorityStatus int

const (
    // Default. Indicates certification authority has an unknown or invalid status.
    UNKNOWN_CLOUDCERTIFICATIONAUTHORITYSTATUS CloudCertificationAuthorityStatus = iota
    // Indicates certification authority is active and can issue certificates.
    ACTIVE_CLOUDCERTIFICATIONAUTHORITYSTATUS
    // Indicates certification authority has been paused from issuing certificates. Paused certification authorities can be put back in an active status to continue issuing certificates.
    PAUSED_CLOUDCERTIFICATIONAUTHORITYSTATUS
    // Indicates certification authority has been revoked. This is a permanent state that cannot be changed.
    REVOKED_CLOUDCERTIFICATIONAUTHORITYSTATUS
    // Indicates certification authority certificate signing request has been created and can be downloaded for signing and then be uploaded.
    SIGNINGPENDING_CLOUDCERTIFICATIONAUTHORITYSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYSTATUS
)

func (i CloudCertificationAuthorityStatus) String() string {
    return []string{"unknown", "active", "paused", "revoked", "signingPending", "unknownFutureValue"}[i]
}
func ParseCloudCertificationAuthorityStatus(v string) (any, error) {
    result := UNKNOWN_CLOUDCERTIFICATIONAUTHORITYSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDCERTIFICATIONAUTHORITYSTATUS
        case "active":
            result = ACTIVE_CLOUDCERTIFICATIONAUTHORITYSTATUS
        case "paused":
            result = PAUSED_CLOUDCERTIFICATIONAUTHORITYSTATUS
        case "revoked":
            result = REVOKED_CLOUDCERTIFICATIONAUTHORITYSTATUS
        case "signingPending":
            result = SIGNINGPENDING_CLOUDCERTIFICATIONAUTHORITYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudCertificationAuthorityStatus(values []CloudCertificationAuthorityStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudCertificationAuthorityStatus) isMultiValue() bool {
    return false
}
