package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody 
type LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data property
    data i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CustomTaskExtensionCallbackDataable
    // The source property
    source *string
    // The type property
    type_escaped *string
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody instantiates a new LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody()(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The data property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetData()(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CustomTaskExtensionCallbackDataable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateCustomTaskExtensionCallbackDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CustomTaskExtensionCallbackDataable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The source property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetSource()(*string) {
    return m.source
}
// GetType gets the type property value. The type property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetData(value i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CustomTaskExtensionCallbackDataable)() {
    m.data = value
}
// SetSource sets the source property value. The source property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetSource(value *string)() {
    m.source = value
}
// SetType sets the type property value. The type property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetType(value *string)() {
    m.type_escaped = value
}
