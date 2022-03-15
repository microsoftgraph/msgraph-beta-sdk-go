package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365ServicesUserCountsable 
type Office365ServicesUserCountsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExchangeActive()(*int32)
    GetExchangeInactive()(*int32)
    GetOffice365Active()(*int32)
    GetOffice365Inactive()(*int32)
    GetOneDriveActive()(*int32)
    GetOneDriveInactive()(*int32)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetSharePointActive()(*int32)
    GetSharePointInactive()(*int32)
    GetSkypeForBusinessActive()(*int32)
    GetSkypeForBusinessInactive()(*int32)
    GetTeamsActive()(*int32)
    GetTeamsInactive()(*int32)
    GetYammerActive()(*int32)
    GetYammerInactive()(*int32)
    SetExchangeActive(value *int32)()
    SetExchangeInactive(value *int32)()
    SetOffice365Active(value *int32)()
    SetOffice365Inactive(value *int32)()
    SetOneDriveActive(value *int32)()
    SetOneDriveInactive(value *int32)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetSharePointActive(value *int32)()
    SetSharePointInactive(value *int32)()
    SetSkypeForBusinessActive(value *int32)()
    SetSkypeForBusinessInactive(value *int32)()
    SetTeamsActive(value *int32)()
    SetTeamsInactive(value *int32)()
    SetYammerActive(value *int32)()
    SetYammerInactive(value *int32)()
}
