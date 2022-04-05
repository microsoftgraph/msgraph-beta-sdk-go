package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyConfigurationBase 
type CrossTenantAccessPolicyConfigurationBase struct {
    Entity
    // Defines your configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
    b2bCollaborationInbound CrossTenantAccessPolicyB2BSettingable;
    // Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
    b2bCollaborationOutbound CrossTenantAccessPolicyB2BSettingable;
    // Defines your configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
    b2bDirectConnectInbound CrossTenantAccessPolicyB2BSettingable;
    // Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
    b2bDirectConnectOutbound CrossTenantAccessPolicyB2BSettingable;
    // Determines the configuration for trusting other Conditional Access claims from external Azure AD organizations.
    inboundTrust CrossTenantAccessPolicyInboundTrustable;
}
// NewCrossTenantAccessPolicyConfigurationBase instantiates a new crossTenantAccessPolicyConfigurationBase and sets the default values.
func NewCrossTenantAccessPolicyConfigurationBase()(*CrossTenantAccessPolicyConfigurationBase) {
    m := &CrossTenantAccessPolicyConfigurationBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCrossTenantAccessPolicyConfigurationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyConfigurationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyConfigurationBase(), nil
}
// GetB2bCollaborationInbound gets the b2bCollaborationInbound property value. Defines your configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bCollaborationInbound()(CrossTenantAccessPolicyB2BSettingable) {
    if m == nil {
        return nil
    } else {
        return m.b2bCollaborationInbound
    }
}
// GetB2bCollaborationOutbound gets the b2bCollaborationOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bCollaborationOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    if m == nil {
        return nil
    } else {
        return m.b2bCollaborationOutbound
    }
}
// GetB2bDirectConnectInbound gets the b2bDirectConnectInbound property value. Defines your configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bDirectConnectInbound()(CrossTenantAccessPolicyB2BSettingable) {
    if m == nil {
        return nil
    } else {
        return m.b2bDirectConnectInbound
    }
}
// GetB2bDirectConnectOutbound gets the b2bDirectConnectOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bDirectConnectOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    if m == nil {
        return nil
    } else {
        return m.b2bDirectConnectOutbound
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyConfigurationBase) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["b2bCollaborationInbound"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationInbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bCollaborationOutbound"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationOutbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bDirectConnectInbound"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectInbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bDirectConnectOutbound"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectOutbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["inboundTrust"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyInboundTrustFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundTrust(val.(CrossTenantAccessPolicyInboundTrustable))
        }
        return nil
    }
    return res
}
// GetInboundTrust gets the inboundTrust property value. Determines the configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationBase) GetInboundTrust()(CrossTenantAccessPolicyInboundTrustable) {
    if m == nil {
        return nil
    } else {
        return m.inboundTrust
    }
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyConfigurationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("b2bCollaborationInbound", m.GetB2bCollaborationInbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bCollaborationOutbound", m.GetB2bCollaborationOutbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bDirectConnectInbound", m.GetB2bDirectConnectInbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bDirectConnectOutbound", m.GetB2bDirectConnectOutbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inboundTrust", m.GetInboundTrust())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetB2bCollaborationInbound sets the b2bCollaborationInbound property value. Defines your configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bCollaborationInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    if m != nil {
        m.b2bCollaborationInbound = value
    }
}
// SetB2bCollaborationOutbound sets the b2bCollaborationOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bCollaborationOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    if m != nil {
        m.b2bCollaborationOutbound = value
    }
}
// SetB2bDirectConnectInbound sets the b2bDirectConnectInbound property value. Defines your configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bDirectConnectInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    if m != nil {
        m.b2bDirectConnectInbound = value
    }
}
// SetB2bDirectConnectOutbound sets the b2bDirectConnectOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bDirectConnectOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    if m != nil {
        m.b2bDirectConnectOutbound = value
    }
}
// SetInboundTrust sets the inboundTrust property value. Determines the configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationBase) SetInboundTrust(value CrossTenantAccessPolicyInboundTrustable)() {
    if m != nil {
        m.inboundTrust = value
    }
}
