package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingInstance 
type DeviceManagementConfigurationSimpleSettingInstance struct {
    DeviceManagementConfigurationSettingInstance
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The simpleSettingValue property
    simpleSettingValue DeviceManagementConfigurationSimpleSettingValueable
}
// NewDeviceManagementConfigurationSimpleSettingInstance instantiates a new DeviceManagementConfigurationSimpleSettingInstance and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingInstance()(*DeviceManagementConfigurationSimpleSettingInstance) {
    m := &DeviceManagementConfigurationSimpleSettingInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingInstance(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
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
// GetSimpleSettingValue gets the simpleSettingValue property value. The simpleSettingValue property
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetSimpleSettingValue()(DeviceManagementConfigurationSimpleSettingValueable) {
    if m == nil {
        return nil
    } else {
        return m.simpleSettingValue
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("simpleSettingValue", m.GetSimpleSettingValue())
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
func (m *DeviceManagementConfigurationSimpleSettingInstance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSimpleSettingValue sets the simpleSettingValue property value. The simpleSettingValue property
func (m *DeviceManagementConfigurationSimpleSettingInstance) SetSimpleSettingValue(value DeviceManagementConfigurationSimpleSettingValueable)() {
    if m != nil {
        m.simpleSettingValue = value
    }
}
