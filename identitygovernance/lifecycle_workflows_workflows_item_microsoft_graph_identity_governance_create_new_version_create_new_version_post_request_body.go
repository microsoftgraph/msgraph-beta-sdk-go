package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody 
type LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody instantiates a new LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody()(*LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) {
    m := &LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
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
    val, err := m.GetBackingStore().Get("workflow")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetWorkflow sets the workflow property value. The workflow property
func (m *LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBody) SetWorkflow(value i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable)() {
    err := m.GetBackingStore().Set("workflow", value)
    if err != nil {
        panic(err)
    }
}
// LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBodyable 
type LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionCreateNewVersionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetWorkflow()(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetWorkflow(value i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable)()
}
