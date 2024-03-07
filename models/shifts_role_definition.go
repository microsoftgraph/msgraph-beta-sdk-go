package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ShiftsRoleDefinition struct {
    Entity
}
// NewShiftsRoleDefinition instantiates a new ShiftsRoleDefinition and sets the default values.
func NewShiftsRoleDefinition()(*ShiftsRoleDefinition) {
    m := &ShiftsRoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateShiftsRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateShiftsRoleDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShiftsRoleDefinition(), nil
}
// GetDescription gets the description property value. The description of the role.
// returns a *string when successful
func (m *ShiftsRoleDefinition) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the role.
// returns a *string when successful
func (m *ShiftsRoleDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ShiftsRoleDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["shiftsRolePermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateShiftsRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ShiftsRolePermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ShiftsRolePermissionable)
                }
            }
            m.SetShiftsRolePermissions(res)
        }
        return nil
    }
    return res
}
// GetShiftsRolePermissions gets the shiftsRolePermissions property value. The collection of role permissions within the role.
// returns a []ShiftsRolePermissionable when successful
func (m *ShiftsRoleDefinition) GetShiftsRolePermissions()([]ShiftsRolePermissionable) {
    val, err := m.GetBackingStore().Get("shiftsRolePermissions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ShiftsRolePermissionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ShiftsRoleDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetShiftsRolePermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShiftsRolePermissions()))
        for i, v := range m.GetShiftsRolePermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("shiftsRolePermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the role.
func (m *ShiftsRoleDefinition) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the role.
func (m *ShiftsRoleDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetShiftsRolePermissions sets the shiftsRolePermissions property value. The collection of role permissions within the role.
func (m *ShiftsRoleDefinition) SetShiftsRolePermissions(value []ShiftsRolePermissionable)() {
    err := m.GetBackingStore().Set("shiftsRolePermissions", value)
    if err != nil {
        panic(err)
    }
}
type ShiftsRoleDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetShiftsRolePermissions()([]ShiftsRolePermissionable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetShiftsRolePermissions(value []ShiftsRolePermissionable)()
}
