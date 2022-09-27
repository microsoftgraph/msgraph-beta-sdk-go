package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalWorkflowProvider provides operations to manage the collection of activityStatistics entities.
type ApprovalWorkflowProvider struct {
    Entity
    // The businessFlows property
    businessFlows []BusinessFlowable
    // The businessFlowsWithRequestsAwaitingMyDecision property
    businessFlowsWithRequestsAwaitingMyDecision []BusinessFlowable
    // The displayName property
    displayName *string
    // The policyTemplates property
    policyTemplates []GovernancePolicyTemplateable
}
// NewApprovalWorkflowProvider instantiates a new approvalWorkflowProvider and sets the default values.
func NewApprovalWorkflowProvider()(*ApprovalWorkflowProvider) {
    m := &ApprovalWorkflowProvider{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.approvalWorkflowProvider";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateApprovalWorkflowProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalWorkflowProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalWorkflowProvider(), nil
}
// GetBusinessFlows gets the businessFlows property value. The businessFlows property
func (m *ApprovalWorkflowProvider) GetBusinessFlows()([]BusinessFlowable) {
    return m.businessFlows
}
// GetBusinessFlowsWithRequestsAwaitingMyDecision gets the businessFlowsWithRequestsAwaitingMyDecision property value. The businessFlowsWithRequestsAwaitingMyDecision property
func (m *ApprovalWorkflowProvider) GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlowable) {
    return m.businessFlowsWithRequestsAwaitingMyDecision
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ApprovalWorkflowProvider) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalWorkflowProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBusinessFlowFromDiscriminatorValue , m.SetBusinessFlows)
    res["businessFlowsWithRequestsAwaitingMyDecision"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBusinessFlowFromDiscriminatorValue , m.SetBusinessFlowsWithRequestsAwaitingMyDecision)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["policyTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernancePolicyTemplateFromDiscriminatorValue , m.SetPolicyTemplates)
    return res
}
// GetPolicyTemplates gets the policyTemplates property value. The policyTemplates property
func (m *ApprovalWorkflowProvider) GetPolicyTemplates()([]GovernancePolicyTemplateable) {
    return m.policyTemplates
}
// Serialize serializes information the current object
func (m *ApprovalWorkflowProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBusinessFlows() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetBusinessFlows())
        err = writer.WriteCollectionOfObjectValues("businessFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessFlowsWithRequestsAwaitingMyDecision() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetBusinessFlowsWithRequestsAwaitingMyDecision())
        err = writer.WriteCollectionOfObjectValues("businessFlowsWithRequestsAwaitingMyDecision", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPolicyTemplates())
        err = writer.WriteCollectionOfObjectValues("policyTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBusinessFlows sets the businessFlows property value. The businessFlows property
func (m *ApprovalWorkflowProvider) SetBusinessFlows(value []BusinessFlowable)() {
    m.businessFlows = value
}
// SetBusinessFlowsWithRequestsAwaitingMyDecision sets the businessFlowsWithRequestsAwaitingMyDecision property value. The businessFlowsWithRequestsAwaitingMyDecision property
func (m *ApprovalWorkflowProvider) SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlowable)() {
    m.businessFlowsWithRequestsAwaitingMyDecision = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ApprovalWorkflowProvider) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPolicyTemplates sets the policyTemplates property value. The policyTemplates property
func (m *ApprovalWorkflowProvider) SetPolicyTemplates(value []GovernancePolicyTemplateable)() {
    m.policyTemplates = value
}
