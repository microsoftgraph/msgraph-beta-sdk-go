package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceRequestable 
type AccessPackageResourceRequestable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageResource()(AccessPackageResourceable)
    GetCatalogId()(*string)
    GetExecuteImmediately()(*bool)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsValidationOnly()(*bool)
    GetJustification()(*string)
    GetRequestor()(AccessPackageSubjectable)
    GetRequestState()(*string)
    GetRequestStatus()(*string)
    GetRequestType()(*string)
    SetAccessPackageResource(value AccessPackageResourceable)()
    SetCatalogId(value *string)()
    SetExecuteImmediately(value *bool)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsValidationOnly(value *bool)()
    SetJustification(value *string)()
    SetRequestor(value AccessPackageSubjectable)()
    SetRequestState(value *string)()
    SetRequestStatus(value *string)()
    SetRequestType(value *string)()
}
