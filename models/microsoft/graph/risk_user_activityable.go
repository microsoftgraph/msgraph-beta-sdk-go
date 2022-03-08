package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskUserActivityable 
type RiskUserActivityable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDetail()(*RiskDetail)
    GetEventTypes()([]RiskEventType)
    GetRiskEventTypes()([]string)
    SetDetail(value *RiskDetail)()
    SetEventTypes(value []RiskEventType)()
    SetRiskEventTypes(value []string)()
}
