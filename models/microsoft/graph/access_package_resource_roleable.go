package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceRoleable 
type AccessPackageResourceRoleable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageResource()(AccessPackageResourceable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    SetAccessPackageResource(value AccessPackageResourceable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
}
