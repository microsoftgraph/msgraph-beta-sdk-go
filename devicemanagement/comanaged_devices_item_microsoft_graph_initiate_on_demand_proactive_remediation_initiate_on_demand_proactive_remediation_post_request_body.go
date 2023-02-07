package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody 
type ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The scriptPolicyId property
    scriptPolicyId *string
}
// NewComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody instantiates a new ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody()(*ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) {
    m := &ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["scriptPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptPolicyId(val)
        }
        return nil
    }
    return res
}
// GetScriptPolicyId gets the scriptPolicyId property value. The scriptPolicyId property
func (m *ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) GetScriptPolicyId()(*string) {
    return m.scriptPolicyId
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("scriptPolicyId", m.GetScriptPolicyId())
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
func (m *ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetScriptPolicyId sets the scriptPolicyId property value. The scriptPolicyId property
func (m *ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) SetScriptPolicyId(value *string)() {
    m.scriptPolicyId = value
}
