package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RemoteActionAuditable 
type RemoteActionAuditable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAction()(*RemoteAction)
    GetActionState()(*ActionState)
    GetDeviceDisplayName()(*string)
    GetDeviceIMEI()(*string)
    GetDeviceOwnerUserPrincipalName()(*string)
    GetInitiatedByUserPrincipalName()(*string)
    GetManagedDeviceId()(*string)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUserName()(*string)
    SetAction(value *RemoteAction)()
    SetActionState(value *ActionState)()
    SetDeviceDisplayName(value *string)()
    SetDeviceIMEI(value *string)()
    SetDeviceOwnerUserPrincipalName(value *string)()
    SetInitiatedByUserPrincipalName(value *string)()
    SetManagedDeviceId(value *string)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUserName(value *string)()
}
