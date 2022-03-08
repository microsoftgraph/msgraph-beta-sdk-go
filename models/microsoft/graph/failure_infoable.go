package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/callrecords"
)

// FailureInfoable 
type FailureInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetReason()(*string)
    GetStage()(*if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage)
    SetReason(value *string)()
    SetStage(value *if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage)()
}
