package graph
import (
    "strings"
    "errors"
)
// 
type CrossTenantAccessPolicyTargetConfigurationAccessType int

const (
    ALLOWED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE CrossTenantAccessPolicyTargetConfigurationAccessType = iota
    BLOCKED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
    UNKNOWNFUTUREVALUE_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
)

func (i CrossTenantAccessPolicyTargetConfigurationAccessType) String() string {
    return []string{"ALLOWED", "BLOCKED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCrossTenantAccessPolicyTargetConfigurationAccessType(v string) (interface{}, error) {
    result := ALLOWED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
    switch strings.ToUpper(v) {
        case "ALLOWED":
            result = ALLOWED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
        case "BLOCKED":
            result = BLOCKED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
        default:
            return 0, errors.New("Unknown CrossTenantAccessPolicyTargetConfigurationAccessType value: " + v)
    }
    return &result, nil
}
func SerializeCrossTenantAccessPolicyTargetConfigurationAccessType(values []CrossTenantAccessPolicyTargetConfigurationAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
