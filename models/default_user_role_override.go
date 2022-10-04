package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultUserRoleOverride provides operations to manage the collection of accessReviewDecision entities.
type DefaultUserRoleOverride struct {
    Entity
    // The isDefault property
    isDefault *bool
    // The rolePermissions property
    rolePermissions []UnifiedRolePermissionable
}
// NewDefaultUserRoleOverride instantiates a new defaultUserRoleOverride and sets the default values.
func NewDefaultUserRoleOverride()(*DefaultUserRoleOverride) {
    m := &DefaultUserRoleOverride{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.defaultUserRoleOverride";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDefaultUserRoleOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultUserRoleOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultUserRoleOverride(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultUserRoleOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isDefault"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDefault)
    res["rolePermissions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRolePermissionFromDiscriminatorValue , m.SetRolePermissions)
    return res
}
// GetIsDefault gets the isDefault property value. The isDefault property
func (m *DefaultUserRoleOverride) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetRolePermissions gets the rolePermissions property value. The rolePermissions property
func (m *DefaultUserRoleOverride) GetRolePermissions()([]UnifiedRolePermissionable) {
    return m.rolePermissions
}
// Serialize serializes information the current object
func (m *DefaultUserRoleOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    if m.GetRolePermissions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRolePermissions())
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *DefaultUserRoleOverride) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetRolePermissions sets the rolePermissions property value. The rolePermissions property
func (m *DefaultUserRoleOverride) SetRolePermissions(value []UnifiedRolePermissionable)() {
    m.rolePermissions = value
}
