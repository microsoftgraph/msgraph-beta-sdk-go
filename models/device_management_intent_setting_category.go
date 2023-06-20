package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentSettingCategory 
type DeviceManagementIntentSettingCategory struct {
    DeviceManagementSettingCategory
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
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingInstanceable)
                }
            }
            m.SetSettings(res)
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The settings this category contains
func (m *DeviceManagementIntentSettingCategory) GetSettings()([]DeviceManagementSettingInstanceable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingInstanceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentSettingCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingCategory.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettings sets the settings property value. The settings this category contains
func (m *DeviceManagementIntentSettingCategory) SetSettings(value []DeviceManagementSettingInstanceable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementIntentSettingCategoryable 
type DeviceManagementIntentSettingCategoryable interface {
    DeviceManagementSettingCategoryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettings()([]DeviceManagementSettingInstanceable)
    SetSettings(value []DeviceManagementSettingInstanceable)()
}
