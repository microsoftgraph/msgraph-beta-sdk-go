// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RbacApplication struct {
    Entity
}
// NewRbacApplication instantiates a new RbacApplication and sets the default values.
func NewRbacApplication()(*RbacApplication) {
    m := &RbacApplication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRbacApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRbacApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRbacApplication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RbacApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["roleAssignmentApprovals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approvalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Approvalable)
                }
            }
            m.SetRoleAssignmentApprovals(res)
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
    res["roleAssignmentScheduleInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentScheduleInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleAssignmentScheduleInstanceable)
                }
            }
            m.SetRoleAssignmentScheduleInstances(res)
        }
        return nil
    }
    res["roleAssignmentScheduleRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentScheduleRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleAssignmentScheduleRequestable)
                }
            }
            m.SetRoleAssignmentScheduleRequests(res)
        }
        return nil
    }
    res["roleAssignmentSchedules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentScheduleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleAssignmentScheduleable)
                }
            }
            m.SetRoleAssignmentSchedules(res)
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
    res["roleEligibilityScheduleInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleEligibilityScheduleInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleEligibilityScheduleInstanceable)
                }
            }
            m.SetRoleEligibilityScheduleInstances(res)
        }
        return nil
    }
    res["roleEligibilityScheduleRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleEligibilityScheduleRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleEligibilityScheduleRequestable)
                }
            }
            m.SetRoleEligibilityScheduleRequests(res)
        }
        return nil
    }
    res["roleEligibilitySchedules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleEligibilityScheduleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleEligibilityScheduleable)
                }
            }
            m.SetRoleEligibilitySchedules(res)
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
// GetResourceNamespaces gets the resourceNamespaces property value. The resourceNamespaces property
// returns a []UnifiedRbacResourceNamespaceable when successful
func (m *RbacApplication) GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable) {
    val, err := m.GetBackingStore().Get("resourceNamespaces")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRbacResourceNamespaceable)
    }
    return nil
}
// GetRoleAssignmentApprovals gets the roleAssignmentApprovals property value. The roleAssignmentApprovals property
// returns a []Approvalable when successful
func (m *RbacApplication) GetRoleAssignmentApprovals()([]Approvalable) {
    val, err := m.GetBackingStore().Get("roleAssignmentApprovals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Approvalable)
    }
    return nil
}
// GetRoleAssignments gets the roleAssignments property value. The roleAssignments property
// returns a []UnifiedRoleAssignmentable when successful
func (m *RbacApplication) GetRoleAssignments()([]UnifiedRoleAssignmentable) {
    val, err := m.GetBackingStore().Get("roleAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleAssignmentable)
    }
    return nil
}
// GetRoleAssignmentScheduleInstances gets the roleAssignmentScheduleInstances property value. The roleAssignmentScheduleInstances property
// returns a []UnifiedRoleAssignmentScheduleInstanceable when successful
func (m *RbacApplication) GetRoleAssignmentScheduleInstances()([]UnifiedRoleAssignmentScheduleInstanceable) {
    val, err := m.GetBackingStore().Get("roleAssignmentScheduleInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleAssignmentScheduleInstanceable)
    }
    return nil
}
// GetRoleAssignmentScheduleRequests gets the roleAssignmentScheduleRequests property value. The roleAssignmentScheduleRequests property
// returns a []UnifiedRoleAssignmentScheduleRequestable when successful
func (m *RbacApplication) GetRoleAssignmentScheduleRequests()([]UnifiedRoleAssignmentScheduleRequestable) {
    val, err := m.GetBackingStore().Get("roleAssignmentScheduleRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleAssignmentScheduleRequestable)
    }
    return nil
}
// GetRoleAssignmentSchedules gets the roleAssignmentSchedules property value. The roleAssignmentSchedules property
// returns a []UnifiedRoleAssignmentScheduleable when successful
func (m *RbacApplication) GetRoleAssignmentSchedules()([]UnifiedRoleAssignmentScheduleable) {
    val, err := m.GetBackingStore().Get("roleAssignmentSchedules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleAssignmentScheduleable)
    }
    return nil
}
// GetRoleDefinitions gets the roleDefinitions property value. The roleDefinitions property
// returns a []UnifiedRoleDefinitionable when successful
func (m *RbacApplication) GetRoleDefinitions()([]UnifiedRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleDefinitionable)
    }
    return nil
}
// GetRoleEligibilityScheduleInstances gets the roleEligibilityScheduleInstances property value. The roleEligibilityScheduleInstances property
// returns a []UnifiedRoleEligibilityScheduleInstanceable when successful
func (m *RbacApplication) GetRoleEligibilityScheduleInstances()([]UnifiedRoleEligibilityScheduleInstanceable) {
    val, err := m.GetBackingStore().Get("roleEligibilityScheduleInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleEligibilityScheduleInstanceable)
    }
    return nil
}
// GetRoleEligibilityScheduleRequests gets the roleEligibilityScheduleRequests property value. The roleEligibilityScheduleRequests property
// returns a []UnifiedRoleEligibilityScheduleRequestable when successful
func (m *RbacApplication) GetRoleEligibilityScheduleRequests()([]UnifiedRoleEligibilityScheduleRequestable) {
    val, err := m.GetBackingStore().Get("roleEligibilityScheduleRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleEligibilityScheduleRequestable)
    }
    return nil
}
// GetRoleEligibilitySchedules gets the roleEligibilitySchedules property value. The roleEligibilitySchedules property
// returns a []UnifiedRoleEligibilityScheduleable when successful
func (m *RbacApplication) GetRoleEligibilitySchedules()([]UnifiedRoleEligibilityScheduleable) {
    val, err := m.GetBackingStore().Get("roleEligibilitySchedules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleEligibilityScheduleable)
    }
    return nil
}
// GetTransitiveRoleAssignments gets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
// returns a []UnifiedRoleAssignmentable when successful
func (m *RbacApplication) GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable) {
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
func (m *RbacApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetRoleAssignmentApprovals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentApprovals()))
        for i, v := range m.GetRoleAssignmentApprovals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentApprovals", cast)
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
    if m.GetRoleAssignmentScheduleInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentScheduleInstances()))
        for i, v := range m.GetRoleAssignmentScheduleInstances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentScheduleRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentScheduleRequests()))
        for i, v := range m.GetRoleAssignmentScheduleRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentSchedules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentSchedules()))
        for i, v := range m.GetRoleAssignmentSchedules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentSchedules", cast)
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
    if m.GetRoleEligibilityScheduleInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleEligibilityScheduleInstances()))
        for i, v := range m.GetRoleEligibilityScheduleInstances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleEligibilityScheduleRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleEligibilityScheduleRequests()))
        for i, v := range m.GetRoleEligibilityScheduleRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleEligibilitySchedules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleEligibilitySchedules()))
        for i, v := range m.GetRoleEligibilitySchedules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilitySchedules", cast)
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
// SetResourceNamespaces sets the resourceNamespaces property value. The resourceNamespaces property
func (m *RbacApplication) SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)() {
    err := m.GetBackingStore().Set("resourceNamespaces", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignmentApprovals sets the roleAssignmentApprovals property value. The roleAssignmentApprovals property
func (m *RbacApplication) SetRoleAssignmentApprovals(value []Approvalable)() {
    err := m.GetBackingStore().Set("roleAssignmentApprovals", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignments sets the roleAssignments property value. The roleAssignments property
func (m *RbacApplication) SetRoleAssignments(value []UnifiedRoleAssignmentable)() {
    err := m.GetBackingStore().Set("roleAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignmentScheduleInstances sets the roleAssignmentScheduleInstances property value. The roleAssignmentScheduleInstances property
func (m *RbacApplication) SetRoleAssignmentScheduleInstances(value []UnifiedRoleAssignmentScheduleInstanceable)() {
    err := m.GetBackingStore().Set("roleAssignmentScheduleInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignmentScheduleRequests sets the roleAssignmentScheduleRequests property value. The roleAssignmentScheduleRequests property
func (m *RbacApplication) SetRoleAssignmentScheduleRequests(value []UnifiedRoleAssignmentScheduleRequestable)() {
    err := m.GetBackingStore().Set("roleAssignmentScheduleRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignmentSchedules sets the roleAssignmentSchedules property value. The roleAssignmentSchedules property
func (m *RbacApplication) SetRoleAssignmentSchedules(value []UnifiedRoleAssignmentScheduleable)() {
    err := m.GetBackingStore().Set("roleAssignmentSchedules", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. The roleDefinitions property
func (m *RbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleEligibilityScheduleInstances sets the roleEligibilityScheduleInstances property value. The roleEligibilityScheduleInstances property
func (m *RbacApplication) SetRoleEligibilityScheduleInstances(value []UnifiedRoleEligibilityScheduleInstanceable)() {
    err := m.GetBackingStore().Set("roleEligibilityScheduleInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleEligibilityScheduleRequests sets the roleEligibilityScheduleRequests property value. The roleEligibilityScheduleRequests property
func (m *RbacApplication) SetRoleEligibilityScheduleRequests(value []UnifiedRoleEligibilityScheduleRequestable)() {
    err := m.GetBackingStore().Set("roleEligibilityScheduleRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleEligibilitySchedules sets the roleEligibilitySchedules property value. The roleEligibilitySchedules property
func (m *RbacApplication) SetRoleEligibilitySchedules(value []UnifiedRoleEligibilityScheduleable)() {
    err := m.GetBackingStore().Set("roleEligibilitySchedules", value)
    if err != nil {
        panic(err)
    }
}
// SetTransitiveRoleAssignments sets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
func (m *RbacApplication) SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)() {
    err := m.GetBackingStore().Set("transitiveRoleAssignments", value)
    if err != nil {
        panic(err)
    }
}
type RbacApplicationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable)
    GetRoleAssignmentApprovals()([]Approvalable)
    GetRoleAssignments()([]UnifiedRoleAssignmentable)
    GetRoleAssignmentScheduleInstances()([]UnifiedRoleAssignmentScheduleInstanceable)
    GetRoleAssignmentScheduleRequests()([]UnifiedRoleAssignmentScheduleRequestable)
    GetRoleAssignmentSchedules()([]UnifiedRoleAssignmentScheduleable)
    GetRoleDefinitions()([]UnifiedRoleDefinitionable)
    GetRoleEligibilityScheduleInstances()([]UnifiedRoleEligibilityScheduleInstanceable)
    GetRoleEligibilityScheduleRequests()([]UnifiedRoleEligibilityScheduleRequestable)
    GetRoleEligibilitySchedules()([]UnifiedRoleEligibilityScheduleable)
    GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable)
    SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)()
    SetRoleAssignmentApprovals(value []Approvalable)()
    SetRoleAssignments(value []UnifiedRoleAssignmentable)()
    SetRoleAssignmentScheduleInstances(value []UnifiedRoleAssignmentScheduleInstanceable)()
    SetRoleAssignmentScheduleRequests(value []UnifiedRoleAssignmentScheduleRequestable)()
    SetRoleAssignmentSchedules(value []UnifiedRoleAssignmentScheduleable)()
    SetRoleDefinitions(value []UnifiedRoleDefinitionable)()
    SetRoleEligibilityScheduleInstances(value []UnifiedRoleEligibilityScheduleInstanceable)()
    SetRoleEligibilityScheduleRequests(value []UnifiedRoleEligibilityScheduleRequestable)()
    SetRoleEligibilitySchedules(value []UnifiedRoleEligibilityScheduleable)()
    SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)()
}
