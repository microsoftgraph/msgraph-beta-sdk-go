package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingComparisonable 
type DeviceManagementSettingComparisonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComparisonResult()(*DeviceManagementComparisonResult)
    GetCurrentValueJson()(*string)
    GetDefinitionId()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    GetNewValueJson()(*string)
    SetComparisonResult(value *DeviceManagementComparisonResult)()
    SetCurrentValueJson(value *string)()
    SetDefinitionId(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetNewValueJson(value *string)()
}
