package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceManagementConfigurationSettingInstance setting instance within policy
type DeviceManagementConfigurationSettingInstance struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementConfigurationSettingInstance instantiates a new DeviceManagementConfigurationSettingInstance and sets the default values.
func NewDeviceManagementConfigurationSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    m := &DeviceManagementConfigurationSettingInstance{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstance":
                        return NewDeviceManagementConfigurationChoiceSettingCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance":
                        return NewDeviceManagementConfigurationChoiceSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstance":
                        return NewDeviceManagementConfigurationGroupSettingCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationGroupSettingInstance":
                        return NewDeviceManagementConfigurationGroupSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionInstance":
                        return NewDeviceManagementConfigurationSettingGroupCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingGroupInstance":
                        return NewDeviceManagementConfigurationSettingGroupInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstance":
                        return NewDeviceManagementConfigurationSimpleSettingCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance":
                        return NewDeviceManagementConfigurationSimpleSettingInstance(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingInstance(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstance) GetAdditionalData()(map[string]any) {
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
func (m *DeviceManagementConfigurationSettingInstance) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["settingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionId(val)
        }
        return nil
    }
    res["settingInstanceTemplateReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingInstanceTemplateReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplateReference(val.(DeviceManagementConfigurationSettingInstanceTemplateReferenceable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingInstance) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingDefinitionId gets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstance) GetSettingDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("settingDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingInstanceTemplateReference gets the settingInstanceTemplateReference property value. Setting Instance Template Reference
func (m *DeviceManagementConfigurationSettingInstance) GetSettingInstanceTemplateReference()(DeviceManagementConfigurationSettingInstanceTemplateReferenceable) {
    val, err := m.GetBackingStore().Get("settingInstanceTemplateReference")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSettingInstanceTemplateReferenceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settingInstanceTemplateReference", m.GetSettingInstanceTemplateReference())
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
func (m *DeviceManagementConfigurationSettingInstance) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceManagementConfigurationSettingInstance) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingInstance) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDefinitionId sets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstance) SetSettingDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("settingDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingInstanceTemplateReference sets the settingInstanceTemplateReference property value. Setting Instance Template Reference
func (m *DeviceManagementConfigurationSettingInstance) SetSettingInstanceTemplateReference(value DeviceManagementConfigurationSettingInstanceTemplateReferenceable)() {
    err := m.GetBackingStore().Set("settingInstanceTemplateReference", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSettingInstanceable 
type DeviceManagementConfigurationSettingInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetSettingDefinitionId()(*string)
    GetSettingInstanceTemplateReference()(DeviceManagementConfigurationSettingInstanceTemplateReferenceable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetSettingDefinitionId(value *string)()
    SetSettingInstanceTemplateReference(value DeviceManagementConfigurationSettingInstanceTemplateReferenceable)()
}
