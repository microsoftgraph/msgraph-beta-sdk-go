package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AndroidDeviceOwnerKioskModeAppPositionItem an item in the list of app positions that sets the order of items on the Managed Home Screen
type AndroidDeviceOwnerKioskModeAppPositionItem struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAndroidDeviceOwnerKioskModeAppPositionItem instantiates a new androidDeviceOwnerKioskModeAppPositionItem and sets the default values.
func NewAndroidDeviceOwnerKioskModeAppPositionItem()(*AndroidDeviceOwnerKioskModeAppPositionItem) {
    m := &AndroidDeviceOwnerKioskModeAppPositionItem{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidDeviceOwnerKioskModeAppPositionItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeAppPositionItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeAppPositionItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetAdditionalData()(map[string]any) {
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
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["item"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerKioskModeHomeScreenItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItem(val.(AndroidDeviceOwnerKioskModeHomeScreenItemable))
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
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    return res
}
// GetItem gets the item property value. Represents an item on the Android Device Owner Managed Home Screen (application, weblink or folder
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetItem()(AndroidDeviceOwnerKioskModeHomeScreenItemable) {
    val, err := m.GetBackingStore().Get("item")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidDeviceOwnerKioskModeHomeScreenItemable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPosition gets the position property value. Position of the item on the grid. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetPosition()(*int32) {
    val, err := m.GetBackingStore().Get("position")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("item", m.GetItem())
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
        err := writer.WriteInt32Value("position", m.GetPosition())
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
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetItem sets the item property value. Represents an item on the Android Device Owner Managed Home Screen (application, weblink or folder
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetItem(value AndroidDeviceOwnerKioskModeHomeScreenItemable)() {
    err := m.GetBackingStore().Set("item", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPosition sets the position property value. Position of the item on the grid. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetPosition(value *int32)() {
    err := m.GetBackingStore().Set("position", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerKioskModeAppPositionItemable 
type AndroidDeviceOwnerKioskModeAppPositionItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetItem()(AndroidDeviceOwnerKioskModeHomeScreenItemable)
    GetOdataType()(*string)
    GetPosition()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetItem(value AndroidDeviceOwnerKioskModeHomeScreenItemable)()
    SetOdataType(value *string)()
    SetPosition(value *int32)()
}
