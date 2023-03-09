package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WebPartPosition 
type WebPartPosition struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWebPartPosition instantiates a new webPartPosition and sets the default values.
func NewWebPartPosition()(*WebPartPosition) {
    m := &WebPartPosition{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebPartPositionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebPartPositionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebPartPosition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebPartPosition) GetAdditionalData()(map[string]any) {
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
func (m *WebPartPosition) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetColumnId gets the columnId property value. Indicates the identifier of the column where the web part is located.
func (m *WebPartPosition) GetColumnId()(*float64) {
    val, err := m.GetBackingStore().Get("columnId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebPartPosition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["columnId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnId(val)
        }
        return nil
    }
    res["horizontalSectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHorizontalSectionId(val)
        }
        return nil
    }
    res["isInVerticalSection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInVerticalSection(val)
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
    res["webPartIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebPartIndex(val)
        }
        return nil
    }
    return res
}
// GetHorizontalSectionId gets the horizontalSectionId property value. Indicates the horizontal section where the web part is located.
func (m *WebPartPosition) GetHorizontalSectionId()(*float64) {
    val, err := m.GetBackingStore().Get("horizontalSectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetIsInVerticalSection gets the isInVerticalSection property value. Indicates whether the web part is located in the vertical section.
func (m *WebPartPosition) GetIsInVerticalSection()(*bool) {
    val, err := m.GetBackingStore().Get("isInVerticalSection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WebPartPosition) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWebPartIndex gets the webPartIndex property value. Index of the current web part. Represents the order of the web part in this column or section.
func (m *WebPartPosition) GetWebPartIndex()(*float64) {
    val, err := m.GetBackingStore().Get("webPartIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WebPartPosition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("columnId", m.GetColumnId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("horizontalSectionId", m.GetHorizontalSectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInVerticalSection", m.GetIsInVerticalSection())
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
        err := writer.WriteFloat64Value("webPartIndex", m.GetWebPartIndex())
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
func (m *WebPartPosition) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WebPartPosition) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetColumnId sets the columnId property value. Indicates the identifier of the column where the web part is located.
func (m *WebPartPosition) SetColumnId(value *float64)() {
    err := m.GetBackingStore().Set("columnId", value)
    if err != nil {
        panic(err)
    }
}
// SetHorizontalSectionId sets the horizontalSectionId property value. Indicates the horizontal section where the web part is located.
func (m *WebPartPosition) SetHorizontalSectionId(value *float64)() {
    err := m.GetBackingStore().Set("horizontalSectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetIsInVerticalSection sets the isInVerticalSection property value. Indicates whether the web part is located in the vertical section.
func (m *WebPartPosition) SetIsInVerticalSection(value *bool)() {
    err := m.GetBackingStore().Set("isInVerticalSection", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebPartPosition) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetWebPartIndex sets the webPartIndex property value. Index of the current web part. Represents the order of the web part in this column or section.
func (m *WebPartPosition) SetWebPartIndex(value *float64)() {
    err := m.GetBackingStore().Set("webPartIndex", value)
    if err != nil {
        panic(err)
    }
}
// WebPartPositionable 
type WebPartPositionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetColumnId()(*float64)
    GetHorizontalSectionId()(*float64)
    GetIsInVerticalSection()(*bool)
    GetOdataType()(*string)
    GetWebPartIndex()(*float64)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetColumnId(value *float64)()
    SetHorizontalSectionId(value *float64)()
    SetIsInVerticalSection(value *bool)()
    SetOdataType(value *string)()
    SetWebPartIndex(value *float64)()
}
