package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermsAndConditionsGroupAssignment 
type TermsAndConditionsGroupAssignment struct {
    Entity
    // Unique identifier of a group that the T&C policy is assigned to.
    targetGroupId *string;
    // Navigation link to the terms and conditions that are assigned.
    termsAndConditions *TermsAndConditions;
}
// NewTermsAndConditionsGroupAssignment instantiates a new termsAndConditionsGroupAssignment and sets the default values.
func NewTermsAndConditionsGroupAssignment()(*TermsAndConditionsGroupAssignment) {
    m := &TermsAndConditionsGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetTargetGroupId gets the targetGroupId property value. Unique identifier of a group that the T&C policy is assigned to.
func (m *TermsAndConditionsGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
// GetTermsAndConditions gets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsGroupAssignment) GetTermsAndConditions()(*TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermsAndConditionsGroupAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    res["termsAndConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsAndConditions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsAndConditions(val.(*TermsAndConditions))
        }
        return nil
    }
    return res
}
func (m *TermsAndConditionsGroupAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TermsAndConditionsGroupAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termsAndConditions", m.GetTermsAndConditions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetGroupId sets the targetGroupId property value. Unique identifier of a group that the T&C policy is assigned to.
func (m *TermsAndConditionsGroupAssignment) SetTargetGroupId(value *string)() {
    if m != nil {
        m.targetGroupId = value
    }
}
// SetTermsAndConditions sets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsGroupAssignment) SetTermsAndConditions(value *TermsAndConditions)() {
    if m != nil {
        m.termsAndConditions = value
    }
}
