package graph
import (
    "strings"
    "errors"
)
type SensitiveTypeSource int

const (
    OUTOFBOX_SENSITIVETYPESOURCE SensitiveTypeSource = iota
    TENANT_SENSITIVETYPESOURCE
)

func (i SensitiveTypeSource) String() string {
    return []string{"OUTOFBOX", "TENANT"}[i]
}
func ParseSensitiveTypeSource(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "OUTOFBOX":
            return OUTOFBOX_SENSITIVETYPESOURCE, nil
        case "TENANT":
            return TENANT_SENSITIVETYPESOURCE, nil
    }
    return 0, errors.New("Unknown SensitiveTypeSource value: " + v)
}
func SerializeSensitiveTypeSource(values []SensitiveTypeSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
