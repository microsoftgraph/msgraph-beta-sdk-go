package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrivilegedAccess struct {
    Entity
    // The display name of the provider managed by PIM.
    displayName *string;
    // A collection of resources for the provider.
    resources []GovernanceResource;
    // A collection of role assignment requests for the provider.
    roleAssignmentRequests []GovernanceRoleAssignmentRequest;
    // A collection of role assignments for the provider.
    roleAssignments []GovernanceRoleAssignment;
    // A collection of role defintions for the provider.
    roleDefinitions []GovernanceRoleDefinition;
    // A collection of role settings for the provider.
    roleSettings []GovernanceRoleSetting;
}
// Instantiates a new privilegedAccess and sets the default values.
func NewPrivilegedAccess()(*PrivilegedAccess) {
    m := &PrivilegedAccess{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) GetResources()([]GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// Gets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentRequests
    }
}
// Gets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) GetRoleAssignments()([]GovernanceRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// Gets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) GetRoleDefinitions()([]GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// Gets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) GetRoleSettings()([]GovernanceRoleSetting) {
    if m == nil {
        return nil
    } else {
        return m.roleSettings
    }
}
// The deserialization information for the current model
func (m *PrivilegedAccess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        res := make([]GovernanceResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceResource))
        }
        m.SetResources(res)
        return nil
    }
    res["roleAssignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignmentRequest() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRoleAssignmentRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRoleAssignmentRequest))
        }
        m.SetRoleAssignmentRequests(res)
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignment() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRoleAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRoleAssignment))
        }
        m.SetRoleAssignments(res)
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleDefinition() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRoleDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRoleDefinition))
        }
        m.SetRoleDefinitions(res)
        return nil
    }
    res["roleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleSetting() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRoleSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRoleSetting))
        }
        m.SetRoleSettings(res)
        return nil
    }
    return res
}
func (m *PrivilegedAccess) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrivilegedAccess) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignmentRequests()))
        for i, v := range m.GetRoleAssignmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentRequests", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleSettings()))
        for i, v := range m.GetRoleSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The display name of the provider managed by PIM.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PrivilegedAccess) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the resources property value. A collection of resources for the provider.
// Parameters:
//  - value : Value to set for the resources property.
func (m *PrivilegedAccess) SetResources(value []GovernanceResource)() {
    m.resources = value
}
// Sets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
// Parameters:
//  - value : Value to set for the roleAssignmentRequests property.
func (m *PrivilegedAccess) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequest)() {
    m.roleAssignmentRequests = value
}
// Sets the roleAssignments property value. A collection of role assignments for the provider.
// Parameters:
//  - value : Value to set for the roleAssignments property.
func (m *PrivilegedAccess) SetRoleAssignments(value []GovernanceRoleAssignment)() {
    m.roleAssignments = value
}
// Sets the roleDefinitions property value. A collection of role defintions for the provider.
// Parameters:
//  - value : Value to set for the roleDefinitions property.
func (m *PrivilegedAccess) SetRoleDefinitions(value []GovernanceRoleDefinition)() {
    m.roleDefinitions = value
}
// Sets the roleSettings property value. A collection of role settings for the provider.
// Parameters:
//  - value : Value to set for the roleSettings property.
func (m *PrivilegedAccess) SetRoleSettings(value []GovernanceRoleSetting)() {
    m.roleSettings = value
}
