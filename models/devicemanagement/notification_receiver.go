package devicemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationReceiver 
type NotificationReceiver struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contact information about the notification receivers, such as an email address. Currently, only email and portal notifications are supported. For portal notifications, contactInformation can be left blank. For email notifications, contactInformation consists of an email address such as serena.davis@contoso.com.
    contactInformation *string
    // Defines the language and format in which the notification will be sent. Supported locale values are: en-us, cs-cz, de-de, es-es, fr-fr, hu-hu, it-it, ja-jp, ko-kr, nl-nl, pl-pl, pt-br, pt-pt, ru-ru, sv-se, tr-tr, zh-cn, zh-tw.
    locale *string
    // The OdataType property
    odataType *string
}
// NewNotificationReceiver instantiates a new notificationReceiver and sets the default values.
func NewNotificationReceiver()(*NotificationReceiver) {
    m := &NotificationReceiver{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateNotificationReceiverFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationReceiverFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationReceiver(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NotificationReceiver) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContactInformation gets the contactInformation property value. The contact information about the notification receivers, such as an email address. Currently, only email and portal notifications are supported. For portal notifications, contactInformation can be left blank. For email notifications, contactInformation consists of an email address such as serena.davis@contoso.com.
func (m *NotificationReceiver) GetContactInformation()(*string) {
    return m.contactInformation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationReceiver) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contactInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContactInformation)
    res["locale"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLocale)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetLocale gets the locale property value. Defines the language and format in which the notification will be sent. Supported locale values are: en-us, cs-cz, de-de, es-es, fr-fr, hu-hu, it-it, ja-jp, ko-kr, nl-nl, pl-pl, pt-br, pt-pt, ru-ru, sv-se, tr-tr, zh-cn, zh-tw.
func (m *NotificationReceiver) GetLocale()(*string) {
    return m.locale
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NotificationReceiver) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *NotificationReceiver) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contactInformation", m.GetContactInformation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
func (m *NotificationReceiver) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContactInformation sets the contactInformation property value. The contact information about the notification receivers, such as an email address. Currently, only email and portal notifications are supported. For portal notifications, contactInformation can be left blank. For email notifications, contactInformation consists of an email address such as serena.davis@contoso.com.
func (m *NotificationReceiver) SetContactInformation(value *string)() {
    m.contactInformation = value
}
// SetLocale sets the locale property value. Defines the language and format in which the notification will be sent. Supported locale values are: en-us, cs-cz, de-de, es-es, fr-fr, hu-hu, it-it, ja-jp, ko-kr, nl-nl, pl-pl, pt-br, pt-pt, ru-ru, sv-se, tr-tr, zh-cn, zh-tw.
func (m *NotificationReceiver) SetLocale(value *string)() {
    m.locale = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NotificationReceiver) SetOdataType(value *string)() {
    m.odataType = value
}
