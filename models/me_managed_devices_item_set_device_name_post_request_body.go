package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesItemSetDeviceNamePostRequestBody provides operations to call the setDeviceName method.
type MeManagedDevicesItemSetDeviceNamePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceName property
    deviceName *string
}
// NewMeManagedDevicesItemSetDeviceNamePostRequestBody instantiates a new MeManagedDevicesItemSetDeviceNamePostRequestBody and sets the default values.
func NewMeManagedDevicesItemSetDeviceNamePostRequestBody()(*MeManagedDevicesItemSetDeviceNamePostRequestBody) {
    m := &MeManagedDevicesItemSetDeviceNamePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesItemSetDeviceNamePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesItemSetDeviceNamePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesItemSetDeviceNamePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesItemSetDeviceNamePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *MeManagedDevicesItemSetDeviceNamePostRequestBody) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesItemSetDeviceNamePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MeManagedDevicesItemSetDeviceNamePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeManagedDevicesItemSetDeviceNamePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *MeManagedDevicesItemSetDeviceNamePostRequestBody) SetDeviceName(value *string)() {
    m.deviceName = value
}
