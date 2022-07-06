package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingPersonable 
type BookingPersonable interface {
    BookingNamedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailAddress()(*string)
    GetType()(*string)
    SetEmailAddress(value *string)()
    SetType(value *string)()
}
