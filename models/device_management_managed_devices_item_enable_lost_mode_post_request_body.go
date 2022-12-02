package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementManagedDevicesItemEnableLostModePostRequestBody provides operations to call the enableLostMode method.
type DeviceManagementManagedDevicesItemEnableLostModePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The footer property
    footer *string
    // The message property
    message *string
    // The phoneNumber property
    phoneNumber *string
}
// NewDeviceManagementManagedDevicesItemEnableLostModePostRequestBody instantiates a new DeviceManagementManagedDevicesItemEnableLostModePostRequestBody and sets the default values.
func NewDeviceManagementManagedDevicesItemEnableLostModePostRequestBody()(*DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) {
    m := &DeviceManagementManagedDevicesItemEnableLostModePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementManagedDevicesItemEnableLostModePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementManagedDevicesItemEnableLostModePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementManagedDevicesItemEnableLostModePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["footer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFooter(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    return res
}
// GetFooter gets the footer property value. The footer property
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) GetFooter()(*string) {
    return m.footer
}
// GetMessage gets the message property value. The message property
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// Serialize serializes information the current object
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("footer", m.GetFooter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
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
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFooter sets the footer property value. The footer property
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) SetFooter(value *string)() {
    m.footer = value
}
// SetMessage sets the message property value. The message property
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *DeviceManagementManagedDevicesItemEnableLostModePostRequestBody) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
