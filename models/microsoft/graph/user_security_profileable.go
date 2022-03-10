package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSecurityProfileable 
type UserSecurityProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccounts()([]UserAccountable)
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRiskScore()(*string)
    GetTags()([]string)
    GetUserPrincipalName()(*string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetAccounts(value []UserAccountable)()
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRiskScore(value *string)()
    SetTags(value []string)()
    SetUserPrincipalName(value *string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
