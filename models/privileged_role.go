package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRole 
type PrivilegedRole struct {
    Entity
    // The assignments for this role. Read-only. Nullable.
    assignments []PrivilegedRoleAssignmentable
    // Role name.
    name *string
    // The settings for this role. Read-only. Nullable.
    settings PrivilegedRoleSettingsable
    // The summary information for this role. Read-only. Nullable.
    summary PrivilegedRoleSummaryable
}
// NewPrivilegedRole instantiates a new privilegedRole and sets the default values.
func NewPrivilegedRole()(*PrivilegedRole) {
    m := &PrivilegedRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRole(), nil
}
// GetAssignments gets the assignments property value. The assignments for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetAssignments()([]PrivilegedRoleAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrivilegedRoleAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePrivilegedRoleSettingsFromDiscriminatorValue , m.SetSettings)
    res["summary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePrivilegedRoleSummaryFromDiscriminatorValue , m.SetSummary)
    return res
}
// GetName gets the name property value. Role name.
func (m *PrivilegedRole) GetName()(*string) {
    return m.name
}
// GetSettings gets the settings property value. The settings for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetSettings()(PrivilegedRoleSettingsable) {
    return m.settings
}
// GetSummary gets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetSummary()(PrivilegedRoleSummaryable) {
    return m.summary
}
// Serialize serializes information the current object
func (m *PrivilegedRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The assignments for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetAssignments(value []PrivilegedRoleAssignmentable)() {
    m.assignments = value
}
// SetName sets the name property value. Role name.
func (m *PrivilegedRole) SetName(value *string)() {
    m.name = value
}
// SetSettings sets the settings property value. The settings for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetSettings(value PrivilegedRoleSettingsable)() {
    m.settings = value
}
// SetSummary sets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetSummary(value PrivilegedRoleSummaryable)() {
    m.summary = value
}
