package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IpSecurityProfileable 
type IpSecurityProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivityGroupNames()([]string)
    GetAddress()(*string)
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCountHits()(*int32)
    GetCountHosts()(*int32)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIpCategories()([]IpCategoryable)
    GetIpReferenceData()([]IpReferenceDataable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRiskScore()(*string)
    GetTags()([]string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetActivityGroupNames(value []string)()
    SetAddress(value *string)()
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCountHits(value *int32)()
    SetCountHosts(value *int32)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIpCategories(value []IpCategoryable)()
    SetIpReferenceData(value []IpReferenceDataable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRiskScore(value *string)()
    SetTags(value []string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
