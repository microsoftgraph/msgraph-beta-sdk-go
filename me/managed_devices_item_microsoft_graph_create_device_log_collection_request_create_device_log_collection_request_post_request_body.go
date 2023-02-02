package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody 
type ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The templateType property
    templateType ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLogCollectionRequestable
}
// NewManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody instantiates a new ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody and sets the default values.
func NewManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody()(*ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody) {
    m := &ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["templateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceLogCollectionRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLogCollectionRequestable))
        }
        return nil
    }
    return res
}
// GetTemplateType gets the templateType property value. The templateType property
func (m *ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody) GetTemplateType()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLogCollectionRequestable) {
    return m.templateType
}
// Serialize serializes information the current object
func (m *ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("templateType", m.GetTemplateType())
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
func (m *ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTemplateType sets the templateType property value. The templateType property
func (m *ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestPostRequestBody) SetTemplateType(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLogCollectionRequestable)() {
    m.templateType = value
}
