package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MeetingRegistrationQuestion struct {
    Entity
    // Answer input type of the custom registration question.
    answerInputType *AnswerInputType;
    // Answer options when answerInputType is radioButton.
    answerOptions []string;
    // Display name of the custom registration question.
    displayName *string;
    // Indicates whether the question is required. Default value is false.
    isRequired *bool;
}
// Instantiates a new meetingRegistrationQuestion and sets the default values.
func NewMeetingRegistrationQuestion()(*MeetingRegistrationQuestion) {
    m := &MeetingRegistrationQuestion{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the answerInputType property value. Answer input type of the custom registration question.
func (m *MeetingRegistrationQuestion) GetAnswerInputType()(*AnswerInputType) {
    if m == nil {
        return nil
    } else {
        return m.answerInputType
    }
}
// Gets the answerOptions property value. Answer options when answerInputType is radioButton.
func (m *MeetingRegistrationQuestion) GetAnswerOptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.answerOptions
    }
}
// Gets the displayName property value. Display name of the custom registration question.
func (m *MeetingRegistrationQuestion) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isRequired property value. Indicates whether the question is required. Default value is false.
func (m *MeetingRegistrationQuestion) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// The deserialization information for the current model
func (m *MeetingRegistrationQuestion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["answerInputType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnswerInputType)
        if err != nil {
            return err
        }
        cast := val.(AnswerInputType)
        m.SetAnswerInputType(&cast)
        return nil
    }
    res["answerOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAnswerOptions(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRequired(val)
        return nil
    }
    return res
}
func (m *MeetingRegistrationQuestion) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MeetingRegistrationQuestion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnswerInputType() != nil {
        cast := m.GetAnswerInputType().String()
        err = writer.WriteStringValue("answerInputType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("answerOptions", m.GetAnswerOptions())
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
        err = writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the answerInputType property value. Answer input type of the custom registration question.
// Parameters:
//  - value : Value to set for the answerInputType property.
func (m *MeetingRegistrationQuestion) SetAnswerInputType(value *AnswerInputType)() {
    m.answerInputType = value
}
// Sets the answerOptions property value. Answer options when answerInputType is radioButton.
// Parameters:
//  - value : Value to set for the answerOptions property.
func (m *MeetingRegistrationQuestion) SetAnswerOptions(value []string)() {
    m.answerOptions = value
}
// Sets the displayName property value. Display name of the custom registration question.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MeetingRegistrationQuestion) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isRequired property value. Indicates whether the question is required. Default value is false.
// Parameters:
//  - value : Value to set for the isRequired property.
func (m *MeetingRegistrationQuestion) SetIsRequired(value *bool)() {
    m.isRequired = value
}
