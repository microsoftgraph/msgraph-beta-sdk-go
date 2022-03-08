package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DomainSecurityProfileable 
type DomainSecurityProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivityGroupNames()([]string)
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCountHits()(*int32)
    GetCountInOrg()(*int32)
    GetDomainCategories()([]ReputationCategoryable)
    GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetRegistrant()(DomainRegistrantable)
    GetRiskScore()(*string)
    GetTags()([]string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetActivityGroupNames(value []string)()
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCountHits(value *int32)()
    SetCountInOrg(value *int32)()
    SetDomainCategories(value []ReputationCategoryable)()
    SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetRegistrant(value DomainRegistrantable)()
    SetRiskScore(value *string)()
    SetTags(value []string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
