package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingServiceable 
type BookingServiceable interface {
    BookingNamedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()(*string)
    GetCustomQuestions()([]BookingQuestionAssignmentable)
    GetDefaultDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDefaultLocation()(Locationable)
    GetDefaultPrice()(*float64)
    GetDefaultPriceType()(*BookingPriceType)
    GetDefaultReminders()([]BookingReminderable)
    GetDescription()(*string)
    GetIsHiddenFromCustomers()(*bool)
    GetIsLocationOnline()(*bool)
    GetMaximumAttendeesCount()(*int32)
    GetNotes()(*string)
    GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetSchedulingPolicy()(BookingSchedulingPolicyable)
    GetSmsNotificationsEnabled()(*bool)
    GetStaffMemberIds()([]string)
    GetWebUrl()(*string)
    SetAdditionalInformation(value *string)()
    SetCustomQuestions(value []BookingQuestionAssignmentable)()
    SetDefaultDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDefaultLocation(value Locationable)()
    SetDefaultPrice(value *float64)()
    SetDefaultPriceType(value *BookingPriceType)()
    SetDefaultReminders(value []BookingReminderable)()
    SetDescription(value *string)()
    SetIsHiddenFromCustomers(value *bool)()
    SetIsLocationOnline(value *bool)()
    SetMaximumAttendeesCount(value *int32)()
    SetNotes(value *string)()
    SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetSchedulingPolicy(value BookingSchedulingPolicyable)()
    SetSmsNotificationsEnabled(value *bool)()
    SetStaffMemberIds(value []string)()
    SetWebUrl(value *string)()
}