package models
// Android action taken when an app requests a dangerous permission.
type AndroidPermissionActionType int

const (
    PROMPT_ANDROIDPERMISSIONACTIONTYPE AndroidPermissionActionType = iota
    AUTOGRANT_ANDROIDPERMISSIONACTIONTYPE
    AUTODENY_ANDROIDPERMISSIONACTIONTYPE
)

func (i AndroidPermissionActionType) String() string {
    return []string{"prompt", "autoGrant", "autoDeny"}[i]
}
func ParseAndroidPermissionActionType(v string) (any, error) {
    result := PROMPT_ANDROIDPERMISSIONACTIONTYPE
    switch v {
        case "prompt":
            result = PROMPT_ANDROIDPERMISSIONACTIONTYPE
        case "autoGrant":
            result = AUTOGRANT_ANDROIDPERMISSIONACTIONTYPE
        case "autoDeny":
            result = AUTODENY_ANDROIDPERMISSIONACTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidPermissionActionType(values []AndroidPermissionActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidPermissionActionType) isMultiValue() bool {
    return false
}
