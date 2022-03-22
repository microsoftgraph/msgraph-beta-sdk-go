package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RbacApplicationMultiple 
type RbacApplicationMultiple struct {
    Entity
    // 
    resourceNamespaces []UnifiedRbacResourceNamespaceable;
    // 
    roleAssignments []UnifiedRoleAssignmentMultipleable;
    // 
    roleDefinitions []UnifiedRoleDefinitionable;
}
// NewRbacApplicationMultiple instantiates a new rbacApplicationMultiple and sets the default values.
func NewRbacApplicationMultiple()(*RbacApplicationMultiple) {
    m := &RbacApplicationMultiple{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRbacApplicationMultipleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRbacApplicationMultipleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRbacApplicationMultiple(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RbacApplicationMultiple) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["resourceNamespaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentMultipleable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleAssignmentMultipleable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    return res
}
// GetResourceNamespaces gets the resourceNamespaces property value. 
func (m *RbacApplicationMultiple) GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable) {
    if m == nil {
        return nil
    } else {
        return m.resourceNamespaces
    }
}
// GetRoleAssignments gets the roleAssignments property value. 
func (m *RbacApplicationMultiple) GetRoleAssignments()([]UnifiedRoleAssignmentMultipleable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. 
func (m *RbacApplicationMultiple) GetRoleDefinitions()([]UnifiedRoleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceNamespaces", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceNamespaces sets the resourceNamespaces property value. 
func (m *RbacApplicationMultiple) SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)() {
    if m != nil {
        m.resourceNamespaces = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. 
func (m *RbacApplicationMultiple) SetRoleAssignments(value []UnifiedRoleAssignmentMultipleable)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. 
func (m *RbacApplicationMultiple) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
