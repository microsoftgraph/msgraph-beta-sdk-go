package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_OWNERTYPE, nil
        case "COMPANY":
            return COMPANY_OWNERTYPE, nil
        case "PERSONAL":
            return PERSONAL_OWNERTYPE, nil
    }
    return 0, errors.New("Unknown OwnerType value: " + v)
}
func SerializeOwnerType(values []OwnerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
