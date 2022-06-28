package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementActionable 
type ManagementActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetReferenceTemplateId()(*string)
    GetReferenceTemplateVersion()(*int32)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetReferenceTemplateId(value *string)()
    SetReferenceTemplateVersion(value *int32)()
}
