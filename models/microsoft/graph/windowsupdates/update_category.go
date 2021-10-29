package windowsupdates
import (
    "strings"
    "errors"
)
// 
type UpdateCategory int

const (
    FEATURE_UPDATECATEGORY UpdateCategory = iota
    QUALITY_UPDATECATEGORY
    UNKNOWNFUTUREVALUE_UPDATECATEGORY
)

func (i UpdateCategory) String() string {
    return []string{"FEATURE", "QUALITY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUpdateCategory(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "FEATURE":
            return FEATURE_UPDATECATEGORY, nil
        case "QUALITY":
            return QUALITY_UPDATECATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_UPDATECATEGORY, nil
    }
    return 0, errors.New("Unknown UpdateCategory value: " + v)
}
func SerializeUpdateCategory(values []UpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
