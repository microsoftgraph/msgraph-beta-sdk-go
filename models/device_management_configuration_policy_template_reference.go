package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceManagementConfigurationPolicyTemplateReference policy template reference information
type DeviceManagementConfigurationPolicyTemplateReference struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementConfigurationPolicyTemplateReference instantiates a new deviceManagementConfigurationPolicyTemplateReference and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateReference()(*DeviceManagementConfigurationPolicyTemplateReference) {
    m := &DeviceManagementConfigurationPolicyTemplateReference{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementConfigurationPolicyTemplateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationPolicyTemplateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationPolicyTemplateReference(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["templateDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateDisplayName(val)
        }
        return nil
    }
    res["templateDisplayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateDisplayVersion(val)
        }
        return nil
    }
    res["templateFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTemplateFamily)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateFamily(val.(*DeviceManagementConfigurationTemplateFamily))
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTemplateDisplayName gets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("templateDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTemplateDisplayVersion gets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayVersion()(*string) {
    val, err := m.GetBackingStore().Get("templateDisplayVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTemplateFamily gets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    val, err := m.GetBackingStore().Get("templateFamily")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationTemplateFamily)
    }
    return nil
}
// GetTemplateId gets the templateId property value. Template id
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("templateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicyTemplateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateFamily() != nil {
        cast := (*m.GetTemplateFamily()).String()
        err := writer.WriteStringValue("templateFamily", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateId", m.GetTemplateId())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateDisplayName sets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayName(value *string)() {
    err := m.GetBackingStore().Set("templateDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateDisplayVersion sets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("templateDisplayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateFamily sets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    err := m.GetBackingStore().Set("templateFamily", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateId sets the templateId property value. Template id
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateId(value *string)() {
    err := m.GetBackingStore().Set("templateId", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationPolicyTemplateReferenceable 
type DeviceManagementConfigurationPolicyTemplateReferenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetTemplateDisplayName()(*string)
    GetTemplateDisplayVersion()(*string)
    GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily)
    GetTemplateId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetTemplateDisplayName(value *string)()
    SetTemplateDisplayVersion(value *string)()
    SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)()
    SetTemplateId(value *string)()
}
