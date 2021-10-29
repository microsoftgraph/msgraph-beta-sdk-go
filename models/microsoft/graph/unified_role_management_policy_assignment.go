package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleManagementPolicyAssignment struct {
    Entity
    // The policy for the assignment.
    policy *UnifiedRoleManagementPolicy;
    // The id of the policy.
    policyId *string;
    // The id of the role definition where the policy applies. If not specified, the policy applies to all roles.
    roleDefinitionId *string;
    // The id of the scope where the policy is assigned. E.g. '/', groupId, etc.
    scopeId *string;
    // The type of the scope where the policy is assigned. One of Directory, DirectoryRole, Group.
    scopeType *string;
}
// Instantiates a new unifiedRoleManagementPolicyAssignment and sets the default values.
func NewUnifiedRoleManagementPolicyAssignment()(*UnifiedRoleManagementPolicyAssignment) {
    m := &UnifiedRoleManagementPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the policy property value. The policy for the assignment.
func (m *UnifiedRoleManagementPolicyAssignment) GetPolicy()(*UnifiedRoleManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// Gets the policyId property value. The id of the policy.
func (m *UnifiedRoleManagementPolicyAssignment) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// Gets the roleDefinitionId property value. The id of the role definition where the policy applies. If not specified, the policy applies to all roles.
func (m *UnifiedRoleManagementPolicyAssignment) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// Gets the scopeId property value. The id of the scope where the policy is assigned. E.g. '/', groupId, etc.
func (m *UnifiedRoleManagementPolicyAssignment) GetScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeId
    }
}
// Gets the scopeType property value. The type of the scope where the policy is assigned. One of Directory, DirectoryRole, Group.
func (m *UnifiedRoleManagementPolicyAssignment) GetScopeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeType
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleManagementPolicyAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicy() })
        if err != nil {
            return err
        }
        m.SetPolicy(val.(*UnifiedRoleManagementPolicy))
        return nil
    }
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyId(val)
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
func (m *UnifiedRoleManagementPolicyAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleManagementPolicyAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyId", m.GetPolicyId())
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
// Sets the policy property value. The policy for the assignment.
// Parameters:
//  - value : Value to set for the policy property.
func (m *UnifiedRoleManagementPolicyAssignment) SetPolicy(value *UnifiedRoleManagementPolicy)() {
    m.policy = value
}
// Sets the policyId property value. The id of the policy.
// Parameters:
//  - value : Value to set for the policyId property.
func (m *UnifiedRoleManagementPolicyAssignment) SetPolicyId(value *string)() {
    m.policyId = value
}
// Sets the roleDefinitionId property value. The id of the role definition where the policy applies. If not specified, the policy applies to all roles.
// Parameters:
//  - value : Value to set for the roleDefinitionId property.
func (m *UnifiedRoleManagementPolicyAssignment) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// Sets the scopeId property value. The id of the scope where the policy is assigned. E.g. '/', groupId, etc.
// Parameters:
//  - value : Value to set for the scopeId property.
func (m *UnifiedRoleManagementPolicyAssignment) SetScopeId(value *string)() {
    m.scopeId = value
}
// Sets the scopeType property value. The type of the scope where the policy is assigned. One of Directory, DirectoryRole, Group.
// Parameters:
//  - value : Value to set for the scopeType property.
func (m *UnifiedRoleManagementPolicyAssignment) SetScopeType(value *string)() {
    m.scopeType = value
}
