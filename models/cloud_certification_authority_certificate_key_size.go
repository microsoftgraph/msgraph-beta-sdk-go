package models
// Enum of possible cloud certification authority certificate cryptography and key size combinations.
type CloudCertificationAuthorityCertificateKeySize int

const (
    // Default. Unknown or invalid value.
    UNKNOWN_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE CloudCertificationAuthorityCertificateKeySize = iota
    // A certificate generated using RSA cryptography and a key size of 2048 bits.
    RSA2048_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    // A certificate generated using RSA cryptography and a key size of 3072 bits.
    RSA3072_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    // A certificate generated using RSA cryptography and a key size of 4096 bits.
    RSA4096_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    // A certificate generated using Elliptic Curve cryptography and a key size of 256 bits.
    ECP256_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    // A certificate generated using Elliptic Curve cryptography and a key size of 256 bits with a Koblitz curve.
    ECP256K_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    // A certificate generated using Elliptic Curve cryptography and a key size of 384 bits.
    ECP384_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    // A certificate generated using Elliptic Curve cryptography and a key size of 521 bits.
    ECP521_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
)

func (i CloudCertificationAuthorityCertificateKeySize) String() string {
    return []string{"unknown", "rsa2048", "rsa3072", "rsa4096", "eCP256", "eCP256k", "eCP384", "eCP521", "unknownFutureValue"}[i]
}
func ParseCloudCertificationAuthorityCertificateKeySize(v string) (any, error) {
    result := UNKNOWN_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "rsa2048":
            result = RSA2048_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "rsa3072":
            result = RSA3072_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "rsa4096":
            result = RSA4096_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "eCP256":
            result = ECP256_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "eCP256k":
            result = ECP256K_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "eCP384":
            result = ECP384_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "eCP521":
            result = ECP521_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDCERTIFICATIONAUTHORITYCERTIFICATEKEYSIZE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudCertificationAuthorityCertificateKeySize(values []CloudCertificationAuthorityCertificateKeySize) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudCertificationAuthorityCertificateKeySize) isMultiValue() bool {
    return false
}
