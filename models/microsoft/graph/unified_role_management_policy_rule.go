package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleManagementPolicyRule struct {
    Entity
    // The target for the policy rule.
    target *UnifiedRoleManagementPolicyRuleTarget;
}
// Instantiates a new unifiedRoleManagementPolicyRule and sets the default values.
func NewUnifiedRoleManagementPolicyRule()(*UnifiedRoleManagementPolicyRule) {
    m := &UnifiedRoleManagementPolicyRule{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the target property value. The target for the policy rule.
func (m *UnifiedRoleManagementPolicyRule) GetTarget()(*UnifiedRoleManagementPolicyRuleTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleManagementPolicyRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicyRuleTarget() })
        if err != nil {
            return err
        }
        m.SetTarget(val.(*UnifiedRoleManagementPolicyRuleTarget))
        return nil
    }
    return res
}
func (m *UnifiedRoleManagementPolicyRule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the target property value. The target for the policy rule.
// Parameters:
//  - value : Value to set for the target property.
func (m *UnifiedRoleManagementPolicyRule) SetTarget(value *UnifiedRoleManagementPolicyRuleTarget)() {
    m.target = value
}
