package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FileSecurityProfileable 
type FileSecurityProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivityGroupNames()([]string)
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCertificateThumbprint()(*string)
    GetExtensions()([]string)
    GetFileType()(*string)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHashes()([]FileHashable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareStates()([]MalwareStateable)
    GetNames()([]string)
    GetRiskScore()(*string)
    GetSize()(*int64)
    GetTags()([]string)
    GetVendorInformation()(SecurityVendorInformationable)
    GetVulnerabilityStates()([]VulnerabilityStateable)
    SetActivityGroupNames(value []string)()
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCertificateThumbprint(value *string)()
    SetExtensions(value []string)()
    SetFileType(value *string)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHashes(value []FileHashable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareStates(value []MalwareStateable)()
    SetNames(value []string)()
    SetRiskScore(value *string)()
    SetSize(value *int64)()
    SetTags(value []string)()
    SetVendorInformation(value SecurityVendorInformationable)()
    SetVulnerabilityStates(value []VulnerabilityStateable)()
}
