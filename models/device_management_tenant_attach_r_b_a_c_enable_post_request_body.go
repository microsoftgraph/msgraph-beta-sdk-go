package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTenantAttachRBACEnablePostRequestBody provides operations to call the enable method.
type DeviceManagementTenantAttachRBACEnablePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enable property
    enable *bool
}
// NewDeviceManagementTenantAttachRBACEnablePostRequestBody instantiates a new DeviceManagementTenantAttachRBACEnablePostRequestBody and sets the default values.
func NewDeviceManagementTenantAttachRBACEnablePostRequestBody()(*DeviceManagementTenantAttachRBACEnablePostRequestBody) {
    m := &DeviceManagementTenantAttachRBACEnablePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementTenantAttachRBACEnablePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTenantAttachRBACEnablePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementTenantAttachRBACEnablePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTenantAttachRBACEnablePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnable gets the enable property value. The enable property
func (m *DeviceManagementTenantAttachRBACEnablePostRequestBody) GetEnable()(*bool) {
    return m.enable
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTenantAttachRBACEnablePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnable(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementTenantAttachRBACEnablePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enable", m.GetEnable())
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
func (m *DeviceManagementTenantAttachRBACEnablePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnable sets the enable property value. The enable property
func (m *DeviceManagementTenantAttachRBACEnablePostRequestBody) SetEnable(value *bool)() {
    m.enable = value
}
