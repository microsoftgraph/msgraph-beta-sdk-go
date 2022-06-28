package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentitySecurityDefaultsEnforcementPolicy 
type IdentitySecurityDefaultsEnforcementPolicy struct {
    PolicyBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If set to true, Azure Active Directory security defaults is enabled for the tenant.
    isEnabled *bool
}
// NewIdentitySecurityDefaultsEnforcementPolicy instantiates a new IdentitySecurityDefaultsEnforcementPolicy and sets the default values.
func NewIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    m := &IdentitySecurityDefaultsEnforcementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentitySecurityDefaultsEnforcementPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
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
    return res
}
// GetIsEnabled gets the isEnabled property value. If set to true, Azure Active Directory security defaults is enabled for the tenant.
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Serialize serializes information the current object
func (m *IdentitySecurityDefaultsEnforcementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
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
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsEnabled sets the isEnabled property value. If set to true, Azure Active Directory security defaults is enabled for the tenant.
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
