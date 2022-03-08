package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSettings provides operations to manage the compliance singleton.
type UserSettings struct {
    Entity
    // The user's settings for the visibility of merge suggestion for the duplicate contacts in the user's contact list.
    contactMergeSuggestions ContactMergeSuggestionsable;
    // Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
    contributionToContentDiscoveryAsOrganizationDisabled *bool;
    // When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
    contributionToContentDiscoveryDisabled *bool;
    // The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
    itemInsights UserInsightsSettingsable;
    // The user's preferences for languages, regional locale and date/time formatting.
    regionalAndLanguageSettings RegionalAndLanguageSettingsable;
    // The shift preferences for the user.
    shiftPreferences ShiftPreferencesable;
}
// NewUserSettings instantiates a new userSettings and sets the default values.
func NewUserSettings()(*UserSettings) {
    m := &UserSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserSettings(), nil
}
// GetContactMergeSuggestions gets the contactMergeSuggestions property value. The user's settings for the visibility of merge suggestion for the duplicate contacts in the user's contact list.
func (m *UserSettings) GetContactMergeSuggestions()(ContactMergeSuggestionsable) {
    if m == nil {
        return nil
    } else {
        return m.contactMergeSuggestions
    }
}
// GetContributionToContentDiscoveryAsOrganizationDisabled gets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
func (m *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contributionToContentDiscoveryAsOrganizationDisabled
    }
}
// GetContributionToContentDiscoveryDisabled gets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
func (m *UserSettings) GetContributionToContentDiscoveryDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contributionToContentDiscoveryDisabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contactMergeSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateContactMergeSuggestionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactMergeSuggestions(val.(ContactMergeSuggestionsable))
        }
        return nil
    }
    res["contributionToContentDiscoveryAsOrganizationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryAsOrganizationDisabled(val)
        }
        return nil
    }
    res["contributionToContentDiscoveryDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryDisabled(val)
        }
        return nil
    }
    res["itemInsights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(UserInsightsSettingsable))
        }
        return nil
    }
    res["regionalAndLanguageSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRegionalAndLanguageSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionalAndLanguageSettings(val.(RegionalAndLanguageSettingsable))
        }
        return nil
    }
    res["shiftPreferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateShiftPreferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShiftPreferences(val.(ShiftPreferencesable))
        }
        return nil
    }
    return res
}
// GetItemInsights gets the itemInsights property value. The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
func (m *UserSettings) GetItemInsights()(UserInsightsSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.itemInsights
    }
}
// GetRegionalAndLanguageSettings gets the regionalAndLanguageSettings property value. The user's preferences for languages, regional locale and date/time formatting.
func (m *UserSettings) GetRegionalAndLanguageSettings()(RegionalAndLanguageSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.regionalAndLanguageSettings
    }
}
// GetShiftPreferences gets the shiftPreferences property value. The shift preferences for the user.
func (m *UserSettings) GetShiftPreferences()(ShiftPreferencesable) {
    if m == nil {
        return nil
    } else {
        return m.shiftPreferences
    }
}
func (m *UserSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contactMergeSuggestions", m.GetContactMergeSuggestions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryAsOrganizationDisabled", m.GetContributionToContentDiscoveryAsOrganizationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryDisabled", m.GetContributionToContentDiscoveryDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("itemInsights", m.GetItemInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("regionalAndLanguageSettings", m.GetRegionalAndLanguageSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shiftPreferences", m.GetShiftPreferences())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContactMergeSuggestions sets the contactMergeSuggestions property value. The user's settings for the visibility of merge suggestion for the duplicate contacts in the user's contact list.
func (m *UserSettings) SetContactMergeSuggestions(value ContactMergeSuggestionsable)() {
    if m != nil {
        m.contactMergeSuggestions = value
    }
}
// SetContributionToContentDiscoveryAsOrganizationDisabled sets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
func (m *UserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)() {
    if m != nil {
        m.contributionToContentDiscoveryAsOrganizationDisabled = value
    }
}
// SetContributionToContentDiscoveryDisabled sets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
func (m *UserSettings) SetContributionToContentDiscoveryDisabled(value *bool)() {
    if m != nil {
        m.contributionToContentDiscoveryDisabled = value
    }
}
// SetItemInsights sets the itemInsights property value. The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
func (m *UserSettings) SetItemInsights(value UserInsightsSettingsable)() {
    if m != nil {
        m.itemInsights = value
    }
}
// SetRegionalAndLanguageSettings sets the regionalAndLanguageSettings property value. The user's preferences for languages, regional locale and date/time formatting.
func (m *UserSettings) SetRegionalAndLanguageSettings(value RegionalAndLanguageSettingsable)() {
    if m != nil {
        m.regionalAndLanguageSettings = value
    }
}
// SetShiftPreferences sets the shiftPreferences property value. The shift preferences for the user.
func (m *UserSettings) SetShiftPreferences(value ShiftPreferencesable)() {
    if m != nil {
        m.shiftPreferences = value
    }
}
