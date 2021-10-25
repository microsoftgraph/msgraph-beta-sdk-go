package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EnrollmentProfile struct {
    Entity
    configurationEndpointUrl *string;
    description *string;
    displayName *string;
    enableAuthenticationViaCompanyPortal *bool;
    requireCompanyPortalOnSetupAssistantEnrolledDevices *bool;
    requiresUserAuthentication *bool;
}
func NewEnrollmentProfile()(*EnrollmentProfile) {
    m := &EnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EnrollmentProfile) GetConfigurationEndpointUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.configurationEndpointUrl
    }
}
func (m *EnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *EnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *EnrollmentProfile) GetEnableAuthenticationViaCompanyPortal()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableAuthenticationViaCompanyPortal
    }
}
func (m *EnrollmentProfile) GetRequireCompanyPortalOnSetupAssistantEnrolledDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireCompanyPortalOnSetupAssistantEnrolledDevices
    }
}
func (m *EnrollmentProfile) GetRequiresUserAuthentication()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requiresUserAuthentication
    }
}
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
func (m *EnrollmentProfile) SetConfigurationEndpointUrl(value *string)() {
    m.configurationEndpointUrl = value
}
func (m *EnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
func (m *EnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *EnrollmentProfile) SetEnableAuthenticationViaCompanyPortal(value *bool)() {
    m.enableAuthenticationViaCompanyPortal = value
}
func (m *EnrollmentProfile) SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(value *bool)() {
    m.requireCompanyPortalOnSetupAssistantEnrolledDevices = value
}
func (m *EnrollmentProfile) SetRequiresUserAuthentication(value *bool)() {
    m.requiresUserAuthentication = value
}
