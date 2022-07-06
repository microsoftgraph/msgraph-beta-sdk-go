package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrationBaseable 
type MeetingRegistrationBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedRegistrant()(*MeetingAudience)
    GetRegistrants()([]MeetingRegistrantBaseable)
    GetType()(*string)
    SetAllowedRegistrant(value *MeetingAudience)()
    SetRegistrants(value []MeetingRegistrantBaseable)()
    SetType(value *string)()
}
