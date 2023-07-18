package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate choice Setting Collection Instance Template
type DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
    // The OdataType property
    OdataType *string
}
// NewDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate instantiates a new deviceManagementConfigurationChoiceSettingCollectionInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate()(*DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstanceTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate(), nil
}
// GetAllowUnmanagedValues gets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) GetAllowUnmanagedValues()(*bool) {
    val, err := m.GetBackingStore().Get("allowUnmanagedValues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetChoiceSettingCollectionValueTemplate gets the choiceSettingCollectionValueTemplate property value. Choice Setting Collection Value Template
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) GetChoiceSettingCollectionValueTemplate()([]DeviceManagementConfigurationChoiceSettingValueTemplateable) {
    val, err := m.GetBackingStore().Get("choiceSettingCollectionValueTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationChoiceSettingValueTemplateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["choiceSettingCollectionValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationChoiceSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationChoiceSettingValueTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationChoiceSettingValueTemplateable)
                }
            }
            m.SetChoiceSettingCollectionValueTemplate(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetChoiceSettingCollectionValueTemplate() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChoiceSettingCollectionValueTemplate()))
        for i, v := range m.GetChoiceSettingCollectionValueTemplate() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("choiceSettingCollectionValueTemplate", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUnmanagedValues sets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) SetAllowUnmanagedValues(value *bool)() {
    err := m.GetBackingStore().Set("allowUnmanagedValues", value)
    if err != nil {
        panic(err)
    }
}
// SetChoiceSettingCollectionValueTemplate sets the choiceSettingCollectionValueTemplate property value. Choice Setting Collection Value Template
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) SetChoiceSettingCollectionValueTemplate(value []DeviceManagementConfigurationChoiceSettingValueTemplateable)() {
    err := m.GetBackingStore().Set("choiceSettingCollectionValueTemplate", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplateable 
type DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplateable interface {
    DeviceManagementConfigurationSettingInstanceTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowUnmanagedValues()(*bool)
    GetChoiceSettingCollectionValueTemplate()([]DeviceManagementConfigurationChoiceSettingValueTemplateable)
    SetAllowUnmanagedValues(value *bool)()
    SetChoiceSettingCollectionValueTemplate(value []DeviceManagementConfigurationChoiceSettingValueTemplateable)()
}
