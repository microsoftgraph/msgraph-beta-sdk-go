package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// NotificationChannel 
type NotificationChannel struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewNotificationChannel instantiates a new notificationChannel and sets the default values.
func NewNotificationChannel()(*NotificationChannel) {
    m := &NotificationChannel{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNotificationChannelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationChannelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationChannel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NotificationChannel) GetAdditionalData()(map[string]any) {
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
func (m *NotificationChannel) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationChannel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notificationChannelType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationChannelType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationChannelType(val.(*NotificationChannelType))
        }
        return nil
    }
    res["notificationReceivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationReceiverFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationReceiverable, len(val))
            for i, v := range val {
                res[i] = v.(NotificationReceiverable)
            }
            m.SetNotificationReceivers(res)
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
    res["receivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetReceivers(res)
        }
        return nil
    }
    return res
}
// GetNotificationChannelType gets the notificationChannelType property value. The type of the notification channel. The possible values are: portal, email, phoneCall, sms, unknownFutureValue.
func (m *NotificationChannel) GetNotificationChannelType()(*NotificationChannelType) {
    val, err := m.GetBackingStore().Get("notificationChannelType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NotificationChannelType)
    }
    return nil
}
// GetNotificationReceivers gets the notificationReceivers property value. Information about the notification receivers, such as locale and contact information. For example, en-us for locale and serena.davis@contoso.com for contact information.
func (m *NotificationChannel) GetNotificationReceivers()([]NotificationReceiverable) {
    val, err := m.GetBackingStore().Get("notificationReceivers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NotificationReceiverable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NotificationChannel) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReceivers gets the receivers property value. The contact information about the notification receivers, such as email addresses. For portal notifications, receivers can be left blank. For email notifications, receivers consists of email addresses such as serena.davis@contoso.com.
func (m *NotificationChannel) GetReceivers()([]string) {
    val, err := m.GetBackingStore().Get("receivers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *NotificationChannel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNotificationChannelType() != nil {
        cast := (*m.GetNotificationChannelType()).String()
        err := writer.WriteStringValue("notificationChannelType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationReceivers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationReceivers()))
        for i, v := range m.GetNotificationReceivers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("notificationReceivers", cast)
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
    if m.GetReceivers() != nil {
        err := writer.WriteCollectionOfStringValues("receivers", m.GetReceivers())
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
func (m *NotificationChannel) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *NotificationChannel) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetNotificationChannelType sets the notificationChannelType property value. The type of the notification channel. The possible values are: portal, email, phoneCall, sms, unknownFutureValue.
func (m *NotificationChannel) SetNotificationChannelType(value *NotificationChannelType)() {
    err := m.GetBackingStore().Set("notificationChannelType", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationReceivers sets the notificationReceivers property value. Information about the notification receivers, such as locale and contact information. For example, en-us for locale and serena.davis@contoso.com for contact information.
func (m *NotificationChannel) SetNotificationReceivers(value []NotificationReceiverable)() {
    err := m.GetBackingStore().Set("notificationReceivers", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NotificationChannel) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivers sets the receivers property value. The contact information about the notification receivers, such as email addresses. For portal notifications, receivers can be left blank. For email notifications, receivers consists of email addresses such as serena.davis@contoso.com.
func (m *NotificationChannel) SetReceivers(value []string)() {
    err := m.GetBackingStore().Set("receivers", value)
    if err != nil {
        panic(err)
    }
}
// NotificationChannelable 
type NotificationChannelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetNotificationChannelType()(*NotificationChannelType)
    GetNotificationReceivers()([]NotificationReceiverable)
    GetOdataType()(*string)
    GetReceivers()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetNotificationChannelType(value *NotificationChannelType)()
    SetNotificationReceivers(value []NotificationReceiverable)()
    SetOdataType(value *string)()
    SetReceivers(value []string)()
}
