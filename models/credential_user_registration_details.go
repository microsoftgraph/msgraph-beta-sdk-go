package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUserRegistrationDetails 
type CredentialUserRegistrationDetails struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
    authMethods []string
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCredentialUserRegistrationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUserRegistrationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUserRegistrationDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CredentialUserRegistrationDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthMethods gets the authMethods property value. Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
func (m *CredentialUserRegistrationDetails) GetAuthMethods()([]string) {
    if m == nil {
        return nil
    } else {
        return m.authMethods
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
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
    if m == nil {
        return nil
    } else {
        return m.isCapable
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetIsMfaRegistered gets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
func (m *CredentialUserRegistrationDetails) GetIsMfaRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaRegistered
    }
}
// GetIsRegistered gets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
func (m *CredentialUserRegistrationDetails) GetIsRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRegistered
    }
}
// GetUserDisplayName gets the userDisplayName property value. Provides the user name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. Provides the user principal name of the corresponding user.
func (m *CredentialUserRegistrationDetails) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Serialize serializes information the current object
func (m *CredentialUserRegistrationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthMethods() != nil {
        err = writer.WriteCollectionOfStringValues("authMethods", m.GetAuthMethods())
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CredentialUserRegistrationDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthMethods sets the authMethods property value. Represents the authentication method that the user has registered. Possible values are: email, mobilePhone, officePhone,  securityQuestion (only used for self-service password reset), appNotification,  appCode, alternateMobilePhone (supported only in registration),  fido,  appPassword,  unknownFutureValue.
func (m *CredentialUserRegistrationDetails) SetAuthMethods(value []string)() {
    if m != nil {
        m.authMethods = value
    }
}
// SetIsCapable sets the isCapable property value. Indicates whether the user is ready to perform self-service password reset or MFA.
func (m *CredentialUserRegistrationDetails) SetIsCapable(value *bool)() {
    if m != nil {
        m.isCapable = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the user enabled to perform self-service password reset.
func (m *CredentialUserRegistrationDetails) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetIsMfaRegistered sets the isMfaRegistered property value. Indicates whether the user is registered for MFA.
func (m *CredentialUserRegistrationDetails) SetIsMfaRegistered(value *bool)() {
    if m != nil {
        m.isMfaRegistered = value
    }
}
// SetIsRegistered sets the isRegistered property value. Indicates whether the user has registered any authentication methods for self-service password reset.
func (m *CredentialUserRegistrationDetails) SetIsRegistered(value *bool)() {
    if m != nil {
        m.isRegistered = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. Provides the user name of the corresponding user.
func (m *CredentialUserRegistrationDetails) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Provides the user principal name of the corresponding user.
func (m *CredentialUserRegistrationDetails) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
