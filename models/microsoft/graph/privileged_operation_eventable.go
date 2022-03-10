package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedOperationEventable 
type PrivilegedOperationEventable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalInformation()(*string)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferenceKey()(*string)
    GetReferenceSystem()(*string)
    GetRequestorId()(*string)
    GetRequestorName()(*string)
    GetRequestType()(*string)
    GetRoleId()(*string)
    GetRoleName()(*string)
    GetTenantId()(*string)
    GetUserId()(*string)
    GetUserMail()(*string)
    GetUserName()(*string)
    SetAdditionalInformation(value *string)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReferenceKey(value *string)()
    SetReferenceSystem(value *string)()
    SetRequestorId(value *string)()
    SetRequestorName(value *string)()
    SetRequestType(value *string)()
    SetRoleId(value *string)()
    SetRoleName(value *string)()
    SetTenantId(value *string)()
    SetUserId(value *string)()
    SetUserMail(value *string)()
    SetUserName(value *string)()
}
