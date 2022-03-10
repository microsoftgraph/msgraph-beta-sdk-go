package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceScopeable 
type AccessPackageResourceScopeable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageResource()(AccessPackageResourceable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsRootScope()(*bool)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetRoleOriginId()(*string)
    GetUrl()(*string)
    SetAccessPackageResource(value AccessPackageResourceable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsRootScope(value *bool)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetRoleOriginId(value *string)()
    SetUrl(value *string)()
}
