package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalWorkflowProvider 
type ApprovalWorkflowProvider struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewApprovalWorkflowProvider instantiates a new approvalWorkflowProvider and sets the default values.
func NewApprovalWorkflowProvider()(*ApprovalWorkflowProvider) {
    m := &ApprovalWorkflowProvider{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalWorkflowProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalWorkflowProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalWorkflowProvider(), nil
}
// GetBusinessFlows gets the businessFlows property value. The businessFlows property
func (m *ApprovalWorkflowProvider) GetBusinessFlows()([]BusinessFlowable) {
    val, err := m.GetBackingStore().Get("businessFlows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]BusinessFlowable)
    }
    return nil
}
// GetBusinessFlowsWithRequestsAwaitingMyDecision gets the businessFlowsWithRequestsAwaitingMyDecision property value. The businessFlowsWithRequestsAwaitingMyDecision property
func (m *ApprovalWorkflowProvider) GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlowable) {
    val, err := m.GetBackingStore().Get("businessFlowsWithRequestsAwaitingMyDecision")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]BusinessFlowable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ApprovalWorkflowProvider) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalWorkflowProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBusinessFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessFlowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BusinessFlowable)
                }
            }
            m.SetBusinessFlows(res)
        }
        return nil
    }
    res["businessFlowsWithRequestsAwaitingMyDecision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBusinessFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessFlowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BusinessFlowable)
                }
            }
            m.SetBusinessFlowsWithRequestsAwaitingMyDecision(res)
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
    res["policyTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernancePolicyTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernancePolicyTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GovernancePolicyTemplateable)
                }
            }
            m.SetPolicyTemplates(res)
        }
        return nil
    }
    return res
}
// GetPolicyTemplates gets the policyTemplates property value. The policyTemplates property
func (m *ApprovalWorkflowProvider) GetPolicyTemplates()([]GovernancePolicyTemplateable) {
    val, err := m.GetBackingStore().Get("policyTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernancePolicyTemplateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApprovalWorkflowProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBusinessFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBusinessFlows()))
        for i, v := range m.GetBusinessFlows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("businessFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessFlowsWithRequestsAwaitingMyDecision() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBusinessFlowsWithRequestsAwaitingMyDecision()))
        for i, v := range m.GetBusinessFlowsWithRequestsAwaitingMyDecision() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicyTemplates()))
        for i, v := range m.GetPolicyTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("policyTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBusinessFlows sets the businessFlows property value. The businessFlows property
func (m *ApprovalWorkflowProvider) SetBusinessFlows(value []BusinessFlowable)() {
    err := m.GetBackingStore().Set("businessFlows", value)
    if err != nil {
        panic(err)
    }
}
// SetBusinessFlowsWithRequestsAwaitingMyDecision sets the businessFlowsWithRequestsAwaitingMyDecision property value. The businessFlowsWithRequestsAwaitingMyDecision property
func (m *ApprovalWorkflowProvider) SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlowable)() {
    err := m.GetBackingStore().Set("businessFlowsWithRequestsAwaitingMyDecision", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ApprovalWorkflowProvider) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyTemplates sets the policyTemplates property value. The policyTemplates property
func (m *ApprovalWorkflowProvider) SetPolicyTemplates(value []GovernancePolicyTemplateable)() {
    err := m.GetBackingStore().Set("policyTemplates", value)
    if err != nil {
        panic(err)
    }
}
// ApprovalWorkflowProviderable 
type ApprovalWorkflowProviderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessFlows()([]BusinessFlowable)
    GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlowable)
    GetDisplayName()(*string)
    GetPolicyTemplates()([]GovernancePolicyTemplateable)
    SetBusinessFlows(value []BusinessFlowable)()
    SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlowable)()
    SetDisplayName(value *string)()
    SetPolicyTemplates(value []GovernancePolicyTemplateable)()
}
