package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleManagementPolicyRule 
type UnifiedRoleManagementPolicyRule struct {
    Entity
    // The target for the policy rule.
    target *UnifiedRoleManagementPolicyRuleTarget;
}
// NewUnifiedRoleManagementPolicyRule instantiates a new unifiedRoleManagementPolicyRule and sets the default values.
func NewUnifiedRoleManagementPolicyRule()(*UnifiedRoleManagementPolicyRule) {
    m := &UnifiedRoleManagementPolicyRule{
        Entity: *NewEntity(),
    }
    return m
}
// GetTarget gets the target property value. The target for the policy rule.
func (m *UnifiedRoleManagementPolicyRule) GetTarget()(*UnifiedRoleManagementPolicyRuleTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicyRuleTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*UnifiedRoleManagementPolicyRuleTarget))
        }
        return nil
    }
    return res
}
func (m *UnifiedRoleManagementPolicyRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. The target for the policy rule.
func (m *UnifiedRoleManagementPolicyRule) SetTarget(value *UnifiedRoleManagementPolicyRuleTarget)() {
    if m != nil {
        m.target = value
    }
}
