package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSecretSettingValue graph model for a secret setting value
type DeviceManagementConfigurationSecretSettingValue struct {
    DeviceManagementConfigurationSimpleSettingValue
}
// NewDeviceManagementConfigurationSecretSettingValue instantiates a new deviceManagementConfigurationSecretSettingValue and sets the default values.
func NewDeviceManagementConfigurationSecretSettingValue()(*DeviceManagementConfigurationSecretSettingValue) {
    m := &DeviceManagementConfigurationSecretSettingValue{
        DeviceManagementConfigurationSimpleSettingValue: *NewDeviceManagementConfigurationSimpleSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSecretSettingValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationSecretSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSecretSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSecretSettingValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSecretSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingValue.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["valueState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSecretSettingValueState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueState(val.(*DeviceManagementConfigurationSecretSettingValueState))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSecretSettingValue) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValue gets the value property value. Value of the secret setting.
func (m *DeviceManagementConfigurationSecretSettingValue) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValueState gets the valueState property value. type tracking the encryption state of a secret setting value
func (m *DeviceManagementConfigurationSecretSettingValue) GetValueState()(*DeviceManagementConfigurationSecretSettingValueState) {
    val, err := m.GetBackingStore().Get("valueState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationSecretSettingValueState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSecretSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSimpleSettingValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    if m.GetValueState() != nil {
        cast := (*m.GetValueState()).String()
        err = writer.WriteStringValue("valueState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSecretSettingValue) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. Value of the secret setting.
func (m *DeviceManagementConfigurationSecretSettingValue) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SetValueState sets the valueState property value. type tracking the encryption state of a secret setting value
func (m *DeviceManagementConfigurationSecretSettingValue) SetValueState(value *DeviceManagementConfigurationSecretSettingValueState)() {
    err := m.GetBackingStore().Set("valueState", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSecretSettingValueable 
type DeviceManagementConfigurationSecretSettingValueable interface {
    DeviceManagementConfigurationSimpleSettingValueable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetValue()(*string)
    GetValueState()(*DeviceManagementConfigurationSecretSettingValueState)
    SetOdataType(value *string)()
    SetValue(value *string)()
    SetValueState(value *DeviceManagementConfigurationSecretSettingValueState)()
}
