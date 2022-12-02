package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody provides operations to call the activateDeviceEsim method.
type DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The carrierUrl property
    carrierUrl *string
}
// NewDeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody instantiates a new DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody and sets the default values.
func NewDeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody()(*DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody) {
    m := &DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCarrierUrl gets the carrierUrl property value. The carrierUrl property
func (m *DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody) GetCarrierUrl()(*string) {
    return m.carrierUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["carrierUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("carrierUrl", m.GetCarrierUrl())
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
func (m *DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCarrierUrl sets the carrierUrl property value. The carrierUrl property
func (m *DeviceManagementComanagedDevicesItemActivateDeviceEsimPostRequestBody) SetCarrierUrl(value *string)() {
    m.carrierUrl = value
}
