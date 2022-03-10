package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSettingsable 
type UserSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContactMergeSuggestions()(ContactMergeSuggestionsable)
    GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool)
    GetContributionToContentDiscoveryDisabled()(*bool)
    GetItemInsights()(UserInsightsSettingsable)
    GetRegionalAndLanguageSettings()(RegionalAndLanguageSettingsable)
    GetShiftPreferences()(ShiftPreferencesable)
    SetContactMergeSuggestions(value ContactMergeSuggestionsable)()
    SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)()
    SetContributionToContentDiscoveryDisabled(value *bool)()
    SetItemInsights(value UserInsightsSettingsable)()
    SetRegionalAndLanguageSettings(value RegionalAndLanguageSettingsable)()
    SetShiftPreferences(value ShiftPreferencesable)()
}
