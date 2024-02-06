package models
import (
    "errors"
    "math"
    "strings"
)
// 
type SensitiveTypeScope int

const (
    FULLDOCUMENT_SENSITIVETYPESCOPE = 1
    PARTIALDOCUMENT_SENSITIVETYPESCOPE = 2
)

func (i SensitiveTypeScope) String() string {
    var values []string
    options := []string{"fullDocument", "partialDocument"}
    for p := 0; p < 2; p++ {
        mantis := SensitiveTypeScope(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSensitiveTypeScope(v string) (any, error) {
    var result SensitiveTypeScope
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "fullDocument":
                result |= FULLDOCUMENT_SENSITIVETYPESCOPE
            case "partialDocument":
                result |= PARTIALDOCUMENT_SENSITIVETYPESCOPE
            default:
                return 0, errors.New("Unknown SensitiveTypeScope value: " + v)
        }
    }
    return &result, nil
}
func SerializeSensitiveTypeScope(values []SensitiveTypeScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SensitiveTypeScope) isMultiValue() bool {
    return true
}
