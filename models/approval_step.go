package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalStep provides operations to manage the collection of accessReviewDecision entities.
type ApprovalStep struct {
    Entity
    // Indicates whether the step is assigned to the calling user to review. Read-only.
    assignedToMe *bool
    // The label provided by the policy creator to identify an approval step. Read-only.
    displayName *string
    // The justification associated with the approval step decision.
    justification *string
    // The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Read-only.
    reviewedBy Identityable
    // The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
    reviewResult *string
    // The step status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
    status *string
}
// NewApprovalStep instantiates a new approvalStep and sets the default values.
func NewApprovalStep()(*ApprovalStep) {
    m := &ApprovalStep{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.approvalStep";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateApprovalStepFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalStepFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalStep(), nil
}
// GetAssignedToMe gets the assignedToMe property value. Indicates whether the step is assigned to the calling user to review. Read-only.
func (m *ApprovalStep) GetAssignedToMe()(*bool) {
    return m.assignedToMe
}
// GetDisplayName gets the displayName property value. The label provided by the policy creator to identify an approval step. Read-only.
func (m *ApprovalStep) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalStep) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedToMe"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAssignedToMe)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["justification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJustification)
    res["reviewedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentityFromDiscriminatorValue , m.SetReviewedBy)
    res["reviewedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetReviewedDateTime)
    res["reviewResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetReviewResult)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    return res
}
// GetJustification gets the justification property value. The justification associated with the approval step decision.
func (m *ApprovalStep) GetJustification()(*string) {
    return m.justification
}
// GetReviewedBy gets the reviewedBy property value. The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Read-only.
func (m *ApprovalStep) GetReviewedBy()(Identityable) {
    return m.reviewedBy
}
// GetReviewedDateTime gets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ApprovalStep) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewedDateTime
}
// GetReviewResult gets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
func (m *ApprovalStep) GetReviewResult()(*string) {
    return m.reviewResult
}
// GetStatus gets the status property value. The step status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
func (m *ApprovalStep) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *ApprovalStep) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("assignedToMe", m.GetAssignedToMe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedBy", m.GetReviewedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewedDateTime", m.GetReviewedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewResult", m.GetReviewResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedToMe sets the assignedToMe property value. Indicates whether the step is assigned to the calling user to review. Read-only.
func (m *ApprovalStep) SetAssignedToMe(value *bool)() {
    m.assignedToMe = value
}
// SetDisplayName sets the displayName property value. The label provided by the policy creator to identify an approval step. Read-only.
func (m *ApprovalStep) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJustification sets the justification property value. The justification associated with the approval step decision.
func (m *ApprovalStep) SetJustification(value *string)() {
    m.justification = value
}
// SetReviewedBy sets the reviewedBy property value. The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Read-only.
func (m *ApprovalStep) SetReviewedBy(value Identityable)() {
    m.reviewedBy = value
}
// SetReviewedDateTime sets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ApprovalStep) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
// SetReviewResult sets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
func (m *ApprovalStep) SetReviewResult(value *string)() {
    m.reviewResult = value
}
// SetStatus sets the status property value. The step status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
func (m *ApprovalStep) SetStatus(value *string)() {
    m.status = value
}
