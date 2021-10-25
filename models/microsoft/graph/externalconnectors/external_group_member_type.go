package externalconnectors
import (
    "strings"
    "errors"
)
type ExternalGroupMemberType int

const (
    USER_EXTERNALGROUPMEMBERTYPE ExternalGroupMemberType = iota
    GROUP_EXTERNALGROUPMEMBERTYPE
    UNKNOWNFUTUREVALUE_EXTERNALGROUPMEMBERTYPE
)

func (i ExternalGroupMemberType) String() string {
    return []string{"USER", "GROUP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseExternalGroupMemberType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "USER":
            return USER_EXTERNALGROUPMEMBERTYPE, nil
        case "GROUP":
            return GROUP_EXTERNALGROUPMEMBERTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EXTERNALGROUPMEMBERTYPE, nil
    }
    return 0, errors.New("Unknown ExternalGroupMemberType value: " + v)
}
func SerializeExternalGroupMemberType(values []ExternalGroupMemberType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
