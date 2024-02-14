package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationReferenceSettingValue model for ReferenceSettingValue
type DeviceManagementConfigurationReferenceSettingValue struct {
    DeviceManagementConfigurationStringSettingValue
}
// NewDeviceManagementConfigurationReferenceSettingValue instantiates a new DeviceManagementConfigurationReferenceSettingValue and sets the default values.
func NewDeviceManagementConfigurationReferenceSettingValue()(*DeviceManagementConfigurationReferenceSettingValue) {
    m := &DeviceManagementConfigurationReferenceSettingValue{
        DeviceManagementConfigurationStringSettingValue: *NewDeviceManagementConfigurationStringSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationReferenceSettingValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationReferenceSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementConfigurationReferenceSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationReferenceSettingValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementConfigurationReferenceSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationStringSettingValue.GetFieldDeserializers()
    res["note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    return res
}
// GetNote gets the note property value. A note that admin can use to put some contextual information
// returns a *string when successful
func (m *DeviceManagementConfigurationReferenceSettingValue) GetNote()(*string) {
    val, err := m.GetBackingStore().Get("note")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationReferenceSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationStringSettingValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("note", m.GetNote())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNote sets the note property value. A note that admin can use to put some contextual information
func (m *DeviceManagementConfigurationReferenceSettingValue) SetNote(value *string)() {
    err := m.GetBackingStore().Set("note", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementConfigurationReferenceSettingValueable interface {
    DeviceManagementConfigurationStringSettingValueable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNote()(*string)
    SetNote(value *string)()
}
