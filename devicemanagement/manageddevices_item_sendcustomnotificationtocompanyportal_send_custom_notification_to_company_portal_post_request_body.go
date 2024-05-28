package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody instantiates a new ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody and sets the default values.
func NewManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody()(*ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) {
    m := &ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notificationBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationBody(val)
        }
        return nil
    }
    res["notificationTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTitle(val)
        }
        return nil
    }
    return res
}
// GetNotificationBody gets the notificationBody property value. The notificationBody property
// returns a *string when successful
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationBody()(*string) {
    val, err := m.GetBackingStore().Get("notificationBody")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNotificationTitle gets the notificationTitle property value. The notificationTitle property
// returns a *string when successful
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationTitle()(*string) {
    val, err := m.GetBackingStore().Get("notificationTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("notificationBody", m.GetNotificationBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTitle", m.GetNotificationTitle())
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
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationBody(value *string)() {
    err := m.GetBackingStore().Set("notificationBody", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationTitle(value *string)() {
    err := m.GetBackingStore().Set("notificationTitle", value)
    if err != nil {
        panic(err)
    }
}
type ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetNotificationBody()(*string)
    GetNotificationTitle()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetNotificationBody(value *string)()
    SetNotificationTitle(value *string)()
}
