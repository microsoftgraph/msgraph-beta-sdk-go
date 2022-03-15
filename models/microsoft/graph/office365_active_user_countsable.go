package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365ActiveUserCountsable 
type Office365ActiveUserCountsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExchange()(*int32)
    GetOffice365()(*int32)
    GetOneDrive()(*int32)
    GetReportDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetSharePoint()(*int32)
    GetSkypeForBusiness()(*int32)
    GetTeams()(*int32)
    GetYammer()(*int32)
    SetExchange(value *int32)()
    SetOffice365(value *int32)()
    SetOneDrive(value *int32)()
    SetReportDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetSharePoint(value *int32)()
    SetSkypeForBusiness(value *int32)()
    SetTeams(value *int32)()
    SetYammer(value *int32)()
}
