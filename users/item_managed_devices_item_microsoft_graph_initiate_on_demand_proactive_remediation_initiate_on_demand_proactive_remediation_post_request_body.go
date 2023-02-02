package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody 
type ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The scriptPolicyId property
    scriptPolicyId *string
}
// NewItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody instantiates a new ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody()(*ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) {
    m := &ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) GetScriptPolicyId()(*string) {
    return m.scriptPolicyId
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetScriptPolicyId sets the scriptPolicyId property value. The scriptPolicyId property
func (m *ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationPostRequestBody) SetScriptPolicyId(value *string)() {
    m.scriptPolicyId = value
}
