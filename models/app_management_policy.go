package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppManagementPolicy 
type AppManagementPolicy struct {
    PolicyBase
    // The appliesTo property
    appliesTo []DirectoryObjectable
    // The isEnabled property
    isEnabled *bool
    // The restrictions property
    restrictions AppManagementConfigurationable
}
// NewAppManagementPolicy instantiates a new appManagementPolicy and sets the default values.
func NewAppManagementPolicy()(*AppManagementPolicy) {
    m := &AppManagementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreateAppManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppManagementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppManagementPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. The appliesTo property
func (m *AppManagementPolicy) GetAppliesTo()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppManagementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetAppliesTo(res)
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
    res["restrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppManagementConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictions(val.(AppManagementConfigurationable))
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
func (m *AppManagementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetRestrictions gets the restrictions property value. The restrictions property
func (m *AppManagementPolicy) GetRestrictions()(AppManagementConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.restrictions
    }
}
// Serialize serializes information the current object
func (m *AppManagementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliesTo", cast)
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
        err = writer.WriteObjectValue("restrictions", m.GetRestrictions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. The appliesTo property
func (m *AppManagementPolicy) SetAppliesTo(value []DirectoryObjectable)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *AppManagementPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetRestrictions sets the restrictions property value. The restrictions property
func (m *AppManagementPolicy) SetRestrictions(value AppManagementConfigurationable)() {
    if m != nil {
        m.restrictions = value
    }
}
