package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRoleManagementPolicyRuleTarget struct {
    additionalData map[string]interface{};
    caller *string;
    enforcedSettings []string;
    inheritableSettings []string;
    level *string;
    operations []string;
    targetObjects []DirectoryObject;
}
func NewUnifiedRoleManagementPolicyRuleTarget()(*UnifiedRoleManagementPolicyRuleTarget) {
    m := &UnifiedRoleManagementPolicyRuleTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetCaller()(*string) {
    if m == nil {
        return nil
    } else {
        return m.caller
    }
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetEnforcedSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSettings
    }
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetInheritableSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inheritableSettings
    }
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.level
    }
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetOperations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetTargetObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.targetObjects
    }
}
func (m *UnifiedRoleManagementPolicyRuleTarget) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["caller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCaller(val)
        return nil
    }
    res["enforcedSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEnforcedSettings(res)
        return nil
    }
    res["inheritableSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetInheritableSettings(res)
        return nil
    }
    res["level"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLevel(val)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOperations(res)
        return nil
    }
    res["targetObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetTargetObjects(res)
        return nil
    }
    return res
}
func (m *UnifiedRoleManagementPolicyRuleTarget) IsNil()(bool) {
    return m == nil
}
func (m *UnifiedRoleManagementPolicyRuleTarget) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("caller", m.GetCaller())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("enforcedSettings", m.GetEnforcedSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("inheritableSettings", m.GetInheritableSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("level", m.GetLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("operations", m.GetOperations())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargetObjects()))
        for i, v := range m.GetTargetObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("targetObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UnifiedRoleManagementPolicyRuleTarget) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UnifiedRoleManagementPolicyRuleTarget) SetCaller(value *string)() {
    m.caller = value
}
func (m *UnifiedRoleManagementPolicyRuleTarget) SetEnforcedSettings(value []string)() {
    m.enforcedSettings = value
}
func (m *UnifiedRoleManagementPolicyRuleTarget) SetInheritableSettings(value []string)() {
    m.inheritableSettings = value
}
func (m *UnifiedRoleManagementPolicyRuleTarget) SetLevel(value *string)() {
    m.level = value
}
func (m *UnifiedRoleManagementPolicyRuleTarget) SetOperations(value []string)() {
    m.operations = value
}
func (m *UnifiedRoleManagementPolicyRuleTarget) SetTargetObjects(value []DirectoryObject)() {
    m.targetObjects = value
}
