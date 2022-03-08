package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/callrecords"
)

// CallRecordable 
type CallRecordable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetJoinWebUrl()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetModalities()([]if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.Modality)
    GetOrganizer()(IdentitySetable)
    GetParticipants()([]IdentitySetable)
    GetSessions()([]Sessionable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetType()(*if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.CallType)
    GetVersion()(*int64)
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetJoinWebUrl(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetModalities(value []if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.Modality)()
    SetOrganizer(value IdentitySetable)()
    SetParticipants(value []IdentitySetable)()
    SetSessions(value []Sessionable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetType(value *if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.CallType)()
    SetVersion(value *int64)()
}
