package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyConfigurationPartner 
type CrossTenantAccessPolicyConfigurationPartner struct {
    CrossTenantAccessPolicyConfigurationBase
    // Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
    isServiceProvider *bool;
    // The tenant identifier for the partner Azure AD organization. Read-only. Key.
    tenantId *string;
}
// NewCrossTenantAccessPolicyConfigurationPartner instantiates a new crossTenantAccessPolicyConfigurationPartner and sets the default values.
func NewCrossTenantAccessPolicyConfigurationPartner()(*CrossTenantAccessPolicyConfigurationPartner) {
    m := &CrossTenantAccessPolicyConfigurationPartner{
        CrossTenantAccessPolicyConfigurationBase: *NewCrossTenantAccessPolicyConfigurationBase(),
    }
    return m
}
// CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyConfigurationPartner(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyConfigurationPartner) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CrossTenantAccessPolicyConfigurationBase.GetFieldDeserializers()
    res["isServiceProvider"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsServiceProvider(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetIsServiceProvider gets the isServiceProvider property value. Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetIsServiceProvider()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isServiceProvider
    }
}
// GetTenantId gets the tenantId property value. The tenant identifier for the partner Azure AD organization. Read-only. Key.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyConfigurationPartner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CrossTenantAccessPolicyConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isServiceProvider", m.GetIsServiceProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsServiceProvider sets the isServiceProvider property value. Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetIsServiceProvider(value *bool)() {
    if m != nil {
        m.isServiceProvider = value
    }
}
// SetTenantId sets the tenantId property value. The tenant identifier for the partner Azure AD organization. Read-only. Key.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
