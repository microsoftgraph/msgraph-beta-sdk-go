package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TermsAndConditionsGroupAssignment struct {
    Entity
    // Unique identifier of a group that the T&C policy is assigned to.
    targetGroupId *string;
    // Navigation link to the terms and conditions that are assigned.
    termsAndConditions *TermsAndConditions;
}
// Instantiates a new termsAndConditionsGroupAssignment and sets the default values.
func NewTermsAndConditionsGroupAssignment()(*TermsAndConditionsGroupAssignment) {
    m := &TermsAndConditionsGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the targetGroupId property value. Unique identifier of a group that the T&C policy is assigned to.
func (m *TermsAndConditionsGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
// Gets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsGroupAssignment) GetTermsAndConditions()(*TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the targetGroupId property value. Unique identifier of a group that the T&C policy is assigned to.
// Parameters:
//  - value : Value to set for the targetGroupId property.
func (m *TermsAndConditionsGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
// Sets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
// Parameters:
//  - value : Value to set for the termsAndConditions property.
func (m *TermsAndConditionsGroupAssignment) SetTermsAndConditions(value *TermsAndConditions)() {
    m.termsAndConditions = value
}
