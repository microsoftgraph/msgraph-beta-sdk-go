package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VppTokenable 
type VppTokenable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppleId()(*string)
    GetAutomaticallyUpdateApps()(*bool)
    GetClaimTokenManagementFromExternalMdm()(*bool)
    GetCountryOrRegion()(*string)
    GetDataSharingConsentGranted()(*bool)
    GetDisplayName()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncStatus()(*VppTokenSyncStatus)
    GetLocationName()(*string)
    GetOrganizationName()(*string)
    GetRoleScopeTagIds()([]string)
    GetState()(*VppTokenState)
    GetToken()(*string)
    GetTokenActionResults()([]VppTokenActionResultable)
    GetVppTokenAccountType()(*VppTokenAccountType)
    SetAppleId(value *string)()
    SetAutomaticallyUpdateApps(value *bool)()
    SetClaimTokenManagementFromExternalMdm(value *bool)()
    SetCountryOrRegion(value *string)()
    SetDataSharingConsentGranted(value *bool)()
    SetDisplayName(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncStatus(value *VppTokenSyncStatus)()
    SetLocationName(value *string)()
    SetOrganizationName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetState(value *VppTokenState)()
    SetToken(value *string)()
    SetTokenActionResults(value []VppTokenActionResultable)()
    SetVppTokenAccountType(value *VppTokenAccountType)()
}
