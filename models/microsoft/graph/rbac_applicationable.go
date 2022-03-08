package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RbacApplicationable 
type RbacApplicationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
