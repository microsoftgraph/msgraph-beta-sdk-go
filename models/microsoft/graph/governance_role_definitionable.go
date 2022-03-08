package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceRoleDefinitionable 
type GovernanceRoleDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetResource()(GovernanceResourceable)
    GetResourceId()(*string)
    GetRoleSetting()(GovernanceRoleSettingable)
    GetTemplateId()(*string)
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetResource(value GovernanceResourceable)()
    SetResourceId(value *string)()
    SetRoleSetting(value GovernanceRoleSettingable)()
    SetTemplateId(value *string)()
}
