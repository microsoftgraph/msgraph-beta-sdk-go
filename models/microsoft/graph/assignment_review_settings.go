package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AssignmentReviewSettings struct {
    // The default decision to apply if the request is not reviewed within the period specified in durationInDays. The possible values are: acceptAccessRecommendation, keepAccess, removeAccess, and unknownFutureValue.
    accessReviewTimeoutBehavior *AccessReviewTimeoutBehavior;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of days within which reviewers should provide input.
    durationInDays *int32;
    // Specifies whether to display recommendations to the reviewer. The default value is true
    isAccessRecommendationEnabled *bool;
    // Specifies whether the reviewer must provide justification for the approval. The default value is true.
    isApprovalJustificationRequired *bool;
    // If true, access reviews are required for assignments from this policy.
    isEnabled *bool;
    // The interval for recurrence, such as monthly or quarterly.
    recurrenceType *string;
    // If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers.
    reviewers []UserSet;
    // Who should be asked to do the review, either Self or Reviewers.
    reviewerType *string;
    // When the first review should start.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new assignmentReviewSettings and sets the default values.
func NewAssignmentReviewSettings()(*AssignmentReviewSettings) {
    m := &AssignmentReviewSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the accessReviewTimeoutBehavior property value. The default decision to apply if the request is not reviewed within the period specified in durationInDays. The possible values are: acceptAccessRecommendation, keepAccess, removeAccess, and unknownFutureValue.
func (m *AssignmentReviewSettings) GetAccessReviewTimeoutBehavior()(*AccessReviewTimeoutBehavior) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewTimeoutBehavior
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentReviewSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the durationInDays property value. The number of days within which reviewers should provide input.
func (m *AssignmentReviewSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// Gets the isAccessRecommendationEnabled property value. Specifies whether to display recommendations to the reviewer. The default value is true
func (m *AssignmentReviewSettings) GetIsAccessRecommendationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAccessRecommendationEnabled
    }
}
// Gets the isApprovalJustificationRequired property value. Specifies whether the reviewer must provide justification for the approval. The default value is true.
func (m *AssignmentReviewSettings) GetIsApprovalJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalJustificationRequired
    }
}
// Gets the isEnabled property value. If true, access reviews are required for assignments from this policy.
func (m *AssignmentReviewSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the recurrenceType property value. The interval for recurrence, such as monthly or quarterly.
func (m *AssignmentReviewSettings) GetRecurrenceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceType
    }
}
// Gets the reviewers property value. If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers.
func (m *AssignmentReviewSettings) GetReviewers()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// Gets the reviewerType property value. Who should be asked to do the review, either Self or Reviewers.
func (m *AssignmentReviewSettings) GetReviewerType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewerType
    }
}
// Gets the startDateTime property value. When the first review should start.
func (m *AssignmentReviewSettings) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// The deserialization information for the current model
func (m *AssignmentReviewSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessReviewTimeoutBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessReviewTimeoutBehavior)
        if err != nil {
            return err
        }
        cast := val.(AccessReviewTimeoutBehavior)
        m.SetAccessReviewTimeoutBehavior(&cast)
        return nil
    }
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDurationInDays(val)
        return nil
    }
    res["isAccessRecommendationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAccessRecommendationEnabled(val)
        return nil
    }
    res["isApprovalJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsApprovalJustificationRequired(val)
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["recurrenceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecurrenceType(val)
        return nil
    }
    res["reviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        res := make([]UserSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSet))
        }
        m.SetReviewers(res)
        return nil
    }
    res["reviewerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReviewerType(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    return res
}
func (m *AssignmentReviewSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AssignmentReviewSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessReviewTimeoutBehavior() != nil {
        cast := m.GetAccessReviewTimeoutBehavior().String()
        err := writer.WriteStringValue("accessReviewTimeoutBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAccessRecommendationEnabled", m.GetIsAccessRecommendationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalJustificationRequired", m.GetIsApprovalJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceType", m.GetRecurrenceType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reviewerType", m.GetReviewerType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// Sets the accessReviewTimeoutBehavior property value. The default decision to apply if the request is not reviewed within the period specified in durationInDays. The possible values are: acceptAccessRecommendation, keepAccess, removeAccess, and unknownFutureValue.
// Parameters:
//  - value : Value to set for the accessReviewTimeoutBehavior property.
func (m *AssignmentReviewSettings) SetAccessReviewTimeoutBehavior(value *AccessReviewTimeoutBehavior)() {
    m.accessReviewTimeoutBehavior = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AssignmentReviewSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the durationInDays property value. The number of days within which reviewers should provide input.
// Parameters:
//  - value : Value to set for the durationInDays property.
func (m *AssignmentReviewSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// Sets the isAccessRecommendationEnabled property value. Specifies whether to display recommendations to the reviewer. The default value is true
// Parameters:
//  - value : Value to set for the isAccessRecommendationEnabled property.
func (m *AssignmentReviewSettings) SetIsAccessRecommendationEnabled(value *bool)() {
    m.isAccessRecommendationEnabled = value
}
// Sets the isApprovalJustificationRequired property value. Specifies whether the reviewer must provide justification for the approval. The default value is true.
// Parameters:
//  - value : Value to set for the isApprovalJustificationRequired property.
func (m *AssignmentReviewSettings) SetIsApprovalJustificationRequired(value *bool)() {
    m.isApprovalJustificationRequired = value
}
// Sets the isEnabled property value. If true, access reviews are required for assignments from this policy.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *AssignmentReviewSettings) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the recurrenceType property value. The interval for recurrence, such as monthly or quarterly.
// Parameters:
//  - value : Value to set for the recurrenceType property.
func (m *AssignmentReviewSettings) SetRecurrenceType(value *string)() {
    m.recurrenceType = value
}
// Sets the reviewers property value. If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers.
// Parameters:
//  - value : Value to set for the reviewers property.
func (m *AssignmentReviewSettings) SetReviewers(value []UserSet)() {
    m.reviewers = value
}
// Sets the reviewerType property value. Who should be asked to do the review, either Self or Reviewers.
// Parameters:
//  - value : Value to set for the reviewerType property.
func (m *AssignmentReviewSettings) SetReviewerType(value *string)() {
    m.reviewerType = value
}
// Sets the startDateTime property value. When the first review should start.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *AssignmentReviewSettings) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
