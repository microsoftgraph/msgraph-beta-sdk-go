package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody 
type ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceName property
    deviceName *string
}
// NewComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody instantiates a new ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody()(*ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody) {
    m := &ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
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
func (m *ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNamePostRequestBody) SetDeviceName(value *string)() {
    m.deviceName = value
}
