package graph
import (
    "strings"
    "errors"
)
// 
type SensitiveTypeScope int

const (
    FULLDOCUMENT_SENSITIVETYPESCOPE SensitiveTypeScope = iota
    PARTIALDOCUMENT_SENSITIVETYPESCOPE
)

func (i SensitiveTypeScope) String() string {
    return []string{"FULLDOCUMENT", "PARTIALDOCUMENT"}[i]
}
func ParseSensitiveTypeScope(v string) (interface{}, error) {
    result := FULLDOCUMENT_SENSITIVETYPESCOPE
    switch strings.ToUpper(v) {
        case "FULLDOCUMENT":
            result = FULLDOCUMENT_SENSITIVETYPESCOPE
        case "PARTIALDOCUMENT":
            result = PARTIALDOCUMENT_SENSITIVETYPESCOPE
        default:
            return 0, errors.New("Unknown SensitiveTypeScope value: " + v)
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
