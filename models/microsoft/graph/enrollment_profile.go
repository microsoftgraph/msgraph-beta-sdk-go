package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new enrollmentProfile and sets the default values.
func NewEnrollmentProfile()(*EnrollmentProfile) {
    m := &EnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the configurationEndpointUrl property value. Configuration endpoint url to use for Enrollment
func (m *EnrollmentProfile) GetConfigurationEndpointUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.configurationEndpointUrl
    }
}
// Gets the description property value. Description of the profile
func (m *EnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Name of the profile
func (m *EnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enableAuthenticationViaCompanyPortal property value. Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
func (m *EnrollmentProfile) GetEnableAuthenticationViaCompanyPortal()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableAuthenticationViaCompanyPortal
    }
}
// Gets the requireCompanyPortalOnSetupAssistantEnrolledDevices property value. Indicates that Company Portal is required on setup assistant enrolled devices
func (m *EnrollmentProfile) GetRequireCompanyPortalOnSetupAssistantEnrolledDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireCompanyPortalOnSetupAssistantEnrolledDevices
    }
}
// Gets the requiresUserAuthentication property value. Indicates if the profile requires user authentication
func (m *EnrollmentProfile) GetRequiresUserAuthentication()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requiresUserAuthentication
    }
}
// The deserialization information for the current model
func (m *EnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationEndpointUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConfigurationEndpointUrl(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["enableAuthenticationViaCompanyPortal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableAuthenticationViaCompanyPortal(val)
        return nil
    }
    res["requireCompanyPortalOnSetupAssistantEnrolledDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(val)
        return nil
    }
    res["requiresUserAuthentication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequiresUserAuthentication(val)
        return nil
    }
    return res
}
func (m *EnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the configurationEndpointUrl property value. Configuration endpoint url to use for Enrollment
// Parameters:
//  - value : Value to set for the configurationEndpointUrl property.
func (m *EnrollmentProfile) SetConfigurationEndpointUrl(value *string)() {
    m.configurationEndpointUrl = value
}
// Sets the description property value. Description of the profile
// Parameters:
//  - value : Value to set for the description property.
func (m *EnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Name of the profile
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enableAuthenticationViaCompanyPortal property value. Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
// Parameters:
//  - value : Value to set for the enableAuthenticationViaCompanyPortal property.
func (m *EnrollmentProfile) SetEnableAuthenticationViaCompanyPortal(value *bool)() {
    m.enableAuthenticationViaCompanyPortal = value
}
// Sets the requireCompanyPortalOnSetupAssistantEnrolledDevices property value. Indicates that Company Portal is required on setup assistant enrolled devices
// Parameters:
//  - value : Value to set for the requireCompanyPortalOnSetupAssistantEnrolledDevices property.
func (m *EnrollmentProfile) SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(value *bool)() {
    m.requireCompanyPortalOnSetupAssistantEnrolledDevices = value
}
// Sets the requiresUserAuthentication property value. Indicates if the profile requires user authentication
// Parameters:
//  - value : Value to set for the requiresUserAuthentication property.
func (m *EnrollmentProfile) SetRequiresUserAuthentication(value *bool)() {
    m.requiresUserAuthentication = value
}
