package models
import (
    "errors"
)
type PermissionsModificationCapability int

const (
    ENABLED_PERMISSIONSMODIFICATIONCAPABILITY PermissionsModificationCapability = iota
    NOTCONFIGURED_PERMISSIONSMODIFICATIONCAPABILITY
    NORECENTDATACOLLECTED_PERMISSIONSMODIFICATIONCAPABILITY
    UNKNOWNFUTUREVALUE_PERMISSIONSMODIFICATIONCAPABILITY
)

func (i PermissionsModificationCapability) String() string {
    return []string{"enabled", "notConfigured", "noRecentDataCollected", "unknownFutureValue"}[i]
}
func ParsePermissionsModificationCapability(v string) (any, error) {
    result := ENABLED_PERMISSIONSMODIFICATIONCAPABILITY
    switch v {
        case "enabled":
            result = ENABLED_PERMISSIONSMODIFICATIONCAPABILITY
        case "notConfigured":
            result = NOTCONFIGURED_PERMISSIONSMODIFICATIONCAPABILITY
        case "noRecentDataCollected":
            result = NORECENTDATACOLLECTED_PERMISSIONSMODIFICATIONCAPABILITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERMISSIONSMODIFICATIONCAPABILITY
        default:
            return 0, errors.New("Unknown PermissionsModificationCapability value: " + v)
    }
    return &result, nil
}
func SerializePermissionsModificationCapability(values []PermissionsModificationCapability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PermissionsModificationCapability) isMultiValue() bool {
    return false
}
