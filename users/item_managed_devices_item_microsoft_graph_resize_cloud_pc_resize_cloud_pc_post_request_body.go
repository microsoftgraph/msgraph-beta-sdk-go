package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody 
type ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The targetServicePlanId property
    targetServicePlanId *string
}
// NewItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody instantiates a new ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody()(*ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody) {
    m := &ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["targetServicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetServicePlanId(val)
        }
        return nil
    }
    return res
}
// GetTargetServicePlanId gets the targetServicePlanId property value. The targetServicePlanId property
func (m *ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody) GetTargetServicePlanId()(*string) {
    return m.targetServicePlanId
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("targetServicePlanId", m.GetTargetServicePlanId())
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
func (m *ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTargetServicePlanId sets the targetServicePlanId property value. The targetServicePlanId property
func (m *ItemManagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcPostRequestBody) SetTargetServicePlanId(value *string)() {
    m.targetServicePlanId = value
}
