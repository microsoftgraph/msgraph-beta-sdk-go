package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/search"
)

// Acronym 
type Acronym struct {
    SearchAnswer
    // 
    standsFor *string;
    // 
    state *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState;
}
// NewAcronym instantiates a new acronym and sets the default values.
func NewAcronym()(*Acronym) {
    m := &Acronym{
        SearchAnswer: *NewSearchAnswer(),
    }
    return m
}
// GetStandsFor gets the standsFor property value. 
func (m *Acronym) GetStandsFor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.standsFor
    }
}
// GetState gets the state property value. 
func (m *Acronym) GetState()(*id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Acronym) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.SearchAnswer.GetFieldDeserializers()
    res["standsFor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStandsFor(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.ParseAnswerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState))
        }
        return nil
    }
    return res
}
func (m *Acronym) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Acronym) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.SearchAnswer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("standsFor", m.GetStandsFor())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStandsFor sets the standsFor property value. 
func (m *Acronym) SetStandsFor(value *string)() {
    if m != nil {
        m.standsFor = value
    }
}
// SetState sets the state property value. 
func (m *Acronym) SetState(value *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)() {
    if m != nil {
        m.state = value
    }
}
