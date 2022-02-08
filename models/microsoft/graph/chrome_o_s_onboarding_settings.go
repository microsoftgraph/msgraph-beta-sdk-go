package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChromeOSOnboardingSettings 
type ChromeOSOnboardingSettings struct {
    Entity
    // The ChromebookTenant's LastDirectorySyncDateTime
    lastDirectorySyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The ChromebookTenant's LastModifiedDateTime
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The ChromebookTenant's OnboardingStatus. Possible values are: unknown, inprogress, onboarded, failed, offboarding, unknownFutureValue.
    onboardingStatus *OnboardingStatus;
    // The ChromebookTenant's OwnerUserPrincipalName
    ownerUserPrincipalName *string;
}
// NewChromeOSOnboardingSettings instantiates a new chromeOSOnboardingSettings and sets the default values.
func NewChromeOSOnboardingSettings()(*ChromeOSOnboardingSettings) {
    m := &ChromeOSOnboardingSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetLastDirectorySyncDateTime gets the lastDirectorySyncDateTime property value. The ChromebookTenant's LastDirectorySyncDateTime
func (m *ChromeOSOnboardingSettings) GetLastDirectorySyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDirectorySyncDateTime
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The ChromebookTenant's LastModifiedDateTime
func (m *ChromeOSOnboardingSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetOnboardingStatus gets the onboardingStatus property value. The ChromebookTenant's OnboardingStatus. Possible values are: unknown, inprogress, onboarded, failed, offboarding, unknownFutureValue.
func (m *ChromeOSOnboardingSettings) GetOnboardingStatus()(*OnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. The ChromebookTenant's OwnerUserPrincipalName
func (m *ChromeOSOnboardingSettings) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChromeOSOnboardingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastDirectorySyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDirectorySyncDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*OnboardingStatus))
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *ChromeOSOnboardingSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetOnboardingStatus()).String()
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
// SetLastDirectorySyncDateTime sets the lastDirectorySyncDateTime property value. The ChromebookTenant's LastDirectorySyncDateTime
func (m *ChromeOSOnboardingSettings) SetLastDirectorySyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastDirectorySyncDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The ChromebookTenant's LastModifiedDateTime
func (m *ChromeOSOnboardingSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetOnboardingStatus sets the onboardingStatus property value. The ChromebookTenant's OnboardingStatus. Possible values are: unknown, inprogress, onboarded, failed, offboarding, unknownFutureValue.
func (m *ChromeOSOnboardingSettings) SetOnboardingStatus(value *OnboardingStatus)() {
    if m != nil {
        m.onboardingStatus = value
    }
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. The ChromebookTenant's OwnerUserPrincipalName
func (m *ChromeOSOnboardingSettings) SetOwnerUserPrincipalName(value *string)() {
    if m != nil {
        m.ownerUserPrincipalName = value
    }
}
