package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestRequirementsable 
type AccessPackageAssignmentRequestRequirementsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
