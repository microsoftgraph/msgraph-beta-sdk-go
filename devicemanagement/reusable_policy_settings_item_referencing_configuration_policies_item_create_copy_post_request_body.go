package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody 
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The description property
    description *string
    // The displayName property
    displayName *string
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody instantiates a new ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
