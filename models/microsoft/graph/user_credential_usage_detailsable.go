package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserCredentialUsageDetailsable 
type UserCredentialUsageDetailsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthMethod()(*UsageAuthMethod)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailureReason()(*string)
    GetFeature()(*FeatureType)
    GetIsSuccess()(*bool)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetAuthMethod(value *UsageAuthMethod)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailureReason(value *string)()
    SetFeature(value *FeatureType)()
    SetIsSuccess(value *bool)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
