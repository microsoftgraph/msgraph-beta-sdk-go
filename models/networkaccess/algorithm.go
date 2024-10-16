package networkaccess
type Algorithm int

const (
    MD5_ALGORITHM Algorithm = iota
    SHA1_ALGORITHM
    SHA256_ALGORITHM
    SHA256AC_ALGORITHM
    UNKNOWNFUTUREVALUE_ALGORITHM
)

func (i Algorithm) String() string {
    return []string{"md5", "sha1", "sha256", "sha256ac", "unknownFutureValue"}[i]
}
func ParseAlgorithm(v string) (any, error) {
    result := MD5_ALGORITHM
    switch v {
        case "md5":
            result = MD5_ALGORITHM
        case "sha1":
            result = SHA1_ALGORITHM
        case "sha256":
            result = SHA256_ALGORITHM
        case "sha256ac":
            result = SHA256AC_ALGORITHM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALGORITHM
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlgorithm(values []Algorithm) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Algorithm) isMultiValue() bool {
    return false
}
