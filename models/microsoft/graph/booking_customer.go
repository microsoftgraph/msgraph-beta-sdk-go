package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingCustomer struct {
    BookingPerson
}
func NewBookingCustomer()(*BookingCustomer) {
    m := &BookingCustomer{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
func (m *BookingCustomer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    return res
}
func (m *BookingCustomer) IsNil()(bool) {
    return m == nil
}
func (m *BookingCustomer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BookingPerson.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
