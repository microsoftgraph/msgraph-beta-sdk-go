package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "EVERYONE":
            return EVERYONE_ACCESSLEVEL, nil
        case "INVITED":
            return INVITED_ACCESSLEVEL, nil
        case "LOCKED":
            return LOCKED_ACCESSLEVEL, nil
        case "SAMEENTERPRISE":
            return SAMEENTERPRISE_ACCESSLEVEL, nil
        case "SAMEENTERPRISEANDFEDERATED":
            return SAMEENTERPRISEANDFEDERATED_ACCESSLEVEL, nil
    }
    return 0, errors.New("Unknown AccessLevel value: " + v)
}
func SerializeAccessLevel(values []AccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
