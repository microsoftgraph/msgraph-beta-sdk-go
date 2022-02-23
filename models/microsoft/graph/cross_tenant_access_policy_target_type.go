package graph
import (
    "strings"
    "errors"
)
// 
type CrossTenantAccessPolicyTargetType int

const (
    USER_CROSSTENANTACCESSPOLICYTARGETTYPE CrossTenantAccessPolicyTargetType = iota
    GROUP_CROSSTENANTACCESSPOLICYTARGETTYPE
    APPLICATION_CROSSTENANTACCESSPOLICYTARGETTYPE
    UNKNOWNFUTUREVALUE_CROSSTENANTACCESSPOLICYTARGETTYPE
)

func (i CrossTenantAccessPolicyTargetType) String() string {
    return []string{"USER", "GROUP", "APPLICATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCrossTenantAccessPolicyTargetType(v string) (interface{}, error) {
    result := USER_CROSSTENANTACCESSPOLICYTARGETTYPE
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_CROSSTENANTACCESSPOLICYTARGETTYPE
        case "GROUP":
            result = GROUP_CROSSTENANTACCESSPOLICYTARGETTYPE
        case "APPLICATION":
            result = APPLICATION_CROSSTENANTACCESSPOLICYTARGETTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CROSSTENANTACCESSPOLICYTARGETTYPE
        default:
            return 0, errors.New("Unknown CrossTenantAccessPolicyTargetType value: " + v)
    }
    return &result, nil
}
func SerializeCrossTenantAccessPolicyTargetType(values []CrossTenantAccessPolicyTargetType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
