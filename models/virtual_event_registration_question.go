package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventRegistrationQuestion 
type VirtualEventRegistrationQuestion struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewVirtualEventRegistrationQuestion instantiates a new virtualEventRegistrationQuestion and sets the default values.
func NewVirtualEventRegistrationQuestion()(*VirtualEventRegistrationQuestion) {
    m := &VirtualEventRegistrationQuestion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventRegistrationQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventRegistrationQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventRegistrationQuestion(), nil
}
// GetAnswerChoices gets the answerChoices property value. Answer choices when answerInputType is singleChoice or multiChoice.
func (m *VirtualEventRegistrationQuestion) GetAnswerChoices()([]string) {
    val, err := m.GetBackingStore().Get("answerChoices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAnswerInputType gets the answerInputType property value. Input type of the registration question answer.
func (m *VirtualEventRegistrationQuestion) GetAnswerInputType()(*VirtualEventRegistrationQuestionAnswerInputType) {
    val, err := m.GetBackingStore().Get("answerInputType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VirtualEventRegistrationQuestionAnswerInputType)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display name of the registration question.
func (m *VirtualEventRegistrationQuestion) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventRegistrationQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIsRequired gets the isRequired property value. Indicates whether the question is required to answer. Default value is false.
func (m *VirtualEventRegistrationQuestion) GetIsRequired()(*bool) {
    val, err := m.GetBackingStore().Get("isRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventRegistrationQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
// SetAnswerChoices sets the answerChoices property value. Answer choices when answerInputType is singleChoice or multiChoice.
func (m *VirtualEventRegistrationQuestion) SetAnswerChoices(value []string)() {
    err := m.GetBackingStore().Set("answerChoices", value)
    if err != nil {
        panic(err)
    }
}
// SetAnswerInputType sets the answerInputType property value. Input type of the registration question answer.
func (m *VirtualEventRegistrationQuestion) SetAnswerInputType(value *VirtualEventRegistrationQuestionAnswerInputType)() {
    err := m.GetBackingStore().Set("answerInputType", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name of the registration question.
func (m *VirtualEventRegistrationQuestion) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRequired sets the isRequired property value. Indicates whether the question is required to answer. Default value is false.
func (m *VirtualEventRegistrationQuestion) SetIsRequired(value *bool)() {
    err := m.GetBackingStore().Set("isRequired", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventRegistrationQuestionable 
type VirtualEventRegistrationQuestionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnswerChoices()([]string)
    GetAnswerInputType()(*VirtualEventRegistrationQuestionAnswerInputType)
    GetDisplayName()(*string)
    GetIsRequired()(*bool)
    SetAnswerChoices(value []string)()
    SetAnswerInputType(value *VirtualEventRegistrationQuestionAnswerInputType)()
    SetDisplayName(value *string)()
    SetIsRequired(value *bool)()
}
