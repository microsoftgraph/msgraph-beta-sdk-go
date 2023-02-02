package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody 
type LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The workflow property
    workflow i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable
}
// NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody instantiates a new LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody()(*LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) {
    m := &LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["workflow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateWorkflowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflow(val.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable))
        }
        return nil
    }
    return res
}
// GetWorkflow gets the workflow property value. The workflow property
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) GetWorkflow()(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable) {
    return m.workflow
}
// Serialize serializes information the current object
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("workflow", m.GetWorkflow())
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
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWorkflow sets the workflow property value. The workflow property
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) SetWorkflow(value i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable)() {
    m.workflow = value
}
