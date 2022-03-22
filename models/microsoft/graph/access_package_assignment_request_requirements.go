package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequestRequirements 
type AccessPackageAssignmentRequestRequirements struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Answers that have already been provided.
    existingAnswers []AccessPackageAnswerable;
    // Indicates whether a request must be approved by an approver.
    isApprovalRequired *bool;
    // Indicates whether approval is required when a user tries to extend their access.
    isApprovalRequiredForExtension *bool;
    // Indicates whether the requestor is allowed to set a custom schedule.
    isCustomAssignmentScheduleAllowed *bool;
    // Indicates whether a requestor must supply justification when submitting an assignment request.
    isRequestorJustificationRequired *bool;
    // The description of the policy that the user is trying to request access using.
    policyDescription *string;
    // The display name of the policy that the user is trying to request access using.
    policyDisplayName *string;
    // The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
    policyId *string;
    // Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
    questions []AccessPackageQuestionable;
    // Schedule restrictions enforced, if any.
    schedule RequestScheduleable;
}
// NewAccessPackageAssignmentRequestRequirements instantiates a new accessPackageAssignmentRequestRequirements and sets the default values.
func NewAccessPackageAssignmentRequestRequirements()(*AccessPackageAssignmentRequestRequirements) {
    m := &AccessPackageAssignmentRequestRequirements{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessPackageAssignmentRequestRequirements(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentRequestRequirements) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExistingAnswers gets the existingAnswers property value. Answers that have already been provided.
func (m *AccessPackageAssignmentRequestRequirements) GetExistingAnswers()([]AccessPackageAnswerable) {
    if m == nil {
        return nil
    } else {
        return m.existingAnswers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequestRequirements) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["existingAnswers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAnswerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAnswerable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAnswerable)
            }
            m.SetExistingAnswers(res)
        }
        return nil
    }
    res["isApprovalRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequired(val)
        }
        return nil
    }
    res["isApprovalRequiredForExtension"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForExtension(val)
        }
        return nil
    }
    res["isCustomAssignmentScheduleAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomAssignmentScheduleAllowed(val)
        }
        return nil
    }
    res["isRequestorJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequestorJustificationRequired(val)
        }
        return nil
    }
    res["policyDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyDescription(val)
        }
        return nil
    }
    res["policyDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyDisplayName(val)
        }
        return nil
    }
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["questions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageQuestionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageQuestionable)
            }
            m.SetQuestions(res)
        }
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(RequestScheduleable))
        }
        return nil
    }
    return res
}
// GetIsApprovalRequired gets the isApprovalRequired property value. Indicates whether a request must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequired
    }
}
// GetIsApprovalRequiredForExtension gets the isApprovalRequiredForExtension property value. Indicates whether approval is required when a user tries to extend their access.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForExtension()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForExtension
    }
}
// GetIsCustomAssignmentScheduleAllowed gets the isCustomAssignmentScheduleAllowed property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) GetIsCustomAssignmentScheduleAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCustomAssignmentScheduleAllowed
    }
}
// GetIsRequestorJustificationRequired gets the isRequestorJustificationRequired property value. Indicates whether a requestor must supply justification when submitting an assignment request.
func (m *AccessPackageAssignmentRequestRequirements) GetIsRequestorJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequestorJustificationRequired
    }
}
// GetPolicyDescription gets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyDescription
    }
}
// GetPolicyDisplayName gets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyDisplayName
    }
}
// GetPolicyId gets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// GetQuestions gets the questions property value. Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
func (m *AccessPackageAssignmentRequestRequirements) GetQuestions()([]AccessPackageQuestionable) {
    if m == nil {
        return nil
    } else {
        return m.questions
    }
}
// GetSchedule gets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) GetSchedule()(RequestScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestRequirements) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetExistingAnswers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExistingAnswers()))
        for i, v := range m.GetExistingAnswers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetQuestions()))
        for i, v := range m.GetQuestions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
    if m != nil {
        m.additionalData = value
    }
}
// SetExistingAnswers sets the existingAnswers property value. Answers that have already been provided.
func (m *AccessPackageAssignmentRequestRequirements) SetExistingAnswers(value []AccessPackageAnswerable)() {
    if m != nil {
        m.existingAnswers = value
    }
}
// SetIsApprovalRequired sets the isApprovalRequired property value. Indicates whether a request must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequired(value *bool)() {
    if m != nil {
        m.isApprovalRequired = value
    }
}
// SetIsApprovalRequiredForExtension sets the isApprovalRequiredForExtension property value. Indicates whether approval is required when a user tries to extend their access.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForExtension(value *bool)() {
    if m != nil {
        m.isApprovalRequiredForExtension = value
    }
}
// SetIsCustomAssignmentScheduleAllowed sets the isCustomAssignmentScheduleAllowed property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) SetIsCustomAssignmentScheduleAllowed(value *bool)() {
    if m != nil {
        m.isCustomAssignmentScheduleAllowed = value
    }
}
// SetIsRequestorJustificationRequired sets the isRequestorJustificationRequired property value. Indicates whether a requestor must supply justification when submitting an assignment request.
func (m *AccessPackageAssignmentRequestRequirements) SetIsRequestorJustificationRequired(value *bool)() {
    if m != nil {
        m.isRequestorJustificationRequired = value
    }
}
// SetPolicyDescription sets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDescription(value *string)() {
    if m != nil {
        m.policyDescription = value
    }
}
// SetPolicyDisplayName sets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDisplayName(value *string)() {
    if m != nil {
        m.policyDisplayName = value
    }
}
// SetPolicyId sets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyId(value *string)() {
    if m != nil {
        m.policyId = value
    }
}
// SetQuestions sets the questions property value. Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
func (m *AccessPackageAssignmentRequestRequirements) SetQuestions(value []AccessPackageQuestionable)() {
    if m != nil {
        m.questions = value
    }
}
// SetSchedule sets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) SetSchedule(value RequestScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
