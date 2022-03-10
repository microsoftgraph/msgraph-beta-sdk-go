package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalSettingsable 
type ApprovalSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApprovalMode()(*string)
    GetApprovalStages()([]ApprovalStageable)
    GetIsApprovalRequired()(*bool)
    GetIsApprovalRequiredForExtension()(*bool)
    GetIsRequestorJustificationRequired()(*bool)
    SetApprovalMode(value *string)()
    SetApprovalStages(value []ApprovalStageable)()
    SetIsApprovalRequired(value *bool)()
    SetIsApprovalRequiredForExtension(value *bool)()
    SetIsRequestorJustificationRequired(value *bool)()
}
