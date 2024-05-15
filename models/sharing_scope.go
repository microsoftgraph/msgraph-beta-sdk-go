package models
type SharingScope int

const (
    ANYONE_SHARINGSCOPE SharingScope = iota
    ORGANIZATION_SHARINGSCOPE
    SPECIFICPEOPLE_SHARINGSCOPE
    ANONYMOUS_SHARINGSCOPE
    USERS_SHARINGSCOPE
    UNKNOWNFUTUREVALUE_SHARINGSCOPE
)

func (i SharingScope) String() string {
    return []string{"anyone", "organization", "specificPeople", "anonymous", "users", "unknownFutureValue"}[i]
}
func ParseSharingScope(v string) (any, error) {
    result := ANYONE_SHARINGSCOPE
    switch v {
        case "anyone":
            result = ANYONE_SHARINGSCOPE
        case "organization":
            result = ORGANIZATION_SHARINGSCOPE
        case "specificPeople":
            result = SPECIFICPEOPLE_SHARINGSCOPE
        case "anonymous":
            result = ANONYMOUS_SHARINGSCOPE
        case "users":
            result = USERS_SHARINGSCOPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SHARINGSCOPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSharingScope(values []SharingScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SharingScope) isMultiValue() bool {
    return false
}
