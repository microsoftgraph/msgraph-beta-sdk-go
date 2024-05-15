package models
type ResourceScopeType int

const (
    GROUP_RESOURCESCOPETYPE ResourceScopeType = iota
    CHAT_RESOURCESCOPETYPE
    TENANT_RESOURCESCOPETYPE
    UNKNOWNFUTUREVALUE_RESOURCESCOPETYPE
    TEAM_RESOURCESCOPETYPE
)

func (i ResourceScopeType) String() string {
    return []string{"group", "chat", "tenant", "unknownFutureValue", "team"}[i]
}
func ParseResourceScopeType(v string) (any, error) {
    result := GROUP_RESOURCESCOPETYPE
    switch v {
        case "group":
            result = GROUP_RESOURCESCOPETYPE
        case "chat":
            result = CHAT_RESOURCESCOPETYPE
        case "tenant":
            result = TENANT_RESOURCESCOPETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RESOURCESCOPETYPE
        case "team":
            result = TEAM_RESOURCESCOPETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeResourceScopeType(values []ResourceScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ResourceScopeType) isMultiValue() bool {
    return false
}
