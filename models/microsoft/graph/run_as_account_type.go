package graph
import (
    "strings"
    "errors"
)
type RunAsAccountType int

const (
    SYSTEM_RUNASACCOUNTTYPE RunAsAccountType = iota
    USER_RUNASACCOUNTTYPE
)

func (i RunAsAccountType) String() string {
    return []string{"SYSTEM", "USER"}[i]
}
func ParseRunAsAccountType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SYSTEM":
            return SYSTEM_RUNASACCOUNTTYPE, nil
        case "USER":
            return USER_RUNASACCOUNTTYPE, nil
    }
    return 0, errors.New("Unknown RunAsAccountType value: " + v)
}
func SerializeRunAsAccountType(values []RunAsAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
