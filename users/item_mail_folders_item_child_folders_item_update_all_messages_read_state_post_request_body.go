package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody instantiates a new ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody and sets the default values.
func NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody()(*ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) {
    m := &ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRead(val)
        }
        return nil
    }
    res["suppressReadReceipts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuppressReadReceipts(val)
        }
        return nil
    }
    return res
}
// GetIsRead gets the isRead property value. The isRead property
// returns a *bool when successful
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) GetIsRead()(*bool) {
    val, err := m.GetBackingStore().Get("isRead")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSuppressReadReceipts gets the suppressReadReceipts property value. The suppressReadReceipts property
// returns a *bool when successful
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) GetSuppressReadReceipts()(*bool) {
    val, err := m.GetBackingStore().Get("suppressReadReceipts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRead", m.GetIsRead())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("suppressReadReceipts", m.GetSuppressReadReceipts())
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
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsRead sets the isRead property value. The isRead property
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) SetIsRead(value *bool)() {
    err := m.GetBackingStore().Set("isRead", value)
    if err != nil {
        panic(err)
    }
}
// SetSuppressReadReceipts sets the suppressReadReceipts property value. The suppressReadReceipts property
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBody) SetSuppressReadReceipts(value *bool)() {
    err := m.GetBackingStore().Set("suppressReadReceipts", value)
    if err != nil {
        panic(err)
    }
}
type ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsRead()(*bool)
    GetSuppressReadReceipts()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsRead(value *bool)()
    SetSuppressReadReceipts(value *bool)()
}
