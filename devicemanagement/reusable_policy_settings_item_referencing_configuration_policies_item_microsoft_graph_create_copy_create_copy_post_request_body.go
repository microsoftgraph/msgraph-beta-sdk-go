package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody 
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // The displayName property
    displayName *string
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody instantiates a new ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemMicrosoftGraphCreateCopyCreateCopyPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
