package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewDecision 
type AccessReviewDecision struct {
    Entity
    // The feature- generated recommendation shown to the reviewer, one of Approve, Deny or NotAvailable.
    accessRecommendation *string;
    // The feature-generated id of the access review.
    accessReviewId *string;
    // When the review completes, if the results were manually applied, the user identity of the user who applied the decision. If the review was auto-applied, the userPrincipalName is empty.
    appliedBy *UserIdentity;
    // The date and time when the review decision was applied.
    appliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The outcome of applying the decision, one of NotApplied, Success, Failed, NotFound or NotSupported.
    applyResult *string;
    // The reviewer's business justification, if supplied.
    justification *string;
    // The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
    reviewedBy *UserIdentity;
    // 
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
    reviewResult *string;
}
// NewAccessReviewDecision instantiates a new accessReviewDecision and sets the default values.
func NewAccessReviewDecision()(*AccessReviewDecision) {
    m := &AccessReviewDecision{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessRecommendation gets the accessRecommendation property value. The feature- generated recommendation shown to the reviewer, one of Approve, Deny or NotAvailable.
func (m *AccessReviewDecision) GetAccessRecommendation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessRecommendation
    }
}
// GetAccessReviewId gets the accessReviewId property value. The feature-generated id of the access review.
func (m *AccessReviewDecision) GetAccessReviewId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewId
    }
}
// GetAppliedBy gets the appliedBy property value. When the review completes, if the results were manually applied, the user identity of the user who applied the decision. If the review was auto-applied, the userPrincipalName is empty.
func (m *AccessReviewDecision) GetAppliedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.appliedBy
    }
}
// GetAppliedDateTime gets the appliedDateTime property value. The date and time when the review decision was applied.
func (m *AccessReviewDecision) GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.appliedDateTime
    }
}
// GetApplyResult gets the applyResult property value. The outcome of applying the decision, one of NotApplied, Success, Failed, NotFound or NotSupported.
func (m *AccessReviewDecision) GetApplyResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applyResult
    }
}
// GetJustification gets the justification property value. The reviewer's business justification, if supplied.
func (m *AccessReviewDecision) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// GetReviewedBy gets the reviewedBy property value. The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
func (m *AccessReviewDecision) GetReviewedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.reviewedBy
    }
}
// GetReviewedDateTime gets the reviewedDateTime property value. 
func (m *AccessReviewDecision) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewedDateTime
    }
}
// GetReviewResult gets the reviewResult property value. The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
func (m *AccessReviewDecision) GetReviewResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewDecision) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessRecommendation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessRecommendation(val)
        }
        return nil
    }
    res["accessReviewId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewId(val)
        }
        return nil
    }
    res["appliedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedBy(val.(*UserIdentity))
        }
        return nil
    }
    res["appliedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedDateTime(val)
        }
        return nil
    }
    res["applyResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyResult(val)
        }
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["reviewedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedBy(val.(*UserIdentity))
        }
        return nil
    }
    res["reviewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedDateTime(val)
        }
        return nil
    }
    res["reviewResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *AccessReviewDecision) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewDecision) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.accessRecommendation = value
    }
}
// SetAccessReviewId sets the accessReviewId property value. The feature-generated id of the access review.
func (m *AccessReviewDecision) SetAccessReviewId(value *string)() {
    if m != nil {
        m.accessReviewId = value
    }
}
// SetAppliedBy sets the appliedBy property value. When the review completes, if the results were manually applied, the user identity of the user who applied the decision. If the review was auto-applied, the userPrincipalName is empty.
func (m *AccessReviewDecision) SetAppliedBy(value *UserIdentity)() {
    if m != nil {
        m.appliedBy = value
    }
}
// SetAppliedDateTime sets the appliedDateTime property value. The date and time when the review decision was applied.
func (m *AccessReviewDecision) SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.appliedDateTime = value
    }
}
// SetApplyResult sets the applyResult property value. The outcome of applying the decision, one of NotApplied, Success, Failed, NotFound or NotSupported.
func (m *AccessReviewDecision) SetApplyResult(value *string)() {
    if m != nil {
        m.applyResult = value
    }
}
// SetJustification sets the justification property value. The reviewer's business justification, if supplied.
func (m *AccessReviewDecision) SetJustification(value *string)() {
    if m != nil {
        m.justification = value
    }
}
// SetReviewedBy sets the reviewedBy property value. The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
func (m *AccessReviewDecision) SetReviewedBy(value *UserIdentity)() {
    if m != nil {
        m.reviewedBy = value
    }
}
// SetReviewedDateTime sets the reviewedDateTime property value. 
func (m *AccessReviewDecision) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewedDateTime = value
    }
}
// SetReviewResult sets the reviewResult property value. The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
func (m *AccessReviewDecision) SetReviewResult(value *string)() {
    if m != nil {
        m.reviewResult = value
    }
}
