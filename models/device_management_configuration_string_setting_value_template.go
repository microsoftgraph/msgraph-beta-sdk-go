package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationStringSettingValueTemplate 
type DeviceManagementConfigurationStringSettingValueTemplate struct {
    DeviceManagementConfigurationSimpleSettingValueTemplate
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // String Setting Value Default Template.
    defaultValue DeviceManagementConfigurationStringSettingValueDefaultTemplateable
}
// NewDeviceManagementConfigurationStringSettingValueTemplate instantiates a new DeviceManagementConfigurationStringSettingValueTemplate and sets the default values.
func NewDeviceManagementConfigurationStringSettingValueTemplate()(*DeviceManagementConfigurationStringSettingValueTemplate) {
    m := &DeviceManagementConfigurationStringSettingValueTemplate{
        DeviceManagementConfigurationSimpleSettingValueTemplate: *NewDeviceManagementConfigurationSimpleSettingValueTemplate(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationStringSettingValueTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationStringSettingValueTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationStringSettingValueTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationStringSettingValueTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultValue gets the defaultValue property value. String Setting Value Default Template.
func (m *DeviceManagementConfigurationStringSettingValueTemplate) GetDefaultValue()(DeviceManagementConfigurationStringSettingValueDefaultTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationStringSettingValueTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingValueTemplate.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationStringSettingValueDefaultTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(DeviceManagementConfigurationStringSettingValueDefaultTemplateable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationStringSettingValueTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSimpleSettingValueTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultValue", m.GetDefaultValue())
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
func (m *DeviceManagementConfigurationStringSettingValueTemplate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultValue sets the defaultValue property value. String Setting Value Default Template.
func (m *DeviceManagementConfigurationStringSettingValueTemplate) SetDefaultValue(value DeviceManagementConfigurationStringSettingValueDefaultTemplateable)() {
    if m != nil {
        m.defaultValue = value
    }
}
