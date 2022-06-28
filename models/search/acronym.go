package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Acronym 
type Acronym struct {
    SearchAnswer
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // What the acronym stands for.
    standsFor *string
    // State of the acronym. Possible values are: published, draft, excluded, or unknownFutureValue.
    state *AnswerState
}
// NewAcronym instantiates a new Acronym and sets the default values.
func NewAcronym()(*Acronym) {
    m := &Acronym{
        SearchAnswer: *NewSearchAnswer(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAcronymFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAcronymFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAcronym(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Acronym) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Acronym) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SearchAnswer.GetFieldDeserializers()
    res["standsFor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStandsFor(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *Acronym) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Acronym) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
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
