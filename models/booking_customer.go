package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomer represents a customer of the business.
type BookingCustomer struct {
    BookingPerson
}
// NewBookingCustomer instantiates a new BookingCustomer and sets the default values.
func NewBookingCustomer()(*BookingCustomer) {
    m := &BookingCustomer{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
// CreateBookingCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingCustomerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingCustomer(), nil
}
// GetAddresses gets the addresses property value. Addresses associated with the customer, including home, business and other addresses.
// returns a []PhysicalAddressable when successful
func (m *BookingCustomer) GetAddresses()([]PhysicalAddressable) {
    val, err := m.GetBackingStore().Get("addresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PhysicalAddressable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date, time and timezone when the customer was created.
// returns a *Time when successful
func (m *BookingCustomer) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingCustomer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhysicalAddressable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PhysicalAddressable)
                }
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phoneable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Phoneable)
                }
            }
            m.SetPhones(res)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date, time and timezone when the customer was last updated.
// returns a *Time when successful
func (m *BookingCustomer) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPhones gets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
// returns a []Phoneable when successful
func (m *BookingCustomer) GetPhones()([]Phoneable) {
    val, err := m.GetBackingStore().Get("phones")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Phoneable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BookingCustomer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingPerson.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddresses sets the addresses property value. Addresses associated with the customer, including home, business and other addresses.
func (m *BookingCustomer) SetAddresses(value []PhysicalAddressable)() {
    err := m.GetBackingStore().Set("addresses", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date, time and timezone when the customer was created.
func (m *BookingCustomer) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date, time and timezone when the customer was last updated.
func (m *BookingCustomer) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPhones sets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
func (m *BookingCustomer) SetPhones(value []Phoneable)() {
    err := m.GetBackingStore().Set("phones", value)
    if err != nil {
        panic(err)
    }
}
type BookingCustomerable interface {
    BookingPersonable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddresses()([]PhysicalAddressable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPhones()([]Phoneable)
    SetAddresses(value []PhysicalAddressable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPhones(value []Phoneable)()
}
