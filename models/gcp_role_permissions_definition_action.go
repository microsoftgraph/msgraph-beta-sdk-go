package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GcpRolePermissionsDefinitionAction struct {
    GcpPermissionsDefinitionAction
}
// NewGcpRolePermissionsDefinitionAction instantiates a new GcpRolePermissionsDefinitionAction and sets the default values.
func NewGcpRolePermissionsDefinitionAction()(*GcpRolePermissionsDefinitionAction) {
    m := &GcpRolePermissionsDefinitionAction{
        GcpPermissionsDefinitionAction: *NewGcpPermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.gcpRolePermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpRolePermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGcpRolePermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpRolePermissionsDefinitionAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GcpRolePermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GcpPermissionsDefinitionAction.GetFieldDeserializers()
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionsDefinitionGcpRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionsDefinitionGcpRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionsDefinitionGcpRoleable)
                }
            }
            m.SetRoles(res)
        }
        return nil
    }
    return res
}
// GetRoles gets the roles property value. The roles property
// returns a []PermissionsDefinitionGcpRoleable when successful
func (m *GcpRolePermissionsDefinitionAction) GetRoles()([]PermissionsDefinitionGcpRoleable) {
    val, err := m.GetBackingStore().Get("roles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PermissionsDefinitionGcpRoleable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GcpRolePermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GcpPermissionsDefinitionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoles()))
        for i, v := range m.GetRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roles", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRoles sets the roles property value. The roles property
func (m *GcpRolePermissionsDefinitionAction) SetRoles(value []PermissionsDefinitionGcpRoleable)() {
    err := m.GetBackingStore().Set("roles", value)
    if err != nil {
        panic(err)
    }
}
type GcpRolePermissionsDefinitionActionable interface {
    GcpPermissionsDefinitionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoles()([]PermissionsDefinitionGcpRoleable)
    SetRoles(value []PermissionsDefinitionGcpRoleable)()
}
