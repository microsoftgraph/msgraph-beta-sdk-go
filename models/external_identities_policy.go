package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalIdentitiesPolicy 
type ExternalIdentitiesPolicy struct {
    PolicyBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allowDeletedIdentitiesDataRemoval property
    allowDeletedIdentitiesDataRemoval *bool
    // The allowExternalIdentitiesToLeave property
    allowExternalIdentitiesToLeave *bool
}
// NewExternalIdentitiesPolicy instantiates a new ExternalIdentitiesPolicy and sets the default values.
func NewExternalIdentitiesPolicy()(*ExternalIdentitiesPolicy) {
    m := &ExternalIdentitiesPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExternalIdentitiesPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalIdentitiesPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalIdentitiesPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExternalIdentitiesPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowDeletedIdentitiesDataRemoval gets the allowDeletedIdentitiesDataRemoval property value. The allowDeletedIdentitiesDataRemoval property
func (m *ExternalIdentitiesPolicy) GetAllowDeletedIdentitiesDataRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeletedIdentitiesDataRemoval
    }
}
// GetAllowExternalIdentitiesToLeave gets the allowExternalIdentitiesToLeave property value. The allowExternalIdentitiesToLeave property
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExternalIdentitiesPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowDeletedIdentitiesDataRemoval sets the allowDeletedIdentitiesDataRemoval property value. The allowDeletedIdentitiesDataRemoval property
func (m *ExternalIdentitiesPolicy) SetAllowDeletedIdentitiesDataRemoval(value *bool)() {
    if m != nil {
        m.allowDeletedIdentitiesDataRemoval = value
    }
}
// SetAllowExternalIdentitiesToLeave sets the allowExternalIdentitiesToLeave property value. The allowExternalIdentitiesToLeave property
func (m *ExternalIdentitiesPolicy) SetAllowExternalIdentitiesToLeave(value *bool)() {
    if m != nil {
        m.allowExternalIdentitiesToLeave = value
    }
}
