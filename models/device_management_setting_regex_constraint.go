package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingRegexConstraint 
type DeviceManagementSettingRegexConstraint struct {
    DeviceManagementConstraint
    // The RegEx pattern to match against
    regex *string
}
// NewDeviceManagementSettingRegexConstraint instantiates a new DeviceManagementSettingRegexConstraint and sets the default values.
func NewDeviceManagementSettingRegexConstraint()(*DeviceManagementSettingRegexConstraint) {
    m := &DeviceManagementSettingRegexConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingRegexConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingRegexConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingRegexConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingRegexConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingRegexConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["regex"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegex)
    return res
}
// GetRegex gets the regex property value. The RegEx pattern to match against
func (m *DeviceManagementSettingRegexConstraint) GetRegex()(*string) {
    return m.regex
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingRegexConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("regex", m.GetRegex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRegex sets the regex property value. The RegEx pattern to match against
func (m *DeviceManagementSettingRegexConstraint) SetRegex(value *string)() {
    m.regex = value
}
