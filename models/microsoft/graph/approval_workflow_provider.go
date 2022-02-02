package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalWorkflowProvider 
type ApprovalWorkflowProvider struct {
    Entity
    // 
    businessFlows []BusinessFlow;
    // 
    businessFlowsWithRequestsAwaitingMyDecision []BusinessFlow;
    // 
    displayName *string;
    // 
    policyTemplates []GovernancePolicyTemplate;
}
// NewApprovalWorkflowProvider instantiates a new approvalWorkflowProvider and sets the default values.
func NewApprovalWorkflowProvider()(*ApprovalWorkflowProvider) {
    m := &ApprovalWorkflowProvider{
        Entity: *NewEntity(),
    }
    return m
}
// GetBusinessFlows gets the businessFlows property value. 
func (m *ApprovalWorkflowProvider) GetBusinessFlows()([]BusinessFlow) {
    if m == nil {
        return nil
    } else {
        return m.businessFlows
    }
}
// GetBusinessFlowsWithRequestsAwaitingMyDecision gets the businessFlowsWithRequestsAwaitingMyDecision property value. 
func (m *ApprovalWorkflowProvider) GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlow) {
    if m == nil {
        return nil
    } else {
        return m.businessFlowsWithRequestsAwaitingMyDecision
    }
}
// GetDisplayName gets the displayName property value. 
func (m *ApprovalWorkflowProvider) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetPolicyTemplates gets the policyTemplates property value. 
func (m *ApprovalWorkflowProvider) GetPolicyTemplates()([]GovernancePolicyTemplate) {
    if m == nil {
        return nil
    } else {
        return m.policyTemplates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalWorkflowProvider) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlow() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessFlow, len(val))
            for i, v := range val {
                res[i] = *(v.(*BusinessFlow))
            }
            m.SetBusinessFlows(res)
        }
        return nil
    }
    res["businessFlowsWithRequestsAwaitingMyDecision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlow() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessFlow, len(val))
            for i, v := range val {
                res[i] = *(v.(*BusinessFlow))
            }
            m.SetBusinessFlowsWithRequestsAwaitingMyDecision(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["policyTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernancePolicyTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernancePolicyTemplate, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernancePolicyTemplate))
            }
            m.SetPolicyTemplates(res)
        }
        return nil
    }
    return res
}
func (m *ApprovalWorkflowProvider) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApprovalWorkflowProvider) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBusinessFlows() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBusinessFlows()))
        for i, v := range m.GetBusinessFlows() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("businessFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessFlowsWithRequestsAwaitingMyDecision() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBusinessFlowsWithRequestsAwaitingMyDecision()))
        for i, v := range m.GetBusinessFlowsWithRequestsAwaitingMyDecision() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicyTemplates()))
        for i, v := range m.GetPolicyTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("policyTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBusinessFlows sets the businessFlows property value. 
func (m *ApprovalWorkflowProvider) SetBusinessFlows(value []BusinessFlow)() {
    if m != nil {
        m.businessFlows = value
    }
}
// SetBusinessFlowsWithRequestsAwaitingMyDecision sets the businessFlowsWithRequestsAwaitingMyDecision property value. 
func (m *ApprovalWorkflowProvider) SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlow)() {
    if m != nil {
        m.businessFlowsWithRequestsAwaitingMyDecision = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *ApprovalWorkflowProvider) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPolicyTemplates sets the policyTemplates property value. 
func (m *ApprovalWorkflowProvider) SetPolicyTemplates(value []GovernancePolicyTemplate)() {
    if m != nil {
        m.policyTemplates = value
    }
}
