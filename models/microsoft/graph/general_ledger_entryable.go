package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GeneralLedgerEntryable 
type GeneralLedgerEntryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccount()(Accountable)
    GetAccountId()(*string)
    GetAccountNumber()(*string)
    GetCreditAmount()(*float64)
    GetDebitAmount()(*float64)
    GetDescription()(*string)
    GetDocumentNumber()(*string)
    GetDocumentType()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPostingDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    SetAccount(value Accountable)()
    SetAccountId(value *string)()
    SetAccountNumber(value *string)()
    SetCreditAmount(value *float64)()
    SetDebitAmount(value *float64)()
    SetDescription(value *string)()
    SetDocumentNumber(value *string)()
    SetDocumentType(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPostingDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
}
