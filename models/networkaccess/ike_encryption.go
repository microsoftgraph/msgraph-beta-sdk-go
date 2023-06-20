package networkaccess
import (
    "errors"
)
// 
type IkeEncryption int

const (
    AES128_IKEENCRYPTION IkeEncryption = iota
    AES192_IKEENCRYPTION
    AES256_IKEENCRYPTION
    GCMAES128_IKEENCRYPTION
    GCMAES256_IKEENCRYPTION
    UNKNOWNFUTUREVALUE_IKEENCRYPTION
)

func (i IkeEncryption) String() string {
    return []string{"aes128", "aes192", "aes256", "gcmAes128", "gcmAes256", "unknownFutureValue"}[i]
}
func ParseIkeEncryption(v string) (any, error) {
    result := AES128_IKEENCRYPTION
    switch v {
        case "aes128":
            result = AES128_IKEENCRYPTION
        case "aes192":
            result = AES192_IKEENCRYPTION
        case "aes256":
            result = AES256_IKEENCRYPTION
        case "gcmAes128":
            result = GCMAES128_IKEENCRYPTION
        case "gcmAes256":
            result = GCMAES256_IKEENCRYPTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IKEENCRYPTION
        default:
            return 0, errors.New("Unknown IkeEncryption value: " + v)
    }
    return &result, nil
}
func SerializeIkeEncryption(values []IkeEncryption) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
