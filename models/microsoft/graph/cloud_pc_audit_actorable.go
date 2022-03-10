package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcAuditActorable 
type CloudPcAuditActorable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationDisplayName()(*string)
    GetApplicationId()(*string)
    GetIpAddress()(*string)
    GetRemoteTenantId()(*string)
    GetRemoteUserId()(*string)
    GetServicePrincipalName()(*string)
    GetType()(*CloudPcAuditActorType)
    GetUserId()(*string)
    GetUserPermissions()([]string)
    GetUserPrincipalName()(*string)
    GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfoable)
    SetApplicationDisplayName(value *string)()
    SetApplicationId(value *string)()
    SetIpAddress(value *string)()
    SetRemoteTenantId(value *string)()
    SetRemoteUserId(value *string)()
    SetServicePrincipalName(value *string)()
    SetType(value *CloudPcAuditActorType)()
    SetUserId(value *string)()
    SetUserPermissions(value []string)()
    SetUserPrincipalName(value *string)()
    SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfoable)()
}
