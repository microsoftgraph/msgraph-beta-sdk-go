package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemActivateDeviceEsimPostRequestBody 
type ItemManagedDevicesItemActivateDeviceEsimPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The carrierUrl property
    carrierUrl *string
}
// NewItemManagedDevicesItemActivateDeviceEsimPostRequestBody instantiates a new ItemManagedDevicesItemActivateDeviceEsimPostRequestBody and sets the default values.
func NewItemManagedDevicesItemActivateDeviceEsimPostRequestBody()(*ItemManagedDevicesItemActivateDeviceEsimPostRequestBody) {
    m := &ItemManagedDevicesItemActivateDeviceEsimPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemManagedDevicesItemActivateDeviceEsimPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemActivateDeviceEsimPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemActivateDeviceEsimPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemActivateDeviceEsimPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCarrierUrl gets the carrierUrl property value. The carrierUrl property
func (m *ItemManagedDevicesItemActivateDeviceEsimPostRequestBody) GetCarrierUrl()(*string) {
    return m.carrierUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemActivateDeviceEsimPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemManagedDevicesItemActivateDeviceEsimPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemManagedDevicesItemActivateDeviceEsimPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCarrierUrl sets the carrierUrl property value. The carrierUrl property
func (m *ItemManagedDevicesItemActivateDeviceEsimPostRequestBody) SetCarrierUrl(value *string)() {
    m.carrierUrl = value
}
