package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUserRegistrationDetails 
type CredentialUserRegistrationDetails struct {
    Entity
}
// NewCredentialUserRegistrationDetails instantiates a new credentialUserRegistrationDetails and sets the default values.
func NewCredentialUserRegistrationDetails()(*CredentialUserRegistrationDetails) {
    m := &CredentialUserRegistrationDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCredentialUserRegistrationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUserRegistrationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUserRegistrationDetails(), nil
}
// GetAuthMethods gets the authMethods property value. Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
func (m *CredentialUserRegistrationDetails) GetAuthMethods()([]RegistrationAuthMethod) {
    val, err := m.GetBackingStore().Get("authMethods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RegistrationAuthMethod)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRegistrationAuthMethod)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RegistrationAuthMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*RegistrationAuthMethod))
            }
            m.SetAuthMethods(res)
        }
        return nil
    }
    res["isCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCapable(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isMfaRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaRegistered(val)
        }
        return nil
    }
    res["isRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRegistered(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetIsCapable gets the isCapable property value. Indicates whether the user is ready to perform self-service password reset or MFA.
func (m *CredentialUserRegistrationDetails) GetIsCapable()(*bool) {
    val, err := m.GetBackingStore().Get("isCapable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsMfaRegistered gets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
func (m *CredentialUserRegistrationDetails) GetIsMfaRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("isMfaRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRegistered gets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("isRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. Provides the user name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. Provides the user principal name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("authMethods", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCapable sets the isCapable property value. Indicates whether the user is ready to perform self-service password reset or MFA.
func (m *CredentialUserRegistrationDetails) SetIsCapable(value *bool)() {
    err := m.GetBackingStore().Set("isCapable", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
func (m *CredentialUserRegistrationDetails) SetIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMfaRegistered sets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
func (m *CredentialUserRegistrationDetails) SetIsMfaRegistered(value *bool)() {
    err := m.GetBackingStore().Set("isMfaRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRegistered sets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
func (m *CredentialUserRegistrationDetails) SetIsRegistered(value *bool)() {
    err := m.GetBackingStore().Set("isRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. Provides the user name of the corresponding user.
func (m *CredentialUserRegistrationDetails) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Provides the user principal name of the corresponding user.
func (m *CredentialUserRegistrationDetails) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// CredentialUserRegistrationDetailsable 
type CredentialUserRegistrationDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthMethods()([]RegistrationAuthMethod)
    GetIsCapable()(*bool)
    GetIsEnabled()(*bool)
    GetIsMfaRegistered()(*bool)
    GetIsRegistered()(*bool)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetAuthMethods(value []RegistrationAuthMethod)()
    SetIsCapable(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsMfaRegistered(value *bool)()
    SetIsRegistered(value *bool)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
