package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EnrollmentProfile 
type EnrollmentProfile struct {
    Entity
    // Configuration endpoint url to use for Enrollment
    configurationEndpointUrl *string;
    // Description of the profile
    description *string;
    // Name of the profile
    displayName *string;
    // Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
    enableAuthenticationViaCompanyPortal *bool;
    // Indicates that Company Portal is required on setup assistant enrolled devices
    requireCompanyPortalOnSetupAssistantEnrolledDevices *bool;
    // Indicates if the profile requires user authentication
    requiresUserAuthentication *bool;
}
// NewEnrollmentProfile instantiates a new enrollmentProfile and sets the default values.
func NewEnrollmentProfile()(*EnrollmentProfile) {
    m := &EnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetConfigurationEndpointUrl gets the configurationEndpointUrl property value. Configuration endpoint url to use for Enrollment
func (m *EnrollmentProfile) GetConfigurationEndpointUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.configurationEndpointUrl
    }
}
// GetDescription gets the description property value. Description of the profile
func (m *EnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Name of the profile
func (m *EnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnableAuthenticationViaCompanyPortal gets the enableAuthenticationViaCompanyPortal property value. Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
func (m *EnrollmentProfile) GetEnableAuthenticationViaCompanyPortal()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableAuthenticationViaCompanyPortal
    }
}
// GetRequireCompanyPortalOnSetupAssistantEnrolledDevices gets the requireCompanyPortalOnSetupAssistantEnrolledDevices property value. Indicates that Company Portal is required on setup assistant enrolled devices
func (m *EnrollmentProfile) GetRequireCompanyPortalOnSetupAssistantEnrolledDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireCompanyPortalOnSetupAssistantEnrolledDevices
    }
}
// GetRequiresUserAuthentication gets the requiresUserAuthentication property value. Indicates if the profile requires user authentication
func (m *EnrollmentProfile) GetRequiresUserAuthentication()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requiresUserAuthentication
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationEndpointUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationEndpointUrl(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enableAuthenticationViaCompanyPortal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAuthenticationViaCompanyPortal(val)
        }
        return nil
    }
    res["requireCompanyPortalOnSetupAssistantEnrolledDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(val)
        }
        return nil
    }
    res["requiresUserAuthentication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresUserAuthentication(val)
        }
        return nil
    }
    return res
}
func (m *EnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EnrollmentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("configurationEndpointUrl", m.GetConfigurationEndpointUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableAuthenticationViaCompanyPortal", m.GetEnableAuthenticationViaCompanyPortal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireCompanyPortalOnSetupAssistantEnrolledDevices", m.GetRequireCompanyPortalOnSetupAssistantEnrolledDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresUserAuthentication", m.GetRequiresUserAuthentication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationEndpointUrl sets the configurationEndpointUrl property value. Configuration endpoint url to use for Enrollment
func (m *EnrollmentProfile) SetConfigurationEndpointUrl(value *string)() {
    if m != nil {
        m.configurationEndpointUrl = value
    }
}
// SetDescription sets the description property value. Description of the profile
func (m *EnrollmentProfile) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Name of the profile
func (m *EnrollmentProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnableAuthenticationViaCompanyPortal sets the enableAuthenticationViaCompanyPortal property value. Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
func (m *EnrollmentProfile) SetEnableAuthenticationViaCompanyPortal(value *bool)() {
    if m != nil {
        m.enableAuthenticationViaCompanyPortal = value
    }
}
// SetRequireCompanyPortalOnSetupAssistantEnrolledDevices sets the requireCompanyPortalOnSetupAssistantEnrolledDevices property value. Indicates that Company Portal is required on setup assistant enrolled devices
func (m *EnrollmentProfile) SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(value *bool)() {
    if m != nil {
        m.requireCompanyPortalOnSetupAssistantEnrolledDevices = value
    }
}
// SetRequiresUserAuthentication sets the requiresUserAuthentication property value. Indicates if the profile requires user authentication
func (m *EnrollmentProfile) SetRequiresUserAuthentication(value *bool)() {
    if m != nil {
        m.requiresUserAuthentication = value
    }
}
