package networkaccess
import (
    "errors"
)
type IpSecEncryption int

const (
    NONE_IPSECENCRYPTION IpSecEncryption = iota
    GCMAES128_IPSECENCRYPTION
    GCMAES192_IPSECENCRYPTION
    GCMAES256_IPSECENCRYPTION
    UNKNOWNFUTUREVALUE_IPSECENCRYPTION
)

func (i IpSecEncryption) String() string {
    return []string{"none", "gcmAes128", "gcmAes192", "gcmAes256", "unknownFutureValue"}[i]
}
func ParseIpSecEncryption(v string) (any, error) {
    result := NONE_IPSECENCRYPTION
    switch v {
        case "none":
            result = NONE_IPSECENCRYPTION
        case "gcmAes128":
            result = GCMAES128_IPSECENCRYPTION
        case "gcmAes192":
            result = GCMAES192_IPSECENCRYPTION
        case "gcmAes256":
            result = GCMAES256_IPSECENCRYPTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IPSECENCRYPTION
        default:
            return 0, errors.New("Unknown IpSecEncryption value: " + v)
    }
    return &result, nil
}
func SerializeIpSecEncryption(values []IpSecEncryption) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IpSecEncryption) isMultiValue() bool {
    return false
}
