package models
// Authentication transform values for Windows10 VPN policies with IKEv2 connection type.
type AuthenticationTransformConstant int

const (
    // MD596
    MD5_96_AUTHENTICATIONTRANSFORMCONSTANT AuthenticationTransformConstant = iota
    // SHA196
    SHA1_96_AUTHENTICATIONTRANSFORMCONSTANT
    // SHA256128
    SHA_256_128_AUTHENTICATIONTRANSFORMCONSTANT
    // GCMAES128
    AES128GCM_AUTHENTICATIONTRANSFORMCONSTANT
    // GCMAES192
    AES192GCM_AUTHENTICATIONTRANSFORMCONSTANT
    // GCMAES256
    AES256GCM_AUTHENTICATIONTRANSFORMCONSTANT
)

func (i AuthenticationTransformConstant) String() string {
    return []string{"md5_96", "sha1_96", "sha_256_128", "aes128Gcm", "aes192Gcm", "aes256Gcm"}[i]
}
func ParseAuthenticationTransformConstant(v string) (any, error) {
    result := MD5_96_AUTHENTICATIONTRANSFORMCONSTANT
    switch v {
        case "md5_96":
            result = MD5_96_AUTHENTICATIONTRANSFORMCONSTANT
        case "sha1_96":
            result = SHA1_96_AUTHENTICATIONTRANSFORMCONSTANT
        case "sha_256_128":
            result = SHA_256_128_AUTHENTICATIONTRANSFORMCONSTANT
        case "aes128Gcm":
            result = AES128GCM_AUTHENTICATIONTRANSFORMCONSTANT
        case "aes192Gcm":
            result = AES192GCM_AUTHENTICATIONTRANSFORMCONSTANT
        case "aes256Gcm":
            result = AES256GCM_AUTHENTICATIONTRANSFORMCONSTANT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthenticationTransformConstant(values []AuthenticationTransformConstant) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationTransformConstant) isMultiValue() bool {
    return false
}
