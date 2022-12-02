package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTemplatesItemMigratableToItemCreateInstancePostRequestBodyable 
type DeviceManagementTemplatesItemMigratableToItemCreateInstancePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetRoleScopeTagIds()([]string)
    GetSettingsDelta()([]DeviceManagementSettingInstanceable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetSettingsDelta(value []DeviceManagementSettingInstanceable)()
}
