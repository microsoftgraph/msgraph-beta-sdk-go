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
    result := NONE_DELEGATEDPRIVILEGESTATUS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DELEGATEDPRIVILEGESTATUS
        case "DELEGATEDADMINPRIVILEGES":
            result = DELEGATEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DELEGATEDPRIVILEGESTATUS
        default:
            return 0, errors.New("Unknown DelegatedPrivilegeStatus value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedPrivilegeStatus(values []DelegatedPrivilegeStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
