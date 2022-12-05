package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody provides operations to call the resume method.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The source property
    source *string
    // The type property
    type_escaped *string
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody instantiates a new IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) {
    m := &IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) GetSource()(*string) {
    return m.source
}
// GetType gets the type property value. The type property
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSource sets the source property value. The source property
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) SetSource(value *string)() {
    m.source = value
}
// SetType sets the type property value. The type property
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBody) SetType(value *string)() {
    m.type_escaped = value
}
