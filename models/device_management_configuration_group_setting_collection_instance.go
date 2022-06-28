package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingCollectionInstance 
type DeviceManagementConfigurationGroupSettingCollectionInstance struct {
    DeviceManagementConfigurationSettingInstance
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of GroupSetting values
    groupSettingCollectionValue []DeviceManagementConfigurationGroupSettingValueable
}
// NewDeviceManagementConfigurationGroupSettingCollectionInstance instantiates a new DeviceManagementConfigurationGroupSettingCollectionInstance and sets the default values.
func NewDeviceManagementConfigurationGroupSettingCollectionInstance()(*DeviceManagementConfigurationGroupSettingCollectionInstance) {
    m := &DeviceManagementConfigurationGroupSettingCollectionInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationGroupSettingCollectionInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationGroupSettingCollectionInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingCollectionInstance(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationGroupSettingCollectionInstance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationGroupSettingCollectionInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["groupSettingCollectionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationGroupSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationGroupSettingValueable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationGroupSettingValueable)
            }
            m.SetGroupSettingCollectionValue(res)
        }
        return nil
    }
    return res
}
// GetGroupSettingCollectionValue gets the groupSettingCollectionValue property value. A collection of GroupSetting values
func (m *DeviceManagementConfigurationGroupSettingCollectionInstance) GetGroupSettingCollectionValue()([]DeviceManagementConfigurationGroupSettingValueable) {
    if m == nil {
        return nil
    } else {
        return m.groupSettingCollectionValue
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationGroupSettingCollectionInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetGroupSettingCollectionValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupSettingCollectionValue()))
        for i, v := range m.GetGroupSettingCollectionValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupSettingCollectionValue", cast)
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
func (m *DeviceManagementConfigurationGroupSettingCollectionInstance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupSettingCollectionValue sets the groupSettingCollectionValue property value. A collection of GroupSetting values
func (m *DeviceManagementConfigurationGroupSettingCollectionInstance) SetGroupSettingCollectionValue(value []DeviceManagementConfigurationGroupSettingValueable)() {
    if m != nil {
        m.groupSettingCollectionValue = value
    }
}
