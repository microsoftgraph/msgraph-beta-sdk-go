package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceManagementConfigurationChoiceSettingValueTemplate choice Setting Value Template
type DeviceManagementConfigurationChoiceSettingValueTemplate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementConfigurationChoiceSettingValueTemplate instantiates a new DeviceManagementConfigurationChoiceSettingValueTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingValueTemplate()(*DeviceManagementConfigurationChoiceSettingValueTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingValueTemplate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingValueTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingValueTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingValueTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDefaultValue gets the defaultValue property value. Choice Setting Value Default Template.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetDefaultValue()(DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable) {
    val, err := m.GetBackingStore().Get("defaultValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueDefaultTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable))
        }
        return nil
    }
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
    res["recommendedValueDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueDefinitionTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedValueDefinition(val.(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable))
        }
        return nil
    }
    res["requiredValueDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueDefinitionTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredValueDefinition(val.(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable))
        }
        return nil
    }
    res["settingValueTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueTemplateId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecommendedValueDefinition gets the recommendedValueDefinition property value. Recommended definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetRecommendedValueDefinition()(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable) {
    val, err := m.GetBackingStore().Get("recommendedValueDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)
    }
    return nil
}
// GetRequiredValueDefinition gets the requiredValueDefinition property value. Required definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetRequiredValueDefinition()(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable) {
    val, err := m.GetBackingStore().Get("requiredValueDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)
    }
    return nil
}
// GetSettingValueTemplateId gets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetSettingValueTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("settingValueTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recommendedValueDefinition", m.GetRecommendedValueDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requiredValueDefinition", m.GetRequiredValueDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingValueTemplateId", m.GetSettingValueTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDefaultValue sets the defaultValue property value. Choice Setting Value Default Template.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetDefaultValue(value DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable)() {
    err := m.GetBackingStore().Set("defaultValue", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommendedValueDefinition sets the recommendedValueDefinition property value. Recommended definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetRecommendedValueDefinition(value DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)() {
    err := m.GetBackingStore().Set("recommendedValueDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredValueDefinition sets the requiredValueDefinition property value. Required definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetRequiredValueDefinition(value DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)() {
    err := m.GetBackingStore().Set("requiredValueDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingValueTemplateId sets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetSettingValueTemplateId(value *string)() {
    err := m.GetBackingStore().Set("settingValueTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationChoiceSettingValueTemplateable 
type DeviceManagementConfigurationChoiceSettingValueTemplateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDefaultValue()(DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable)
    GetOdataType()(*string)
    GetRecommendedValueDefinition()(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)
    GetRequiredValueDefinition()(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)
    GetSettingValueTemplateId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDefaultValue(value DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable)()
    SetOdataType(value *string)()
    SetRecommendedValueDefinition(value DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)()
    SetRequiredValueDefinition(value DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)()
    SetSettingValueTemplateId(value *string)()
}
