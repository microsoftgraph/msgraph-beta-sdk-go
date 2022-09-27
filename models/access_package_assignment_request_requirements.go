package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestRequirements 
type AccessPackageAssignmentRequestRequirements struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Answers that have already been provided.
    existingAnswers []AccessPackageAnswerable
    // Indicates whether a request must be approved by an approver.
    isApprovalRequired *bool
    // Indicates whether approval is required when a user tries to extend their access.
    isApprovalRequiredForExtension *bool
    // Indicates whether the requestor is allowed to set a custom schedule.
    isCustomAssignmentScheduleAllowed *bool
    // Indicates whether a requestor must supply justification when submitting an assignment request.
    isRequestorJustificationRequired *bool
    // The OdataType property
    odataType *string
    // The description of the policy that the user is trying to request access using.
    policyDescription *string
    // The display name of the policy that the user is trying to request access using.
    policyDisplayName *string
    // The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
    policyId *string
    // Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
    questions []AccessPackageQuestionable
    // Schedule restrictions enforced, if any.
    schedule RequestScheduleable
}
// NewAccessPackageAssignmentRequestRequirements instantiates a new accessPackageAssignmentRequestRequirements and sets the default values.
func NewAccessPackageAssignmentRequestRequirements()(*AccessPackageAssignmentRequestRequirements) {
    m := &AccessPackageAssignmentRequestRequirements{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.accessPackageAssignmentRequestRequirements";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequestRequirements(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentRequestRequirements) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExistingAnswers gets the existingAnswers property value. Answers that have already been provided.
func (m *AccessPackageAssignmentRequestRequirements) GetExistingAnswers()([]AccessPackageAnswerable) {
    return m.existingAnswers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequestRequirements) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["existingAnswers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageAnswerFromDiscriminatorValue , m.SetExistingAnswers)
    res["isApprovalRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsApprovalRequired)
    res["isApprovalRequiredForExtension"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsApprovalRequiredForExtension)
    res["isCustomAssignmentScheduleAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCustomAssignmentScheduleAllowed)
    res["isRequestorJustificationRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRequestorJustificationRequired)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["policyDescription"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPolicyDescription)
    res["policyDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPolicyDisplayName)
    res["policyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPolicyId)
    res["questions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageQuestionFromDiscriminatorValue , m.SetQuestions)
    res["schedule"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRequestScheduleFromDiscriminatorValue , m.SetSchedule)
    return res
}
// GetIsApprovalRequired gets the isApprovalRequired property value. Indicates whether a request must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequired()(*bool) {
    return m.isApprovalRequired
}
// GetIsApprovalRequiredForExtension gets the isApprovalRequiredForExtension property value. Indicates whether approval is required when a user tries to extend their access.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForExtension()(*bool) {
    return m.isApprovalRequiredForExtension
}
// GetIsCustomAssignmentScheduleAllowed gets the isCustomAssignmentScheduleAllowed property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) GetIsCustomAssignmentScheduleAllowed()(*bool) {
    return m.isCustomAssignmentScheduleAllowed
}
// GetIsRequestorJustificationRequired gets the isRequestorJustificationRequired property value. Indicates whether a requestor must supply justification when submitting an assignment request.
func (m *AccessPackageAssignmentRequestRequirements) GetIsRequestorJustificationRequired()(*bool) {
    return m.isRequestorJustificationRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentRequestRequirements) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyDescription gets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDescription()(*string) {
    return m.policyDescription
}
// GetPolicyDisplayName gets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDisplayName()(*string) {
    return m.policyDisplayName
}
// GetPolicyId gets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyId()(*string) {
    return m.policyId
}
// GetQuestions gets the questions property value. Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
func (m *AccessPackageAssignmentRequestRequirements) GetQuestions()([]AccessPackageQuestionable) {
    return m.questions
}
// GetSchedule gets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) GetSchedule()(RequestScheduleable) {
    return m.schedule
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestRequirements) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExistingAnswers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExistingAnswers())
        err := writer.WriteCollectionOfObjectValues("existingAnswers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequired", m.GetIsApprovalRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForExtension", m.GetIsApprovalRequiredForExtension())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCustomAssignmentScheduleAllowed", m.GetIsCustomAssignmentScheduleAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequestorJustificationRequired", m.GetIsRequestorJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyDescription", m.GetPolicyDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyDisplayName", m.GetPolicyDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    if m.GetQuestions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetQuestions())
        err := writer.WriteCollectionOfObjectValues("questions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("schedule", m.GetSchedule())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentRequestRequirements) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExistingAnswers sets the existingAnswers property value. Answers that have already been provided.
func (m *AccessPackageAssignmentRequestRequirements) SetExistingAnswers(value []AccessPackageAnswerable)() {
    m.existingAnswers = value
}
// SetIsApprovalRequired sets the isApprovalRequired property value. Indicates whether a request must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequired(value *bool)() {
    m.isApprovalRequired = value
}
// SetIsApprovalRequiredForExtension sets the isApprovalRequiredForExtension property value. Indicates whether approval is required when a user tries to extend their access.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForExtension(value *bool)() {
    m.isApprovalRequiredForExtension = value
}
// SetIsCustomAssignmentScheduleAllowed sets the isCustomAssignmentScheduleAllowed property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) SetIsCustomAssignmentScheduleAllowed(value *bool)() {
    m.isCustomAssignmentScheduleAllowed = value
}
// SetIsRequestorJustificationRequired sets the isRequestorJustificationRequired property value. Indicates whether a requestor must supply justification when submitting an assignment request.
func (m *AccessPackageAssignmentRequestRequirements) SetIsRequestorJustificationRequired(value *bool)() {
    m.isRequestorJustificationRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentRequestRequirements) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyDescription sets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDescription(value *string)() {
    m.policyDescription = value
}
// SetPolicyDisplayName sets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDisplayName(value *string)() {
    m.policyDisplayName = value
}
// SetPolicyId sets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyId(value *string)() {
    m.policyId = value
}
// SetQuestions sets the questions property value. Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
func (m *AccessPackageAssignmentRequestRequirements) SetQuestions(value []AccessPackageQuestionable)() {
    m.questions = value
}
// SetSchedule sets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) SetSchedule(value RequestScheduleable)() {
    m.schedule = value
}
