package search

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AnswerKeywordable 
type AnswerKeywordable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetKeywords()([]string)
    GetMatchSimilarKeywords()(*bool)
    GetReservedKeywords()([]string)
    SetKeywords(value []string)()
    SetMatchSimilarKeywords(value *bool)()
    SetReservedKeywords(value []string)()
}
