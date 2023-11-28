package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody 
type OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody instantiates a new OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody and sets the default values.
func NewOnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody()(*OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) {
    m := &OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["phoneNumbers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPhoneNumbers(res)
        }
        return nil
    }
    res["smsType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseVirtualAppointmentSmsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentSmsType))
        }
        return nil
    }
    return res
}
// GetPhoneNumbers gets the phoneNumbers property value. The phoneNumbers property
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) GetPhoneNumbers()([]string) {
    val, err := m.GetBackingStore().Get("phoneNumbers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSmsType gets the smsType property value. The smsType property
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) GetSmsType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentSmsType) {
    val, err := m.GetBackingStore().Get("smsType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentSmsType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPhoneNumbers() != nil {
        err := writer.WriteCollectionOfStringValues("phoneNumbers", m.GetPhoneNumbers())
        if err != nil {
            return err
        }
    }
    if m.GetSmsType() != nil {
        cast := (*m.GetSmsType()).String()
        err := writer.WriteStringValue("smsType", &cast)
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
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetPhoneNumbers sets the phoneNumbers property value. The phoneNumbers property
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) SetPhoneNumbers(value []string)() {
    err := m.GetBackingStore().Set("phoneNumbers", value)
    if err != nil {
        panic(err)
    }
}
// SetSmsType sets the smsType property value. The smsType property
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBody) SetSmsType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentSmsType)() {
    err := m.GetBackingStore().Set("smsType", value)
    if err != nil {
        panic(err)
    }
}
// OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBodyable 
type OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetPhoneNumbers()([]string)
    GetSmsType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentSmsType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetPhoneNumbers(value []string)()
    SetSmsType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentSmsType)()
}
