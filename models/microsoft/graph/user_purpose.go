package graph
import (
    "strings"
    "errors"
)
// 
type UserPurpose int

const (
    UNKNOWN_USERPURPOSE UserPurpose = iota
    USER_USERPURPOSE
    LINKED_USERPURPOSE
    SHARED_USERPURPOSE
    ROOM_USERPURPOSE
    EQUIPMENT_USERPURPOSE
    OTHERS_USERPURPOSE
    UNKNOWNFUTUREVALUE_USERPURPOSE
)

func (i UserPurpose) String() string {
    return []string{"UNKNOWN", "USER", "LINKED", "SHARED", "ROOM", "EQUIPMENT", "OTHERS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUserPurpose(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_USERPURPOSE, nil
        case "USER":
            return USER_USERPURPOSE, nil
        case "LINKED":
            return LINKED_USERPURPOSE, nil
        case "SHARED":
            return SHARED_USERPURPOSE, nil
        case "ROOM":
            return ROOM_USERPURPOSE, nil
        case "EQUIPMENT":
            return EQUIPMENT_USERPURPOSE, nil
        case "OTHERS":
            return OTHERS_USERPURPOSE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_USERPURPOSE, nil
    }
    return 0, errors.New("Unknown UserPurpose value: " + v)
}
func SerializeUserPurpose(values []UserPurpose) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
