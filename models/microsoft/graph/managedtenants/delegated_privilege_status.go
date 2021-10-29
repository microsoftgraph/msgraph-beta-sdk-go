package managedtenants
import (
    "strings"
    "errors"
)
// 
type DelegatedPrivilegeStatus int

const (
    NONE_DELEGATEDPRIVILEGESTATUS DelegatedPrivilegeStatus = iota
    DELEGATEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
    UNKNOWNFUTUREVALUE_DELEGATEDPRIVILEGESTATUS
)

func (i DelegatedPrivilegeStatus) String() string {
    return []string{"NONE", "DELEGATEDADMINPRIVILEGES", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedPrivilegeStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DELEGATEDPRIVILEGESTATUS, nil
        case "DELEGATEDADMINPRIVILEGES":
            return DELEGATEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DELEGATEDPRIVILEGESTATUS, nil
    }
    return 0, errors.New("Unknown DelegatedPrivilegeStatus value: " + v)
}
func SerializeDelegatedPrivilegeStatus(values []DelegatedPrivilegeStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
