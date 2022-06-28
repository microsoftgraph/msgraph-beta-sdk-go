package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthorizationPolicyable 
type AuthorizationPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
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
