package search

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Bookmarkable 
type Bookmarkable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    SearchAnswerable
    GetAvailabilityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAvailabilityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCategories()([]string)
    GetGroupIds()([]string)
    GetIsSuggested()(*bool)
    GetKeywords()(AnswerKeywordable)
    GetLanguageTags()([]string)
    GetPlatforms()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DevicePlatformType)
    GetPowerAppIds()([]string)
    GetState()(*AnswerState)
    GetTargetedVariations()([]AnswerVariantable)
    SetAvailabilityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAvailabilityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCategories(value []string)()
    SetGroupIds(value []string)()
    SetIsSuggested(value *bool)()
    SetKeywords(value AnswerKeywordable)()
    SetLanguageTags(value []string)()
    SetPlatforms(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DevicePlatformType)()
    SetPowerAppIds(value []string)()
    SetState(value *AnswerState)()
    SetTargetedVariations(value []AnswerVariantable)()
}
