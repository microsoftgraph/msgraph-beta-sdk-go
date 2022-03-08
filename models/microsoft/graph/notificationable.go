package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Notificationable 
type Notificationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayTimeToLive()(*int32)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupName()(*string)
    GetPayload()(PayloadTypesable)
    GetPriority()(*Priority)
    GetTargetHostName()(*string)
    GetTargetPolicy()(TargetPolicyEndpointsable)
    SetDisplayTimeToLive(value *int32)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupName(value *string)()
    SetPayload(value PayloadTypesable)()
    SetPriority(value *Priority)()
    SetTargetHostName(value *string)()
    SetTargetPolicy(value TargetPolicyEndpointsable)()
}
