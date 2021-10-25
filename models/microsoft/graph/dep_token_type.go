package graph
import (
    "strings"
    "errors"
)
type DepTokenType int

const (
    NONE_DEPTOKENTYPE DepTokenType = iota
    DEP_DEPTOKENTYPE
    APPLESCHOOLMANAGER_DEPTOKENTYPE
)

func (i DepTokenType) String() string {
    return []string{"NONE", "DEP", "APPLESCHOOLMANAGER"}[i]
}
func ParseDepTokenType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEPTOKENTYPE, nil
        case "DEP":
            return DEP_DEPTOKENTYPE, nil
        case "APPLESCHOOLMANAGER":
            return APPLESCHOOLMANAGER_DEPTOKENTYPE, nil
    }
    return 0, errors.New("Unknown DepTokenType value: " + v)
}
func SerializeDepTokenType(values []DepTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
