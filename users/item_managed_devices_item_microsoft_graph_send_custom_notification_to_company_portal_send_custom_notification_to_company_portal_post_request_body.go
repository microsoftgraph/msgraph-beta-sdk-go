package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody 
type ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The notificationBody property
    notificationBody *string
    // The notificationTitle property
    notificationTitle *string
}
// NewItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody instantiates a new ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody()(*ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) {
    m := &ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationBody()(*string) {
    return m.notificationBody
}
// GetNotificationTitle gets the notificationTitle property value. The notificationTitle property
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationTitle()(*string) {
    return m.notificationTitle
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationBody(value *string)() {
    m.notificationBody = value
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationTitle(value *string)() {
    m.notificationTitle = value
}
