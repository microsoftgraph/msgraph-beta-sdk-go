package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MacOSAppleEventReceiver represents a process that can receive an Apple Event notification.
type MacOSAppleEventReceiver struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMacOSAppleEventReceiver instantiates a new macOSAppleEventReceiver and sets the default values.
func NewMacOSAppleEventReceiver()(*MacOSAppleEventReceiver) {
    m := &MacOSAppleEventReceiver{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMacOSAppleEventReceiverFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSAppleEventReceiverFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSAppleEventReceiver(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSAppleEventReceiver) GetAdditionalData()(map[string]any) {
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
// GetAllowed gets the allowed property value. Allow or block this app from receiving Apple events.
func (m *MacOSAppleEventReceiver) GetAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("allowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *MacOSAppleEventReceiver) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCodeRequirement gets the codeRequirement property value. Code requirement for the app or binary that receives the Apple Event.
func (m *MacOSAppleEventReceiver) GetCodeRequirement()(*string) {
    val, err := m.GetBackingStore().Get("codeRequirement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSAppleEventReceiver) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowed(val)
        }
        return nil
    }
    res["codeRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeRequirement(val)
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["identifierType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSProcessIdentifierType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifierType(val.(*MacOSProcessIdentifierType))
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
// GetIdentifier gets the identifier property value. Bundle ID of the app or file path of the process or executable that receives the Apple Event.
func (m *MacOSAppleEventReceiver) GetIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentifierType gets the identifierType property value. Process identifier types for MacOS Privacy Preferences
func (m *MacOSAppleEventReceiver) GetIdentifierType()(*MacOSProcessIdentifierType) {
    val, err := m.GetBackingStore().Get("identifierType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSProcessIdentifierType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSAppleEventReceiver) GetOdataType()(*string) {
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
func (m *MacOSAppleEventReceiver) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowed", m.GetAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeRequirement", m.GetCodeRequirement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetIdentifierType() != nil {
        cast := (*m.GetIdentifierType()).String()
        err := writer.WriteStringValue("identifierType", &cast)
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
func (m *MacOSAppleEventReceiver) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowed sets the allowed property value. Allow or block this app from receiving Apple events.
func (m *MacOSAppleEventReceiver) SetAllowed(value *bool)() {
    err := m.GetBackingStore().Set("allowed", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *MacOSAppleEventReceiver) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCodeRequirement sets the codeRequirement property value. Code requirement for the app or binary that receives the Apple Event.
func (m *MacOSAppleEventReceiver) SetCodeRequirement(value *string)() {
    err := m.GetBackingStore().Set("codeRequirement", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentifier sets the identifier property value. Bundle ID of the app or file path of the process or executable that receives the Apple Event.
func (m *MacOSAppleEventReceiver) SetIdentifier(value *string)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentifierType sets the identifierType property value. Process identifier types for MacOS Privacy Preferences
func (m *MacOSAppleEventReceiver) SetIdentifierType(value *MacOSProcessIdentifierType)() {
    err := m.GetBackingStore().Set("identifierType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSAppleEventReceiver) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// MacOSAppleEventReceiverable 
type MacOSAppleEventReceiverable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowed()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCodeRequirement()(*string)
    GetIdentifier()(*string)
    GetIdentifierType()(*MacOSProcessIdentifierType)
    GetOdataType()(*string)
    SetAllowed(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCodeRequirement(value *string)()
    SetIdentifier(value *string)()
    SetIdentifierType(value *MacOSProcessIdentifierType)()
    SetOdataType(value *string)()
}
