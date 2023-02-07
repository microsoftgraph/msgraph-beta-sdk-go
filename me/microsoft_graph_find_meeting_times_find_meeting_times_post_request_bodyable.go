package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// MicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBodyable 
type MicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)
    GetIsOrganizerOptional()(*bool)
    GetLocationConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)
    GetMaxCandidates()(*int32)
    GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMinimumAttendeePercentage()(*float64)
    GetReturnSuggestionReasons()(*bool)
    GetTimeConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)
    SetAttendees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)()
    SetIsOrganizerOptional(value *bool)()
    SetLocationConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)()
    SetMaxCandidates(value *int32)()
    SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMinimumAttendeePercentage(value *float64)()
    SetReturnSuggestionReasons(value *bool)()
    SetTimeConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)()
}
