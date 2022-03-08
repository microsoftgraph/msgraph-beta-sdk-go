package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgedAccountsReceivableable 
type AgedAccountsReceivableable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAgedAsOfDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetBalanceDue()(*float64)
    GetCurrencyCode()(*string)
    GetCurrentAmount()(*float64)
    GetCustomerNumber()(*string)
    GetName()(*string)
    GetPeriod1Amount()(*float64)
    GetPeriod2Amount()(*float64)
    GetPeriod3Amount()(*float64)
    GetPeriodLengthFilter()(*string)
    SetAgedAsOfDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetBalanceDue(value *float64)()
    SetCurrencyCode(value *string)()
    SetCurrentAmount(value *float64)()
    SetCustomerNumber(value *string)()
    SetName(value *string)()
    SetPeriod1Amount(value *float64)()
    SetPeriod2Amount(value *float64)()
    SetPeriod3Amount(value *float64)()
    SetPeriodLengthFilter(value *string)()
}
