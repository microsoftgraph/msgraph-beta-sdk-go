package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkContentCameraConfiguration 
type TeamworkContentCameraConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkContentCameraConfiguration instantiates a new teamworkContentCameraConfiguration and sets the default values.
func NewTeamworkContentCameraConfiguration()(*TeamworkContentCameraConfiguration) {
    m := &TeamworkContentCameraConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkContentCameraConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkContentCameraConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkContentCameraConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkContentCameraConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkContentCameraConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkContentCameraConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isContentCameraInverted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentCameraInverted(val)
        }
        return nil
    }
    res["isContentCameraOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentCameraOptional(val)
        }
        return nil
    }
    res["isContentEnhancementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentEnhancementEnabled(val)
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
// GetIsContentCameraInverted gets the isContentCameraInverted property value. True if the content camera is inverted.
func (m *TeamworkContentCameraConfiguration) GetIsContentCameraInverted()(*bool) {
    val, err := m.GetBackingStore().Get("isContentCameraInverted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsContentCameraOptional gets the isContentCameraOptional property value. True if the content camera is optional.
func (m *TeamworkContentCameraConfiguration) GetIsContentCameraOptional()(*bool) {
    val, err := m.GetBackingStore().Get("isContentCameraOptional")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsContentEnhancementEnabled gets the isContentEnhancementEnabled property value. True if the content enhancement is enabled.
func (m *TeamworkContentCameraConfiguration) GetIsContentEnhancementEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isContentEnhancementEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkContentCameraConfiguration) GetOdataType()(*string) {
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
func (m *TeamworkContentCameraConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isContentCameraInverted", m.GetIsContentCameraInverted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentCameraOptional", m.GetIsContentCameraOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentEnhancementEnabled", m.GetIsContentEnhancementEnabled())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkContentCameraConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *TeamworkContentCameraConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsContentCameraInverted sets the isContentCameraInverted property value. True if the content camera is inverted.
func (m *TeamworkContentCameraConfiguration) SetIsContentCameraInverted(value *bool)() {
    err := m.GetBackingStore().Set("isContentCameraInverted", value)
    if err != nil {
        panic(err)
    }
}
// SetIsContentCameraOptional sets the isContentCameraOptional property value. True if the content camera is optional.
func (m *TeamworkContentCameraConfiguration) SetIsContentCameraOptional(value *bool)() {
    err := m.GetBackingStore().Set("isContentCameraOptional", value)
    if err != nil {
        panic(err)
    }
}
// SetIsContentEnhancementEnabled sets the isContentEnhancementEnabled property value. True if the content enhancement is enabled.
func (m *TeamworkContentCameraConfiguration) SetIsContentEnhancementEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isContentEnhancementEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkContentCameraConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkContentCameraConfigurationable 
type TeamworkContentCameraConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsContentCameraInverted()(*bool)
    GetIsContentCameraOptional()(*bool)
    GetIsContentEnhancementEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsContentCameraInverted(value *bool)()
    SetIsContentCameraOptional(value *bool)()
    SetIsContentEnhancementEnabled(value *bool)()
    SetOdataType(value *string)()
}
