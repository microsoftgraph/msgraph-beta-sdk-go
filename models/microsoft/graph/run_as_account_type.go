package graph
import (
    "strings"
    "errors"
)
// 
type RunAsAccountType int

const (
    SYSTEM_RUNASACCOUNTTYPE RunAsAccountType = iota
    USER_RUNASACCOUNTTYPE
)

func (i RunAsAccountType) String() string {
    return []string{"SYSTEM", "USER"}[i]
}
func ParseRunAsAccountType(v string) (interface{}, error) {
    result := SYSTEM_RUNASACCOUNTTYPE
    switch strings.ToUpper(v) {
        case "SYSTEM":
            result = SYSTEM_RUNASACCOUNTTYPE
        case "USER":
            result = USER_RUNASACCOUNTTYPE
        default:
            return 0, errors.New("Unknown RunAsAccountType value: " + v)
    }
    return &result, nil
}
func SerializeRunAsAccountType(values []RunAsAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
