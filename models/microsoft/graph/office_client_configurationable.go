package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeClientConfigurationable 
type OfficeClientConfigurationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]OfficeClientConfigurationAssignmentable)
    GetCheckinStatuses()([]OfficeClientCheckinStatusable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetPolicyPayload()([]byte)
    GetPriority()(*int32)
    GetUserCheckinSummary()(OfficeUserCheckinSummaryable)
    GetUserPreferencePayload()([]byte)
    SetAssignments(value []OfficeClientConfigurationAssignmentable)()
    SetCheckinStatuses(value []OfficeClientCheckinStatusable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetPolicyPayload(value []byte)()
    SetPriority(value *int32)()
    SetUserCheckinSummary(value OfficeUserCheckinSummaryable)()
    SetUserPreferencePayload(value []byte)()
}
