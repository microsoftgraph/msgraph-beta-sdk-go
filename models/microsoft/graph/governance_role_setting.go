package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GovernanceRoleSetting struct {
    Entity
    adminEligibleSettings []GovernanceRuleSetting;
    adminMemberSettings []GovernanceRuleSetting;
    isDefault *bool;
    lastUpdatedBy *string;
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    resource *GovernanceResource;
    resourceId *string;
    roleDefinition *GovernanceRoleDefinition;
    roleDefinitionId *string;
    userEligibleSettings []GovernanceRuleSetting;
    userMemberSettings []GovernanceRuleSetting;
}
func NewGovernanceRoleSetting()(*GovernanceRoleSetting) {
    m := &GovernanceRoleSetting{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GovernanceRoleSetting) GetAdminEligibleSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.adminEligibleSettings
    }
}
func (m *GovernanceRoleSetting) GetAdminMemberSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.adminMemberSettings
    }
}
func (m *GovernanceRoleSetting) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
func (m *GovernanceRoleSetting) GetLastUpdatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedBy
    }
}
func (m *GovernanceRoleSetting) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
func (m *GovernanceRoleSetting) GetResource()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *GovernanceRoleSetting) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
func (m *GovernanceRoleSetting) GetRoleDefinition()(*GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
func (m *GovernanceRoleSetting) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
func (m *GovernanceRoleSetting) GetUserEligibleSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.userEligibleSettings
    }
}
func (m *GovernanceRoleSetting) GetUserMemberSettings()([]GovernanceRuleSetting) {
    if m == nil {
        return nil
    } else {
        return m.userMemberSettings
    }
}
func (m *GovernanceRoleSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adminEligibleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRuleSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRuleSetting))
        }
        m.SetAdminEligibleSettings(res)
        return nil
    }
    res["adminMemberSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRuleSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRuleSetting))
        }
        m.SetAdminMemberSettings(res)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["lastUpdatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedBy(val)
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*GovernanceResource))
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    res["roleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleDefinition() })
        if err != nil {
            return err
        }
        m.SetRoleDefinition(val.(*GovernanceRoleDefinition))
        return nil
    }
    res["roleDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleDefinitionId(val)
        return nil
    }
    res["userEligibleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRuleSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRuleSetting))
        }
        m.SetUserEligibleSettings(res)
        return nil
    }
    res["userMemberSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRuleSetting() })
        if err != nil {
            return err
        }
        res := make([]GovernanceRuleSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceRuleSetting))
        }
        m.SetUserMemberSettings(res)
        return nil
    }
    return res
}
func (m *GovernanceRoleSetting) IsNil()(bool) {
    return m == nil
}
func (m *GovernanceRoleSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
func (m *GovernanceRoleSetting) SetAdminEligibleSettings(value []GovernanceRuleSetting)() {
    m.adminEligibleSettings = value
}
func (m *GovernanceRoleSetting) SetAdminMemberSettings(value []GovernanceRuleSetting)() {
    m.adminMemberSettings = value
}
func (m *GovernanceRoleSetting) SetIsDefault(value *bool)() {
    m.isDefault = value
}
func (m *GovernanceRoleSetting) SetLastUpdatedBy(value *string)() {
    m.lastUpdatedBy = value
}
func (m *GovernanceRoleSetting) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
func (m *GovernanceRoleSetting) SetResource(value *GovernanceResource)() {
    m.resource = value
}
func (m *GovernanceRoleSetting) SetResourceId(value *string)() {
    m.resourceId = value
}
func (m *GovernanceRoleSetting) SetRoleDefinition(value *GovernanceRoleDefinition)() {
    m.roleDefinition = value
}
func (m *GovernanceRoleSetting) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
func (m *GovernanceRoleSetting) SetUserEligibleSettings(value []GovernanceRuleSetting)() {
    m.userEligibleSettings = value
}
func (m *GovernanceRoleSetting) SetUserMemberSettings(value []GovernanceRuleSetting)() {
    m.userMemberSettings = value
}
