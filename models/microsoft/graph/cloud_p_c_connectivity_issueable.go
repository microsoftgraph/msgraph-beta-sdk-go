package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPCConnectivityIssueable 
type CloudPCConnectivityIssueable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceId()(*string)
    GetErrorCode()(*string)
    GetErrorDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetErrorDescription()(*string)
    GetRecommendedAction()(*string)
    GetUserId()(*string)
    SetDeviceId(value *string)()
    SetErrorCode(value *string)()
    SetErrorDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetErrorDescription(value *string)()
    SetRecommendedAction(value *string)()
    SetUserId(value *string)()
}
