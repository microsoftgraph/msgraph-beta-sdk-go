package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthenticationDetailable 
type AuthenticationDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthenticationMethod()(*string)
    GetAuthenticationMethodDetail()(*string)
    GetAuthenticationStepDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAuthenticationStepRequirement()(*string)
    GetAuthenticationStepResultDetail()(*string)
    GetSucceeded()(*bool)
    SetAuthenticationMethod(value *string)()
    SetAuthenticationMethodDetail(value *string)()
    SetAuthenticationStepDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAuthenticationStepRequirement(value *string)()
    SetAuthenticationStepResultDetail(value *string)()
    SetSucceeded(value *bool)()
}
