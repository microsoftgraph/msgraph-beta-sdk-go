package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingPerson struct {
    BookingNamedEntity
    emailAddress *string;
}
func NewBookingPerson()(*BookingPerson) {
    m := &BookingPerson{
        BookingNamedEntity: *NewBookingNamedEntity(),
    }
    return m
}
func (m *BookingPerson) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
func (m *BookingPerson) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingNamedEntity.GetFieldDeserializers()
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailAddress(val)
        return nil
    }
    return res
}
func (m *BookingPerson) IsNil()(bool) {
    return m == nil
}
func (m *BookingPerson) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BookingNamedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *BookingPerson) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
