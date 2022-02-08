package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingRegistrationQuestion 
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
// NewMeetingRegistrationQuestion instantiates a new meetingRegistrationQuestion and sets the default values.
func NewMeetingRegistrationQuestion()(*MeetingRegistrationQuestion) {
    m := &MeetingRegistrationQuestion{
        Entity: *NewEntity(),
    }
    return m
}
// GetAnswerInputType gets the answerInputType property value. Answer input type of the custom registration question.
func (m *MeetingRegistrationQuestion) GetAnswerInputType()(*AnswerInputType) {
    if m == nil {
        return nil
    } else {
        return m.answerInputType
    }
}
// GetAnswerOptions gets the answerOptions property value. Answer options when answerInputType is radioButton.
func (m *MeetingRegistrationQuestion) GetAnswerOptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.answerOptions
    }
}
// GetDisplayName gets the displayName property value. Display name of the custom registration question.
func (m *MeetingRegistrationQuestion) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsRequired gets the isRequired property value. Indicates whether the question is required. Default value is false.
func (m *MeetingRegistrationQuestion) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrationQuestion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["answerInputType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnswerInputType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnswerInputType(val.(*AnswerInputType))
        }
        return nil
    }
    res["answerOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAnswerOptions(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    return res
}
func (m *MeetingRegistrationQuestion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingRegistrationQuestion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnswerInputType() != nil {
        cast := (*m.GetAnswerInputType()).String()
        err = writer.WriteStringValue("answerInputType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnswerOptions() != nil {
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
// SetAnswerInputType sets the answerInputType property value. Answer input type of the custom registration question.
func (m *MeetingRegistrationQuestion) SetAnswerInputType(value *AnswerInputType)() {
    if m != nil {
        m.answerInputType = value
    }
}
// SetAnswerOptions sets the answerOptions property value. Answer options when answerInputType is radioButton.
func (m *MeetingRegistrationQuestion) SetAnswerOptions(value []string)() {
    if m != nil {
        m.answerOptions = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the custom registration question.
func (m *MeetingRegistrationQuestion) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsRequired sets the isRequired property value. Indicates whether the question is required. Default value is false.
func (m *MeetingRegistrationQuestion) SetIsRequired(value *bool)() {
    if m != nil {
        m.isRequired = value
    }
}
