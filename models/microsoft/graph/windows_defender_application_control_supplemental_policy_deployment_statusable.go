package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable 
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyStatuses)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetPolicy()(WindowsDefenderApplicationControlSupplementalPolicyable)
    GetPolicyVersion()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetDeploymentStatus(value *WindowsDefenderApplicationControlSupplementalPolicyStatuses)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetPolicy(value WindowsDefenderApplicationControlSupplementalPolicyable)()
    SetPolicyVersion(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
