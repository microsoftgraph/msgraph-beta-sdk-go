package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JournalLineable 
type JournalLineable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(Accountable)
    GetAccountId()(*string)
    GetAccountNumber()(*string)
    GetAmount()(*float64)
    GetComment()(*string)
    GetDescription()(*string)
    GetDocumentNumber()(*string)
    GetExternalDocumentNumber()(*string)
    GetJournalDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLineNumber()(*int32)
    GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAccount(value Accountable)()
    SetAccountId(value *string)()
    SetAccountNumber(value *string)()
    SetAmount(value *float64)()
    SetComment(value *string)()
    SetDescription(value *string)()
    SetDocumentNumber(value *string)()
    SetExternalDocumentNumber(value *string)()
    SetJournalDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLineNumber(value *int32)()
    SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
