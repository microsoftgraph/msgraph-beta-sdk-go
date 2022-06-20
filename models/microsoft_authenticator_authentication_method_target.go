package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorAuthenticationMethodTarget 
type MicrosoftAuthenticatorAuthenticationMethodTarget struct {
    AuthenticationMethodTarget
    // Determines which types of notifications can be used for sign-in. The possible values are: deviceBasedPush (passwordless only), push, and any.
    authenticationMode *MicrosoftAuthenticatorAuthenticationMode
    // Determines whether the user is shown additional context in their Authenticator app notification. In the body of the Authenticator notification, the user will be shown the app they are signing into along with the location that the authentication request originated from. Possible values are: enabled, disabled, default.
    displayAppInformationRequiredState *AdvancedConfigState
    // Requires number matching for MFA notifications. Value is ignored for phone sign-in notifications. Possible values are: enabled, disabled, default.
    numberMatchingRequiredState *AdvancedConfigState
}
// NewMicrosoftAuthenticatorAuthenticationMethodTarget instantiates a new MicrosoftAuthenticatorAuthenticationMethodTarget and sets the default values.
func NewMicrosoftAuthenticatorAuthenticationMethodTarget()(*MicrosoftAuthenticatorAuthenticationMethodTarget) {
    m := &MicrosoftAuthenticatorAuthenticationMethodTarget{
        AuthenticationMethodTarget: *NewAuthenticationMethodTarget(),
    }
    return m
}
// CreateMicrosoftAuthenticatorAuthenticationMethodTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAuthenticatorAuthenticationMethodTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorAuthenticationMethodTarget(), nil
}
// GetAuthenticationMode gets the authenticationMode property value. Determines which types of notifications can be used for sign-in. The possible values are: deviceBasedPush (passwordless only), push, and any.
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) GetAuthenticationMode()(*MicrosoftAuthenticatorAuthenticationMode) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMode
    }
}
// GetDisplayAppInformationRequiredState gets the displayAppInformationRequiredState property value. Determines whether the user is shown additional context in their Authenticator app notification. In the body of the Authenticator notification, the user will be shown the app they are signing into along with the location that the authentication request originated from. Possible values are: enabled, disabled, default.
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) GetDisplayAppInformationRequiredState()(*AdvancedConfigState) {
    if m == nil {
        return nil
    } else {
        return m.displayAppInformationRequiredState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodTarget.GetFieldDeserializers()
    res["authenticationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftAuthenticatorAuthenticationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMode(val.(*MicrosoftAuthenticatorAuthenticationMode))
        }
        return nil
    }
    res["displayAppInformationRequiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAppInformationRequiredState(val.(*AdvancedConfigState))
        }
        return nil
    }
    res["numberMatchingRequiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberMatchingRequiredState(val.(*AdvancedConfigState))
        }
        return nil
    }
    return res
}
// GetNumberMatchingRequiredState gets the numberMatchingRequiredState property value. Requires number matching for MFA notifications. Value is ignored for phone sign-in notifications. Possible values are: enabled, disabled, default.
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) GetNumberMatchingRequiredState()(*AdvancedConfigState) {
    if m == nil {
        return nil
    } else {
        return m.numberMatchingRequiredState
    }
}
// Serialize serializes information the current object
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodTarget.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMode() != nil {
        cast := (*m.GetAuthenticationMode()).String()
        err = writer.WriteStringValue("authenticationMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDisplayAppInformationRequiredState() != nil {
        cast := (*m.GetDisplayAppInformationRequiredState()).String()
        err = writer.WriteStringValue("displayAppInformationRequiredState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNumberMatchingRequiredState() != nil {
        cast := (*m.GetNumberMatchingRequiredState()).String()
        err = writer.WriteStringValue("numberMatchingRequiredState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMode sets the authenticationMode property value. Determines which types of notifications can be used for sign-in. The possible values are: deviceBasedPush (passwordless only), push, and any.
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) SetAuthenticationMode(value *MicrosoftAuthenticatorAuthenticationMode)() {
    if m != nil {
        m.authenticationMode = value
    }
}
// SetDisplayAppInformationRequiredState sets the displayAppInformationRequiredState property value. Determines whether the user is shown additional context in their Authenticator app notification. In the body of the Authenticator notification, the user will be shown the app they are signing into along with the location that the authentication request originated from. Possible values are: enabled, disabled, default.
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) SetDisplayAppInformationRequiredState(value *AdvancedConfigState)() {
    if m != nil {
        m.displayAppInformationRequiredState = value
    }
}
// SetNumberMatchingRequiredState sets the numberMatchingRequiredState property value. Requires number matching for MFA notifications. Value is ignored for phone sign-in notifications. Possible values are: enabled, disabled, default.
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) SetNumberMatchingRequiredState(value *AdvancedConfigState)() {
    if m != nil {
        m.numberMatchingRequiredState = value
    }
}
