package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingInstance 
type DeviceManagementConfigurationGroupSettingInstance struct {
    DeviceManagementConfigurationSettingInstance
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupSettingValue property
    groupSettingValue DeviceManagementConfigurationGroupSettingValueable
}
// NewDeviceManagementConfigurationGroupSettingInstance instantiates a new DeviceManagementConfigurationGroupSettingInstance and sets the default values.
func NewDeviceManagementConfigurationGroupSettingInstance()(*DeviceManagementConfigurationGroupSettingInstance) {
    m := &DeviceManagementConfigurationGroupSettingInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationGroupSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationGroupSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingInstance(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationGroupSettingInstance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationGroupSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["groupSettingValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationGroupSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupSettingValue(val.(DeviceManagementConfigurationGroupSettingValueable))
        }
        return nil
    }
    return res
}
// GetGroupSettingValue gets the groupSettingValue property value. The groupSettingValue property
func (m *DeviceManagementConfigurationGroupSettingInstance) GetGroupSettingValue()(DeviceManagementConfigurationGroupSettingValueable) {
    if m == nil {
        return nil
    } else {
        return m.groupSettingValue
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationGroupSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("groupSettingValue", m.GetGroupSettingValue())
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
func (m *DeviceManagementConfigurationGroupSettingInstance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupSettingValue sets the groupSettingValue property value. The groupSettingValue property
func (m *DeviceManagementConfigurationGroupSettingInstance) SetGroupSettingValue(value DeviceManagementConfigurationGroupSettingValueable)() {
    if m != nil {
        m.groupSettingValue = value
    }
}
