package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleDefinition provides operations to manage the deviceManagement singleton.
type RoleDefinition struct {
    Entity
    // Description of the Role definition.
    description *string;
    // Display Name of the Role definition.
    displayName *string;
    // Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
    isBuiltIn *bool;
    // Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
    isBuiltInRoleDefinition *bool;
    // List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
    permissions []RolePermissionable;
    // List of Role assignments for this role definition.
    roleAssignments []RoleAssignmentable;
    // List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
    rolePermissions []RolePermissionable;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
}
// NewRoleDefinition instantiates a new roleDefinition and sets the default values.
func NewRoleDefinition()(*RoleDefinition) {
    m := &RoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleDefinitionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRoleDefinition(), nil
}
// GetDescription gets the description property value. Description of the Role definition.
func (m *RoleDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display Name of the Role definition.
func (m *RoleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    res["isBuiltInRoleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltInRoleDefinition(val)
        }
        return nil
    }
    res["permissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RolePermissionable, len(val))
            for i, v := range val {
                res[i] = v.(RolePermissionable)
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["rolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RolePermissionable, len(val))
            for i, v := range val {
                res[i] = v.(RolePermissionable)
            }
            m.SetRolePermissions(res)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// GetIsBuiltInRoleDefinition gets the isBuiltInRoleDefinition property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) GetIsBuiltInRoleDefinition()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltInRoleDefinition
    }
}
// GetPermissions gets the permissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) GetPermissions()([]RolePermissionable) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
// GetRoleAssignments gets the roleAssignments property value. List of Role assignments for this role definition.
func (m *RoleDefinition) GetRoleAssignments()([]RoleAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRolePermissions gets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) GetRolePermissions()([]RolePermissionable) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RoleDefinition) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *RoleDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RoleDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltInRoleDefinition", m.GetIsBuiltInRoleDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissions()))
        for i, v := range m.GetPermissions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("permissions", cast)
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
    if m.GetRolePermissions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the Role definition.
func (m *RoleDefinition) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display Name of the Role definition.
func (m *RoleDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) SetIsBuiltIn(value *bool)() {
    if m != nil {
        m.isBuiltIn = value
    }
}
// SetIsBuiltInRoleDefinition sets the isBuiltInRoleDefinition property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) SetIsBuiltInRoleDefinition(value *bool)() {
    if m != nil {
        m.isBuiltInRoleDefinition = value
    }
}
// SetPermissions sets the permissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) SetPermissions(value []RolePermissionable)() {
    if m != nil {
        m.permissions = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. List of Role assignments for this role definition.
func (m *RoleDefinition) SetRoleAssignments(value []RoleAssignmentable)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRolePermissions sets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) SetRolePermissions(value []RolePermissionable)() {
    if m != nil {
        m.rolePermissions = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RoleDefinition) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
