package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CrossTenantAccessPolicyConfigurationBase 
type CrossTenantAccessPolicyConfigurationBase struct {
    Entity
    // Defines your configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
    b2bCollaborationInbound *CrossTenantAccessPolicyB2BSetting;
    // Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
    b2bCollaborationOutbound *CrossTenantAccessPolicyB2BSetting;
    // Defines your configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
    b2bDirectConnectInbound *CrossTenantAccessPolicyB2BSetting;
    // Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
    b2bDirectConnectOutbound *CrossTenantAccessPolicyB2BSetting;
    // Determines the configuration for trusting other Conditional Access claims from external Azure AD organizations.
    inboundTrust *CrossTenantAccessPolicyInboundTrust;
}
// NewCrossTenantAccessPolicyConfigurationBase instantiates a new crossTenantAccessPolicyConfigurationBase and sets the default values.
func NewCrossTenantAccessPolicyConfigurationBase()(*CrossTenantAccessPolicyConfigurationBase) {
    m := &CrossTenantAccessPolicyConfigurationBase{
        Entity: *NewEntity(),
    }
    return m
}
// GetB2bCollaborationInbound gets the b2bCollaborationInbound property value. Defines your configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bCollaborationInbound()(*CrossTenantAccessPolicyB2BSetting) {
    if m == nil {
        return nil
    } else {
        return m.b2bCollaborationInbound
    }
}
// GetB2bCollaborationOutbound gets the b2bCollaborationOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bCollaborationOutbound()(*CrossTenantAccessPolicyB2BSetting) {
    if m == nil {
        return nil
    } else {
        return m.b2bCollaborationOutbound
    }
}
// GetB2bDirectConnectInbound gets the b2bDirectConnectInbound property value. Defines your configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bDirectConnectInbound()(*CrossTenantAccessPolicyB2BSetting) {
    if m == nil {
        return nil
    } else {
        return m.b2bDirectConnectInbound
    }
}
// GetB2bDirectConnectOutbound gets the b2bDirectConnectOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) GetB2bDirectConnectOutbound()(*CrossTenantAccessPolicyB2BSetting) {
    if m == nil {
        return nil
    } else {
        return m.b2bDirectConnectOutbound
    }
}
// GetInboundTrust gets the inboundTrust property value. Determines the configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationBase) GetInboundTrust()(*CrossTenantAccessPolicyInboundTrust) {
    if m == nil {
        return nil
    } else {
        return m.inboundTrust
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyConfigurationBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["b2bCollaborationInbound"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyB2BSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationInbound(val.(*CrossTenantAccessPolicyB2BSetting))
        }
        return nil
    }
    res["b2bCollaborationOutbound"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyB2BSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationOutbound(val.(*CrossTenantAccessPolicyB2BSetting))
        }
        return nil
    }
    res["b2bDirectConnectInbound"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyB2BSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectInbound(val.(*CrossTenantAccessPolicyB2BSetting))
        }
        return nil
    }
    res["b2bDirectConnectOutbound"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyB2BSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectOutbound(val.(*CrossTenantAccessPolicyB2BSetting))
        }
        return nil
    }
    res["inboundTrust"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyInboundTrust() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundTrust(val.(*CrossTenantAccessPolicyInboundTrust))
        }
        return nil
    }
    return res
}
func (m *CrossTenantAccessPolicyConfigurationBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyConfigurationBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bCollaborationInbound(value *CrossTenantAccessPolicyB2BSetting)() {
    if m != nil {
        m.b2bCollaborationInbound = value
    }
}
// SetB2bCollaborationOutbound sets the b2bCollaborationOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bCollaborationOutbound(value *CrossTenantAccessPolicyB2BSetting)() {
    if m != nil {
        m.b2bCollaborationOutbound = value
    }
}
// SetB2bDirectConnectInbound sets the b2bDirectConnectInbound property value. Defines your configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bDirectConnectInbound(value *CrossTenantAccessPolicyB2BSetting)() {
    if m != nil {
        m.b2bDirectConnectInbound = value
    }
}
// SetB2bDirectConnectOutbound sets the b2bDirectConnectOutbound property value. Defines your configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationBase) SetB2bDirectConnectOutbound(value *CrossTenantAccessPolicyB2BSetting)() {
    if m != nil {
        m.b2bDirectConnectOutbound = value
    }
}
// SetInboundTrust sets the inboundTrust property value. Determines the configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationBase) SetInboundTrust(value *CrossTenantAccessPolicyInboundTrust)() {
    if m != nil {
        m.inboundTrust = value
    }
}
