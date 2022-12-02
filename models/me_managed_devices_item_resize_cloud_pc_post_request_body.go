package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesItemResizeCloudPcPostRequestBody provides operations to call the resizeCloudPc method.
type MeManagedDevicesItemResizeCloudPcPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The targetServicePlanId property
    targetServicePlanId *string
}
// NewMeManagedDevicesItemResizeCloudPcPostRequestBody instantiates a new MeManagedDevicesItemResizeCloudPcPostRequestBody and sets the default values.
func NewMeManagedDevicesItemResizeCloudPcPostRequestBody()(*MeManagedDevicesItemResizeCloudPcPostRequestBody) {
    m := &MeManagedDevicesItemResizeCloudPcPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesItemResizeCloudPcPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesItemResizeCloudPcPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesItemResizeCloudPcPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesItemResizeCloudPcPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesItemResizeCloudPcPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MeManagedDevicesItemResizeCloudPcPostRequestBody) GetTargetServicePlanId()(*string) {
    return m.targetServicePlanId
}
// Serialize serializes information the current object
func (m *MeManagedDevicesItemResizeCloudPcPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeManagedDevicesItemResizeCloudPcPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetTargetServicePlanId sets the targetServicePlanId property value. The targetServicePlanId property
func (m *MeManagedDevicesItemResizeCloudPcPostRequestBody) SetTargetServicePlanId(value *string)() {
    m.targetServicePlanId = value
}
