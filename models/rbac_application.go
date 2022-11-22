package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RbacApplication 
type RbacApplication struct {
    Entity
    // The resourceNamespaces property
    resourceNamespaces []UnifiedRbacResourceNamespaceable
    // The roleAssignmentApprovals property
    roleAssignmentApprovals []Approvalable
    // The roleAssignments property
    roleAssignments []UnifiedRoleAssignmentable
    // The roleAssignmentScheduleInstances property
    roleAssignmentScheduleInstances []UnifiedRoleAssignmentScheduleInstanceable
    // The roleAssignmentScheduleRequests property
    roleAssignmentScheduleRequests []UnifiedRoleAssignmentScheduleRequestable
    // The roleAssignmentSchedules property
    roleAssignmentSchedules []UnifiedRoleAssignmentScheduleable
    // The roleDefinitions property
    roleDefinitions []UnifiedRoleDefinitionable
    // The roleEligibilityScheduleInstances property
    roleEligibilityScheduleInstances []UnifiedRoleEligibilityScheduleInstanceable
    // The roleEligibilityScheduleRequests property
    roleEligibilityScheduleRequests []UnifiedRoleEligibilityScheduleRequestable
    // The roleEligibilitySchedules property
    roleEligibilitySchedules []UnifiedRoleEligibilityScheduleable
    // The transitiveRoleAssignments property
    transitiveRoleAssignments []UnifiedRoleAssignmentable
}
// NewRbacApplication instantiates a new RbacApplication and sets the default values.
func NewRbacApplication()(*RbacApplication) {
    m := &RbacApplication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRbacApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRbacApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRbacApplication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RbacApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["resourceNamespaces"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue , m.SetResourceNamespaces)
    res["roleAssignmentApprovals"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue , m.SetRoleAssignmentApprovals)
    res["roleAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleAssignmentFromDiscriminatorValue , m.SetRoleAssignments)
    res["roleAssignmentScheduleInstances"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue , m.SetRoleAssignmentScheduleInstances)
    res["roleAssignmentScheduleRequests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue , m.SetRoleAssignmentScheduleRequests)
    res["roleAssignmentSchedules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue , m.SetRoleAssignmentSchedules)
    res["roleDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleDefinitionFromDiscriminatorValue , m.SetRoleDefinitions)
    res["roleEligibilityScheduleInstances"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue , m.SetRoleEligibilityScheduleInstances)
    res["roleEligibilityScheduleRequests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue , m.SetRoleEligibilityScheduleRequests)
    res["roleEligibilitySchedules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue , m.SetRoleEligibilitySchedules)
    res["transitiveRoleAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleAssignmentFromDiscriminatorValue , m.SetTransitiveRoleAssignments)
    return res
}
// GetResourceNamespaces gets the resourceNamespaces property value. The resourceNamespaces property
func (m *RbacApplication) GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable) {
    return m.resourceNamespaces
}
// GetRoleAssignmentApprovals gets the roleAssignmentApprovals property value. The roleAssignmentApprovals property
func (m *RbacApplication) GetRoleAssignmentApprovals()([]Approvalable) {
    return m.roleAssignmentApprovals
}
// GetRoleAssignments gets the roleAssignments property value. The roleAssignments property
func (m *RbacApplication) GetRoleAssignments()([]UnifiedRoleAssignmentable) {
    return m.roleAssignments
}
// GetRoleAssignmentScheduleInstances gets the roleAssignmentScheduleInstances property value. The roleAssignmentScheduleInstances property
func (m *RbacApplication) GetRoleAssignmentScheduleInstances()([]UnifiedRoleAssignmentScheduleInstanceable) {
    return m.roleAssignmentScheduleInstances
}
// GetRoleAssignmentScheduleRequests gets the roleAssignmentScheduleRequests property value. The roleAssignmentScheduleRequests property
func (m *RbacApplication) GetRoleAssignmentScheduleRequests()([]UnifiedRoleAssignmentScheduleRequestable) {
    return m.roleAssignmentScheduleRequests
}
// GetRoleAssignmentSchedules gets the roleAssignmentSchedules property value. The roleAssignmentSchedules property
func (m *RbacApplication) GetRoleAssignmentSchedules()([]UnifiedRoleAssignmentScheduleable) {
    return m.roleAssignmentSchedules
}
// GetRoleDefinitions gets the roleDefinitions property value. The roleDefinitions property
func (m *RbacApplication) GetRoleDefinitions()([]UnifiedRoleDefinitionable) {
    return m.roleDefinitions
}
// GetRoleEligibilityScheduleInstances gets the roleEligibilityScheduleInstances property value. The roleEligibilityScheduleInstances property
func (m *RbacApplication) GetRoleEligibilityScheduleInstances()([]UnifiedRoleEligibilityScheduleInstanceable) {
    return m.roleEligibilityScheduleInstances
}
// GetRoleEligibilityScheduleRequests gets the roleEligibilityScheduleRequests property value. The roleEligibilityScheduleRequests property
func (m *RbacApplication) GetRoleEligibilityScheduleRequests()([]UnifiedRoleEligibilityScheduleRequestable) {
    return m.roleEligibilityScheduleRequests
}
// GetRoleEligibilitySchedules gets the roleEligibilitySchedules property value. The roleEligibilitySchedules property
func (m *RbacApplication) GetRoleEligibilitySchedules()([]UnifiedRoleEligibilityScheduleable) {
    return m.roleEligibilitySchedules
}
// GetTransitiveRoleAssignments gets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
func (m *RbacApplication) GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable) {
    return m.transitiveRoleAssignments
}
// Serialize serializes information the current object
func (m *RbacApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetResourceNamespaces() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetResourceNamespaces())
        err = writer.WriteCollectionOfObjectValues("resourceNamespaces", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentApprovals() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignmentApprovals())
        err = writer.WriteCollectionOfObjectValues("roleAssignmentApprovals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignments())
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentScheduleInstances() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignmentScheduleInstances())
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentScheduleRequests() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignmentScheduleRequests())
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentSchedules() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignmentSchedules())
        err = writer.WriteCollectionOfObjectValues("roleAssignmentSchedules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleDefinitions())
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleEligibilityScheduleInstances() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleEligibilityScheduleInstances())
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleEligibilityScheduleRequests() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleEligibilityScheduleRequests())
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleEligibilitySchedules() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleEligibilitySchedules())
        err = writer.WriteCollectionOfObjectValues("roleEligibilitySchedules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveRoleAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTransitiveRoleAssignments())
        err = writer.WriteCollectionOfObjectValues("transitiveRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceNamespaces sets the resourceNamespaces property value. The resourceNamespaces property
func (m *RbacApplication) SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)() {
    m.resourceNamespaces = value
}
// SetRoleAssignmentApprovals sets the roleAssignmentApprovals property value. The roleAssignmentApprovals property
func (m *RbacApplication) SetRoleAssignmentApprovals(value []Approvalable)() {
    m.roleAssignmentApprovals = value
}
// SetRoleAssignments sets the roleAssignments property value. The roleAssignments property
func (m *RbacApplication) SetRoleAssignments(value []UnifiedRoleAssignmentable)() {
    m.roleAssignments = value
}
// SetRoleAssignmentScheduleInstances sets the roleAssignmentScheduleInstances property value. The roleAssignmentScheduleInstances property
func (m *RbacApplication) SetRoleAssignmentScheduleInstances(value []UnifiedRoleAssignmentScheduleInstanceable)() {
    m.roleAssignmentScheduleInstances = value
}
// SetRoleAssignmentScheduleRequests sets the roleAssignmentScheduleRequests property value. The roleAssignmentScheduleRequests property
func (m *RbacApplication) SetRoleAssignmentScheduleRequests(value []UnifiedRoleAssignmentScheduleRequestable)() {
    m.roleAssignmentScheduleRequests = value
}
// SetRoleAssignmentSchedules sets the roleAssignmentSchedules property value. The roleAssignmentSchedules property
func (m *RbacApplication) SetRoleAssignmentSchedules(value []UnifiedRoleAssignmentScheduleable)() {
    m.roleAssignmentSchedules = value
}
// SetRoleDefinitions sets the roleDefinitions property value. The roleDefinitions property
func (m *RbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    m.roleDefinitions = value
}
// SetRoleEligibilityScheduleInstances sets the roleEligibilityScheduleInstances property value. The roleEligibilityScheduleInstances property
func (m *RbacApplication) SetRoleEligibilityScheduleInstances(value []UnifiedRoleEligibilityScheduleInstanceable)() {
    m.roleEligibilityScheduleInstances = value
}
// SetRoleEligibilityScheduleRequests sets the roleEligibilityScheduleRequests property value. The roleEligibilityScheduleRequests property
func (m *RbacApplication) SetRoleEligibilityScheduleRequests(value []UnifiedRoleEligibilityScheduleRequestable)() {
    m.roleEligibilityScheduleRequests = value
}
// SetRoleEligibilitySchedules sets the roleEligibilitySchedules property value. The roleEligibilitySchedules property
func (m *RbacApplication) SetRoleEligibilitySchedules(value []UnifiedRoleEligibilityScheduleable)() {
    m.roleEligibilitySchedules = value
}
// SetTransitiveRoleAssignments sets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
func (m *RbacApplication) SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)() {
    m.transitiveRoleAssignments = value
}
