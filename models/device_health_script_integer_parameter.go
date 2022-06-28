package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptIntegerParameter 
type DeviceHealthScriptIntegerParameter struct {
    DeviceHealthScriptParameter
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The default value of Integer param. Valid values -2147483648 to 2147483647
    defaultValue *int32
}
// NewDeviceHealthScriptIntegerParameter instantiates a new DeviceHealthScriptIntegerParameter and sets the default values.
func NewDeviceHealthScriptIntegerParameter()(*DeviceHealthScriptIntegerParameter) {
    m := &DeviceHealthScriptIntegerParameter{
        DeviceHealthScriptParameter: *NewDeviceHealthScriptParameter(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceHealthScriptIntegerParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptIntegerParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptIntegerParameter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptIntegerParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultValue gets the defaultValue property value. The default value of Integer param. Valid values -2147483648 to 2147483647
func (m *DeviceHealthScriptIntegerParameter) GetDefaultValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptIntegerParameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceHealthScriptParameter.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptIntegerParameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceHealthScriptParameter.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptIntegerParameter) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultValue sets the defaultValue property value. The default value of Integer param. Valid values -2147483648 to 2147483647
func (m *DeviceHealthScriptIntegerParameter) SetDefaultValue(value *int32)() {
    if m != nil {
        m.defaultValue = value
    }
}
