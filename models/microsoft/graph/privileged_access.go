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
    resources []GovernanceResourceable;
    // A collection of role assignment requests for the provider.
    roleAssignmentRequests []GovernanceRoleAssignmentRequestable;
    // A collection of role assignments for the provider.
    roleAssignments []GovernanceRoleAssignmentable;
    // A collection of role defintions for the provider.
    roleDefinitions []GovernanceRoleDefinitionable;
    // A collection of role settings for the provider.
    roleSettings []GovernanceRoleSettingable;
}
// NewPrivilegedAccess instantiates a new privilegedAccess and sets the default values.
func NewPrivilegedAccess()(*PrivilegedAccess) {
    m := &PrivilegedAccess{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrivilegedAccess(), nil
}
// GetDisplayName gets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
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
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceResourceable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    res["roleAssignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleAssignmentRequestable)
            }
            m.SetRoleAssignmentRequests(res)
        }
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleAssignmentable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleDefinitionable)
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleSettingable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleSettingable)
            }
            m.SetRoleSettings(res)
        }
        return nil
    }
    return res
}
// GetResources gets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) GetResources()([]GovernanceResourceable) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// GetRoleAssignmentRequests gets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequestable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentRequests
    }
}
// GetRoleAssignments gets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) GetRoleAssignments()([]GovernanceRoleAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) GetRoleDefinitions()([]GovernanceRoleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// GetRoleSettings gets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) GetRoleSettings()([]GovernanceRoleSettingable) {
    if m == nil {
        return nil
    } else {
        return m.roleSettings
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignmentRequests()))
        for i, v := range m.GetRoleAssignmentRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentRequests", cast)
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
    if m.GetRoleSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleSettings()))
        for i, v := range m.GetRoleSettings() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *PrivilegedAccess) SetResources(value []GovernanceResourceable)() {
    if m != nil {
        m.resources = value
    }
}
// SetRoleAssignmentRequests sets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequestable)() {
    if m != nil {
        m.roleAssignmentRequests = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) SetRoleAssignments(value []GovernanceRoleAssignmentable)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) SetRoleDefinitions(value []GovernanceRoleDefinitionable)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
// SetRoleSettings sets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) SetRoleSettings(value []GovernanceRoleSettingable)() {
    if m != nil {
        m.roleSettings = value
    }
}
