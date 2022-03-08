package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuditActorable 
type AuditActorable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationDisplayName()(*string)
    GetApplicationId()(*string)
    GetIpAddress()(*string)
    GetRemoteTenantId()(*string)
    GetRemoteUserId()(*string)
    GetServicePrincipalName()(*string)
    GetType()(*string)
    GetUserId()(*string)
    GetUserPermissions()([]string)
    GetUserPrincipalName()(*string)
    GetUserRoleScopeTags()([]RoleScopeTagInfoable)
    SetApplicationDisplayName(value *string)()
    SetApplicationId(value *string)()
    SetIpAddress(value *string)()
    SetRemoteTenantId(value *string)()
    SetRemoteUserId(value *string)()
    SetServicePrincipalName(value *string)()
    SetType(value *string)()
    SetUserId(value *string)()
    SetUserPermissions(value []string)()
    SetUserPrincipalName(value *string)()
    SetUserRoleScopeTags(value []RoleScopeTagInfoable)()
}
