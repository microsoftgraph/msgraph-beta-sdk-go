package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Calendarable 
type Calendarable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedOnlineMeetingProviders()([]string)
    GetCalendarGroupId()(*string)
    GetCalendarPermissions()([]CalendarPermissionable)
    GetCalendarView()([]Eventable)
    GetCanEdit()(*bool)
    GetCanShare()(*bool)
    GetCanViewPrivateItems()(*bool)
    GetChangeKey()(*string)
    GetColor()(*CalendarColor)
    GetDefaultOnlineMeetingProvider()(*OnlineMeetingProviderType)
    GetEvents()([]Eventable)
    GetHexColor()(*string)
    GetIsDefaultCalendar()(*bool)
    GetIsRemovable()(*bool)
    GetIsShared()(*bool)
    GetIsSharedWithMe()(*bool)
    GetIsTallyingResponses()(*bool)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetName()(*string)
    GetOwner()(EmailAddressable)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    SetAllowedOnlineMeetingProviders(value []string)()
    SetCalendarGroupId(value *string)()
    SetCalendarPermissions(value []CalendarPermissionable)()
    SetCalendarView(value []Eventable)()
    SetCanEdit(value *bool)()
    SetCanShare(value *bool)()
    SetCanViewPrivateItems(value *bool)()
    SetChangeKey(value *string)()
    SetColor(value *CalendarColor)()
    SetDefaultOnlineMeetingProvider(value *OnlineMeetingProviderType)()
    SetEvents(value []Eventable)()
    SetHexColor(value *string)()
    SetIsDefaultCalendar(value *bool)()
    SetIsRemovable(value *bool)()
    SetIsShared(value *bool)()
    SetIsSharedWithMe(value *bool)()
    SetIsTallyingResponses(value *bool)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetName(value *string)()
    SetOwner(value EmailAddressable)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
}
