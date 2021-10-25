package graph
import (
    "strings"
    "errors"
)
type GroupPolicyType int

const (
    ADMXBACKED_GROUPPOLICYTYPE GroupPolicyType = iota
    ADMXINGESTED_GROUPPOLICYTYPE
)

func (i GroupPolicyType) String() string {
    return []string{"ADMXBACKED", "ADMXINGESTED"}[i]
}
func ParseGroupPolicyType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ADMXBACKED":
            return ADMXBACKED_GROUPPOLICYTYPE, nil
        case "ADMXINGESTED":
            return ADMXINGESTED_GROUPPOLICYTYPE, nil
    }
    return 0, errors.New("Unknown GroupPolicyType value: " + v)
}
func SerializeGroupPolicyType(values []GroupPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
