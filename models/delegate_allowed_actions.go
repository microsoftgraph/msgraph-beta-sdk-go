package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DelegateAllowedActions struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDelegateAllowedActions instantiates a new DelegateAllowedActions and sets the default values.
func NewDelegateAllowedActions()(*DelegateAllowedActions) {
    m := &DelegateAllowedActions{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDelegateAllowedActionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDelegateAllowedActionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegateAllowedActions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DelegateAllowedActions) GetAdditionalData()(map[string]any) {
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
// returns a BackingStore when successful
func (m *DelegateAllowedActions) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DelegateAllowedActions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["joinActiveCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinActiveCalls(val)
        }
        return nil
    }
    res["makeCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMakeCalls(val)
        }
        return nil
    }
    res["manageCallAndDelegateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManageCallAndDelegateSettings(val)
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
    res["pickUpHeldCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPickUpHeldCalls(val)
        }
        return nil
    }
    res["receiveCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceiveCalls(val)
        }
        return nil
    }
    return res
}
// GetJoinActiveCalls gets the joinActiveCalls property value. The joinActiveCalls property
// returns a *bool when successful
func (m *DelegateAllowedActions) GetJoinActiveCalls()(*bool) {
    val, err := m.GetBackingStore().Get("joinActiveCalls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMakeCalls gets the makeCalls property value. The makeCalls property
// returns a *bool when successful
func (m *DelegateAllowedActions) GetMakeCalls()(*bool) {
    val, err := m.GetBackingStore().Get("makeCalls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManageCallAndDelegateSettings gets the manageCallAndDelegateSettings property value. The manageCallAndDelegateSettings property
// returns a *bool when successful
func (m *DelegateAllowedActions) GetManageCallAndDelegateSettings()(*bool) {
    val, err := m.GetBackingStore().Get("manageCallAndDelegateSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DelegateAllowedActions) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPickUpHeldCalls gets the pickUpHeldCalls property value. The pickUpHeldCalls property
// returns a *bool when successful
func (m *DelegateAllowedActions) GetPickUpHeldCalls()(*bool) {
    val, err := m.GetBackingStore().Get("pickUpHeldCalls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetReceiveCalls gets the receiveCalls property value. The receiveCalls property
// returns a *bool when successful
func (m *DelegateAllowedActions) GetReceiveCalls()(*bool) {
    val, err := m.GetBackingStore().Get("receiveCalls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DelegateAllowedActions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("joinActiveCalls", m.GetJoinActiveCalls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("makeCalls", m.GetMakeCalls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("manageCallAndDelegateSettings", m.GetManageCallAndDelegateSettings())
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
        err := writer.WriteBoolValue("pickUpHeldCalls", m.GetPickUpHeldCalls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("receiveCalls", m.GetReceiveCalls())
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
func (m *DelegateAllowedActions) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DelegateAllowedActions) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetJoinActiveCalls sets the joinActiveCalls property value. The joinActiveCalls property
func (m *DelegateAllowedActions) SetJoinActiveCalls(value *bool)() {
    err := m.GetBackingStore().Set("joinActiveCalls", value)
    if err != nil {
        panic(err)
    }
}
// SetMakeCalls sets the makeCalls property value. The makeCalls property
func (m *DelegateAllowedActions) SetMakeCalls(value *bool)() {
    err := m.GetBackingStore().Set("makeCalls", value)
    if err != nil {
        panic(err)
    }
}
// SetManageCallAndDelegateSettings sets the manageCallAndDelegateSettings property value. The manageCallAndDelegateSettings property
func (m *DelegateAllowedActions) SetManageCallAndDelegateSettings(value *bool)() {
    err := m.GetBackingStore().Set("manageCallAndDelegateSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DelegateAllowedActions) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPickUpHeldCalls sets the pickUpHeldCalls property value. The pickUpHeldCalls property
func (m *DelegateAllowedActions) SetPickUpHeldCalls(value *bool)() {
    err := m.GetBackingStore().Set("pickUpHeldCalls", value)
    if err != nil {
        panic(err)
    }
}
// SetReceiveCalls sets the receiveCalls property value. The receiveCalls property
func (m *DelegateAllowedActions) SetReceiveCalls(value *bool)() {
    err := m.GetBackingStore().Set("receiveCalls", value)
    if err != nil {
        panic(err)
    }
}
type DelegateAllowedActionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetJoinActiveCalls()(*bool)
    GetMakeCalls()(*bool)
    GetManageCallAndDelegateSettings()(*bool)
    GetOdataType()(*string)
    GetPickUpHeldCalls()(*bool)
    GetReceiveCalls()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetJoinActiveCalls(value *bool)()
    SetMakeCalls(value *bool)()
    SetManageCallAndDelegateSettings(value *bool)()
    SetOdataType(value *string)()
    SetPickUpHeldCalls(value *bool)()
    SetReceiveCalls(value *bool)()
}
