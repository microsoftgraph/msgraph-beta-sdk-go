package graph
import (
    "strings"
    "errors"
)
// 
type AutoAdmittedUsersType int

const (
    EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE AutoAdmittedUsersType = iota
    EVERYONE_AUTOADMITTEDUSERSTYPE
)

func (i AutoAdmittedUsersType) String() string {
    return []string{"EVERYONEINCOMPANY", "EVERYONE"}[i]
}
func ParseAutoAdmittedUsersType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "EVERYONEINCOMPANY":
            return EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE, nil
        case "EVERYONE":
            return EVERYONE_AUTOADMITTEDUSERSTYPE, nil
    }
    return 0, errors.New("Unknown AutoAdmittedUsersType value: " + v)
}
func SerializeAutoAdmittedUsersType(values []AutoAdmittedUsersType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
