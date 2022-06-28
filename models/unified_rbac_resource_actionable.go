package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRbacResourceActionable 
type UnifiedRbacResourceActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
