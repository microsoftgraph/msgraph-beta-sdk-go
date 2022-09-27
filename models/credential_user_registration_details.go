package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUserRegistrationDetails 
type CredentialUserRegistrationDetails struct {
    Entity
    // Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
    authMethods []RegistrationAuthMethod
    // Indicates whether the user is ready to perform self-service password reset or MFA.
    isCapable *bool
    // Indicates whether the user enabled to perform self-service password reset.
    isEnabled *bool
    // Indicates whether the user is registered for MFA.
    isMfaRegistered *bool
    // Indicates whether the user has registered any authentication methods for self-service password reset.
    isRegistered *bool
    // Provides the user name of the corresponding user.
    userDisplayName *string
    // Provides the user principal name of the corresponding user.
    userPrincipalName *string
}
// NewCredentialUserRegistrationDetails instantiates a new CredentialUserRegistrationDetails and sets the default values.
func NewCredentialUserRegistrationDetails()(*CredentialUserRegistrationDetails) {
    m := &CredentialUserRegistrationDetails{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.credentialUserRegistrationDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCredentialUserRegistrationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUserRegistrationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUserRegistrationDetails(), nil
}
// GetAuthMethods gets the authMethods property value. Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
func (m *CredentialUserRegistrationDetails) GetAuthMethods()([]RegistrationAuthMethod) {
    return m.authMethods
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethods"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseRegistrationAuthMethod , m.SetAuthMethods)
    res["isCapable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCapable)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["isMfaRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsMfaRegistered)
    res["isRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRegistered)
    res["userDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserDisplayName)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetIsCapable gets the isCapable property value. Indicates whether the user is ready to perform self-service password reset or MFA.
func (m *CredentialUserRegistrationDetails) GetIsCapable()(*bool) {
    return m.isCapable
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetIsMfaRegistered gets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
func (m *CredentialUserRegistrationDetails) GetIsMfaRegistered()(*bool) {
    return m.isMfaRegistered
}
// GetIsRegistered gets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsRegistered()(*bool) {
    return m.isRegistered
}
// GetUserDisplayName gets the userDisplayName property value. Provides the user name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserPrincipalName gets the userPrincipalName property value. Provides the user principal name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *CredentialUserRegistrationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthMethods() != nil {
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
// SetAuthMethods sets the authMethods property value. Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
func (m *CredentialUserRegistrationDetails) SetAuthMethods(value []RegistrationAuthMethod)() {
    m.authMethods = value
}
// SetIsCapable sets the isCapable property value. Indicates whether the user is ready to perform self-service password reset or MFA.
func (m *CredentialUserRegistrationDetails) SetIsCapable(value *bool)() {
    m.isCapable = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
func (m *CredentialUserRegistrationDetails) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetIsMfaRegistered sets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
func (m *CredentialUserRegistrationDetails) SetIsMfaRegistered(value *bool)() {
    m.isMfaRegistered = value
}
// SetIsRegistered sets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
func (m *CredentialUserRegistrationDetails) SetIsRegistered(value *bool)() {
    m.isRegistered = value
}
// SetUserDisplayName sets the userDisplayName property value. Provides the user name of the corresponding user.
func (m *CredentialUserRegistrationDetails) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. Provides the user principal name of the corresponding user.
func (m *CredentialUserRegistrationDetails) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
