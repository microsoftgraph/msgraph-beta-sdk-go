package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttendeeBase struct {
    Recipient
    type_escaped *AttendeeType;
}
func NewAttendeeBase()(*AttendeeBase) {
    m := &AttendeeBase{
        Recipient: *NewRecipient(),
    }
    return m
}
func (m *AttendeeBase) GetType_escaped()(*AttendeeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *AttendeeBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Recipient.GetFieldDeserializers()
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttendeeType)
        if err != nil {
            return err
        }
        cast := val.(AttendeeType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *AttendeeBase) IsNil()(bool) {
    return m == nil
}
func (m *AttendeeBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Recipient.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AttendeeBase) SetType_escaped(value *AttendeeType)() {
    m.type_escaped = value
}
