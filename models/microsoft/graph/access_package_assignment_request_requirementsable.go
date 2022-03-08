package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequestRequirementsable 
type AccessPackageAssignmentRequestRequirementsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExistingAnswers()([]AccessPackageAnswerable)
    GetIsApprovalRequired()(*bool)
    GetIsApprovalRequiredForExtension()(*bool)
    GetIsCustomAssignmentScheduleAllowed()(*bool)
    GetIsRequestorJustificationRequired()(*bool)
    GetPolicyDescription()(*string)
    GetPolicyDisplayName()(*string)
    GetPolicyId()(*string)
    GetQuestions()([]AccessPackageQuestionable)
    GetSchedule()(RequestScheduleable)
    SetExistingAnswers(value []AccessPackageAnswerable)()
    SetIsApprovalRequired(value *bool)()
    SetIsApprovalRequiredForExtension(value *bool)()
    SetIsCustomAssignmentScheduleAllowed(value *bool)()
    SetIsRequestorJustificationRequired(value *bool)()
    SetPolicyDescription(value *string)()
    SetPolicyDisplayName(value *string)()
    SetPolicyId(value *string)()
    SetQuestions(value []AccessPackageQuestionable)()
    SetSchedule(value RequestScheduleable)()
}
