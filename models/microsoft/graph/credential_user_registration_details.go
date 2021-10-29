package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CredentialUserRegistrationDetails struct {
    Entity
    // Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
    authMethods []RegistrationAuthMethod;
    // Indicates whether the user is ready to perform self-service password reset or MFA.
    isCapable *bool;
    // Indicates whether the user enabled to perform self-service password reset.
    isEnabled *bool;
    // Indicates whether the user is registered for MFA.
    isMfaRegistered *bool;
    // Indicates whether the user has registered any authentication methods for self-service password reset.
    isRegistered *bool;
    // Provides the user name of the corresponding user.
    userDisplayName *string;
    // Provides the user principal name of the corresponding user.
    userPrincipalName *string;
}
// Instantiates a new credentialUserRegistrationDetails and sets the default values.
func NewCredentialUserRegistrationDetails()(*CredentialUserRegistrationDetails) {
    m := &CredentialUserRegistrationDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the authMethods property value. Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
func (m *CredentialUserRegistrationDetails) GetAuthMethods()([]RegistrationAuthMethod) {
    if m == nil {
        return nil
    } else {
        return m.authMethods
    }
}
// Gets the isCapable property value. Indicates whether the user is ready to perform self-service password reset or MFA.
func (m *CredentialUserRegistrationDetails) GetIsCapable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCapable
    }
}
// Gets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
func (m *CredentialUserRegistrationDetails) GetIsMfaRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaRegistered
    }
}
// Gets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRegistered
    }
}
// Gets the userDisplayName property value. Provides the user name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userPrincipalName property value. Provides the user principal name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *CredentialUserRegistrationDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRegistrationAuthMethod)
        if err != nil {
            return err
        }
        res := make([]RegistrationAuthMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*RegistrationAuthMethod))
        }
        m.SetAuthMethods(res)
        return nil
    }
    res["isCapable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCapable(val)
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["isMfaRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMfaRegistered(val)
        return nil
    }
    res["isRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRegistered(val)
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserDisplayName(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *CredentialUserRegistrationDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CredentialUserRegistrationDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("authMethods", SerializeRegistrationAuthMethod(m.GetAuthMethods()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCapable", m.GetIsCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaRegistered", m.GetIsMfaRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRegistered", m.GetIsRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the authMethods property value. Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
// Parameters:
//  - value : Value to set for the authMethods property.
func (m *CredentialUserRegistrationDetails) SetAuthMethods(value []RegistrationAuthMethod)() {
    m.authMethods = value
}
// Sets the isCapable property value. Indicates whether the user is ready to perform self-service password reset or MFA.
// Parameters:
//  - value : Value to set for the isCapable property.
func (m *CredentialUserRegistrationDetails) SetIsCapable(value *bool)() {
    m.isCapable = value
}
// Sets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *CredentialUserRegistrationDetails) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
// Parameters:
//  - value : Value to set for the isMfaRegistered property.
func (m *CredentialUserRegistrationDetails) SetIsMfaRegistered(value *bool)() {
    m.isMfaRegistered = value
}
// Sets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
// Parameters:
//  - value : Value to set for the isRegistered property.
func (m *CredentialUserRegistrationDetails) SetIsRegistered(value *bool)() {
    m.isRegistered = value
}
// Sets the userDisplayName property value. Provides the user name of the corresponding user.
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *CredentialUserRegistrationDetails) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userPrincipalName property value. Provides the user principal name of the corresponding user.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *CredentialUserRegistrationDetails) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
