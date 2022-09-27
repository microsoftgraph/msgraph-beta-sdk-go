package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptStringParameter 
type DeviceHealthScriptStringParameter struct {
    DeviceHealthScriptParameter
    // The default value of string param
    defaultValue *string
}
// NewDeviceHealthScriptStringParameter instantiates a new DeviceHealthScriptStringParameter and sets the default values.
func NewDeviceHealthScriptStringParameter()(*DeviceHealthScriptStringParameter) {
    m := &DeviceHealthScriptStringParameter{
        DeviceHealthScriptParameter: *NewDeviceHealthScriptParameter(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptStringParameter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptStringParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptStringParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptStringParameter(), nil
}
// GetDefaultValue gets the defaultValue property value. The default value of string param
func (m *DeviceHealthScriptStringParameter) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptStringParameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceHealthScriptParameter.GetFieldDeserializers()
    res["defaultValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultValue)
    return res
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptStringParameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceHealthScriptParameter.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. The default value of string param
func (m *DeviceHealthScriptStringParameter) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
