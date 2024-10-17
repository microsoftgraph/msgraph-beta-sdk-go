package models
// Enum type of possible leaf certificate statuses. These statuses indicate whether certificates are active and usable or unusable if they have been revoked or expired.
type CloudCertificationAuthorityLeafCertificateStatus int

const (
    // Default. Unknown or invalid status.
    UNKNOWN_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS CloudCertificationAuthorityLeafCertificateStatus = iota
    // Certificate is active, indicating it is in its validity period and not revoked.
    ACTIVE_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
    // Certificate has been revoked by its issuing certification authority.
    REVOKED_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
    // Certificate has expired.
    EXPIRED_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
)

func (i CloudCertificationAuthorityLeafCertificateStatus) String() string {
    return []string{"unknown", "active", "revoked", "expired", "unknownFutureValue"}[i]
}
func ParseCloudCertificationAuthorityLeafCertificateStatus(v string) (any, error) {
    result := UNKNOWN_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
        case "active":
            result = ACTIVE_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
        case "revoked":
            result = REVOKED_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
        case "expired":
            result = EXPIRED_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYLEAFCERTIFICATESTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudCertificationAuthorityLeafCertificateStatus(values []CloudCertificationAuthorityLeafCertificateStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudCertificationAuthorityLeafCertificateStatus) isMultiValue() bool {
    return false
}
