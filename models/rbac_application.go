package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RbacApplication provides operations to manage the roleManagement singleton.
type RbacApplication struct {
    Entity
    // The resourceNamespaces property
    resourceNamespaces []UnifiedRbacResourceNamespaceable
    // The roleAssignmentApprovals property
    roleAssignmentApprovals []Approvalable
    // Resource to grant access to users or groups.
    roleAssignments []UnifiedRoleAssignmentable
    // Instances for active role assignments.
    roleAssignmentScheduleInstances []UnifiedRoleAssignmentScheduleInstanceable
    // Requests for active role assignments to principals through PIM.
    roleAssignmentScheduleRequests []UnifiedRoleAssignmentScheduleRequestable
    // Schedules for active role assignment operations.
    roleAssignmentSchedules []UnifiedRoleAssignmentScheduleable
    // Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
    roleDefinitions []UnifiedRoleDefinitionable
    // Instances for role eligibility requests.
    roleEligibilityScheduleInstances []UnifiedRoleEligibilityScheduleInstanceable
    // Requests for role eligibilities for principals through PIM.
    roleEligibilityScheduleRequests []UnifiedRoleEligibilityScheduleRequestable
    // Schedules for role eligibility operations.
    roleEligibilitySchedules []UnifiedRoleEligibilityScheduleable
    // The transitiveRoleAssignments property
    transitiveRoleAssignments []UnifiedRoleAssignmentable
}
// NewRbacApplication instantiates a new rbacApplication and sets the default values.
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
    res["roleAssignmentApprovals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approvalable, len(val))
            for i, v := range val {
                res[i] = v.(Approvalable)
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
                res[i] = v.(UnifiedRoleAssignmentable)
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
                res[i] = v.(UnifiedRoleAssignmentScheduleInstanceable)
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
                res[i] = v.(UnifiedRoleAssignmentScheduleRequestable)
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
                res[i] = v.(UnifiedRoleAssignmentScheduleable)
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
                res[i] = v.(UnifiedRoleDefinitionable)
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
                res[i] = v.(UnifiedRoleEligibilityScheduleInstanceable)
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
                res[i] = v.(UnifiedRoleEligibilityScheduleRequestable)
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
                res[i] = v.(UnifiedRoleEligibilityScheduleable)
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
                res[i] = v.(UnifiedRoleAssignmentable)
            }
            m.SetTransitiveRoleAssignments(res)
        }
        return nil
    }
    return res
}
// GetResourceNamespaces gets the resourceNamespaces property value. The resourceNamespaces property
func (m *RbacApplication) GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable) {
    if m == nil {
        return nil
    } else {
        return m.resourceNamespaces
    }
}
// GetRoleAssignmentApprovals gets the roleAssignmentApprovals property value. The roleAssignmentApprovals property
func (m *RbacApplication) GetRoleAssignmentApprovals()([]Approvalable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentApprovals
    }
}
// GetRoleAssignments gets the roleAssignments property value. Resource to grant access to users or groups.
func (m *RbacApplication) GetRoleAssignments()([]UnifiedRoleAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleAssignmentScheduleInstances gets the roleAssignmentScheduleInstances property value. Instances for active role assignments.
func (m *RbacApplication) GetRoleAssignmentScheduleInstances()([]UnifiedRoleAssignmentScheduleInstanceable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleInstances
    }
}
// GetRoleAssignmentScheduleRequests gets the roleAssignmentScheduleRequests property value. Requests for active role assignments to principals through PIM.
func (m *RbacApplication) GetRoleAssignmentScheduleRequests()([]UnifiedRoleAssignmentScheduleRequestable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleRequests
    }
}
// GetRoleAssignmentSchedules gets the roleAssignmentSchedules property value. Schedules for active role assignment operations.
func (m *RbacApplication) GetRoleAssignmentSchedules()([]UnifiedRoleAssignmentScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentSchedules
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *RbacApplication) GetRoleDefinitions()([]UnifiedRoleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// GetRoleEligibilityScheduleInstances gets the roleEligibilityScheduleInstances property value. Instances for role eligibility requests.
func (m *RbacApplication) GetRoleEligibilityScheduleInstances()([]UnifiedRoleEligibilityScheduleInstanceable) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilityScheduleInstances
    }
}
// GetRoleEligibilityScheduleRequests gets the roleEligibilityScheduleRequests property value. Requests for role eligibilities for principals through PIM.
func (m *RbacApplication) GetRoleEligibilityScheduleRequests()([]UnifiedRoleEligibilityScheduleRequestable) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilityScheduleRequests
    }
}
// GetRoleEligibilitySchedules gets the roleEligibilitySchedules property value. Schedules for role eligibility operations.
func (m *RbacApplication) GetRoleEligibilitySchedules()([]UnifiedRoleEligibilityScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilitySchedules
    }
}
// GetTransitiveRoleAssignments gets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
func (m *RbacApplication) GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.transitiveRoleAssignments
    }
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceNamespaces", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentApprovals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentApprovals()))
        for i, v := range m.GetRoleAssignmentApprovals() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentApprovals", cast)
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
    if m.GetRoleAssignmentScheduleInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentScheduleInstances()))
        for i, v := range m.GetRoleAssignmentScheduleInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentScheduleRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentScheduleRequests()))
        for i, v := range m.GetRoleAssignmentScheduleRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentSchedules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentSchedules()))
        for i, v := range m.GetRoleAssignmentSchedules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentSchedules", cast)
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
    if m.GetRoleEligibilityScheduleInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleEligibilityScheduleInstances()))
        for i, v := range m.GetRoleEligibilityScheduleInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleEligibilityScheduleRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleEligibilityScheduleRequests()))
        for i, v := range m.GetRoleEligibilityScheduleRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleEligibilitySchedules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleEligibilitySchedules()))
        for i, v := range m.GetRoleEligibilitySchedules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilitySchedules", cast)
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
func (m *RbacApplication) SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)() {
    if m != nil {
        m.resourceNamespaces = value
    }
}
// SetRoleAssignmentApprovals sets the roleAssignmentApprovals property value. The roleAssignmentApprovals property
func (m *RbacApplication) SetRoleAssignmentApprovals(value []Approvalable)() {
    if m != nil {
        m.roleAssignmentApprovals = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. Resource to grant access to users or groups.
func (m *RbacApplication) SetRoleAssignments(value []UnifiedRoleAssignmentable)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleAssignmentScheduleInstances sets the roleAssignmentScheduleInstances property value. Instances for active role assignments.
func (m *RbacApplication) SetRoleAssignmentScheduleInstances(value []UnifiedRoleAssignmentScheduleInstanceable)() {
    if m != nil {
        m.roleAssignmentScheduleInstances = value
    }
}
// SetRoleAssignmentScheduleRequests sets the roleAssignmentScheduleRequests property value. Requests for active role assignments to principals through PIM.
func (m *RbacApplication) SetRoleAssignmentScheduleRequests(value []UnifiedRoleAssignmentScheduleRequestable)() {
    if m != nil {
        m.roleAssignmentScheduleRequests = value
    }
}
// SetRoleAssignmentSchedules sets the roleAssignmentSchedules property value. Schedules for active role assignment operations.
func (m *RbacApplication) SetRoleAssignmentSchedules(value []UnifiedRoleAssignmentScheduleable)() {
    if m != nil {
        m.roleAssignmentSchedules = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *RbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
// SetRoleEligibilityScheduleInstances sets the roleEligibilityScheduleInstances property value. Instances for role eligibility requests.
func (m *RbacApplication) SetRoleEligibilityScheduleInstances(value []UnifiedRoleEligibilityScheduleInstanceable)() {
    if m != nil {
        m.roleEligibilityScheduleInstances = value
    }
}
// SetRoleEligibilityScheduleRequests sets the roleEligibilityScheduleRequests property value. Requests for role eligibilities for principals through PIM.
func (m *RbacApplication) SetRoleEligibilityScheduleRequests(value []UnifiedRoleEligibilityScheduleRequestable)() {
    if m != nil {
        m.roleEligibilityScheduleRequests = value
    }
}
// SetRoleEligibilitySchedules sets the roleEligibilitySchedules property value. Schedules for role eligibility operations.
func (m *RbacApplication) SetRoleEligibilitySchedules(value []UnifiedRoleEligibilityScheduleable)() {
    if m != nil {
        m.roleEligibilitySchedules = value
    }
}
// SetTransitiveRoleAssignments sets the transitiveRoleAssignments property value. The transitiveRoleAssignments property
func (m *RbacApplication) SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)() {
    if m != nil {
        m.transitiveRoleAssignments = value
    }
}
