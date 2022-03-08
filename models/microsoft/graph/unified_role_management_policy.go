package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleManagementPolicy provides operations to manage the policyRoot singleton.
type UnifiedRoleManagementPolicy struct {
    Entity
    // Description for the policy.
    description *string;
    // Display name for the policy.
    displayName *string;
    // The list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
    effectiveRules []UnifiedRoleManagementPolicyRuleable;
    // This can only be set to true for a single tenant wide policy which will apply to all scopes and roles. Set the scopeId to '/' and scopeType to Directory.
    isOrganizationDefault *bool;
    // The identity who last modified the role setting.
    lastModifiedBy Identityable;
    // The time when the role setting was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of rules like approval rule, expiration rule, etc.
    rules []UnifiedRoleManagementPolicyRuleable;
    // The id of the scope where the policy is created. E.g. '/', groupId, etc.
    scopeId *string;
    // The type of the scope where the policy is created. One of Directory, DirectoryRole, Group.
    scopeType *string;
}
// NewUnifiedRoleManagementPolicy instantiates a new unifiedRoleManagementPolicy and sets the default values.
func NewUnifiedRoleManagementPolicy()(*UnifiedRoleManagementPolicy) {
    m := &UnifiedRoleManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUnifiedRoleManagementPolicy(), nil
}
// GetDescription gets the description property value. Description for the policy.
func (m *UnifiedRoleManagementPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name for the policy.
func (m *UnifiedRoleManagementPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEffectiveRules gets the effectiveRules property value. The list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicy) GetEffectiveRules()([]UnifiedRoleManagementPolicyRuleable) {
    if m == nil {
        return nil
    } else {
        return m.effectiveRules
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["effectiveRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyRuleable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyRuleable)
            }
            m.SetEffectiveRules(res)
        }
        return nil
    }
    res["isOrganizationDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizationDefault(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(Identityable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["rules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyRuleable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    res["scopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeId(val)
        }
        return nil
    }
    res["scopeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeType(val)
        }
        return nil
    }
    return res
}
// GetIsOrganizationDefault gets the isOrganizationDefault property value. This can only be set to true for a single tenant wide policy which will apply to all scopes and roles. Set the scopeId to '/' and scopeType to Directory.
func (m *UnifiedRoleManagementPolicy) GetIsOrganizationDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizationDefault
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity who last modified the role setting.
func (m *UnifiedRoleManagementPolicy) GetLastModifiedBy()(Identityable) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The time when the role setting was last modified.
func (m *UnifiedRoleManagementPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetRules gets the rules property value. The collection of rules like approval rule, expiration rule, etc.
func (m *UnifiedRoleManagementPolicy) GetRules()([]UnifiedRoleManagementPolicyRuleable) {
    if m == nil {
        return nil
    } else {
        return m.rules
    }
}
// GetScopeId gets the scopeId property value. The id of the scope where the policy is created. E.g. '/', groupId, etc.
func (m *UnifiedRoleManagementPolicy) GetScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeId
    }
}
// GetScopeType gets the scopeType property value. The type of the scope where the policy is created. One of Directory, DirectoryRole, Group.
func (m *UnifiedRoleManagementPolicy) GetScopeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeType
    }
}
func (m *UnifiedRoleManagementPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetEffectiveRules() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEffectiveRules()))
        for i, v := range m.GetEffectiveRules() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetRules() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetDescription sets the description property value. Description for the policy.
func (m *UnifiedRoleManagementPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name for the policy.
func (m *UnifiedRoleManagementPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEffectiveRules sets the effectiveRules property value. The list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicy) SetEffectiveRules(value []UnifiedRoleManagementPolicyRuleable)() {
    if m != nil {
        m.effectiveRules = value
    }
}
// SetIsOrganizationDefault sets the isOrganizationDefault property value. This can only be set to true for a single tenant wide policy which will apply to all scopes and roles. Set the scopeId to '/' and scopeType to Directory.
func (m *UnifiedRoleManagementPolicy) SetIsOrganizationDefault(value *bool)() {
    if m != nil {
        m.isOrganizationDefault = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity who last modified the role setting.
func (m *UnifiedRoleManagementPolicy) SetLastModifiedBy(value Identityable)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The time when the role setting was last modified.
func (m *UnifiedRoleManagementPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetRules sets the rules property value. The collection of rules like approval rule, expiration rule, etc.
func (m *UnifiedRoleManagementPolicy) SetRules(value []UnifiedRoleManagementPolicyRuleable)() {
    if m != nil {
        m.rules = value
    }
}
// SetScopeId sets the scopeId property value. The id of the scope where the policy is created. E.g. '/', groupId, etc.
func (m *UnifiedRoleManagementPolicy) SetScopeId(value *string)() {
    if m != nil {
        m.scopeId = value
    }
}
// SetScopeType sets the scopeType property value. The type of the scope where the policy is created. One of Directory, DirectoryRole, Group.
func (m *UnifiedRoleManagementPolicy) SetScopeType(value *string)() {
    if m != nil {
        m.scopeType = value
    }
}
