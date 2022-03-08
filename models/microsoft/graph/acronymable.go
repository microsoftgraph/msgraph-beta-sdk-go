package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/search"
)

// Acronymable 
type Acronymable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    SearchAnswerable
    GetStandsFor()(*string)
    GetState()(*id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)
    SetStandsFor(value *string)()
    SetState(value *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)()
}
