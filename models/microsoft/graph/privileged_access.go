package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrivilegedAccess struct {
    Entity
    displayName *string;
    resources []GovernanceResource;
    roleAssignmentRequests []GovernanceRoleAssignmentRequest;
    roleAssignments []GovernanceRoleAssignment;
    roleDefinitions []GovernanceRoleDefinition;
    roleSettings []GovernanceRoleSetting;
}
func NewPrivilegedAccess()(*PrivilegedAccess) {
    m := &PrivilegedAccess{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrivilegedAccess) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *PrivilegedAccess) GetResources()([]GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
func (m *PrivilegedAccess) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentRequests
    }
}
func (m *PrivilegedAccess) GetRoleAssignments()([]GovernanceRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
func (m *PrivilegedAccess) GetRoleDefinitions()([]GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
func (m *PrivilegedAccess) GetRoleSettings()([]GovernanceRoleSetting) {
    if m == nil {
        return nil
    } else {
        return m.roleSettings
    }
}
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
func (m *PrivilegedAccess) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *PrivilegedAccess) SetResources(value []GovernanceResource)() {
    m.resources = value
}
func (m *PrivilegedAccess) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequest)() {
    m.roleAssignmentRequests = value
}
func (m *PrivilegedAccess) SetRoleAssignments(value []GovernanceRoleAssignment)() {
    m.roleAssignments = value
}
func (m *PrivilegedAccess) SetRoleDefinitions(value []GovernanceRoleDefinition)() {
    m.roleDefinitions = value
}
func (m *PrivilegedAccess) SetRoleSettings(value []GovernanceRoleSetting)() {
    m.roleSettings = value
}
