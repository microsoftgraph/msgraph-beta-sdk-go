package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedAccess 
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
// NewPrivilegedAccess instantiates a new privilegedAccess and sets the default values.
func NewPrivilegedAccess()(*PrivilegedAccess) {
    m := &PrivilegedAccess{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetResources gets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) GetResources()([]GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// GetRoleAssignmentRequests gets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentRequests
    }
}
// GetRoleAssignments gets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) GetRoleAssignments()([]GovernanceRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) GetRoleDefinitions()([]GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// GetRoleSettings gets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) GetRoleSettings()([]GovernanceRoleSetting) {
    if m == nil {
        return nil
    } else {
        return m.roleSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceResource))
            }
            m.SetResources(res)
        }
        return nil
    }
    res["roleAssignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleAssignmentRequest))
            }
            m.SetRoleAssignmentRequests(res)
        }
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleAssignment))
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleDefinition))
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleSetting))
            }
            m.SetRoleSettings(res)
        }
        return nil
    }
    return res
}
func (m *PrivilegedAccess) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetResources() != nil {
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
    if m.GetRoleAssignmentRequests() != nil {
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
    if m.GetRoleSettings() != nil {
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
// SetDisplayName sets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetResources sets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) SetResources(value []GovernanceResource)() {
    if m != nil {
        m.resources = value
    }
}
// SetRoleAssignmentRequests sets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequest)() {
    if m != nil {
        m.roleAssignmentRequests = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) SetRoleAssignments(value []GovernanceRoleAssignment)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) SetRoleDefinitions(value []GovernanceRoleDefinition)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
// SetRoleSettings sets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) SetRoleSettings(value []GovernanceRoleSetting)() {
    if m != nil {
        m.roleSettings = value
    }
}
