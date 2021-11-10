package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementDerivedCredentialSettings struct {
    Entity
    // The display name for the profile.
    displayName *string;
    // The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
    helpUrl *string;
    // The derived credential provider to use.
    issuer *DeviceManagementDerivedCredentialIssuer;
    // The methods used to inform the end user to open Company Portal to deliver Wi-Fi, VPN, or email profiles that use certificates to the device.
    notificationType *DeviceManagementDerivedCredentialNotificationType;
}
// Instantiates a new deviceManagementDerivedCredentialSettings and sets the default values.
func NewDeviceManagementDerivedCredentialSettings()(*DeviceManagementDerivedCredentialSettings) {
    m := &DeviceManagementDerivedCredentialSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name for the profile.
func (m *DeviceManagementDerivedCredentialSettings) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the helpUrl property value. The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
func (m *DeviceManagementDerivedCredentialSettings) GetHelpUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpUrl
    }
}
// Gets the issuer property value. The derived credential provider to use.
func (m *DeviceManagementDerivedCredentialSettings) GetIssuer()(*DeviceManagementDerivedCredentialIssuer) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
// Gets the notificationType property value. The methods used to inform the end user to open Company Portal to deliver Wi-Fi, VPN, or email profiles that use certificates to the device.
func (m *DeviceManagementDerivedCredentialSettings) GetNotificationType()(*DeviceManagementDerivedCredentialNotificationType) {
    if m == nil {
        return nil
    } else {
        return m.notificationType
    }
}
// The deserialization information for the current model
func (m *DeviceManagementDerivedCredentialSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["helpUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpUrl(val)
        }
        return nil
    }
    res["issuer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementDerivedCredentialIssuer)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementDerivedCredentialIssuer)
            m.SetIssuer(&cast)
        }
        return nil
    }
    res["notificationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementDerivedCredentialNotificationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementDerivedCredentialNotificationType)
            m.SetNotificationType(&cast)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementDerivedCredentialSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementDerivedCredentialSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("helpUrl", m.GetHelpUrl())
        if err != nil {
            return err
        }
    }
    if m.GetIssuer() != nil {
        cast := m.GetIssuer().String()
        err = writer.WriteStringValue("issuer", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationType() != nil {
        cast := m.GetNotificationType().String()
        err = writer.WriteStringValue("notificationType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The display name for the profile.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementDerivedCredentialSettings) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the helpUrl property value. The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
// Parameters:
//  - value : Value to set for the helpUrl property.
func (m *DeviceManagementDerivedCredentialSettings) SetHelpUrl(value *string)() {
    m.helpUrl = value
}
// Sets the issuer property value. The derived credential provider to use.
// Parameters:
//  - value : Value to set for the issuer property.
func (m *DeviceManagementDerivedCredentialSettings) SetIssuer(value *DeviceManagementDerivedCredentialIssuer)() {
    m.issuer = value
}
// Sets the notificationType property value. The methods used to inform the end user to open Company Portal to deliver Wi-Fi, VPN, or email profiles that use certificates to the device.
// Parameters:
//  - value : Value to set for the notificationType property.
func (m *DeviceManagementDerivedCredentialSettings) SetNotificationType(value *DeviceManagementDerivedCredentialNotificationType)() {
    m.notificationType = value
}
