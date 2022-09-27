package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingFileConstraint 
type DeviceManagementSettingFileConstraint struct {
    DeviceManagementConstraint
    // Acceptable file extensions to upload for this setting
    supportedExtensions []string
}
// NewDeviceManagementSettingFileConstraint instantiates a new DeviceManagementSettingFileConstraint and sets the default values.
func NewDeviceManagementSettingFileConstraint()(*DeviceManagementSettingFileConstraint) {
    m := &DeviceManagementSettingFileConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingFileConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingFileConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingFileConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingFileConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingFileConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["supportedExtensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSupportedExtensions)
    return res
}
// GetSupportedExtensions gets the supportedExtensions property value. Acceptable file extensions to upload for this setting
func (m *DeviceManagementSettingFileConstraint) GetSupportedExtensions()([]string) {
    return m.supportedExtensions
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingFileConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSupportedExtensions() != nil {
        err = writer.WriteCollectionOfStringValues("supportedExtensions", m.GetSupportedExtensions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSupportedExtensions sets the supportedExtensions property value. Acceptable file extensions to upload for this setting
func (m *DeviceManagementSettingFileConstraint) SetSupportedExtensions(value []string)() {
    m.supportedExtensions = value
}
