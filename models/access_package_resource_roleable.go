package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceRoleable 
type AccessPackageResourceRoleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
