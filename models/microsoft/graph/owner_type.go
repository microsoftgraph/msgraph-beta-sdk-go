package graph
import (
    "strings"
    "errors"
)
// 
type OwnerType int

const (
    UNKNOWN_OWNERTYPE OwnerType = iota
    COMPANY_OWNERTYPE
    PERSONAL_OWNERTYPE
)

func (i OwnerType) String() string {
    return []string{"UNKNOWN", "COMPANY", "PERSONAL"}[i]
}
func ParseOwnerType(v string) (interface{}, error) {
    result := UNKNOWN_OWNERTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_OWNERTYPE
        case "COMPANY":
            result = COMPANY_OWNERTYPE
        case "PERSONAL":
            result = PERSONAL_OWNERTYPE
        default:
            return 0, errors.New("Unknown OwnerType value: " + v)
    }
    return &result, nil
}
func SerializeOwnerType(values []OwnerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
