package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentReviewSettings 
type AssignmentReviewSettings struct {
    // The default decision to apply if the request is not reviewed within the period specified in durationInDays. The possible values are: acceptAccessRecommendation, keepAccess, removeAccess, and unknownFutureValue.
    accessReviewTimeoutBehavior *AccessReviewTimeoutBehavior
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of days within which reviewers should provide input.
    durationInDays *int32
    // Specifies whether to display recommendations to the reviewer. The default value is true
    isAccessRecommendationEnabled *bool
    // Specifies whether the reviewer must provide justification for the approval. The default value is true.
    isApprovalJustificationRequired *bool
    // If true, access reviews are required for assignments from this policy.
    isEnabled *bool
    // The OdataType property
    odataType *string
    // The interval for recurrence, such as monthly or quarterly.
    recurrenceType *string
    // If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers.
    reviewers []UserSetable
    // Who should be asked to do the review, either Self or Reviewers.
    reviewerType *string
    // When the first review should start.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewAssignmentReviewSettings instantiates a new assignmentReviewSettings and sets the default values.
func NewAssignmentReviewSettings()(*AssignmentReviewSettings) {
    m := &AssignmentReviewSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.assignmentReviewSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAssignmentReviewSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentReviewSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentReviewSettings(), nil
}
// GetAccessReviewTimeoutBehavior gets the accessReviewTimeoutBehavior property value. The default decision to apply if the request is not reviewed within the period specified in durationInDays. The possible values are: acceptAccessRecommendation, keepAccess, removeAccess, and unknownFutureValue.
func (m *AssignmentReviewSettings) GetAccessReviewTimeoutBehavior()(*AccessReviewTimeoutBehavior) {
    return m.accessReviewTimeoutBehavior
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentReviewSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDurationInDays gets the durationInDays property value. The number of days within which reviewers should provide input.
func (m *AssignmentReviewSettings) GetDurationInDays()(*int32) {
    return m.durationInDays
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentReviewSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessReviewTimeoutBehavior"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAccessReviewTimeoutBehavior , m.SetAccessReviewTimeoutBehavior)
    res["durationInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDurationInDays)
    res["isAccessRecommendationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAccessRecommendationEnabled)
    res["isApprovalJustificationRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsApprovalJustificationRequired)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["recurrenceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRecurrenceType)
    res["reviewers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserSetFromDiscriminatorValue , m.SetReviewers)
    res["reviewerType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetReviewerType)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    return res
}
// GetIsAccessRecommendationEnabled gets the isAccessRecommendationEnabled property value. Specifies whether to display recommendations to the reviewer. The default value is true
func (m *AssignmentReviewSettings) GetIsAccessRecommendationEnabled()(*bool) {
    return m.isAccessRecommendationEnabled
}
// GetIsApprovalJustificationRequired gets the isApprovalJustificationRequired property value. Specifies whether the reviewer must provide justification for the approval. The default value is true.
func (m *AssignmentReviewSettings) GetIsApprovalJustificationRequired()(*bool) {
    return m.isApprovalJustificationRequired
}
// GetIsEnabled gets the isEnabled property value. If true, access reviews are required for assignments from this policy.
func (m *AssignmentReviewSettings) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentReviewSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRecurrenceType gets the recurrenceType property value. The interval for recurrence, such as monthly or quarterly.
func (m *AssignmentReviewSettings) GetRecurrenceType()(*string) {
    return m.recurrenceType
}
// GetReviewers gets the reviewers property value. If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers.
func (m *AssignmentReviewSettings) GetReviewers()([]UserSetable) {
    return m.reviewers
}
// GetReviewerType gets the reviewerType property value. Who should be asked to do the review, either Self or Reviewers.
func (m *AssignmentReviewSettings) GetReviewerType()(*string) {
    return m.reviewerType
}
// GetStartDateTime gets the startDateTime property value. When the first review should start.
func (m *AssignmentReviewSettings) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *AssignmentReviewSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessReviewTimeoutBehavior() != nil {
        cast := (*m.GetAccessReviewTimeoutBehavior()).String()
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    if m.GetReviewers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReviewers())
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
// SetAccessReviewTimeoutBehavior sets the accessReviewTimeoutBehavior property value. The default decision to apply if the request is not reviewed within the period specified in durationInDays. The possible values are: acceptAccessRecommendation, keepAccess, removeAccess, and unknownFutureValue.
func (m *AssignmentReviewSettings) SetAccessReviewTimeoutBehavior(value *AccessReviewTimeoutBehavior)() {
    m.accessReviewTimeoutBehavior = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentReviewSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDurationInDays sets the durationInDays property value. The number of days within which reviewers should provide input.
func (m *AssignmentReviewSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// SetIsAccessRecommendationEnabled sets the isAccessRecommendationEnabled property value. Specifies whether to display recommendations to the reviewer. The default value is true
func (m *AssignmentReviewSettings) SetIsAccessRecommendationEnabled(value *bool)() {
    m.isAccessRecommendationEnabled = value
}
// SetIsApprovalJustificationRequired sets the isApprovalJustificationRequired property value. Specifies whether the reviewer must provide justification for the approval. The default value is true.
func (m *AssignmentReviewSettings) SetIsApprovalJustificationRequired(value *bool)() {
    m.isApprovalJustificationRequired = value
}
// SetIsEnabled sets the isEnabled property value. If true, access reviews are required for assignments from this policy.
func (m *AssignmentReviewSettings) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentReviewSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecurrenceType sets the recurrenceType property value. The interval for recurrence, such as monthly or quarterly.
func (m *AssignmentReviewSettings) SetRecurrenceType(value *string)() {
    m.recurrenceType = value
}
// SetReviewers sets the reviewers property value. If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers.
func (m *AssignmentReviewSettings) SetReviewers(value []UserSetable)() {
    m.reviewers = value
}
// SetReviewerType sets the reviewerType property value. Who should be asked to do the review, either Self or Reviewers.
func (m *AssignmentReviewSettings) SetReviewerType(value *string)() {
    m.reviewerType = value
}
// SetStartDateTime sets the startDateTime property value. When the first review should start.
func (m *AssignmentReviewSettings) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
