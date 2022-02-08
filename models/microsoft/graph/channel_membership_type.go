package graph
import (
    "strings"
    "errors"
)
// 
type ChannelMembershipType int

const (
    STANDARD_CHANNELMEMBERSHIPTYPE ChannelMembershipType = iota
    PRIVATE_CHANNELMEMBERSHIPTYPE
    UNKNOWNFUTUREVALUE_CHANNELMEMBERSHIPTYPE
    SHARED_CHANNELMEMBERSHIPTYPE
)

func (i ChannelMembershipType) String() string {
    return []string{"STANDARD", "PRIVATE", "UNKNOWNFUTUREVALUE", "SHARED"}[i]
}
func ParseChannelMembershipType(v string) (interface{}, error) {
    result := STANDARD_CHANNELMEMBERSHIPTYPE
    switch strings.ToUpper(v) {
        case "STANDARD":
            result = STANDARD_CHANNELMEMBERSHIPTYPE
        case "PRIVATE":
            result = PRIVATE_CHANNELMEMBERSHIPTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CHANNELMEMBERSHIPTYPE
        case "SHARED":
            result = SHARED_CHANNELMEMBERSHIPTYPE
        default:
            return 0, errors.New("Unknown ChannelMembershipType value: " + v)
    }
    return &result, nil
}
func SerializeChannelMembershipType(values []ChannelMembershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
