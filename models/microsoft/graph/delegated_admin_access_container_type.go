package graph
import (
    "strings"
    "errors"
)
// 
type DelegatedAdminAccessContainerType int

const (
    SECURITYGROUP_DELEGATEDADMINACCESSCONTAINERTYPE DelegatedAdminAccessContainerType = iota
    UNKNOWNFUTUREVALUE_DELEGATEDADMINACCESSCONTAINERTYPE
)

func (i DelegatedAdminAccessContainerType) String() string {
    return []string{"SECURITYGROUP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedAdminAccessContainerType(v string) (interface{}, error) {
    result := SECURITYGROUP_DELEGATEDADMINACCESSCONTAINERTYPE
    switch strings.ToUpper(v) {
        case "SECURITYGROUP":
            result = SECURITYGROUP_DELEGATEDADMINACCESSCONTAINERTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINACCESSCONTAINERTYPE
        default:
            return 0, errors.New("Unknown DelegatedAdminAccessContainerType value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminAccessContainerType(values []DelegatedAdminAccessContainerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
