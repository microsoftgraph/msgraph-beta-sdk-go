package search

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Acronym provides operations to manage the searchEntity singleton.
type Acronym struct {
    SearchAnswer
    // What the acronym stands for.
    standsFor *string;
    // State of the acronym. Possible values are: published, draft, excluded, or unknownFutureValue.
    state *AnswerState;
}
// NewAcronym instantiates a new acronym and sets the default values.
func NewAcronym()(*Acronym) {
    m := &Acronym{
        SearchAnswer: *NewSearchAnswer(),
    }
    return m
}
// CreateAcronymFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAcronymFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAcronym(), nil
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
        val, err := n.GetEnumValue(ParseAnswerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AnswerState))
        }
        return nil
    }
    return res
}
// GetStandsFor gets the standsFor property value. What the acronym stands for.
func (m *Acronym) GetStandsFor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.standsFor
    }
}
// GetState gets the state property value. State of the acronym. Possible values are: published, draft, excluded, or unknownFutureValue.
func (m *Acronym) GetState()(*AnswerState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
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
// SetStandsFor sets the standsFor property value. What the acronym stands for.
func (m *Acronym) SetStandsFor(value *string)() {
    if m != nil {
        m.standsFor = value
    }
}
// SetState sets the state property value. State of the acronym. Possible values are: published, draft, excluded, or unknownFutureValue.
func (m *Acronym) SetState(value *AnswerState)() {
    if m != nil {
        m.state = value
    }
}
