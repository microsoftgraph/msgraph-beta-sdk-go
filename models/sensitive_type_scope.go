package models
import (
    "errors"
    "strings"
)
// 
type SensitiveTypeScope int

const (
    FULLDOCUMENT_SENSITIVETYPESCOPE SensitiveTypeScope = iota
    PARTIALDOCUMENT_SENSITIVETYPESCOPE
)

func (i SensitiveTypeScope) String() string {
    var values []string
    for p := SensitiveTypeScope(1); p <= PARTIALDOCUMENT_SENSITIVETYPESCOPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"fullDocument", "partialDocument"}[p])
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
