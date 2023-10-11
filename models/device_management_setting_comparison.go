package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceManagementSettingComparison entity representing setting comparison result
type DeviceManagementSettingComparison struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementSettingComparison instantiates a new deviceManagementSettingComparison and sets the default values.
func NewDeviceManagementSettingComparison()(*DeviceManagementSettingComparison) {
    m := &DeviceManagementSettingComparison{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementSettingComparisonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingComparisonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingComparison(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingComparison) GetAdditionalData()(map[string]any) {
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
func (m *DeviceManagementSettingComparison) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetComparisonResult gets the comparisonResult property value. Setting comparison result type
func (m *DeviceManagementSettingComparison) GetComparisonResult()(*DeviceManagementComparisonResult) {
    val, err := m.GetBackingStore().Get("comparisonResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementComparisonResult)
    }
    return nil
}
// GetCurrentValueJson gets the currentValueJson property value. JSON representation of current intent (or) template setting's value
func (m *DeviceManagementSettingComparison) GetCurrentValueJson()(*string) {
    val, err := m.GetBackingStore().Get("currentValueJson")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDefinitionId gets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingComparison) GetDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("definitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The setting's display name
func (m *DeviceManagementSettingComparison) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingComparison) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comparisonResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementComparisonResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComparisonResult(val.(*DeviceManagementComparisonResult))
        }
        return nil
    }
    res["currentValueJson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValueJson(val)
        }
        return nil
    }
    res["definitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["newValueJson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewValueJson(val)
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
    return res
}
// GetId gets the id property value. The setting ID
func (m *DeviceManagementSettingComparison) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNewValueJson gets the newValueJson property value. JSON representation of new template setting's value
func (m *DeviceManagementSettingComparison) GetNewValueJson()(*string) {
    val, err := m.GetBackingStore().Get("newValueJson")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingComparison) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingComparison) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComparisonResult() != nil {
        cast := (*m.GetComparisonResult()).String()
        err := writer.WriteStringValue("comparisonResult", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currentValueJson", m.GetCurrentValueJson())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("definitionId", m.GetDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newValueJson", m.GetNewValueJson())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingComparison) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceManagementSettingComparison) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetComparisonResult sets the comparisonResult property value. Setting comparison result type
func (m *DeviceManagementSettingComparison) SetComparisonResult(value *DeviceManagementComparisonResult)() {
    err := m.GetBackingStore().Set("comparisonResult", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrentValueJson sets the currentValueJson property value. JSON representation of current intent (or) template setting's value
func (m *DeviceManagementSettingComparison) SetCurrentValueJson(value *string)() {
    err := m.GetBackingStore().Set("currentValueJson", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinitionId sets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingComparison) SetDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("definitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The setting's display name
func (m *DeviceManagementSettingComparison) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. The setting ID
func (m *DeviceManagementSettingComparison) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetNewValueJson sets the newValueJson property value. JSON representation of new template setting's value
func (m *DeviceManagementSettingComparison) SetNewValueJson(value *string)() {
    err := m.GetBackingStore().Set("newValueJson", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingComparison) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementSettingComparisonable 
type DeviceManagementSettingComparisonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetComparisonResult()(*DeviceManagementComparisonResult)
    GetCurrentValueJson()(*string)
    GetDefinitionId()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    GetNewValueJson()(*string)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetComparisonResult(value *DeviceManagementComparisonResult)()
    SetCurrentValueJson(value *string)()
    SetDefinitionId(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetNewValueJson(value *string)()
    SetOdataType(value *string)()
}
