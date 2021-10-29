package graph
import (
    "strings"
    "errors"
)
// 
type WindowsUserType int

const (
    ADMINISTRATOR_WINDOWSUSERTYPE WindowsUserType = iota
    STANDARD_WINDOWSUSERTYPE
)

func (i WindowsUserType) String() string {
    return []string{"ADMINISTRATOR", "STANDARD"}[i]
}
func ParseWindowsUserType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ADMINISTRATOR":
            return ADMINISTRATOR_WINDOWSUSERTYPE, nil
        case "STANDARD":
            return STANDARD_WINDOWSUSERTYPE, nil
    }
    return 0, errors.New("Unknown WindowsUserType value: " + v)
}
func SerializeWindowsUserType(values []WindowsUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
