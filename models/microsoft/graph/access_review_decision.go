package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewDecision struct {
    Entity
    accessRecommendation *string;
    accessReviewId *string;
    appliedBy *UserIdentity;
    appliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    applyResult *string;
    justification *string;
    reviewedBy *UserIdentity;
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    reviewResult *string;
}
func NewAccessReviewDecision()(*AccessReviewDecision) {
    m := &AccessReviewDecision{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AccessReviewDecision) GetAccessRecommendation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessRecommendation
    }
}
func (m *AccessReviewDecision) GetAccessReviewId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewId
    }
}
func (m *AccessReviewDecision) GetAppliedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.appliedBy
    }
}
func (m *AccessReviewDecision) GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.appliedDateTime
    }
}
func (m *AccessReviewDecision) GetApplyResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applyResult
    }
}
func (m *AccessReviewDecision) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
func (m *AccessReviewDecision) GetReviewedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.reviewedBy
    }
}
func (m *AccessReviewDecision) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewedDateTime
    }
}
func (m *AccessReviewDecision) GetReviewResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
func (m *AccessReviewDecision) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessRecommendation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccessRecommendation(val)
        return nil
    }
    res["accessReviewId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccessReviewId(val)
        return nil
    }
    res["appliedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetAppliedBy(val.(*UserIdentity))
        return nil
    }
    res["appliedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAppliedDateTime(val)
        return nil
    }
    res["applyResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplyResult(val)
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["reviewedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetReviewedBy(val.(*UserIdentity))
        return nil
    }
    res["reviewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReviewedDateTime(val)
        return nil
    }
    res["reviewResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReviewResult(val)
        return nil
    }
    return res
}
func (m *AccessReviewDecision) IsNil()(bool) {
    return m == nil
}
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
func (m *AccessReviewDecision) SetAccessRecommendation(value *string)() {
    m.accessRecommendation = value
}
func (m *AccessReviewDecision) SetAccessReviewId(value *string)() {
    m.accessReviewId = value
}
func (m *AccessReviewDecision) SetAppliedBy(value *UserIdentity)() {
    m.appliedBy = value
}
func (m *AccessReviewDecision) SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.appliedDateTime = value
}
func (m *AccessReviewDecision) SetApplyResult(value *string)() {
    m.applyResult = value
}
func (m *AccessReviewDecision) SetJustification(value *string)() {
    m.justification = value
}
func (m *AccessReviewDecision) SetReviewedBy(value *UserIdentity)() {
    m.reviewedBy = value
}
func (m *AccessReviewDecision) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
func (m *AccessReviewDecision) SetReviewResult(value *string)() {
    m.reviewResult = value
}
