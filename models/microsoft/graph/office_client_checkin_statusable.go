package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeClientCheckinStatusable 
type OfficeClientCheckinStatusable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppliedPolicies()([]string)
    GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceName()(*string)
    GetDevicePlatform()(*string)
    GetDevicePlatformVersion()(*string)
    GetErrorMessage()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetWasSuccessful()(*bool)
    SetAppliedPolicies(value []string)()
    SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceName(value *string)()
    SetDevicePlatform(value *string)()
    SetDevicePlatformVersion(value *string)()
    SetErrorMessage(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetWasSuccessful(value *bool)()
}
