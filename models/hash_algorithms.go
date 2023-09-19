package models
import (
    "errors"
    "strings"
)
// Hash Algorithm Options.
type HashAlgorithms int

const (
    // SHA-1 Hash Algorithm.
    SHA1_HASHALGORITHMS HashAlgorithms = iota
    // SHA-2 Hash Algorithm.
    SHA2_HASHALGORITHMS
)

func (i HashAlgorithms) String() string {
    var values []string
    for p := HashAlgorithms(1); p <= SHA2_HASHALGORITHMS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"sha1", "sha2"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseHashAlgorithms(v string) (any, error) {
    var result HashAlgorithms
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "sha1":
                result |= SHA1_HASHALGORITHMS
            case "sha2":
                result |= SHA2_HASHALGORITHMS
            default:
                return 0, errors.New("Unknown HashAlgorithms value: " + v)
        }
    }
    return &result, nil
}
func SerializeHashAlgorithms(values []HashAlgorithms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HashAlgorithms) isMultiValue() bool {
    return true
}
