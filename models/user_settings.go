package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSettings 
type UserSettings struct {
    Entity
}
// NewUserSettings instantiates a new UserSettings and sets the default values.
func NewUserSettings()(*UserSettings) {
    m := &UserSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSettings(), nil
}
// GetContactMergeSuggestions gets the contactMergeSuggestions property value. The user's settings for the visibility of merge suggestion for the duplicate contacts in the user's contact list.
func (m *UserSettings) GetContactMergeSuggestions()(ContactMergeSuggestionsable) {
    val, err := m.GetBackingStore().Get("contactMergeSuggestions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ContactMergeSuggestionsable)
    }
    return nil
}
// GetContributionToContentDiscoveryAsOrganizationDisabled gets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
func (m *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("contributionToContentDiscoveryAsOrganizationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContributionToContentDiscoveryDisabled gets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
func (m *UserSettings) GetContributionToContentDiscoveryDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("contributionToContentDiscoveryDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contactMergeSuggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContactMergeSuggestionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactMergeSuggestions(val.(ContactMergeSuggestionsable))
        }
        return nil
    }
    res["contributionToContentDiscoveryAsOrganizationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryAsOrganizationDisabled(val)
        }
        return nil
    }
    res["contributionToContentDiscoveryDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryDisabled(val)
        }
        return nil
    }
    res["itemInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(UserInsightsSettingsable))
        }
        return nil
    }
    res["regionalAndLanguageSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRegionalAndLanguageSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionalAndLanguageSettings(val.(RegionalAndLanguageSettingsable))
        }
        return nil
    }
    res["shiftPreferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    val, err := m.GetBackingStore().Get("itemInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserInsightsSettingsable)
    }
    return nil
}
// GetRegionalAndLanguageSettings gets the regionalAndLanguageSettings property value. The user's preferences for languages, regional locale and date/time formatting.
func (m *UserSettings) GetRegionalAndLanguageSettings()(RegionalAndLanguageSettingsable) {
    val, err := m.GetBackingStore().Get("regionalAndLanguageSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RegionalAndLanguageSettingsable)
    }
    return nil
}
// GetShiftPreferences gets the shiftPreferences property value. The shift preferences for the user.
func (m *UserSettings) GetShiftPreferences()(ShiftPreferencesable) {
    val, err := m.GetBackingStore().Get("shiftPreferences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ShiftPreferencesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("contactMergeSuggestions", value)
    if err != nil {
        panic(err)
    }
}
// SetContributionToContentDiscoveryAsOrganizationDisabled sets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
func (m *UserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("contributionToContentDiscoveryAsOrganizationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetContributionToContentDiscoveryDisabled sets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
func (m *UserSettings) SetContributionToContentDiscoveryDisabled(value *bool)() {
    err := m.GetBackingStore().Set("contributionToContentDiscoveryDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetItemInsights sets the itemInsights property value. The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
func (m *UserSettings) SetItemInsights(value UserInsightsSettingsable)() {
    err := m.GetBackingStore().Set("itemInsights", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionalAndLanguageSettings sets the regionalAndLanguageSettings property value. The user's preferences for languages, regional locale and date/time formatting.
func (m *UserSettings) SetRegionalAndLanguageSettings(value RegionalAndLanguageSettingsable)() {
    err := m.GetBackingStore().Set("regionalAndLanguageSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetShiftPreferences sets the shiftPreferences property value. The shift preferences for the user.
func (m *UserSettings) SetShiftPreferences(value ShiftPreferencesable)() {
    err := m.GetBackingStore().Set("shiftPreferences", value)
    if err != nil {
        panic(err)
    }
}
// UserSettingsable 
type UserSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
