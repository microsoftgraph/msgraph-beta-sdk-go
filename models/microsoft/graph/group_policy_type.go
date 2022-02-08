package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicyType int

const (
    ADMXBACKED_GROUPPOLICYTYPE GroupPolicyType = iota
    ADMXINGESTED_GROUPPOLICYTYPE
)

func (i GroupPolicyType) String() string {
    return []string{"ADMXBACKED", "ADMXINGESTED"}[i]
}
func ParseGroupPolicyType(v string) (interface{}, error) {
    result := ADMXBACKED_GROUPPOLICYTYPE
    switch strings.ToUpper(v) {
        case "ADMXBACKED":
            result = ADMXBACKED_GROUPPOLICYTYPE
        case "ADMXINGESTED":
            result = ADMXINGESTED_GROUPPOLICYTYPE
        default:
            return 0, errors.New("Unknown GroupPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyType(values []GroupPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
