package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RbacApplicationMultiple 
type RbacApplicationMultiple struct {
    Entity
    // 
    resourceNamespaces []UnifiedRbacResourceNamespace;
    // 
    roleAssignments []UnifiedRoleAssignmentMultiple;
    // 
    roleDefinitions []UnifiedRoleDefinition;
}
// NewRbacApplicationMultiple instantiates a new rbacApplicationMultiple and sets the default values.
func NewRbacApplicationMultiple()(*RbacApplicationMultiple) {
    m := &RbacApplicationMultiple{
        Entity: *NewEntity(),
    }
    return m
}
// GetResourceNamespaces gets the resourceNamespaces property value. 
func (m *RbacApplicationMultiple) GetResourceNamespaces()([]UnifiedRbacResourceNamespace) {
    if m == nil {
        return nil
    } else {
        return m.resourceNamespaces
    }
}
// GetRoleAssignments gets the roleAssignments property value. 
func (m *RbacApplicationMultiple) GetRoleAssignments()([]UnifiedRoleAssignmentMultiple) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. 
func (m *RbacApplicationMultiple) GetRoleDefinitions()([]UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RbacApplicationMultiple) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleAssignmentMultiple() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentMultiple, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleAssignmentMultiple))
            }
            m.SetRoleAssignments(res)
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
    return res
}
func (m *RbacApplicationMultiple) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RbacApplicationMultiple) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// SetResourceNamespaces sets the resourceNamespaces property value. 
func (m *RbacApplicationMultiple) SetResourceNamespaces(value []UnifiedRbacResourceNamespace)() {
    if m != nil {
        m.resourceNamespaces = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. 
func (m *RbacApplicationMultiple) SetRoleAssignments(value []UnifiedRoleAssignmentMultiple)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. 
func (m *RbacApplicationMultiple) SetRoleDefinitions(value []UnifiedRoleDefinition)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
