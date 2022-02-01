package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RbacApplication 
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
// NewRbacApplication instantiates a new rbacApplication and sets the default values.
func NewRbacApplication()(*RbacApplication) {
    m := &RbacApplication{
        Entity: *NewEntity(),
    }
    return m
}
// GetResourceNamespaces gets the resourceNamespaces property value. 
func (m *RbacApplication) GetResourceNamespaces()([]UnifiedRbacResourceNamespace) {
    if m == nil {
        return nil
    } else {
        return m.resourceNamespaces
    }
}
// GetRoleAssignmentApprovals gets the roleAssignmentApprovals property value. 
func (m *RbacApplication) GetRoleAssignmentApprovals()([]Approval) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentApprovals
    }
}
// GetRoleAssignments gets the roleAssignments property value. Resource to grant access to users or groups.
func (m *RbacApplication) GetRoleAssignments()([]UnifiedRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleAssignmentScheduleInstances gets the roleAssignmentScheduleInstances property value. 
func (m *RbacApplication) GetRoleAssignmentScheduleInstances()([]UnifiedRoleAssignmentScheduleInstance) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleInstances
    }
}
// GetRoleAssignmentScheduleRequests gets the roleAssignmentScheduleRequests property value. 
func (m *RbacApplication) GetRoleAssignmentScheduleRequests()([]UnifiedRoleAssignmentScheduleRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleRequests
    }
}
// GetRoleAssignmentSchedules gets the roleAssignmentSchedules property value. 
func (m *RbacApplication) GetRoleAssignmentSchedules()([]UnifiedRoleAssignmentSchedule) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentSchedules
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *RbacApplication) GetRoleDefinitions()([]UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// GetRoleEligibilityScheduleInstances gets the roleEligibilityScheduleInstances property value. 
func (m *RbacApplication) GetRoleEligibilityScheduleInstances()([]UnifiedRoleEligibilityScheduleInstance) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilityScheduleInstances
    }
}
// GetRoleEligibilityScheduleRequests gets the roleEligibilityScheduleRequests property value. 
func (m *RbacApplication) GetRoleEligibilityScheduleRequests()([]UnifiedRoleEligibilityScheduleRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilityScheduleRequests
    }
}
// GetRoleEligibilitySchedules gets the roleEligibilitySchedules property value. 
func (m *RbacApplication) GetRoleEligibilitySchedules()([]UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.roleEligibilitySchedules
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *RbacApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetResourceNamespaces() != nil {
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
    if m.GetRoleAssignmentApprovals() != nil {
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
    if m.GetRoleAssignments() != nil {
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
    if m.GetRoleAssignmentScheduleInstances() != nil {
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
    if m.GetRoleAssignmentScheduleRequests() != nil {
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
    if m.GetRoleAssignmentSchedules() != nil {
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
    if m.GetRoleDefinitions() != nil {
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
    if m.GetRoleEligibilityScheduleInstances() != nil {
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
    if m.GetRoleEligibilityScheduleRequests() != nil {
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
    if m.GetRoleEligibilitySchedules() != nil {
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
// SetResourceNamespaces sets the resourceNamespaces property value. 
func (m *RbacApplication) SetResourceNamespaces(value []UnifiedRbacResourceNamespace)() {
    if m != nil {
        m.resourceNamespaces = value
    }
}
// SetRoleAssignmentApprovals sets the roleAssignmentApprovals property value. 
func (m *RbacApplication) SetRoleAssignmentApprovals(value []Approval)() {
    if m != nil {
        m.roleAssignmentApprovals = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. Resource to grant access to users or groups.
func (m *RbacApplication) SetRoleAssignments(value []UnifiedRoleAssignment)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleAssignmentScheduleInstances sets the roleAssignmentScheduleInstances property value. 
func (m *RbacApplication) SetRoleAssignmentScheduleInstances(value []UnifiedRoleAssignmentScheduleInstance)() {
    if m != nil {
        m.roleAssignmentScheduleInstances = value
    }
}
// SetRoleAssignmentScheduleRequests sets the roleAssignmentScheduleRequests property value. 
func (m *RbacApplication) SetRoleAssignmentScheduleRequests(value []UnifiedRoleAssignmentScheduleRequest)() {
    if m != nil {
        m.roleAssignmentScheduleRequests = value
    }
}
// SetRoleAssignmentSchedules sets the roleAssignmentSchedules property value. 
func (m *RbacApplication) SetRoleAssignmentSchedules(value []UnifiedRoleAssignmentSchedule)() {
    if m != nil {
        m.roleAssignmentSchedules = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *RbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinition)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
// SetRoleEligibilityScheduleInstances sets the roleEligibilityScheduleInstances property value. 
func (m *RbacApplication) SetRoleEligibilityScheduleInstances(value []UnifiedRoleEligibilityScheduleInstance)() {
    if m != nil {
        m.roleEligibilityScheduleInstances = value
    }
}
// SetRoleEligibilityScheduleRequests sets the roleEligibilityScheduleRequests property value. 
func (m *RbacApplication) SetRoleEligibilityScheduleRequests(value []UnifiedRoleEligibilityScheduleRequest)() {
    if m != nil {
        m.roleEligibilityScheduleRequests = value
    }
}
// SetRoleEligibilitySchedules sets the roleEligibilitySchedules property value. 
func (m *RbacApplication) SetRoleEligibilitySchedules(value []UnifiedRoleEligibilitySchedule)() {
    if m != nil {
        m.roleEligibilitySchedules = value
    }
}
