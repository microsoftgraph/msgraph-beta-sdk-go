package models
import (
    "errors"
)
type GcpAccessType int

const (
    PUBLIC_GCPACCESSTYPE GcpAccessType = iota
    SUBJECTTOOBJECTACLS_GCPACCESSTYPE
    PRIVATE_GCPACCESSTYPE
    UNKNOWNFUTUREVALUE_GCPACCESSTYPE
)

func (i GcpAccessType) String() string {
    return []string{"public", "subjectToObjectAcls", "private", "unknownFutureValue"}[i]
}
func ParseGcpAccessType(v string) (any, error) {
    result := PUBLIC_GCPACCESSTYPE
    switch v {
        case "public":
            result = PUBLIC_GCPACCESSTYPE
        case "subjectToObjectAcls":
            result = SUBJECTTOOBJECTACLS_GCPACCESSTYPE
        case "private":
            result = PRIVATE_GCPACCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_GCPACCESSTYPE
        default:
            return 0, errors.New("Unknown GcpAccessType value: " + v)
    }
    return &result, nil
}
func SerializeGcpAccessType(values []GcpAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GcpAccessType) isMultiValue() bool {
    return false
}
