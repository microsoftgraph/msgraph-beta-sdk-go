package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate 
type DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Linked policy may append values which are not present in the template.
    allowUnmanagedValues *bool
    // Simple Setting Collection Value Template
    simpleSettingCollectionValueTemplate []DeviceManagementConfigurationSimpleSettingValueTemplateable
}
// NewDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate instantiates a new DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate()(*DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) {
    m := &DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowUnmanagedValues gets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) GetAllowUnmanagedValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUnmanagedValues
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstanceTemplate.GetFieldDeserializers()
    res["allowUnmanagedValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUnmanagedValues(val)
        }
        return nil
    }
    res["simpleSettingCollectionValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSimpleSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSimpleSettingValueTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSimpleSettingValueTemplateable)
            }
            m.SetSimpleSettingCollectionValueTemplate(res)
        }
        return nil
    }
    return res
}
// GetSimpleSettingCollectionValueTemplate gets the simpleSettingCollectionValueTemplate property value. Simple Setting Collection Value Template
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) GetSimpleSettingCollectionValueTemplate()([]DeviceManagementConfigurationSimpleSettingValueTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.simpleSettingCollectionValueTemplate
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstanceTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowUnmanagedValues", m.GetAllowUnmanagedValues())
        if err != nil {
            return err
        }
    }
    if m.GetSimpleSettingCollectionValueTemplate() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSimpleSettingCollectionValueTemplate()))
        for i, v := range m.GetSimpleSettingCollectionValueTemplate() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("simpleSettingCollectionValueTemplate", cast)
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
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowUnmanagedValues sets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) SetAllowUnmanagedValues(value *bool)() {
    if m != nil {
        m.allowUnmanagedValues = value
    }
}
// SetSimpleSettingCollectionValueTemplate sets the simpleSettingCollectionValueTemplate property value. Simple Setting Collection Value Template
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) SetSimpleSettingCollectionValueTemplate(value []DeviceManagementConfigurationSimpleSettingValueTemplateable)() {
    if m != nil {
        m.simpleSettingCollectionValueTemplate = value
    }
}
