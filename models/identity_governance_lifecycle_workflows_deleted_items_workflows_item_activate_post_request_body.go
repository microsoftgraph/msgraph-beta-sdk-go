package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody provides operations to call the activate method.
type IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The subjects property
    subjects []Userable
}
// NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody instantiates a new IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody()(*IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody) {
    m := &IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["subjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                res[i] = v.(Userable)
            }
            m.SetSubjects(res)
        }
        return nil
    }
    return res
}
// GetSubjects gets the subjects property value. The subjects property
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody) GetSubjects()([]Userable) {
    return m.subjects
}
// Serialize serializes information the current object
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSubjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjects()))
        for i, v := range m.GetSubjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("subjects", cast)
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
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSubjects sets the subjects property value. The subjects property
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemActivatePostRequestBody) SetSubjects(value []Userable)() {
    m.subjects = value
}
