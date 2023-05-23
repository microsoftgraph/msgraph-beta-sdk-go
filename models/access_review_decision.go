package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewDecision 
type AccessReviewDecision struct {
    Entity
}
// NewAccessReviewDecision instantiates a new AccessReviewDecision and sets the default values.
func NewAccessReviewDecision()(*AccessReviewDecision) {
    m := &AccessReviewDecision{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewDecisionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewDecisionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewDecision(), nil
}
// GetAccessRecommendation gets the accessRecommendation property value. The feature- generated recommendation shown to the reviewer, one of Approve, Deny or NotAvailable.
func (m *AccessReviewDecision) GetAccessRecommendation()(*string) {
    val, err := m.GetBackingStore().Get("accessRecommendation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAccessReviewId gets the accessReviewId property value. The feature-generated id of the access review.
func (m *AccessReviewDecision) GetAccessReviewId()(*string) {
    val, err := m.GetBackingStore().Get("accessReviewId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppliedBy gets the appliedBy property value. When the review completes, if the results were manually applied, the user identity of the user who applied the decision. If the review was auto-applied, the userPrincipalName is empty.
func (m *AccessReviewDecision) GetAppliedBy()(UserIdentityable) {
    val, err := m.GetBackingStore().Get("appliedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserIdentityable)
    }
    return nil
}
// GetAppliedDateTime gets the appliedDateTime property value. The date and time when the review decision was applied.
func (m *AccessReviewDecision) GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("appliedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetApplyResult gets the applyResult property value. The outcome of applying the decision, one of NotApplied, Success, Failed, NotFound or NotSupported.
func (m *AccessReviewDecision) GetApplyResult()(*string) {
    val, err := m.GetBackingStore().Get("applyResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewDecision) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessRecommendation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessRecommendation(val)
        }
        return nil
    }
    res["accessReviewId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewId(val)
        }
        return nil
    }
    res["appliedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["appliedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedDateTime(val)
        }
        return nil
    }
    res["applyResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyResult(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["reviewedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["reviewedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedDateTime(val)
        }
        return nil
    }
    res["reviewResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewResult(val)
        }
        return nil
    }
    return res
}
// GetJustification gets the justification property value. The reviewer's business justification, if supplied.
func (m *AccessReviewDecision) GetJustification()(*string) {
    val, err := m.GetBackingStore().Get("justification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReviewedBy gets the reviewedBy property value. The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
func (m *AccessReviewDecision) GetReviewedBy()(UserIdentityable) {
    val, err := m.GetBackingStore().Get("reviewedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserIdentityable)
    }
    return nil
}
// GetReviewedDateTime gets the reviewedDateTime property value. The reviewedDateTime property
func (m *AccessReviewDecision) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("reviewedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReviewResult gets the reviewResult property value. The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
func (m *AccessReviewDecision) GetReviewResult()(*string) {
    val, err := m.GetBackingStore().Get("reviewResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReviewDecision) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accessRecommendation", m.GetAccessRecommendation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("accessReviewId", m.GetAccessReviewId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appliedBy", m.GetAppliedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("appliedDateTime", m.GetAppliedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applyResult", m.GetApplyResult())
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
    return nil
}
// SetAccessRecommendation sets the accessRecommendation property value. The feature- generated recommendation shown to the reviewer, one of Approve, Deny or NotAvailable.
func (m *AccessReviewDecision) SetAccessRecommendation(value *string)() {
    err := m.GetBackingStore().Set("accessRecommendation", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessReviewId sets the accessReviewId property value. The feature-generated id of the access review.
func (m *AccessReviewDecision) SetAccessReviewId(value *string)() {
    err := m.GetBackingStore().Set("accessReviewId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedBy sets the appliedBy property value. When the review completes, if the results were manually applied, the user identity of the user who applied the decision. If the review was auto-applied, the userPrincipalName is empty.
func (m *AccessReviewDecision) SetAppliedBy(value UserIdentityable)() {
    err := m.GetBackingStore().Set("appliedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedDateTime sets the appliedDateTime property value. The date and time when the review decision was applied.
func (m *AccessReviewDecision) SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("appliedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetApplyResult sets the applyResult property value. The outcome of applying the decision, one of NotApplied, Success, Failed, NotFound or NotSupported.
func (m *AccessReviewDecision) SetApplyResult(value *string)() {
    err := m.GetBackingStore().Set("applyResult", value)
    if err != nil {
        panic(err)
    }
}
// SetJustification sets the justification property value. The reviewer's business justification, if supplied.
func (m *AccessReviewDecision) SetJustification(value *string)() {
    err := m.GetBackingStore().Set("justification", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewedBy sets the reviewedBy property value. The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
func (m *AccessReviewDecision) SetReviewedBy(value UserIdentityable)() {
    err := m.GetBackingStore().Set("reviewedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewedDateTime sets the reviewedDateTime property value. The reviewedDateTime property
func (m *AccessReviewDecision) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("reviewedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewResult sets the reviewResult property value. The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
func (m *AccessReviewDecision) SetReviewResult(value *string)() {
    err := m.GetBackingStore().Set("reviewResult", value)
    if err != nil {
        panic(err)
    }
}
// AccessReviewDecisionable 
type AccessReviewDecisionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessRecommendation()(*string)
    GetAccessReviewId()(*string)
    GetAppliedBy()(UserIdentityable)
    GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetApplyResult()(*string)
    GetJustification()(*string)
    GetReviewedBy()(UserIdentityable)
    GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewResult()(*string)
    SetAccessRecommendation(value *string)()
    SetAccessReviewId(value *string)()
    SetAppliedBy(value UserIdentityable)()
    SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetApplyResult(value *string)()
    SetJustification(value *string)()
    SetReviewedBy(value UserIdentityable)()
    SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewResult(value *string)()
}
