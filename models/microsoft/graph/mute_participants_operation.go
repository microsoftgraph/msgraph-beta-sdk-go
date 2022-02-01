package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MuteParticipantsOperation 
type MuteParticipantsOperation struct {
    CommsOperation
    // 
    participants []string;
}
// NewMuteParticipantsOperation instantiates a new muteParticipantsOperation and sets the default values.
func NewMuteParticipantsOperation()(*MuteParticipantsOperation) {
    m := &MuteParticipantsOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// GetParticipants gets the participants property value. 
func (m *MuteParticipantsOperation) GetParticipants()([]string) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MuteParticipantsOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetParticipants(res)
        }
        return nil
    }
    return res
}
func (m *MuteParticipantsOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MuteParticipantsOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetParticipants() != nil {
        err = writer.WriteCollectionOfStringValues("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParticipants sets the participants property value. 
func (m *MuteParticipantsOperation) SetParticipants(value []string)() {
    if m != nil {
        m.participants = value
    }
}
