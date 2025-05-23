// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package networkaccess
type IpSecIntegrity int

const (
    GCMAES128_IPSECINTEGRITY IpSecIntegrity = iota
    GCMAES192_IPSECINTEGRITY
    GCMAES256_IPSECINTEGRITY
    SHA256_IPSECINTEGRITY
    UNKNOWNFUTUREVALUE_IPSECINTEGRITY
)

func (i IpSecIntegrity) String() string {
    return []string{"gcmAes128", "gcmAes192", "gcmAes256", "sha256", "unknownFutureValue"}[i]
}
func ParseIpSecIntegrity(v string) (any, error) {
    result := GCMAES128_IPSECINTEGRITY
    switch v {
        case "gcmAes128":
            result = GCMAES128_IPSECINTEGRITY
        case "gcmAes192":
            result = GCMAES192_IPSECINTEGRITY
        case "gcmAes256":
            result = GCMAES256_IPSECINTEGRITY
        case "sha256":
            result = SHA256_IPSECINTEGRITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IPSECINTEGRITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIpSecIntegrity(values []IpSecIntegrity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IpSecIntegrity) isMultiValue() bool {
    return false
}
