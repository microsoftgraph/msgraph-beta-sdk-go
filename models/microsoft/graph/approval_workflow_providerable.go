package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalWorkflowProviderable 
type ApprovalWorkflowProviderable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBusinessFlows()([]BusinessFlowable)
    GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlowable)
    GetDisplayName()(*string)
    GetPolicyTemplates()([]GovernancePolicyTemplateable)
    SetBusinessFlows(value []BusinessFlowable)()
    SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlowable)()
    SetDisplayName(value *string)()
    SetPolicyTemplates(value []GovernancePolicyTemplateable)()
}
