package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecommendationResourceable 
type RecommendationResourceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAdditionalDetails()([]KeyValueable)
    GetApiUrl()(*string)
    GetDisplayName()(*string)
    GetOwner()(*string)
    GetPortalUrl()(*string)
    GetRank()(*int32)
    GetRecommendationId()(*string)
    GetResourceType()(*string)
    GetStatus()(*RecommendationStatus)
    SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAdditionalDetails(value []KeyValueable)()
    SetApiUrl(value *string)()
    SetDisplayName(value *string)()
    SetOwner(value *string)()
    SetPortalUrl(value *string)()
    SetRank(value *int32)()
    SetRecommendationId(value *string)()
    SetResourceType(value *string)()
    SetStatus(value *RecommendationStatus)()
}
