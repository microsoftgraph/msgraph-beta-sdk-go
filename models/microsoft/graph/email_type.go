package graph
import (
    "strings"
    "errors"
)
// 
type EmailType int

const (
    UNKNOWN_EMAILTYPE EmailType = iota
    WORK_EMAILTYPE
    PERSONAL_EMAILTYPE
    MAIN_EMAILTYPE
    OTHER_EMAILTYPE
)

func (i EmailType) String() string {
    return []string{"UNKNOWN", "WORK", "PERSONAL", "MAIN", "OTHER"}[i]
}
func ParseEmailType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_EMAILTYPE, nil
        case "WORK":
            return WORK_EMAILTYPE, nil
        case "PERSONAL":
            return PERSONAL_EMAILTYPE, nil
        case "MAIN":
            return MAIN_EMAILTYPE, nil
        case "OTHER":
            return OTHER_EMAILTYPE, nil
    }
    return 0, errors.New("Unknown EmailType value: " + v)
}
func SerializeEmailType(values []EmailType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
