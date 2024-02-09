package models
import (
    "errors"
)
type PermissionsDefinitionIdentityType int

const (
    USER_PERMISSIONSDEFINITIONIDENTITYTYPE PermissionsDefinitionIdentityType = iota
    ROLE_PERMISSIONSDEFINITIONIDENTITYTYPE
    APPLICATION_PERMISSIONSDEFINITIONIDENTITYTYPE
    MANAGEDIDENTITY_PERMISSIONSDEFINITIONIDENTITYTYPE
    SERVICEACCOUNT_PERMISSIONSDEFINITIONIDENTITYTYPE
    UNKNOWNFUTUREVALUE_PERMISSIONSDEFINITIONIDENTITYTYPE
)

func (i PermissionsDefinitionIdentityType) String() string {
    return []string{"user", "role", "application", "managedIdentity", "serviceAccount", "unknownFutureValue"}[i]
}
func ParsePermissionsDefinitionIdentityType(v string) (any, error) {
    result := USER_PERMISSIONSDEFINITIONIDENTITYTYPE
    switch v {
        case "user":
            result = USER_PERMISSIONSDEFINITIONIDENTITYTYPE
        case "role":
            result = ROLE_PERMISSIONSDEFINITIONIDENTITYTYPE
        case "application":
            result = APPLICATION_PERMISSIONSDEFINITIONIDENTITYTYPE
        case "managedIdentity":
            result = MANAGEDIDENTITY_PERMISSIONSDEFINITIONIDENTITYTYPE
        case "serviceAccount":
            result = SERVICEACCOUNT_PERMISSIONSDEFINITIONIDENTITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERMISSIONSDEFINITIONIDENTITYTYPE
        default:
            return 0, errors.New("Unknown PermissionsDefinitionIdentityType value: " + v)
    }
    return &result, nil
}
func SerializePermissionsDefinitionIdentityType(values []PermissionsDefinitionIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PermissionsDefinitionIdentityType) isMultiValue() bool {
    return false
}
