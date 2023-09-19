package networkaccess
import (
    "errors"
)
// 
type IkeIntegrity int

const (
    SHA256_IKEINTEGRITY IkeIntegrity = iota
    SHA384_IKEINTEGRITY
    GCMAES128_IKEINTEGRITY
    GCMAES256_IKEINTEGRITY
    UNKNOWNFUTUREVALUE_IKEINTEGRITY
)

func (i IkeIntegrity) String() string {
    return []string{"sha256", "sha384", "gcmAes128", "gcmAes256", "unknownFutureValue"}[i]
}
func ParseIkeIntegrity(v string) (any, error) {
    result := SHA256_IKEINTEGRITY
    switch v {
        case "sha256":
            result = SHA256_IKEINTEGRITY
        case "sha384":
            result = SHA384_IKEINTEGRITY
        case "gcmAes128":
            result = GCMAES128_IKEINTEGRITY
        case "gcmAes256":
            result = GCMAES256_IKEINTEGRITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IKEINTEGRITY
        default:
            return 0, errors.New("Unknown IkeIntegrity value: " + v)
    }
    return &result, nil
}
func SerializeIkeIntegrity(values []IkeIntegrity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IkeIntegrity) isMultiValue() bool {
    return false
}
