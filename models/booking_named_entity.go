package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingNamedEntity 
type BookingNamedEntity struct {
    Entity
    // A name for the derived entity, which interfaces with customers.
    displayName *string
}
// NewBookingNamedEntity instantiates a new BookingNamedEntity and sets the default values.
func NewBookingNamedEntity()(*BookingNamedEntity) {
    m := &BookingNamedEntity{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.bookingNamedEntity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBookingNamedEntityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingNamedEntityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.bookingBusiness":
                        return NewBookingBusiness(), nil
                    case "#microsoft.graph.bookingCustomer":
                        return NewBookingCustomer(), nil
                    case "#microsoft.graph.bookingPerson":
                        return NewBookingPerson(), nil
                    case "#microsoft.graph.bookingService":
                        return NewBookingService(), nil
                    case "#microsoft.graph.bookingStaffMember":
                        return NewBookingStaffMember(), nil
                }
            }
        }
    }
    return NewBookingNamedEntity(), nil
}
// GetDisplayName gets the displayName property value. A name for the derived entity, which interfaces with customers.
func (m *BookingNamedEntity) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingNamedEntity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    return res
}
// Serialize serializes information the current object
func (m *BookingNamedEntity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. A name for the derived entity, which interfaces with customers.
func (m *BookingNamedEntity) SetDisplayName(value *string)() {
    m.displayName = value
}
