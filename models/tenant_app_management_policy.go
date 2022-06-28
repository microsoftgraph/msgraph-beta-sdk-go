package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAppManagementPolicy 
type TenantAppManagementPolicy struct {
    PolicyBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Restrictions that apply as default to all application objects in the tenant.
    applicationRestrictions AppManagementConfigurationable
    // Denotes whether the policy is enabled. Default value is false.
    isEnabled *bool
    // Restrictions that apply as default to all service principal objects in the tenant.
    servicePrincipalRestrictions AppManagementConfigurationable
}
// NewTenantAppManagementPolicy instantiates a new TenantAppManagementPolicy and sets the default values.
func NewTenantAppManagementPolicy()(*TenantAppManagementPolicy) {
    m := &TenantAppManagementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantAppManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAppManagementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAppManagementPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAppManagementPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationRestrictions gets the applicationRestrictions property value. Restrictions that apply as default to all application objects in the tenant.
func (m *TenantAppManagementPolicy) GetApplicationRestrictions()(AppManagementConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.applicationRestrictions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAppManagementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["applicationRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppManagementConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationRestrictions(val.(AppManagementConfigurationable))
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
    res["servicePrincipalRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppManagementConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalRestrictions(val.(AppManagementConfigurationable))
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Denotes whether the policy is enabled. Default value is false.
func (m *TenantAppManagementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetServicePrincipalRestrictions gets the servicePrincipalRestrictions property value. Restrictions that apply as default to all service principal objects in the tenant.
func (m *TenantAppManagementPolicy) GetServicePrincipalRestrictions()(AppManagementConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalRestrictions
    }
}
// Serialize serializes information the current object
func (m *TenantAppManagementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicationRestrictions", m.GetApplicationRestrictions())
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
        err = writer.WriteObjectValue("servicePrincipalRestrictions", m.GetServicePrincipalRestrictions())
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
func (m *TenantAppManagementPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationRestrictions sets the applicationRestrictions property value. Restrictions that apply as default to all application objects in the tenant.
func (m *TenantAppManagementPolicy) SetApplicationRestrictions(value AppManagementConfigurationable)() {
    if m != nil {
        m.applicationRestrictions = value
    }
}
// SetIsEnabled sets the isEnabled property value. Denotes whether the policy is enabled. Default value is false.
func (m *TenantAppManagementPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetServicePrincipalRestrictions sets the servicePrincipalRestrictions property value. Restrictions that apply as default to all service principal objects in the tenant.
func (m *TenantAppManagementPolicy) SetServicePrincipalRestrictions(value AppManagementConfigurationable)() {
    if m != nil {
        m.servicePrincipalRestrictions = value
    }
}
