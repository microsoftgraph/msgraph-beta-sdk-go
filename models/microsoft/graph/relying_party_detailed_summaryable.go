package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RelyingPartyDetailedSummaryable 
type RelyingPartyDetailedSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFailedSignInCount()(*int64)
    GetMigrationStatus()(*MigrationStatus)
    GetMigrationValidationDetails()([]KeyValuePairable)
    GetRelyingPartyId()(*string)
    GetRelyingPartyName()(*string)
    GetReplyUrls()([]string)
    GetServiceId()(*string)
    GetSignInSuccessRate()(*float64)
    GetSuccessfulSignInCount()(*int64)
    GetTotalSignInCount()(*int64)
    GetUniqueUserCount()(*int64)
    SetFailedSignInCount(value *int64)()
    SetMigrationStatus(value *MigrationStatus)()
    SetMigrationValidationDetails(value []KeyValuePairable)()
    SetRelyingPartyId(value *string)()
    SetRelyingPartyName(value *string)()
    SetReplyUrls(value []string)()
    SetServiceId(value *string)()
    SetSignInSuccessRate(value *float64)()
    SetSuccessfulSignInCount(value *int64)()
    SetTotalSignInCount(value *int64)()
    SetUniqueUserCount(value *int64)()
}
