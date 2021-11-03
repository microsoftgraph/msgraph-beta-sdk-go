package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MuteParticipantsOperation struct {
    CommsOperation
    // 
    participants []string;
}
// Instantiates a new muteParticipantsOperation and sets the default values.
func NewMuteParticipantsOperation()(*MuteParticipantsOperation) {
    m := &MuteParticipantsOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// Gets the participants property value. 
func (m *MuteParticipantsOperation) GetParticipants()([]string) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// The deserialization information for the current model
func (m *MuteParticipantsOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetParticipants(res)
        return nil
    }
    return res
}
func (m *MuteParticipantsOperation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MuteParticipantsOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the participants property value. 
// Parameters:
//  - value : Value to set for the participants property.
func (m *MuteParticipantsOperation) SetParticipants(value []string)() {
    m.participants = value
}
