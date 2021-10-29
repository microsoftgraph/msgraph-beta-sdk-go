package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleManagementPolicy struct {
    Entity
    // Description for the policy.
    description *string;
    // Display name for the policy.
    displayName *string;
    // The list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
    effectiveRules []UnifiedRoleManagementPolicyRule;
    // This can only be set to true for a single tenant wide policy which will apply to all scopes and roles. Set the scopeId to '/' and scopeType to Directory.
    isOrganizationDefault *bool;
    // The identity who last modified the role setting.
    lastModifiedBy *Identity;
    // The time when the role setting was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of rules like approval rule, expiration rule, etc.
    rules []UnifiedRoleManagementPolicyRule;
    // The id of the scope where the policy is created. E.g. '/', groupId, etc.
    scopeId *string;
    // The type of the scope where the policy is created. One of Directory, DirectoryRole, Group.
    scopeType *string;
}
// Instantiates a new unifiedRoleManagementPolicy and sets the default values.
func NewUnifiedRoleManagementPolicy()(*UnifiedRoleManagementPolicy) {
    m := &UnifiedRoleManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. Description for the policy.
func (m *UnifiedRoleManagementPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display name for the policy.
func (m *UnifiedRoleManagementPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the effectiveRules property value. The list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicy) GetEffectiveRules()([]UnifiedRoleManagementPolicyRule) {
    if m == nil {
        return nil
    } else {
        return m.effectiveRules
    }
}
// Gets the isOrganizationDefault property value. This can only be set to true for a single tenant wide policy which will apply to all scopes and roles. Set the scopeId to '/' and scopeType to Directory.
func (m *UnifiedRoleManagementPolicy) GetIsOrganizationDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizationDefault
    }
}
// Gets the lastModifiedBy property value. The identity who last modified the role setting.
func (m *UnifiedRoleManagementPolicy) GetLastModifiedBy()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. The time when the role setting was last modified.
func (m *UnifiedRoleManagementPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the rules property value. The collection of rules like approval rule, expiration rule, etc.
func (m *UnifiedRoleManagementPolicy) GetRules()([]UnifiedRoleManagementPolicyRule) {
    if m == nil {
        return nil
    } else {
        return m.rules
    }
}
// Gets the scopeId property value. The id of the scope where the policy is created. E.g. '/', groupId, etc.
func (m *UnifiedRoleManagementPolicy) GetScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeId
    }
}
// Gets the scopeType property value. The type of the scope where the policy is created. One of Directory, DirectoryRole, Group.
func (m *UnifiedRoleManagementPolicy) GetScopeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeType
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["effectiveRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicyRule() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRoleManagementPolicyRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRoleManagementPolicyRule))
        }
        m.SetEffectiveRules(res)
        return nil
    }
    res["isOrganizationDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOrganizationDefault(val)
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetLastModifiedBy(val.(*Identity))
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["rules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicyRule() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRoleManagementPolicyRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRoleManagementPolicyRule))
        }
        m.SetRules(res)
        return nil
    }
    res["scopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScopeId(val)
        return nil
    }
    res["scopeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScopeType(val)
        return nil
    }
    return res
}
func (m *UnifiedRoleManagementPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleManagementPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEffectiveRules()))
        for i, v := range m.GetEffectiveRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("effectiveRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizationDefault", m.GetIsOrganizationDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeId", m.GetScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeType", m.GetScopeType())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. Description for the policy.
// Parameters:
//  - value : Value to set for the description property.
func (m *UnifiedRoleManagementPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display name for the policy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *UnifiedRoleManagementPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the effectiveRules property value. The list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
// Parameters:
//  - value : Value to set for the effectiveRules property.
func (m *UnifiedRoleManagementPolicy) SetEffectiveRules(value []UnifiedRoleManagementPolicyRule)() {
    m.effectiveRules = value
}
// Sets the isOrganizationDefault property value. This can only be set to true for a single tenant wide policy which will apply to all scopes and roles. Set the scopeId to '/' and scopeType to Directory.
// Parameters:
//  - value : Value to set for the isOrganizationDefault property.
func (m *UnifiedRoleManagementPolicy) SetIsOrganizationDefault(value *bool)() {
    m.isOrganizationDefault = value
}
// Sets the lastModifiedBy property value. The identity who last modified the role setting.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *UnifiedRoleManagementPolicy) SetLastModifiedBy(value *Identity)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. The time when the role setting was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *UnifiedRoleManagementPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the rules property value. The collection of rules like approval rule, expiration rule, etc.
// Parameters:
//  - value : Value to set for the rules property.
func (m *UnifiedRoleManagementPolicy) SetRules(value []UnifiedRoleManagementPolicyRule)() {
    m.rules = value
}
// Sets the scopeId property value. The id of the scope where the policy is created. E.g. '/', groupId, etc.
// Parameters:
//  - value : Value to set for the scopeId property.
func (m *UnifiedRoleManagementPolicy) SetScopeId(value *string)() {
    m.scopeId = value
}
// Sets the scopeType property value. The type of the scope where the policy is created. One of Directory, DirectoryRole, Group.
// Parameters:
//  - value : Value to set for the scopeType property.
func (m *UnifiedRoleManagementPolicy) SetScopeType(value *string)() {
    m.scopeType = value
}
