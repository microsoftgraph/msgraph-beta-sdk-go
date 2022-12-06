package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentSettingCategory 
type DeviceManagementIntentSettingCategory struct {
    DeviceManagementSettingCategory
    // The settings this category contains
    settings []DeviceManagementSettingInstanceable
}
// NewDeviceManagementIntentSettingCategory instantiates a new DeviceManagementIntentSettingCategory and sets the default values.
func NewDeviceManagementIntentSettingCategory()(*DeviceManagementIntentSettingCategory) {
    m := &DeviceManagementIntentSettingCategory{
        DeviceManagementSettingCategory: *NewDeviceManagementSettingCategory(),
    }
    return m
}
// CreateDeviceManagementIntentSettingCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentSettingCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentSettingCategory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentSettingCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingCategory.GetFieldDeserializers()
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue , m.SetSettings)
    return res
}
// GetSettings gets the settings property value. The settings this category contains
func (m *DeviceManagementIntentSettingCategory) GetSettings()([]DeviceManagementSettingInstanceable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentSettingCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingCategory.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSettings())
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettings sets the settings property value. The settings this category contains
func (m *DeviceManagementIntentSettingCategory) SetSettings(value []DeviceManagementSettingInstanceable)() {
    m.settings = value
}
