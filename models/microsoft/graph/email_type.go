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
    result := UNKNOWN_EMAILTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_EMAILTYPE
        case "WORK":
            result = WORK_EMAILTYPE
        case "PERSONAL":
            result = PERSONAL_EMAILTYPE
        case "MAIN":
            result = MAIN_EMAILTYPE
        case "OTHER":
            result = OTHER_EMAILTYPE
        default:
            return 0, errors.New("Unknown EmailType value: " + v)
    }
    return &result, nil
}
func SerializeEmailType(values []EmailType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
