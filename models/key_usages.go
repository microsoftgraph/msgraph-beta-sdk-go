package models
import (
    "errors"
    "strings"
)
// Key Usage Options.
type KeyUsages int

const (
    // Key Encipherment Usage.
    KEYENCIPHERMENT_KEYUSAGES KeyUsages = iota
    // Digital Signature Usage.
    DIGITALSIGNATURE_KEYUSAGES
)

func (i KeyUsages) String() string {
    var values []string
    for p := KeyUsages(1); p <= DIGITALSIGNATURE_KEYUSAGES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"keyEncipherment", "digitalSignature"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseKeyUsages(v string) (any, error) {
    var result KeyUsages
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "keyEncipherment":
                result |= KEYENCIPHERMENT_KEYUSAGES
            case "digitalSignature":
                result |= DIGITALSIGNATURE_KEYUSAGES
            default:
                return 0, errors.New("Unknown KeyUsages value: " + v)
        }
    }
    return &result, nil
}
func SerializeKeyUsages(values []KeyUsages) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i KeyUsages) isMultiValue() bool {
    return true
}
