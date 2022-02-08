package graph
import (
    "strings"
    "errors"
)
// 
type SensitiveTypeSource int

const (
    OUTOFBOX_SENSITIVETYPESOURCE SensitiveTypeSource = iota
    TENANT_SENSITIVETYPESOURCE
)

func (i SensitiveTypeSource) String() string {
    return []string{"OUTOFBOX", "TENANT"}[i]
}
func ParseSensitiveTypeSource(v string) (interface{}, error) {
    result := OUTOFBOX_SENSITIVETYPESOURCE
    switch strings.ToUpper(v) {
        case "OUTOFBOX":
            result = OUTOFBOX_SENSITIVETYPESOURCE
        case "TENANT":
            result = TENANT_SENSITIVETYPESOURCE
        default:
            return 0, errors.New("Unknown SensitiveTypeSource value: " + v)
    }
    return &result, nil
}
func SerializeSensitiveTypeSource(values []SensitiveTypeSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
