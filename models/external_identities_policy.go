package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalIdentitiesPolicy 
type ExternalIdentitiesPolicy struct {
    PolicyBase
    // Notifies Azure AD whether to clean up the user information about the external identity, from the guest tenant, when the user is deleted in their home tenant.
    allowDeletedIdentitiesDataRemoval *bool
    // Defines whether external users can leave the guest tenant. If set to false, self-service controls are not enabled, and the admin of the guest tenant must manually remove the external user from the guest tenant.
    allowExternalIdentitiesToLeave *bool
}
// NewExternalIdentitiesPolicy instantiates a new ExternalIdentitiesPolicy and sets the default values.
func NewExternalIdentitiesPolicy()(*ExternalIdentitiesPolicy) {
    m := &ExternalIdentitiesPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreateExternalIdentitiesPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalIdentitiesPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalIdentitiesPolicy(), nil
}
// GetAllowDeletedIdentitiesDataRemoval gets the allowDeletedIdentitiesDataRemoval property value. Notifies Azure AD whether to clean up the user information about the external identity, from the guest tenant, when the user is deleted in their home tenant.
func (m *ExternalIdentitiesPolicy) GetAllowDeletedIdentitiesDataRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeletedIdentitiesDataRemoval
    }
}
// GetAllowExternalIdentitiesToLeave gets the allowExternalIdentitiesToLeave property value. Defines whether external users can leave the guest tenant. If set to false, self-service controls are not enabled, and the admin of the guest tenant must manually remove the external user from the guest tenant.
func (m *ExternalIdentitiesPolicy) GetAllowExternalIdentitiesToLeave()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowExternalIdentitiesToLeave
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalIdentitiesPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["allowDeletedIdentitiesDataRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeletedIdentitiesDataRemoval(val)
        }
        return nil
    }
    res["allowExternalIdentitiesToLeave"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowExternalIdentitiesToLeave(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExternalIdentitiesPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowDeletedIdentitiesDataRemoval", m.GetAllowDeletedIdentitiesDataRemoval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowExternalIdentitiesToLeave", m.GetAllowExternalIdentitiesToLeave())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowDeletedIdentitiesDataRemoval sets the allowDeletedIdentitiesDataRemoval property value. Notifies Azure AD whether to clean up the user information about the external identity, from the guest tenant, when the user is deleted in their home tenant.
func (m *ExternalIdentitiesPolicy) SetAllowDeletedIdentitiesDataRemoval(value *bool)() {
    if m != nil {
        m.allowDeletedIdentitiesDataRemoval = value
    }
}
// SetAllowExternalIdentitiesToLeave sets the allowExternalIdentitiesToLeave property value. Defines whether external users can leave the guest tenant. If set to false, self-service controls are not enabled, and the admin of the guest tenant must manually remove the external user from the guest tenant.
func (m *ExternalIdentitiesPolicy) SetAllowExternalIdentitiesToLeave(value *bool)() {
    if m != nil {
        m.allowExternalIdentitiesToLeave = value
    }
}
