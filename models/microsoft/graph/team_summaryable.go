package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamSummaryable 
type TeamSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetGuestsCount()(*int32)
    GetMembersCount()(*int32)
    GetOwnersCount()(*int32)
    SetGuestsCount(value *int32)()
    SetMembersCount(value *int32)()
    SetOwnersCount(value *int32)()
}
