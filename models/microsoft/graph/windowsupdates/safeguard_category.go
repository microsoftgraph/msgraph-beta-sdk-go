package windowsupdates
import (
    "strings"
    "errors"
)
// 
type SafeguardCategory int

const (
    LIKELYISSUES_SAFEGUARDCATEGORY SafeguardCategory = iota
    UNKNOWNFUTUREVALUE_SAFEGUARDCATEGORY
)

func (i SafeguardCategory) String() string {
    return []string{"LIKELYISSUES", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSafeguardCategory(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "LIKELYISSUES":
            return LIKELYISSUES_SAFEGUARDCATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SAFEGUARDCATEGORY, nil
    }
    return 0, errors.New("Unknown SafeguardCategory value: " + v)
}
func SerializeSafeguardCategory(values []SafeguardCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
