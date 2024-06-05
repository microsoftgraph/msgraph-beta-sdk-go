package models
type PermissionKind int

const (
    ALL_PERMISSIONKIND PermissionKind = iota
    ENUMERATED_PERMISSIONKIND
    ALLPERMISSIONSONRESOURCEAPP_PERMISSIONKIND
    UNKNOWNFUTUREVALUE_PERMISSIONKIND
)

func (i PermissionKind) String() string {
    return []string{"all", "enumerated", "allPermissionsOnResourceApp", "unknownFutureValue"}[i]
}
func ParsePermissionKind(v string) (any, error) {
    result := ALL_PERMISSIONKIND
    switch v {
        case "all":
            result = ALL_PERMISSIONKIND
        case "enumerated":
            result = ENUMERATED_PERMISSIONKIND
        case "allPermissionsOnResourceApp":
            result = ALLPERMISSIONSONRESOURCEAPP_PERMISSIONKIND
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERMISSIONKIND
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePermissionKind(values []PermissionKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PermissionKind) isMultiValue() bool {
    return false
}
