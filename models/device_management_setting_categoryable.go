package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingCategoryable 
type DeviceManagementSettingCategoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetHasRequiredSetting()(*bool)
    GetSettingDefinitions()([]DeviceManagementSettingDefinitionable)
    GetType()(*string)
    SetDisplayName(value *string)()
    SetHasRequiredSetting(value *bool)()
    SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)()
    SetType(value *string)()
}
