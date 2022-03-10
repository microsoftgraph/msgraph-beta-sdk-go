package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalStageable 
type ApprovalStageable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApprovalStageTimeOutInDays()(*int32)
    GetEscalationApprovers()([]UserSetable)
    GetEscalationTimeInMinutes()(*int32)
    GetIsApproverJustificationRequired()(*bool)
    GetIsEscalationEnabled()(*bool)
    GetPrimaryApprovers()([]UserSetable)
    SetApprovalStageTimeOutInDays(value *int32)()
    SetEscalationApprovers(value []UserSetable)()
    SetEscalationTimeInMinutes(value *int32)()
    SetIsApproverJustificationRequired(value *bool)()
    SetIsEscalationEnabled(value *bool)()
    SetPrimaryApprovers(value []UserSetable)()
}
