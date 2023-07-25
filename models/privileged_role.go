package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRole 
type PrivilegedRole struct {
    Entity
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
// GetAssignments gets the assignments property value. The assignments property
func (m *PrivilegedRole) GetAssignments()([]PrivilegedRoleAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegedRoleAssignmentable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedRoleAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrivilegedRoleAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(PrivilegedRoleSettingsable))
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val.(PrivilegedRoleSummaryable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *PrivilegedRole) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettings gets the settings property value. The settings property
func (m *PrivilegedRole) GetSettings()(PrivilegedRoleSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivilegedRoleSettingsable)
    }
    return nil
}
// GetSummary gets the summary property value. The summary property
func (m *PrivilegedRole) GetSummary()(PrivilegedRoleSummaryable) {
    val, err := m.GetBackingStore().Get("summary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivilegedRoleSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegedRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
// SetAssignments sets the assignments property value. The assignments property
func (m *PrivilegedRole) SetAssignments(value []PrivilegedRoleAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *PrivilegedRole) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings property
func (m *PrivilegedRole) SetSettings(value PrivilegedRoleSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetSummary sets the summary property value. The summary property
func (m *PrivilegedRole) SetSummary(value PrivilegedRoleSummaryable)() {
    err := m.GetBackingStore().Set("summary", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegedRoleable 
type PrivilegedRoleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]PrivilegedRoleAssignmentable)
    GetName()(*string)
    GetSettings()(PrivilegedRoleSettingsable)
    GetSummary()(PrivilegedRoleSummaryable)
    SetAssignments(value []PrivilegedRoleAssignmentable)()
    SetName(value *string)()
    SetSettings(value PrivilegedRoleSettingsable)()
    SetSummary(value PrivilegedRoleSummaryable)()
}
