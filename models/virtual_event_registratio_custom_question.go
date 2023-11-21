package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventRegistratioCustomQuestion 
type VirtualEventRegistratioCustomQuestion struct {
    VirtualEventRegistrationQuestionBase
}
// NewVirtualEventRegistratioCustomQuestion instantiates a new virtualEventRegistratioCustomQuestion and sets the default values.
func NewVirtualEventRegistratioCustomQuestion()(*VirtualEventRegistratioCustomQuestion) {
    m := &VirtualEventRegistratioCustomQuestion{
        VirtualEventRegistrationQuestionBase: *NewVirtualEventRegistrationQuestionBase(),
    }
    odataTypeValue := "#microsoft.graph.virtualEventRegistratioCustomQuestion"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVirtualEventRegistratioCustomQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventRegistratioCustomQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventRegistratioCustomQuestion(), nil
}
// GetAnswerChoices gets the answerChoices property value. The answerChoices property
func (m *VirtualEventRegistratioCustomQuestion) GetAnswerChoices()([]string) {
    val, err := m.GetBackingStore().Get("answerChoices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAnswerInputType gets the answerInputType property value. The answerInputType property
func (m *VirtualEventRegistratioCustomQuestion) GetAnswerInputType()(*VirtualEventRegistrationQuestionAnswerInputType) {
    val, err := m.GetBackingStore().Get("answerInputType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VirtualEventRegistrationQuestionAnswerInputType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventRegistratioCustomQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VirtualEventRegistrationQuestionBase.GetFieldDeserializers()
    res["answerChoices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAnswerChoices(res)
        }
        return nil
    }
    res["answerInputType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVirtualEventRegistrationQuestionAnswerInputType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnswerInputType(val.(*VirtualEventRegistrationQuestionAnswerInputType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *VirtualEventRegistratioCustomQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VirtualEventRegistrationQuestionBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnswerChoices() != nil {
        err = writer.WriteCollectionOfStringValues("answerChoices", m.GetAnswerChoices())
        if err != nil {
            return err
        }
    }
    if m.GetAnswerInputType() != nil {
        cast := (*m.GetAnswerInputType()).String()
        err = writer.WriteStringValue("answerInputType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnswerChoices sets the answerChoices property value. The answerChoices property
func (m *VirtualEventRegistratioCustomQuestion) SetAnswerChoices(value []string)() {
    err := m.GetBackingStore().Set("answerChoices", value)
    if err != nil {
        panic(err)
    }
}
// SetAnswerInputType sets the answerInputType property value. The answerInputType property
func (m *VirtualEventRegistratioCustomQuestion) SetAnswerInputType(value *VirtualEventRegistrationQuestionAnswerInputType)() {
    err := m.GetBackingStore().Set("answerInputType", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventRegistratioCustomQuestionable 
type VirtualEventRegistratioCustomQuestionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventRegistrationQuestionBaseable
    GetAnswerChoices()([]string)
    GetAnswerInputType()(*VirtualEventRegistrationQuestionAnswerInputType)
    SetAnswerChoices(value []string)()
    SetAnswerInputType(value *VirtualEventRegistrationQuestionAnswerInputType)()
}
