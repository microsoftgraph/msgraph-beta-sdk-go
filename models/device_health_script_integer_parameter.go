package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptIntegerParameter properties of the  Integer script parameter.
type DeviceHealthScriptIntegerParameter struct {
    DeviceHealthScriptParameter
    // The OdataType property
    OdataType *string
}
// NewDeviceHealthScriptIntegerParameter instantiates a new deviceHealthScriptIntegerParameter and sets the default values.
func NewDeviceHealthScriptIntegerParameter()(*DeviceHealthScriptIntegerParameter) {
    m := &DeviceHealthScriptIntegerParameter{
        DeviceHealthScriptParameter: *NewDeviceHealthScriptParameter(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptIntegerParameter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceHealthScriptIntegerParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptIntegerParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptIntegerParameter(), nil
}
// GetDefaultValue gets the defaultValue property value. The default value of Integer param. Valid values -2147483648 to 2147483647
func (m *DeviceHealthScriptIntegerParameter) GetDefaultValue()(*int32) {
    val, err := m.GetBackingStore().Get("defaultValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
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
    return nil
}
// SetDefaultValue sets the defaultValue property value. The default value of Integer param. Valid values -2147483648 to 2147483647
func (m *DeviceHealthScriptIntegerParameter) SetDefaultValue(value *int32)() {
    err := m.GetBackingStore().Set("defaultValue", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthScriptIntegerParameterable 
type DeviceHealthScriptIntegerParameterable interface {
    DeviceHealthScriptParameterable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*int32)
    SetDefaultValue(value *int32)()
}
