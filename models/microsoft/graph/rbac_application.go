package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RbacApplication struct {
    Entity
    // 
    resourceNamespaces []UnifiedRbacResourceNamespace;
    // 
    roleAssignmentApprovals []Approval;
    // Resource to grant access to users or groups.
    roleAssignments []UnifiedRoleAssignment;
    // 
    roleAssignmentScheduleInstances []UnifiedRoleAssignmentScheduleInstance;
    // 
    roleAssignmentScheduleRequests []UnifiedRoleAssignmentScheduleRequest;
    // 
    roleAssignmentSchedules []UnifiedRoleAssignmentSchedule;
    // Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
    roleDefinitions []UnifiedRoleDefinition;
    // 
    roleEligibilityScheduleInstances []UnifiedRoleEligibilityScheduleInstance;
    // 
    roleEligibilityScheduleRequests []UnifiedRoleEligibilityScheduleRequest;
    // 
    roleEligibilitySchedules []UnifiedRoleEligibilitySchedule;
}
// Instantiates a new rbacApplication and sets the default values.
func NewRbacApplication()(*RbacApplication) {
    m := &RbacApplication{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the resourceNamespaces property value. 
func (m *RbacApplication) GetResourceNamespaces()([]UnifiedRbacResourceNamespace) {
    if m == nil {
        return nil
    } else {
        return m.resourceNamespaces
    }
}
// Gets the roleAssignmentApprovals property value. 
func (m *RbacApplication) GetRoleAssignmentApprovals()([]Approval) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentApprovals
    }
}
// Gets the roleAssignments property value. Resource to grant access to users or groups.
func (m *RbacApplication) GetRoleAssignments()([]UnifiedRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// Gets the roleAssignmentScheduleInstances property value. 
func (m *RbacApplication) GetRoleAssignmentScheduleInstances()([]UnifiedRoleAssignmentScheduleInstance) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleInstances
    }
}
// Gets the roleAssignmentScheduleRequests property value. 
func (m *RbacApplication) GetRoleAssignmentScheduleRequests()([]UnifiedRoleAssignmentScheduleRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleRequests
    }
}
// Gets the roleAssignmentSchedules property value. 
func (m *RbacApplication) GetRoleAssignmentSchedules()([]UnifiedRoleAssignmentSchedule) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentSchedules
    }
}
// Gets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *RbacApplication) GetRoleDefinitions()([]UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// Gets the roleEligibilityScheduleInstances property value. 
func (m *RbacApplication) GetRoleEligibilityScheduleInstances()([]UnifiedRoleEligibilityScheduleInstance) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilityScheduleInstances
    }
}
// Gets the roleEligibilityScheduleRequests property value. 
func (m *RbacApplication) GetRoleEligibilityScheduleRequests()([]UnifiedRoleEligibilityScheduleRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilityScheduleRequests
    }
}
// Gets the roleEligibilitySchedules property value. 
func (m *RbacApplication) GetRoleEligibilitySchedules()([]UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilitySchedules
    }
}
// The deserialization information for the current model
func (m *RbacApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["resourceNamespaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRbacResourceNamespace() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRbacResourceNamespace, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRbacResourceNamespace))
            }
            m.SetResourceNamespaces(res)
        }
        return nil
    }
    res["roleAssignmentApprovals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApproval() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approval, len(val))
            for i, v := range val {
                res[i] = *(v.(*Approval))
            }
            m.SetRoleAssignmentApprovals(res)
        }
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleAssignment))
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleAssignmentScheduleInstances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleAssignmentScheduleInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentScheduleInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleAssignmentScheduleInstance))
            }
            m.SetRoleAssignmentScheduleInstances(res)
        }
        return nil
    }
    res["roleAssignmentScheduleRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleAssignmentScheduleRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentScheduleRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleAssignmentScheduleRequest))
            }
            m.SetRoleAssignmentScheduleRequests(res)
        }
        return nil
    }
    res["roleAssignmentSchedules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleAssignmentSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentSchedule, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleAssignmentSchedule))
            }
            m.SetRoleAssignmentSchedules(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleDefinition))
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleEligibilityScheduleInstances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilityScheduleInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleEligibilityScheduleInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleEligibilityScheduleInstance))
            }
            m.SetRoleEligibilityScheduleInstances(res)
        }
        return nil
    }
    res["roleEligibilityScheduleRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilityScheduleRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleEligibilityScheduleRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleEligibilityScheduleRequest))
            }
            m.SetRoleEligibilityScheduleRequests(res)
        }
        return nil
    }
    res["roleEligibilitySchedules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilitySchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleEligibilitySchedule, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleEligibilitySchedule))
            }
            m.SetRoleEligibilitySchedules(res)
        }
        return nil
    }
    return res
}
func (m *RbacApplication) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RbacApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceNamespaces()))
        for i, v := range m.GetResourceNamespaces() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resourceNamespaces", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignmentApprovals()))
        for i, v := range m.GetRoleAssignmentApprovals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentApprovals", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignmentScheduleInstances()))
        for i, v := range m.GetRoleAssignmentScheduleInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignmentScheduleRequests()))
        for i, v := range m.GetRoleAssignmentScheduleRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignmentSchedules()))
        for i, v := range m.GetRoleAssignmentSchedules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentSchedules", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleEligibilityScheduleInstances()))
        for i, v := range m.GetRoleEligibilityScheduleInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleEligibilityScheduleRequests()))
        for i, v := range m.GetRoleEligibilityScheduleRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilityScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleEligibilitySchedules()))
        for i, v := range m.GetRoleEligibilitySchedules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleEligibilitySchedules", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the resourceNamespaces property value. 
// Parameters:
//  - value : Value to set for the resourceNamespaces property.
func (m *RbacApplication) SetResourceNamespaces(value []UnifiedRbacResourceNamespace)() {
    m.resourceNamespaces = value
}
// Sets the roleAssignmentApprovals property value. 
// Parameters:
//  - value : Value to set for the roleAssignmentApprovals property.
func (m *RbacApplication) SetRoleAssignmentApprovals(value []Approval)() {
    m.roleAssignmentApprovals = value
}
// Sets the roleAssignments property value. Resource to grant access to users or groups.
// Parameters:
//  - value : Value to set for the roleAssignments property.
func (m *RbacApplication) SetRoleAssignments(value []UnifiedRoleAssignment)() {
    m.roleAssignments = value
}
// Sets the roleAssignmentScheduleInstances property value. 
// Parameters:
//  - value : Value to set for the roleAssignmentScheduleInstances property.
func (m *RbacApplication) SetRoleAssignmentScheduleInstances(value []UnifiedRoleAssignmentScheduleInstance)() {
    m.roleAssignmentScheduleInstances = value
}
// Sets the roleAssignmentScheduleRequests property value. 
// Parameters:
//  - value : Value to set for the roleAssignmentScheduleRequests property.
func (m *RbacApplication) SetRoleAssignmentScheduleRequests(value []UnifiedRoleAssignmentScheduleRequest)() {
    m.roleAssignmentScheduleRequests = value
}
// Sets the roleAssignmentSchedules property value. 
// Parameters:
//  - value : Value to set for the roleAssignmentSchedules property.
func (m *RbacApplication) SetRoleAssignmentSchedules(value []UnifiedRoleAssignmentSchedule)() {
    m.roleAssignmentSchedules = value
}
// Sets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
// Parameters:
//  - value : Value to set for the roleDefinitions property.
func (m *RbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinition)() {
    m.roleDefinitions = value
}
// Sets the roleEligibilityScheduleInstances property value. 
// Parameters:
//  - value : Value to set for the roleEligibilityScheduleInstances property.
func (m *RbacApplication) SetRoleEligibilityScheduleInstances(value []UnifiedRoleEligibilityScheduleInstance)() {
    m.roleEligibilityScheduleInstances = value
}
// Sets the roleEligibilityScheduleRequests property value. 
// Parameters:
//  - value : Value to set for the roleEligibilityScheduleRequests property.
func (m *RbacApplication) SetRoleEligibilityScheduleRequests(value []UnifiedRoleEligibilityScheduleRequest)() {
    m.roleEligibilityScheduleRequests = value
}
// Sets the roleEligibilitySchedules property value. 
// Parameters:
//  - value : Value to set for the roleEligibilitySchedules property.
func (m *RbacApplication) SetRoleEligibilitySchedules(value []UnifiedRoleEligibilitySchedule)() {
    m.roleEligibilitySchedules = value
}
