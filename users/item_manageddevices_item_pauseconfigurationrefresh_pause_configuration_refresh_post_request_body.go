package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody instantiates a new ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody and sets the default values.
func NewItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody()(*ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) {
    m := &ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pauseTimePeriodInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPauseTimePeriodInMinutes(val)
        }
        return nil
    }
    return res
}
// GetPauseTimePeriodInMinutes gets the pauseTimePeriodInMinutes property value. The pauseTimePeriodInMinutes property
// returns a *int32 when successful
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) GetPauseTimePeriodInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("pauseTimePeriodInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("pauseTimePeriodInMinutes", m.GetPauseTimePeriodInMinutes())
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
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetPauseTimePeriodInMinutes sets the pauseTimePeriodInMinutes property value. The pauseTimePeriodInMinutes property
func (m *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBody) SetPauseTimePeriodInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("pauseTimePeriodInMinutes", value)
    if err != nil {
        panic(err)
    }
}
type ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetPauseTimePeriodInMinutes()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetPauseTimePeriodInMinutes(value *int32)()
}
