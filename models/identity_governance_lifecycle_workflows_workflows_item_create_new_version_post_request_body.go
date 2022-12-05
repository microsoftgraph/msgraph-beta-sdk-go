package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody provides operations to call the createNewVersion method.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody instantiates a new IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody) {
    m := &IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
