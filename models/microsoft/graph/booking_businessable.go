package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingBusinessable 
type BookingBusinessable interface {
    BookingNamedEntityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddress()(PhysicalAddressable)
    GetAppointments()([]BookingAppointmentable)
    GetBusinessHours()([]BookingWorkHoursable)
    GetBusinessType()(*string)
    GetCalendarView()([]BookingAppointmentable)
    GetCustomers()([]BookingCustomerable)
    GetCustomQuestions()([]BookingCustomQuestionable)
    GetDefaultCurrencyIso()(*string)
    GetEmail()(*string)
    GetIsPublished()(*bool)
    GetPhone()(*string)
    GetPublicUrl()(*string)
    GetSchedulingPolicy()(BookingSchedulingPolicyable)
    GetServices()([]BookingServiceable)
    GetStaffMembers()([]BookingStaffMemberable)
    GetWebSiteUrl()(*string)
    SetAddress(value PhysicalAddressable)()
    SetAppointments(value []BookingAppointmentable)()
    SetBusinessHours(value []BookingWorkHoursable)()
    SetBusinessType(value *string)()
    SetCalendarView(value []BookingAppointmentable)()
    SetCustomers(value []BookingCustomerable)()
    SetCustomQuestions(value []BookingCustomQuestionable)()
    SetDefaultCurrencyIso(value *string)()
    SetEmail(value *string)()
    SetIsPublished(value *bool)()
    SetPhone(value *string)()
    SetPublicUrl(value *string)()
    SetSchedulingPolicy(value BookingSchedulingPolicyable)()
    SetServices(value []BookingServiceable)()
    SetStaffMembers(value []BookingStaffMemberable)()
    SetWebSiteUrl(value *string)()
}
