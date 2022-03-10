package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365ServicesUserCountsable 
type Office365ServicesUserCountsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExchangeActive()(*int64)
    GetExchangeInactive()(*int64)
    GetOffice365Active()(*int64)
    GetOffice365Inactive()(*int64)
    GetOneDriveActive()(*int64)
    GetOneDriveInactive()(*int64)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetSharePointActive()(*int64)
    GetSharePointInactive()(*int64)
    GetSkypeForBusinessActive()(*int64)
    GetSkypeForBusinessInactive()(*int64)
    GetTeamsActive()(*int64)
    GetTeamsInactive()(*int64)
    GetYammerActive()(*int64)
    GetYammerInactive()(*int64)
    SetExchangeActive(value *int64)()
    SetExchangeInactive(value *int64)()
    SetOffice365Active(value *int64)()
    SetOffice365Inactive(value *int64)()
    SetOneDriveActive(value *int64)()
    SetOneDriveInactive(value *int64)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetSharePointActive(value *int64)()
    SetSharePointInactive(value *int64)()
    SetSkypeForBusinessActive(value *int64)()
    SetSkypeForBusinessInactive(value *int64)()
    SetTeamsActive(value *int64)()
    SetTeamsInactive(value *int64)()
    SetYammerActive(value *int64)()
    SetYammerInactive(value *int64)()
}
