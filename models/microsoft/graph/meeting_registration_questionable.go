package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingRegistrationQuestionable 
type MeetingRegistrationQuestionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAnswerInputType()(*AnswerInputType)
    GetAnswerOptions()([]string)
    GetDisplayName()(*string)
    GetIsRequired()(*bool)
    SetAnswerInputType(value *AnswerInputType)()
    SetAnswerOptions(value []string)()
    SetDisplayName(value *string)()
    SetIsRequired(value *bool)()
}
