package graph
import (
    "strings"
    "errors"
)
// 
type AccessLevel int

const (
    EVERYONE_ACCESSLEVEL AccessLevel = iota
    INVITED_ACCESSLEVEL
    LOCKED_ACCESSLEVEL
    SAMEENTERPRISE_ACCESSLEVEL
    SAMEENTERPRISEANDFEDERATED_ACCESSLEVEL
)

func (i AccessLevel) String() string {
    return []string{"EVERYONE", "INVITED", "LOCKED", "SAMEENTERPRISE", "SAMEENTERPRISEANDFEDERATED"}[i]
}
func ParseAccessLevel(v string) (interface{}, error) {
    result := EVERYONE_ACCESSLEVEL
    switch strings.ToUpper(v) {
        case "EVERYONE":
            result = EVERYONE_ACCESSLEVEL
        case "INVITED":
            result = INVITED_ACCESSLEVEL
        case "LOCKED":
            result = LOCKED_ACCESSLEVEL
        case "SAMEENTERPRISE":
            result = SAMEENTERPRISE_ACCESSLEVEL
        case "SAMEENTERPRISEANDFEDERATED":
            result = SAMEENTERPRISEANDFEDERATED_ACCESSLEVEL
        default:
            return 0, errors.New("Unknown AccessLevel value: " + v)
    }
    return &result, nil
}
func SerializeAccessLevel(values []AccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
