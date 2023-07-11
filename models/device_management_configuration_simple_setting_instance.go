package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingInstance setting instance within policy
type DeviceManagementConfigurationSimpleSettingInstance struct {
    DeviceManagementConfigurationSettingInstance
}
// NewDeviceManagementConfigurationSimpleSettingInstance instantiates a new deviceManagementConfigurationSimpleSettingInstance and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingInstance()(*DeviceManagementConfigurationSimpleSettingInstance) {
    m := &DeviceManagementConfigurationSimpleSettingInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
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
    res["simpleSettingValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSimpleSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimpleSettingValue(val.(DeviceManagementConfigurationSimpleSettingValueable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSimpleSettingValue gets the simpleSettingValue property value. The simpleSettingValue property
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetSimpleSettingValue()(DeviceManagementConfigurationSimpleSettingValueable) {
    val, err := m.GetBackingStore().Get("simpleSettingValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSimpleSettingValueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
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
        err = writer.WriteObjectValue("simpleSettingValue", m.GetSimpleSettingValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingInstance) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSimpleSettingValue sets the simpleSettingValue property value. The simpleSettingValue property
func (m *DeviceManagementConfigurationSimpleSettingInstance) SetSimpleSettingValue(value DeviceManagementConfigurationSimpleSettingValueable)() {
    err := m.GetBackingStore().Set("simpleSettingValue", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSimpleSettingInstanceable 
type DeviceManagementConfigurationSimpleSettingInstanceable interface {
    DeviceManagementConfigurationSettingInstanceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetSimpleSettingValue()(DeviceManagementConfigurationSimpleSettingValueable)
    SetOdataType(value *string)()
    SetSimpleSettingValue(value DeviceManagementConfigurationSimpleSettingValueable)()
}
