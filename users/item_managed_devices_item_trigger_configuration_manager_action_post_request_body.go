package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody 
type ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody instantiates a new ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody and sets the default values.
func NewItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody()(*ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) {
    m := &ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConfigurationManagerAction gets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) GetConfigurationManagerAction()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable) {
    val, err := m.GetBackingStore().Get("configurationManagerAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configurationManagerAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConfigurationManagerActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerAction(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configurationManagerAction", m.GetConfigurationManagerAction())
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
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConfigurationManagerAction sets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) SetConfigurationManagerAction(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable)() {
    err := m.GetBackingStore().Set("configurationManagerAction", value)
    if err != nil {
        panic(err)
    }
}
// ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyable 
type ItemManagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConfigurationManagerAction()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConfigurationManagerAction(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable)()
}
