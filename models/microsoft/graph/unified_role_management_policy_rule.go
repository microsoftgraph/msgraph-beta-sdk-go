package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRoleManagementPolicyRule struct {
    Entity
    target *UnifiedRoleManagementPolicyRuleTarget;
}
func NewUnifiedRoleManagementPolicyRule()(*UnifiedRoleManagementPolicyRule) {
    m := &UnifiedRoleManagementPolicyRule{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UnifiedRoleManagementPolicyRule) GetTarget()(*UnifiedRoleManagementPolicyRuleTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
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
func (m *UnifiedRoleManagementPolicyRule) SetTarget(value *UnifiedRoleManagementPolicyRuleTarget)() {
    m.target = value
}
