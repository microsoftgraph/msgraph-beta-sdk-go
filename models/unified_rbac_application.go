package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRbacApplication 
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
func CreateUnifiedRbacApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRbacApplication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["resourceNamespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRbacResourceNamespaceable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRbacResourceNamespaceable)
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
                res[i] = v.(UnifiedRoleAssignmentable)
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
                res[i] = v.(UnifiedRoleDefinitionable)
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
                res[i] = v.(UnifiedRoleAssignmentable)
            }
            m.SetTransitiveRoleAssignments(res)
        }
        return nil
    }
    return res
}
// GetResourceNamespaces gets the resourceNamespaces property value. The resourceNamespaces property
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
// GetRoleAssignments gets the roleAssignments property value. The roleAssignments property
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
// GetRoleDefinitions gets the roleDefinitions property value. The roleDefinitions property
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
// GetTransitiveRoleAssignments gets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
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
    if m.GetResourceNamespaces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceNamespaces()))
        for i, v := range m.GetResourceNamespaces() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceNamespaces", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveRoleAssignments()))
        for i, v := range m.GetTransitiveRoleAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceNamespaces sets the resourceNamespaces property value. The resourceNamespaces property
func (m *UnifiedRbacApplication) SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)() {
    err := m.GetBackingStore().Set("resourceNamespaces", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignments sets the roleAssignments property value. The roleAssignments property
func (m *UnifiedRbacApplication) SetRoleAssignments(value []UnifiedRoleAssignmentable)() {
    err := m.GetBackingStore().Set("roleAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. The roleDefinitions property
func (m *UnifiedRbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetTransitiveRoleAssignments sets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
func (m *UnifiedRbacApplication) SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)() {
    err := m.GetBackingStore().Set("transitiveRoleAssignments", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRbacApplicationable 
type UnifiedRbacApplicationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable)
    GetRoleAssignments()([]UnifiedRoleAssignmentable)
    GetRoleDefinitions()([]UnifiedRoleDefinitionable)
    GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable)
    SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)()
    SetRoleAssignments(value []UnifiedRoleAssignmentable)()
    SetRoleDefinitions(value []UnifiedRoleDefinitionable)()
    SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)()
}
