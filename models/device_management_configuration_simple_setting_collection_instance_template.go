package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate 
type DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
}
// NewDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate instantiates a new DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate()(*DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) {
    m := &DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstanceTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate(), nil
}
// GetAllowUnmanagedValues gets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) GetAllowUnmanagedValues()(*bool) {
    val, err := m.GetBackingStore().Get("allowUnmanagedValues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSimpleSettingValueTemplateable)
                }
            }
            m.SetSimpleSettingCollectionValueTemplate(res)
        }
        return nil
    }
    return res
}
// GetSimpleSettingCollectionValueTemplate gets the simpleSettingCollectionValueTemplate property value. Simple Setting Collection Value Template
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) GetSimpleSettingCollectionValueTemplate()([]DeviceManagementConfigurationSimpleSettingValueTemplateable) {
    val, err := m.GetBackingStore().Get("simpleSettingCollectionValueTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSimpleSettingValueTemplateable)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("simpleSettingCollectionValueTemplate", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUnmanagedValues sets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) SetAllowUnmanagedValues(value *bool)() {
    err := m.GetBackingStore().Set("allowUnmanagedValues", value)
    if err != nil {
        panic(err)
    }
}
// SetSimpleSettingCollectionValueTemplate sets the simpleSettingCollectionValueTemplate property value. Simple Setting Collection Value Template
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) SetSimpleSettingCollectionValueTemplate(value []DeviceManagementConfigurationSimpleSettingValueTemplateable)() {
    err := m.GetBackingStore().Set("simpleSettingCollectionValueTemplate", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplateable 
type DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplateable interface {
    DeviceManagementConfigurationSettingInstanceTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowUnmanagedValues()(*bool)
    GetSimpleSettingCollectionValueTemplate()([]DeviceManagementConfigurationSimpleSettingValueTemplateable)
    SetAllowUnmanagedValues(value *bool)()
    SetSimpleSettingCollectionValueTemplate(value []DeviceManagementConfigurationSimpleSettingValueTemplateable)()
}
