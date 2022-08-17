package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorAuthenticationMethodTarget 
type MicrosoftAuthenticatorAuthenticationMethodTarget struct {
    AuthenticationMethodTarget
    // The authenticationMode property
    authenticationMode *MicrosoftAuthenticatorAuthenticationMode
    // The displayAppInformationRequiredState property
    displayAppInformationRequiredState *AdvancedConfigState
    // The numberMatchingRequiredState property
    numberMatchingRequiredState *AdvancedConfigState
}
// NewMicrosoftAuthenticatorAuthenticationMethodTarget instantiates a new MicrosoftAuthenticatorAuthenticationMethodTarget and sets the default values.
func NewMicrosoftAuthenticatorAuthenticationMethodTarget()(*MicrosoftAuthenticatorAuthenticationMethodTarget) {
    m := &MicrosoftAuthenticatorAuthenticationMethodTarget{
        AuthenticationMethodTarget: *NewAuthenticationMethodTarget(),
    }
    odataTypeValue := "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMicrosoftAuthenticatorAuthenticationMethodTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAuthenticatorAuthenticationMethodTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorAuthenticationMethodTarget(), nil
}
// GetAuthenticationMode gets the authenticationMode property value. The authenticationMode property
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) GetAuthenticationMode()(*MicrosoftAuthenticatorAuthenticationMode) {
    return m.authenticationMode
}
// GetDisplayAppInformationRequiredState gets the displayAppInformationRequiredState property value. The displayAppInformationRequiredState property
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) GetDisplayAppInformationRequiredState()(*AdvancedConfigState) {
    return m.displayAppInformationRequiredState
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
// GetNumberMatchingRequiredState gets the numberMatchingRequiredState property value. The numberMatchingRequiredState property
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) GetNumberMatchingRequiredState()(*AdvancedConfigState) {
    return m.numberMatchingRequiredState
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
// SetAuthenticationMode sets the authenticationMode property value. The authenticationMode property
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) SetAuthenticationMode(value *MicrosoftAuthenticatorAuthenticationMode)() {
    m.authenticationMode = value
}
// SetDisplayAppInformationRequiredState sets the displayAppInformationRequiredState property value. The displayAppInformationRequiredState property
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) SetDisplayAppInformationRequiredState(value *AdvancedConfigState)() {
    m.displayAppInformationRequiredState = value
}
// SetNumberMatchingRequiredState sets the numberMatchingRequiredState property value. The numberMatchingRequiredState property
func (m *MicrosoftAuthenticatorAuthenticationMethodTarget) SetNumberMatchingRequiredState(value *AdvancedConfigState)() {
    m.numberMatchingRequiredState = value
}
