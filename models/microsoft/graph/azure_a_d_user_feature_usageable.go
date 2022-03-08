package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AzureADUserFeatureUsageable 
type AzureADUserFeatureUsageable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFeatureUsageDetails()([]FeatureUsageDetailable)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLicenseAssigned()(*AzureADLicenseType)
    GetLicenseRecommended()(*AzureADLicenseType)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetFeatureUsageDetails(value []FeatureUsageDetailable)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLicenseAssigned(value *AzureADLicenseType)()
    SetLicenseRecommended(value *AzureADLicenseType)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
