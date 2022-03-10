package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PaymentTermable 
type PaymentTermable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCalculateDiscountOnCreditMemos()(*bool)
    GetCode()(*string)
    GetDiscountDateCalculation()(*string)
    GetDiscountPercent()(*float64)
    GetDisplayName()(*string)
    GetDueDateCalculation()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCalculateDiscountOnCreditMemos(value *bool)()
    SetCode(value *string)()
    SetDiscountDateCalculation(value *string)()
    SetDiscountPercent(value *float64)()
    SetDisplayName(value *string)()
    SetDueDateCalculation(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
