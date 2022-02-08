package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InternalDomainFederation 
type InternalDomainFederation struct {
    SamlOrWsFedProvider
    // 
    activeSignInUri *string;
    // 
    federatedIdpMfaBehavior *FederatedIdpMfaBehavior;
    // 
    isSignedAuthenticationRequestRequired *bool;
    // 
    nextSigningCertificate *string;
    // 
    promptLoginBehavior *PromptLoginBehavior;
    // 
    signingCertificateUpdateStatus *SigningCertificateUpdateStatus;
    // 
    signOutUri *string;
}
// NewInternalDomainFederation instantiates a new internalDomainFederation and sets the default values.
func NewInternalDomainFederation()(*InternalDomainFederation) {
    m := &InternalDomainFederation{
        SamlOrWsFedProvider: *NewSamlOrWsFedProvider(),
    }
    return m
}
// GetActiveSignInUri gets the activeSignInUri property value. 
func (m *InternalDomainFederation) GetActiveSignInUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activeSignInUri
    }
}
// GetFederatedIdpMfaBehavior gets the federatedIdpMfaBehavior property value. 
func (m *InternalDomainFederation) GetFederatedIdpMfaBehavior()(*FederatedIdpMfaBehavior) {
    if m == nil {
        return nil
    } else {
        return m.federatedIdpMfaBehavior
    }
}
// GetIsSignedAuthenticationRequestRequired gets the isSignedAuthenticationRequestRequired property value. 
func (m *InternalDomainFederation) GetIsSignedAuthenticationRequestRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSignedAuthenticationRequestRequired
    }
}
// GetNextSigningCertificate gets the nextSigningCertificate property value. 
func (m *InternalDomainFederation) GetNextSigningCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextSigningCertificate
    }
}
// GetPromptLoginBehavior gets the promptLoginBehavior property value. 
func (m *InternalDomainFederation) GetPromptLoginBehavior()(*PromptLoginBehavior) {
    if m == nil {
        return nil
    } else {
        return m.promptLoginBehavior
    }
}
// GetSigningCertificateUpdateStatus gets the signingCertificateUpdateStatus property value. 
func (m *InternalDomainFederation) GetSigningCertificateUpdateStatus()(*SigningCertificateUpdateStatus) {
    if m == nil {
        return nil
    } else {
        return m.signingCertificateUpdateStatus
    }
}
// GetSignOutUri gets the signOutUri property value. 
func (m *InternalDomainFederation) GetSignOutUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signOutUri
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternalDomainFederation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.SamlOrWsFedProvider.GetFieldDeserializers()
    res["activeSignInUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveSignInUri(val)
        }
        return nil
    }
    res["federatedIdpMfaBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFederatedIdpMfaBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedIdpMfaBehavior(val.(*FederatedIdpMfaBehavior))
        }
        return nil
    }
    res["isSignedAuthenticationRequestRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSignedAuthenticationRequestRequired(val)
        }
        return nil
    }
    res["nextSigningCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextSigningCertificate(val)
        }
        return nil
    }
    res["promptLoginBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePromptLoginBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPromptLoginBehavior(val.(*PromptLoginBehavior))
        }
        return nil
    }
    res["signingCertificateUpdateStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSigningCertificateUpdateStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningCertificateUpdateStatus(val.(*SigningCertificateUpdateStatus))
        }
        return nil
    }
    res["signOutUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *InternalDomainFederation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InternalDomainFederation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetActiveSignInUri sets the activeSignInUri property value. 
func (m *InternalDomainFederation) SetActiveSignInUri(value *string)() {
    if m != nil {
        m.activeSignInUri = value
    }
}
// SetFederatedIdpMfaBehavior sets the federatedIdpMfaBehavior property value. 
func (m *InternalDomainFederation) SetFederatedIdpMfaBehavior(value *FederatedIdpMfaBehavior)() {
    if m != nil {
        m.federatedIdpMfaBehavior = value
    }
}
// SetIsSignedAuthenticationRequestRequired sets the isSignedAuthenticationRequestRequired property value. 
func (m *InternalDomainFederation) SetIsSignedAuthenticationRequestRequired(value *bool)() {
    if m != nil {
        m.isSignedAuthenticationRequestRequired = value
    }
}
// SetNextSigningCertificate sets the nextSigningCertificate property value. 
func (m *InternalDomainFederation) SetNextSigningCertificate(value *string)() {
    if m != nil {
        m.nextSigningCertificate = value
    }
}
// SetPromptLoginBehavior sets the promptLoginBehavior property value. 
func (m *InternalDomainFederation) SetPromptLoginBehavior(value *PromptLoginBehavior)() {
    if m != nil {
        m.promptLoginBehavior = value
    }
}
// SetSigningCertificateUpdateStatus sets the signingCertificateUpdateStatus property value. 
func (m *InternalDomainFederation) SetSigningCertificateUpdateStatus(value *SigningCertificateUpdateStatus)() {
    if m != nil {
        m.signingCertificateUpdateStatus = value
    }
}
// SetSignOutUri sets the signOutUri property value. 
func (m *InternalDomainFederation) SetSignOutUri(value *string)() {
    if m != nil {
        m.signOutUri = value
    }
}
