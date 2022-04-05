package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SensitivityLabelable 
type SensitivityLabelable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetContentFormats()([]string)
    GetDescription()(*string)
    GetHasProtection()(*bool)
    GetIsActive()(*bool)
    GetIsAppliable()(*bool)
    GetName()(*string)
    GetParent()(SensitivityLabelable)
    GetSensitivity()(*int32)
    GetTooltip()(*string)
    SetColor(value *string)()
    SetContentFormats(value []string)()
    SetDescription(value *string)()
    SetHasProtection(value *bool)()
    SetIsActive(value *bool)()
    SetIsAppliable(value *bool)()
    SetName(value *string)()
    SetParent(value SensitivityLabelable)()
    SetSensitivity(value *int32)()
    SetTooltip(value *string)()
}
