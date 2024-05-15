package models
import (
    "math"
    "strings"
)
// Key Usage Options.
type KeyUsages int

const (
    // Key Encipherment Usage.
    KEYENCIPHERMENT_KEYUSAGES = 1
    // Digital Signature Usage.
    DIGITALSIGNATURE_KEYUSAGES = 2
)

func (i KeyUsages) String() string {
    var values []string
    options := []string{"keyEncipherment", "digitalSignature"}
    for p := 0; p < 2; p++ {
        mantis := KeyUsages(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
