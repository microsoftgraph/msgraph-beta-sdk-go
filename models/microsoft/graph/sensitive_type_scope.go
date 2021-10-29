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
    switch strings.ToUpper(v) {
        case "FULLDOCUMENT":
            return FULLDOCUMENT_SENSITIVETYPESCOPE, nil
        case "PARTIALDOCUMENT":
            return PARTIALDOCUMENT_SENSITIVETYPESCOPE, nil
    }
    return 0, errors.New("Unknown SensitiveTypeScope value: " + v)
}
func SerializeSensitiveTypeScope(values []SensitiveTypeScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
