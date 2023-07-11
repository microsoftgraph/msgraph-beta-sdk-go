package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationIntegerSettingValueTemplate integer Setting Value Template
type DeviceManagementConfigurationIntegerSettingValueTemplate struct {
    DeviceManagementConfigurationSimpleSettingValueTemplate
    // The OdataType property
    OdataType *string
}
// NewDeviceManagementConfigurationIntegerSettingValueTemplate instantiates a new deviceManagementConfigurationIntegerSettingValueTemplate and sets the default values.
func NewDeviceManagementConfigurationIntegerSettingValueTemplate()(*DeviceManagementConfigurationIntegerSettingValueTemplate) {
    m := &DeviceManagementConfigurationIntegerSettingValueTemplate{
        DeviceManagementConfigurationSimpleSettingValueTemplate: *NewDeviceManagementConfigurationSimpleSettingValueTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationIntegerSettingValueTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationIntegerSettingValueTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationIntegerSettingValueTemplate(), nil
}
// GetDefaultValue gets the defaultValue property value. Integer Setting Value Default Template.
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) GetDefaultValue()(DeviceManagementConfigurationIntegerSettingValueDefaultTemplateable) {
    val, err := m.GetBackingStore().Get("defaultValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationIntegerSettingValueDefaultTemplateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingValueTemplate.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationIntegerSettingValueDefaultTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(DeviceManagementConfigurationIntegerSettingValueDefaultTemplateable))
        }
        return nil
    }
    res["recommendedValueDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationIntegerSettingValueDefinitionTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedValueDefinition(val.(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable))
        }
        return nil
    }
    res["requiredValueDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationIntegerSettingValueDefinitionTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredValueDefinition(val.(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable))
        }
        return nil
    }
    return res
}
// GetRecommendedValueDefinition gets the recommendedValueDefinition property value. Recommended value definition.
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) GetRecommendedValueDefinition()(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable) {
    val, err := m.GetBackingStore().Get("recommendedValueDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)
    }
    return nil
}
// GetRequiredValueDefinition gets the requiredValueDefinition property value. Required value definition.
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) GetRequiredValueDefinition()(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable) {
    val, err := m.GetBackingStore().Get("requiredValueDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("recommendedValueDefinition", m.GetRecommendedValueDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requiredValueDefinition", m.GetRequiredValueDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. Integer Setting Value Default Template.
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) SetDefaultValue(value DeviceManagementConfigurationIntegerSettingValueDefaultTemplateable)() {
    err := m.GetBackingStore().Set("defaultValue", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommendedValueDefinition sets the recommendedValueDefinition property value. Recommended value definition.
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) SetRecommendedValueDefinition(value DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)() {
    err := m.GetBackingStore().Set("recommendedValueDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredValueDefinition sets the requiredValueDefinition property value. Required value definition.
func (m *DeviceManagementConfigurationIntegerSettingValueTemplate) SetRequiredValueDefinition(value DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)() {
    err := m.GetBackingStore().Set("requiredValueDefinition", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationIntegerSettingValueTemplateable 
type DeviceManagementConfigurationIntegerSettingValueTemplateable interface {
    DeviceManagementConfigurationSimpleSettingValueTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(DeviceManagementConfigurationIntegerSettingValueDefaultTemplateable)
    GetRecommendedValueDefinition()(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)
    GetRequiredValueDefinition()(DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)
    SetDefaultValue(value DeviceManagementConfigurationIntegerSettingValueDefaultTemplateable)()
    SetRecommendedValueDefinition(value DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)()
    SetRequiredValueDefinition(value DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable)()
}
