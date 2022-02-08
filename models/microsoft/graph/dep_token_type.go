package graph
import (
    "strings"
    "errors"
)
// 
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
    result := NONE_DEPTOKENTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEPTOKENTYPE
        case "DEP":
            result = DEP_DEPTOKENTYPE
        case "APPLESCHOOLMANAGER":
            result = APPLESCHOOLMANAGER_DEPTOKENTYPE
        default:
            return 0, errors.New("Unknown DepTokenType value: " + v)
    }
    return &result, nil
}
func SerializeDepTokenType(values []DepTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
