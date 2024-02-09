package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoleMembershipGovernanceCriteria struct {
    GovernanceCriteria
}
// NewRoleMembershipGovernanceCriteria instantiates a new RoleMembershipGovernanceCriteria and sets the default values.
func NewRoleMembershipGovernanceCriteria()(*RoleMembershipGovernanceCriteria) {
    m := &RoleMembershipGovernanceCriteria{
        GovernanceCriteria: *NewGovernanceCriteria(),
    }
    odataTypeValue := "#microsoft.graph.roleMembershipGovernanceCriteria"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRoleMembershipGovernanceCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleMembershipGovernanceCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleMembershipGovernanceCriteria(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
// returns a *string when successful
func (m *RoleMembershipGovernanceCriteria) GetRoleId()(*string) {
    val, err := m.GetBackingStore().Get("roleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleTemplateId gets the roleTemplateId property value. The roleTemplateId property
// returns a *string when successful
func (m *RoleMembershipGovernanceCriteria) GetRoleTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("roleTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    return nil
}
// SetRoleId sets the roleId property value. The roleId property
func (m *RoleMembershipGovernanceCriteria) SetRoleId(value *string)() {
    err := m.GetBackingStore().Set("roleId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The roleTemplateId property
func (m *RoleMembershipGovernanceCriteria) SetRoleTemplateId(value *string)() {
    err := m.GetBackingStore().Set("roleTemplateId", value)
    if err != nil {
        panic(err)
    }
}
type RoleMembershipGovernanceCriteriaable interface {
    GovernanceCriteriaable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoleId()(*string)
    GetRoleTemplateId()(*string)
    SetRoleId(value *string)()
    SetRoleTemplateId(value *string)()
}
