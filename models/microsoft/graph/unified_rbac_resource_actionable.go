package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRbacResourceActionable 
type UnifiedRbacResourceActionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActionVerb()(*string)
    GetDescription()(*string)
    GetName()(*string)
    GetResourceScope()(UnifiedRbacResourceScopeable)
    GetResourceScopeId()(*string)
    SetActionVerb(value *string)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetResourceScope(value UnifiedRbacResourceScopeable)()
    SetResourceScopeId(value *string)()
}
