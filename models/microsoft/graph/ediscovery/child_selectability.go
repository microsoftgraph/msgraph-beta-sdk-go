package ediscovery
import (
    "strings"
    "errors"
)
// 
type ChildSelectability int

const (
    ONE_CHILDSELECTABILITY ChildSelectability = iota
    MANY_CHILDSELECTABILITY
)

func (i ChildSelectability) String() string {
    return []string{"ONE", "MANY"}[i]
}
func ParseChildSelectability(v string) (interface{}, error) {
    result := ONE_CHILDSELECTABILITY
    switch strings.ToUpper(v) {
        case "ONE":
            result = ONE_CHILDSELECTABILITY
        case "MANY":
            result = MANY_CHILDSELECTABILITY
        default:
            return 0, errors.New("Unknown ChildSelectability value: " + v)
    }
    return &result, nil
}
func SerializeChildSelectability(values []ChildSelectability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
