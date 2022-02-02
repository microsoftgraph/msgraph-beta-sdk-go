package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceRoleSetting 
type GovernanceRoleSetting struct {
    Entity
    // The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
    adminEligibleSettings []GovernanceRuleSetting;
    // The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
    adminMemberSettings []GovernanceRuleSetting;
    // Read-only. Indicate if the roleSetting is a default roleSetting
    isDefault *bool;
    // Read-only. The display name of the administrator who last updated the roleSetting.
    lastUpdatedBy *string;
    // Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. The associated resource for this role setting.
    resource *GovernanceResource;
    // Required. The id of the resource that the role setting is associated with.
    resourceId *string;
    // Read-only. The role definition that is enforced with this role setting.
    roleDefinition *GovernanceRoleDefinition;
    // Required. The id of the role definition that the role setting is associated with.
    roleDefinitionId *string;
    // The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
    userEligibleSettings []GovernanceRuleSetting;
    // The rule settings that are evaluated when a user tries to activate his role assignment.
    userMemberSettings []GovernanceRuleSetting;
}
// NewGovernanceRoleSetting instantiates a new governanceRoleSetting and sets the default values.
func NewGovernanceRoleSetting()(*GovernanceRoleSetting) {
    m := &GovernanceRoleSetting{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdminEligibleSettings gets the adminEligibleSettings property value. The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
func (m *GovernanceRoleSetting) GetAdminEligibleSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.adminEligibleSettings
    }
}
// GetAdminMemberSettings gets the adminMemberSettings property value. The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
func (m *GovernanceRoleSetting) GetAdminMemberSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.adminMemberSettings
    }
}
// GetIsDefault gets the isDefault property value. Read-only. Indicate if the roleSetting is a default roleSetting
func (m *GovernanceRoleSetting) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetLastUpdatedBy gets the lastUpdatedBy property value. Read-only. The display name of the administrator who last updated the roleSetting.
func (m *GovernanceRoleSetting) GetLastUpdatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedBy
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleSetting) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetResource gets the resource property value. Read-only. The associated resource for this role setting.
func (m *GovernanceRoleSetting) GetResource()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceId gets the resourceId property value. Required. The id of the resource that the role setting is associated with.
func (m *GovernanceRoleSetting) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetRoleDefinition gets the roleDefinition property value. Read-only. The role definition that is enforced with this role setting.
func (m *GovernanceRoleSetting) GetRoleDefinition()(*GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Required. The id of the role definition that the role setting is associated with.
func (m *GovernanceRoleSetting) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// GetUserEligibleSettings gets the userEligibleSettings property value. The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
func (m *GovernanceRoleSetting) GetUserEligibleSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.userEligibleSettings
    }
}
// GetUserMemberSettings gets the userMemberSettings property value. The rule settings that are evaluated when a user tries to activate his role assignment.
func (m *GovernanceRoleSetting) GetUserMemberSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.userMemberSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRoleSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adminEligibleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRuleSetting))
            }
            m.SetAdminEligibleSettings(res)
        }
        return nil
    }
    res["adminMemberSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRuleSetting))
            }
            m.SetAdminMemberSettings(res)
        }
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["lastUpdatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedBy(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(*GovernanceResource))
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(*GovernanceRoleDefinition))
        }
        return nil
    }
    res["roleDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    res["userEligibleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRuleSetting))
            }
            m.SetUserEligibleSettings(res)
        }
        return nil
    }
    res["userMemberSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRuleSetting))
            }
            m.SetUserMemberSettings(res)
        }
        return nil
    }
    return res
}
func (m *GovernanceRoleSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernanceRoleSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdminEligibleSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdminEligibleSettings()))
        for i, v := range m.GetAdminEligibleSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("adminEligibleSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdminMemberSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdminMemberSettings()))
        for i, v := range m.GetAdminMemberSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("adminMemberSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastUpdatedBy", m.GetLastUpdatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetUserEligibleSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserEligibleSettings()))
        for i, v := range m.GetUserEligibleSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userEligibleSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserMemberSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserMemberSettings()))
        for i, v := range m.GetUserMemberSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userMemberSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdminEligibleSettings sets the adminEligibleSettings property value. The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
func (m *GovernanceRoleSetting) SetAdminEligibleSettings(value []GovernanceRuleSetting)() {
    if m != nil {
        m.adminEligibleSettings = value
    }
}
// SetAdminMemberSettings sets the adminMemberSettings property value. The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
func (m *GovernanceRoleSetting) SetAdminMemberSettings(value []GovernanceRuleSetting)() {
    if m != nil {
        m.adminMemberSettings = value
    }
}
// SetIsDefault sets the isDefault property value. Read-only. Indicate if the roleSetting is a default roleSetting
func (m *GovernanceRoleSetting) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetLastUpdatedBy sets the lastUpdatedBy property value. Read-only. The display name of the administrator who last updated the roleSetting.
func (m *GovernanceRoleSetting) SetLastUpdatedBy(value *string)() {
    if m != nil {
        m.lastUpdatedBy = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleSetting) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetResource sets the resource property value. Read-only. The associated resource for this role setting.
func (m *GovernanceRoleSetting) SetResource(value *GovernanceResource)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceId sets the resourceId property value. Required. The id of the resource that the role setting is associated with.
func (m *GovernanceRoleSetting) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetRoleDefinition sets the roleDefinition property value. Read-only. The role definition that is enforced with this role setting.
func (m *GovernanceRoleSetting) SetRoleDefinition(value *GovernanceRoleDefinition)() {
    if m != nil {
        m.roleDefinition = value
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Required. The id of the role definition that the role setting is associated with.
func (m *GovernanceRoleSetting) SetRoleDefinitionId(value *string)() {
    if m != nil {
        m.roleDefinitionId = value
    }
}
// SetUserEligibleSettings sets the userEligibleSettings property value. The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
func (m *GovernanceRoleSetting) SetUserEligibleSettings(value []GovernanceRuleSetting)() {
    if m != nil {
        m.userEligibleSettings = value
    }
}
// SetUserMemberSettings sets the userMemberSettings property value. The rule settings that are evaluated when a user tries to activate his role assignment.
func (m *GovernanceRoleSetting) SetUserMemberSettings(value []GovernanceRuleSetting)() {
    if m != nil {
        m.userMemberSettings = value
    }
}
