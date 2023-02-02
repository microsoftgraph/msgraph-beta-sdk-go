package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody 
type TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enable property
    enable *bool
}
// NewTenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody instantiates a new TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody and sets the default values.
func NewTenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody()(*TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody) {
    m := &TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTenantAttachRBACMicrosoftGraphEnableEnablePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAttachRBACMicrosoftGraphEnableEnablePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnable gets the enable property value. The enable property
func (m *TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody) GetEnable()(*bool) {
    return m.enable
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnable sets the enable property value. The enable property
func (m *TenantAttachRBACMicrosoftGraphEnableEnablePostRequestBody) SetEnable(value *bool)() {
    m.enable = value
}
