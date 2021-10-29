package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChromeOSOnboardingSettings struct {
    Entity
    // The ChromebookTenant's LastDirectorySyncDateTime
    lastDirectorySyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The ChromebookTenant's LastModifiedDateTime
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The ChromebookTenant's OnboardingStatus. Possible values are: unknown, inprogress, onboarded, failed.
    onboardingStatus *OnboardingStatus;
    // The ChromebookTenant's OwnerUserPrincipalName
    ownerUserPrincipalName *string;
}
// Instantiates a new chromeOSOnboardingSettings and sets the default values.
func NewChromeOSOnboardingSettings()(*ChromeOSOnboardingSettings) {
    m := &ChromeOSOnboardingSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the lastDirectorySyncDateTime property value. The ChromebookTenant's LastDirectorySyncDateTime
func (m *ChromeOSOnboardingSettings) GetLastDirectorySyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDirectorySyncDateTime
    }
}
// Gets the lastModifiedDateTime property value. The ChromebookTenant's LastModifiedDateTime
func (m *ChromeOSOnboardingSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the onboardingStatus property value. The ChromebookTenant's OnboardingStatus. Possible values are: unknown, inprogress, onboarded, failed.
func (m *ChromeOSOnboardingSettings) GetOnboardingStatus()(*OnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// Gets the ownerUserPrincipalName property value. The ChromebookTenant's OwnerUserPrincipalName
func (m *ChromeOSOnboardingSettings) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
// The deserialization information for the current model
func (m *ChromeOSOnboardingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastDirectorySyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastDirectorySyncDateTime(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["onboardingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnboardingStatus)
        if err != nil {
            return err
        }
        cast := val.(OnboardingStatus)
        m.SetOnboardingStatus(&cast)
        return nil
    }
    res["ownerUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *ChromeOSOnboardingSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ChromeOSOnboardingSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastDirectorySyncDateTime", m.GetLastDirectorySyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := m.GetOnboardingStatus().String()
        err = writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the lastDirectorySyncDateTime property value. The ChromebookTenant's LastDirectorySyncDateTime
// Parameters:
//  - value : Value to set for the lastDirectorySyncDateTime property.
func (m *ChromeOSOnboardingSettings) SetLastDirectorySyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDirectorySyncDateTime = value
}
// Sets the lastModifiedDateTime property value. The ChromebookTenant's LastModifiedDateTime
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *ChromeOSOnboardingSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the onboardingStatus property value. The ChromebookTenant's OnboardingStatus. Possible values are: unknown, inprogress, onboarded, failed.
// Parameters:
//  - value : Value to set for the onboardingStatus property.
func (m *ChromeOSOnboardingSettings) SetOnboardingStatus(value *OnboardingStatus)() {
    m.onboardingStatus = value
}
// Sets the ownerUserPrincipalName property value. The ChromebookTenant's OwnerUserPrincipalName
// Parameters:
//  - value : Value to set for the ownerUserPrincipalName property.
func (m *ChromeOSOnboardingSettings) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
