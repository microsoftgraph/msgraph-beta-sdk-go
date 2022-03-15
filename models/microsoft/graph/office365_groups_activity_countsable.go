package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365GroupsActivityCountsable 
type Office365GroupsActivityCountsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExchangeEmailsReceived()(*int32)
    GetReportDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetYammerMessagesLiked()(*int32)
    GetYammerMessagesPosted()(*int32)
    GetYammerMessagesRead()(*int32)
    SetExchangeEmailsReceived(value *int32)()
    SetReportDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetYammerMessagesLiked(value *int32)()
    SetYammerMessagesPosted(value *int32)()
    SetYammerMessagesRead(value *int32)()
}
