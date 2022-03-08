package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthorizationPolicyable 
type AuthorizationPolicyable interface {
    PolicyBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowedToSignUpEmailBasedSubscriptions()(*bool)
    GetAllowedToUseSSPR()(*bool)
    GetAllowEmailVerifiedUsersToJoinOrganization()(*bool)
    GetAllowInvitesFrom()(*AllowInvitesFrom)
    GetBlockMsolPowerShell()(*bool)
    GetDefaultUserRoleOverrides()([]DefaultUserRoleOverrideable)
    GetDefaultUserRolePermissions()(DefaultUserRolePermissionsable)
    GetEnabledPreviewFeatures()([]string)
    GetGuestUserRoleId()(*string)
    GetPermissionGrantPolicyIdsAssignedToDefaultUserRole()([]string)
    SetAllowedToSignUpEmailBasedSubscriptions(value *bool)()
    SetAllowedToUseSSPR(value *bool)()
    SetAllowEmailVerifiedUsersToJoinOrganization(value *bool)()
    SetAllowInvitesFrom(value *AllowInvitesFrom)()
    SetBlockMsolPowerShell(value *bool)()
    SetDefaultUserRoleOverrides(value []DefaultUserRoleOverrideable)()
    SetDefaultUserRolePermissions(value DefaultUserRolePermissionsable)()
    SetEnabledPreviewFeatures(value []string)()
    SetGuestUserRoleId(value *string)()
    SetPermissionGrantPolicyIdsAssignedToDefaultUserRole(value []string)()
}
