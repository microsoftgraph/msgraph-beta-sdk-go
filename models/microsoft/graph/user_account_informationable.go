package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserAccountInformationable 
type UserAccountInformationable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAgeGroup()(*string)
    GetCountryCode()(*string)
    GetPreferredLanguageTag()(LocaleInfoable)
    GetUserPrincipalName()(*string)
    SetAgeGroup(value *string)()
    SetCountryCode(value *string)()
    SetPreferredLanguageTag(value LocaleInfoable)()
    SetUserPrincipalName(value *string)()
}
