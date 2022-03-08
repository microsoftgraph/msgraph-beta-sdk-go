package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecommendedActionable 
type RecommendedActionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActionWebUrl()(*string)
    GetPotentialScoreImpact()(*float64)
    GetTitle()(*string)
    SetActionWebUrl(value *string)()
    SetPotentialScoreImpact(value *float64)()
    SetTitle(value *string)()
}
