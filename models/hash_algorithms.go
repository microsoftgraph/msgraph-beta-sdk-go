package models
import (
    "math"
    "strings"
)
// Hash Algorithm Options.
type HashAlgorithms int

const (
    // SHA-1 Hash Algorithm.
    SHA1_HASHALGORITHMS = 1
    // SHA-2 Hash Algorithm.
    SHA2_HASHALGORITHMS = 2
)

func (i HashAlgorithms) String() string {
    var values []string
    options := []string{"sha1", "sha2"}
    for p := 0; p < 2; p++ {
        mantis := HashAlgorithms(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
