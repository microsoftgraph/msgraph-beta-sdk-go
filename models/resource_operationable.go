package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceOperationable 
type ResourceOperationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionName()(*string)
    GetDescription()(*string)
    GetEnabledForScopeValidation()(*bool)
    GetResource()(*string)
    GetResourceName()(*string)
    SetActionName(value *string)()
    SetDescription(value *string)()
    SetEnabledForScopeValidation(value *bool)()
    SetResource(value *string)()
    SetResourceName(value *string)()
}
