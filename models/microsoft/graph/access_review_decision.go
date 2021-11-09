package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new accessReviewDecision and sets the default values.
func NewAccessReviewDecision()(*AccessReviewDecision) {
    m := &AccessReviewDecision{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessRecommendation property value. The feature- generated recommendation shown to the reviewer, one of Approve, Deny or NotAvailable.
func (m *AccessReviewDecision) GetAccessRecommendation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessRecommendation
    }
}
// Gets the accessReviewId property value. The feature-generated id of the access review.
func (m *AccessReviewDecision) GetAccessReviewId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewId
    }
}
// Gets the appliedBy property value. When the review completes, if the results were manually applied, the user identity of the user who applied the decision. If the review was auto-applied, the userPrincipalName is empty.
func (m *AccessReviewDecision) GetAppliedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.appliedBy
    }
}
// Gets the appliedDateTime property value. The date and time when the review decision was applied.
func (m *AccessReviewDecision) GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.appliedDateTime
    }
}
// Gets the applyResult property value. The outcome of applying the decision, one of NotApplied, Success, Failed, NotFound or NotSupported.
func (m *AccessReviewDecision) GetApplyResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applyResult
    }
}
// Gets the justification property value. The reviewer's business justification, if supplied.
func (m *AccessReviewDecision) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// Gets the reviewedBy property value. The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
func (m *AccessReviewDecision) GetReviewedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.reviewedBy
    }
}
// Gets the reviewedDateTime property value. 
func (m *AccessReviewDecision) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewedDateTime
    }
}
// Gets the reviewResult property value. The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
func (m *AccessReviewDecision) GetReviewResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the accessRecommendation property value. The feature- generated recommendation shown to the reviewer, one of Approve, Deny or NotAvailable.
// Parameters:
//  - value : Value to set for the accessRecommendation property.
func (m *AccessReviewDecision) SetAccessRecommendation(value *string)() {
    m.accessRecommendation = value
}
// Sets the accessReviewId property value. The feature-generated id of the access review.
// Parameters:
//  - value : Value to set for the accessReviewId property.
func (m *AccessReviewDecision) SetAccessReviewId(value *string)() {
    m.accessReviewId = value
}
// Sets the appliedBy property value. When the review completes, if the results were manually applied, the user identity of the user who applied the decision. If the review was auto-applied, the userPrincipalName is empty.
// Parameters:
//  - value : Value to set for the appliedBy property.
func (m *AccessReviewDecision) SetAppliedBy(value *UserIdentity)() {
    m.appliedBy = value
}
// Sets the appliedDateTime property value. The date and time when the review decision was applied.
// Parameters:
//  - value : Value to set for the appliedDateTime property.
func (m *AccessReviewDecision) SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.appliedDateTime = value
}
// Sets the applyResult property value. The outcome of applying the decision, one of NotApplied, Success, Failed, NotFound or NotSupported.
// Parameters:
//  - value : Value to set for the applyResult property.
func (m *AccessReviewDecision) SetApplyResult(value *string)() {
    m.applyResult = value
}
// Sets the justification property value. The reviewer's business justification, if supplied.
// Parameters:
//  - value : Value to set for the justification property.
func (m *AccessReviewDecision) SetJustification(value *string)() {
    m.justification = value
}
// Sets the reviewedBy property value. The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
// Parameters:
//  - value : Value to set for the reviewedBy property.
func (m *AccessReviewDecision) SetReviewedBy(value *UserIdentity)() {
    m.reviewedBy = value
}
// Sets the reviewedDateTime property value. 
// Parameters:
//  - value : Value to set for the reviewedDateTime property.
func (m *AccessReviewDecision) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
// Sets the reviewResult property value. The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
// Parameters:
//  - value : Value to set for the reviewResult property.
func (m *AccessReviewDecision) SetReviewResult(value *string)() {
    m.reviewResult = value
}
