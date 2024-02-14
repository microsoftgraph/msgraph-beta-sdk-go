package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UnifiedRbacApplication struct {
    Entity
}
// NewUnifiedRbacApplication instantiates a new UnifiedRbacApplication and sets the default values.
func NewUnifiedRbacApplication()(*UnifiedRbacApplication) {
    m := &UnifiedRbacApplication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRbacApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUnifiedRbacApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRbacApplication(), nil
}
// GetCustomAppScopes gets the customAppScopes property value. Workload-specific scope object that represents the resources for which the principal has been granted access.
// returns a []CustomAppScopeable when successful
func (m *UnifiedRbacApplication) GetCustomAppScopes()([]CustomAppScopeable) {
    val, err := m.GetBackingStore().Get("customAppScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomAppScopeable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UnifiedRbacApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customAppScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomAppScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomAppScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomAppScopeable)
                }
            }
            m.SetCustomAppScopes(res)
        }
        return nil
    }
    res["resourceNamespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRbacResourceNamespaceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRbacResourceNamespaceable)
                }
            }
            m.SetResourceNamespaces(res)
        }
        return nil
    }
    res["roleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleAssignmentable)
                }
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleDefinitionable)
                }
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["transitiveRoleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleAssignmentable)
                }
            }
            m.SetTransitiveRoleAssignments(res)
        }
        return nil
    }
    return res
}
// GetResourceNamespaces gets the resourceNamespaces property value. Resource that represents a collection of related actions.
// returns a []UnifiedRbacResourceNamespaceable when successful
func (m *UnifiedRbacApplication) GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable) {
    val, err := m.GetBackingStore().Get("resourceNamespaces")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRbacResourceNamespaceable)
    }
    return nil
}
// GetRoleAssignments gets the roleAssignments property value. Resource to grant access to users or groups.
// returns a []UnifiedRoleAssignmentable when successful
func (m *UnifiedRbacApplication) GetRoleAssignments()([]UnifiedRoleAssignmentable) {
    val, err := m.GetBackingStore().Get("roleAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleAssignmentable)
    }
    return nil
}
// GetRoleDefinitions gets the roleDefinitions property value. The roles allowed by RBAC providers and the permissions assigned to the roles.
// returns a []UnifiedRoleDefinitionable when successful
func (m *UnifiedRbacApplication) GetRoleDefinitions()([]UnifiedRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleDefinitionable)
    }
    return nil
}
// GetTransitiveRoleAssignments gets the transitiveRoleAssignments property value. Resource to grant access to users or groups that are transitive.
// returns a []UnifiedRoleAssignmentable when successful
func (m *UnifiedRbacApplication) GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable) {
    val, err := m.GetBackingStore().Get("transitiveRoleAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleAssignmentable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRbacApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomAppScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomAppScopes()))
        for i, v := range m.GetCustomAppScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customAppScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceNamespaces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceNamespaces()))
        for i, v := range m.GetResourceNamespaces() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceNamespaces", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveRoleAssignments()))
        for i, v := range m.GetTransitiveRoleAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("transitiveRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomAppScopes sets the customAppScopes property value. Workload-specific scope object that represents the resources for which the principal has been granted access.
func (m *UnifiedRbacApplication) SetCustomAppScopes(value []CustomAppScopeable)() {
    err := m.GetBackingStore().Set("customAppScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceNamespaces sets the resourceNamespaces property value. Resource that represents a collection of related actions.
func (m *UnifiedRbacApplication) SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)() {
    err := m.GetBackingStore().Set("resourceNamespaces", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignments sets the roleAssignments property value. Resource to grant access to users or groups.
func (m *UnifiedRbacApplication) SetRoleAssignments(value []UnifiedRoleAssignmentable)() {
    err := m.GetBackingStore().Set("roleAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. The roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *UnifiedRbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetTransitiveRoleAssignments sets the transitiveRoleAssignments property value. Resource to grant access to users or groups that are transitive.
func (m *UnifiedRbacApplication) SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)() {
    err := m.GetBackingStore().Set("transitiveRoleAssignments", value)
    if err != nil {
        panic(err)
    }
}
type UnifiedRbacApplicationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomAppScopes()([]CustomAppScopeable)
    GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable)
    GetRoleAssignments()([]UnifiedRoleAssignmentable)
    GetRoleDefinitions()([]UnifiedRoleDefinitionable)
    GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable)
    SetCustomAppScopes(value []CustomAppScopeable)()
    SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)()
    SetRoleAssignments(value []UnifiedRoleAssignmentable)()
    SetRoleDefinitions(value []UnifiedRoleDefinitionable)()
    SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)()
}
