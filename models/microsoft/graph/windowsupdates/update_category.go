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
    result := FEATURE_UPDATECATEGORY
    switch strings.ToUpper(v) {
        case "FEATURE":
            result = FEATURE_UPDATECATEGORY
        case "QUALITY":
            result = QUALITY_UPDATECATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_UPDATECATEGORY
        default:
            return 0, errors.New("Unknown UpdateCategory value: " + v)
    }
    return &result, nil
}
func SerializeUpdateCategory(values []UpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
