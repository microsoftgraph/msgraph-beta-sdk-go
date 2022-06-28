package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleMembershipGovernanceCriteria 
type RoleMembershipGovernanceCriteria struct {
    GovernanceCriteria
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The roleId property
    roleId *string
    // The roleTemplateId property
    roleTemplateId *string
}
// NewRoleMembershipGovernanceCriteria instantiates a new RoleMembershipGovernanceCriteria and sets the default values.
func NewRoleMembershipGovernanceCriteria()(*RoleMembershipGovernanceCriteria) {
    m := &RoleMembershipGovernanceCriteria{
        GovernanceCriteria: *NewGovernanceCriteria(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleMembershipGovernanceCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleMembershipGovernanceCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleMembershipGovernanceCriteria(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleMembershipGovernanceCriteria) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleMembershipGovernanceCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GovernanceCriteria.GetFieldDeserializers()
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleTemplateId(val)
        }
        return nil
    }
    return res
}
// GetRoleId gets the roleId property value. The roleId property
func (m *RoleMembershipGovernanceCriteria) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// GetRoleTemplateId gets the roleTemplateId property value. The roleTemplateId property
func (m *RoleMembershipGovernanceCriteria) GetRoleTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleTemplateId
    }
}
// Serialize serializes information the current object
func (m *RoleMembershipGovernanceCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GovernanceCriteria.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleTemplateId", m.GetRoleTemplateId())
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
func (m *RoleMembershipGovernanceCriteria) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRoleId sets the roleId property value. The roleId property
func (m *RoleMembershipGovernanceCriteria) SetRoleId(value *string)() {
    if m != nil {
        m.roleId = value
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The roleTemplateId property
func (m *RoleMembershipGovernanceCriteria) SetRoleTemplateId(value *string)() {
    if m != nil {
        m.roleTemplateId = value
    }
}
