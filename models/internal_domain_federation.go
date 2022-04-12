package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InternalDomainFederation 
type InternalDomainFederation struct {
    SamlOrWsFedProvider
    // The activeSignInUri property
    activeSignInUri *string
    // The federatedIdpMfaBehavior property
    federatedIdpMfaBehavior *FederatedIdpMfaBehavior
    // The isSignedAuthenticationRequestRequired property
    isSignedAuthenticationRequestRequired *bool
    // The nextSigningCertificate property
    nextSigningCertificate *string
    // The promptLoginBehavior property
    promptLoginBehavior *PromptLoginBehavior
    // The signingCertificateUpdateStatus property
    signingCertificateUpdateStatus SigningCertificateUpdateStatusable
    // The signOutUri property
    signOutUri *string
}
// NewInternalDomainFederation instantiates a new internalDomainFederation and sets the default values.
func NewInternalDomainFederation()(*InternalDomainFederation) {
    m := &InternalDomainFederation{
        SamlOrWsFedProvider: *NewSamlOrWsFedProvider(),
    }
    return m
}
// CreateInternalDomainFederationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInternalDomainFederationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInternalDomainFederation(), nil
}
// GetActiveSignInUri gets the activeSignInUri property value. The activeSignInUri property
func (m *InternalDomainFederation) GetActiveSignInUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activeSignInUri
    }
}
// GetFederatedIdpMfaBehavior gets the federatedIdpMfaBehavior property value. The federatedIdpMfaBehavior property
func (m *InternalDomainFederation) GetFederatedIdpMfaBehavior()(*FederatedIdpMfaBehavior) {
    if m == nil {
        return nil
    } else {
        return m.federatedIdpMfaBehavior
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternalDomainFederation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SamlOrWsFedProvider.GetFieldDeserializers()
    res["activeSignInUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveSignInUri(val)
        }
        return nil
    }
    res["federatedIdpMfaBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFederatedIdpMfaBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedIdpMfaBehavior(val.(*FederatedIdpMfaBehavior))
        }
        return nil
    }
    res["isSignedAuthenticationRequestRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSignedAuthenticationRequestRequired(val)
        }
        return nil
    }
    res["nextSigningCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextSigningCertificate(val)
        }
        return nil
    }
    res["promptLoginBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePromptLoginBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPromptLoginBehavior(val.(*PromptLoginBehavior))
        }
        return nil
    }
    res["signingCertificateUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigningCertificateUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningCertificateUpdateStatus(val.(SigningCertificateUpdateStatusable))
        }
        return nil
    }
    res["signOutUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignOutUri(val)
        }
        return nil
    }
    return res
}
// GetIsSignedAuthenticationRequestRequired gets the isSignedAuthenticationRequestRequired property value. The isSignedAuthenticationRequestRequired property
func (m *InternalDomainFederation) GetIsSignedAuthenticationRequestRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSignedAuthenticationRequestRequired
    }
}
// GetNextSigningCertificate gets the nextSigningCertificate property value. The nextSigningCertificate property
func (m *InternalDomainFederation) GetNextSigningCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextSigningCertificate
    }
}
// GetPromptLoginBehavior gets the promptLoginBehavior property value. The promptLoginBehavior property
func (m *InternalDomainFederation) GetPromptLoginBehavior()(*PromptLoginBehavior) {
    if m == nil {
        return nil
    } else {
        return m.promptLoginBehavior
    }
}
// GetSigningCertificateUpdateStatus gets the signingCertificateUpdateStatus property value. The signingCertificateUpdateStatus property
func (m *InternalDomainFederation) GetSigningCertificateUpdateStatus()(SigningCertificateUpdateStatusable) {
    if m == nil {
        return nil
    } else {
        return m.signingCertificateUpdateStatus
    }
}
// GetSignOutUri gets the signOutUri property value. The signOutUri property
func (m *InternalDomainFederation) GetSignOutUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signOutUri
    }
}
// Serialize serializes information the current object
func (m *InternalDomainFederation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SamlOrWsFedProvider.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activeSignInUri", m.GetActiveSignInUri())
        if err != nil {
            return err
        }
    }
    if m.GetFederatedIdpMfaBehavior() != nil {
        cast := (*m.GetFederatedIdpMfaBehavior()).String()
        err = writer.WriteStringValue("federatedIdpMfaBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSignedAuthenticationRequestRequired", m.GetIsSignedAuthenticationRequestRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nextSigningCertificate", m.GetNextSigningCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetPromptLoginBehavior() != nil {
        cast := (*m.GetPromptLoginBehavior()).String()
        err = writer.WriteStringValue("promptLoginBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signingCertificateUpdateStatus", m.GetSigningCertificateUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signOutUri", m.GetSignOutUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveSignInUri sets the activeSignInUri property value. The activeSignInUri property
func (m *InternalDomainFederation) SetActiveSignInUri(value *string)() {
    if m != nil {
        m.activeSignInUri = value
    }
}
// SetFederatedIdpMfaBehavior sets the federatedIdpMfaBehavior property value. The federatedIdpMfaBehavior property
func (m *InternalDomainFederation) SetFederatedIdpMfaBehavior(value *FederatedIdpMfaBehavior)() {
    if m != nil {
        m.federatedIdpMfaBehavior = value
    }
}
// SetIsSignedAuthenticationRequestRequired sets the isSignedAuthenticationRequestRequired property value. The isSignedAuthenticationRequestRequired property
func (m *InternalDomainFederation) SetIsSignedAuthenticationRequestRequired(value *bool)() {
    if m != nil {
        m.isSignedAuthenticationRequestRequired = value
    }
}
// SetNextSigningCertificate sets the nextSigningCertificate property value. The nextSigningCertificate property
func (m *InternalDomainFederation) SetNextSigningCertificate(value *string)() {
    if m != nil {
        m.nextSigningCertificate = value
    }
}
// SetPromptLoginBehavior sets the promptLoginBehavior property value. The promptLoginBehavior property
func (m *InternalDomainFederation) SetPromptLoginBehavior(value *PromptLoginBehavior)() {
    if m != nil {
        m.promptLoginBehavior = value
    }
}
// SetSigningCertificateUpdateStatus sets the signingCertificateUpdateStatus property value. The signingCertificateUpdateStatus property
func (m *InternalDomainFederation) SetSigningCertificateUpdateStatus(value SigningCertificateUpdateStatusable)() {
    if m != nil {
        m.signingCertificateUpdateStatus = value
    }
}
// SetSignOutUri sets the signOutUri property value. The signOutUri property
func (m *InternalDomainFederation) SetSignOutUri(value *string)() {
    if m != nil {
        m.signOutUri = value
    }
}
