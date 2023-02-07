package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody 
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The priority property
    priority *int32
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody instantiates a new ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPriority sets the priority property value. The priority property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphReorderReorderPostRequestBody) SetPriority(value *int32)() {
    m.priority = value
}
