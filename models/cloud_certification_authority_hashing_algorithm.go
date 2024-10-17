package models
// Enum type of possible certificate hashing algorithms used by the certification authority to create certificates.
type CloudCertificationAuthorityHashingAlgorithm int

const (
    // Default. The hashing algorithm is unknown or invalid.
    UNKNOWN_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM CloudCertificationAuthorityHashingAlgorithm = iota
    // The hashing algorithm is SHA-256.
    SHA256_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
    // The hashing algorithm is SHA-384.
    SHA384_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
    // The hashing algorithm is SHA-512.
    SHA512_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
)

func (i CloudCertificationAuthorityHashingAlgorithm) String() string {
    return []string{"unknown", "sha256", "sha384", "sha512", "unknownFutureValue"}[i]
}
func ParseCloudCertificationAuthorityHashingAlgorithm(v string) (any, error) {
    result := UNKNOWN_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
        case "sha256":
            result = SHA256_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
        case "sha384":
            result = SHA384_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
        case "sha512":
            result = SHA512_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYHASHINGALGORITHM
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudCertificationAuthorityHashingAlgorithm(values []CloudCertificationAuthorityHashingAlgorithm) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudCertificationAuthorityHashingAlgorithm) isMultiValue() bool {
    return false
}
