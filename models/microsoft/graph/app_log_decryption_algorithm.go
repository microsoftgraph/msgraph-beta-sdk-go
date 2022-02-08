package graph
import (
    "strings"
    "errors"
)
// 
type AppLogDecryptionAlgorithm int

const (
    AES256_APPLOGDECRYPTIONALGORITHM AppLogDecryptionAlgorithm = iota
)

func (i AppLogDecryptionAlgorithm) String() string {
    return []string{"AES256"}[i]
}
func ParseAppLogDecryptionAlgorithm(v string) (interface{}, error) {
    result := AES256_APPLOGDECRYPTIONALGORITHM
    switch strings.ToUpper(v) {
        case "AES256":
            result = AES256_APPLOGDECRYPTIONALGORITHM
        default:
            return 0, errors.New("Unknown AppLogDecryptionAlgorithm value: " + v)
    }
    return &result, nil
}
func SerializeAppLogDecryptionAlgorithm(values []AppLogDecryptionAlgorithm) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
