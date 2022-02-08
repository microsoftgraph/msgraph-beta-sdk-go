package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingCustomer 
type BookingCustomer struct {
    BookingPerson
    // Addresses associated with the customer. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
    addresses []PhysicalAddress;
    // Phone numbers associated with the customer, including home, business and mobile numbers.
    phones []Phone;
}
// NewBookingCustomer instantiates a new bookingCustomer and sets the default values.
func NewBookingCustomer()(*BookingCustomer) {
    m := &BookingCustomer{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
// GetAddresses gets the addresses property value. Addresses associated with the customer. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingCustomer) GetAddresses()([]PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.addresses
    }
}
// GetPhones gets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
func (m *BookingCustomer) GetPhones()([]Phone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["addresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhysicalAddress, len(val))
            for i, v := range val {
                res[i] = *(v.(*PhysicalAddress))
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhone() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phone, len(val))
            for i, v := range val {
                res[i] = *(v.(*Phone))
            }
            m.SetPhones(res)
        }
        return nil
    }
    return res
}
func (m *BookingCustomer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingCustomer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BookingPerson.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddresses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddresses sets the addresses property value. Addresses associated with the customer. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingCustomer) SetAddresses(value []PhysicalAddress)() {
    if m != nil {
        m.addresses = value
    }
}
// SetPhones sets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
func (m *BookingCustomer) SetPhones(value []Phone)() {
    if m != nil {
        m.phones = value
    }
}
