package graph
import (
    "strings"
    "errors"
)
// 
type SearchAlterationType int

const (
    SUGGESTION_SEARCHALTERATIONTYPE SearchAlterationType = iota
    MODIFICATION_SEARCHALTERATIONTYPE
    UNKNOWNFUTUREVALUE_SEARCHALTERATIONTYPE
)

func (i SearchAlterationType) String() string {
    return []string{"SUGGESTION", "MODIFICATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSearchAlterationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SUGGESTION":
            return SUGGESTION_SEARCHALTERATIONTYPE, nil
        case "MODIFICATION":
            return MODIFICATION_SEARCHALTERATIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SEARCHALTERATIONTYPE, nil
    }
    return 0, errors.New("Unknown SearchAlterationType value: " + v)
}
func SerializeSearchAlterationType(values []SearchAlterationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
