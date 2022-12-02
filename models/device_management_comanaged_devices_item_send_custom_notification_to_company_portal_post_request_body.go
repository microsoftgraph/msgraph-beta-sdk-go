package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody provides operations to call the sendCustomNotificationToCompanyPortal method.
type DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The notificationBody property
    notificationBody *string
    // The notificationTitle property
    notificationTitle *string
}
// NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody instantiates a new DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody and sets the default values.
func NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody()(*DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) {
    m := &DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationBody()(*string) {
    return m.notificationBody
}
// GetNotificationTitle gets the notificationTitle property value. The notificationTitle property
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationTitle()(*string) {
    return m.notificationTitle
}
// Serialize serializes information the current object
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationBody(value *string)() {
    m.notificationBody = value
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationTitle(value *string)() {
    m.notificationTitle = value
}
