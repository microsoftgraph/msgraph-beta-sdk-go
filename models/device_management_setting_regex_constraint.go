package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingRegexConstraint constraint enforcing the setting matches against a given RegEx pattern
type DeviceManagementSettingRegexConstraint struct {
    DeviceManagementConstraint
}
// NewDeviceManagementSettingRegexConstraint instantiates a new deviceManagementSettingRegexConstraint and sets the default values.
func NewDeviceManagementSettingRegexConstraint()(*DeviceManagementSettingRegexConstraint) {
    m := &DeviceManagementSettingRegexConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingRegexConstraint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementSettingRegexConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingRegexConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingRegexConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingRegexConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["regex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegex(val)
        }
        return nil
    }
    return res
}
// GetRegex gets the regex property value. The RegEx pattern to match against
func (m *DeviceManagementSettingRegexConstraint) GetRegex()(*string) {
    val, err := m.GetBackingStore().Get("regex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("regex", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementSettingRegexConstraintable 
type DeviceManagementSettingRegexConstraintable interface {
    DeviceManagementConstraintable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRegex()(*string)
    SetRegex(value *string)()
}
