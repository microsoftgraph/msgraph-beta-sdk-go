package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationalActivityDetailable 
type EducationalActivityDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAbbreviation()(*string)
    GetActivities()([]string)
    GetAwards()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFieldsOfStudy()([]string)
    GetGrade()(*string)
    GetNotes()(*string)
    GetWebUrl()(*string)
    SetAbbreviation(value *string)()
    SetActivities(value []string)()
    SetAwards(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFieldsOfStudy(value []string)()
    SetGrade(value *string)()
    SetNotes(value *string)()
    SetWebUrl(value *string)()
}
