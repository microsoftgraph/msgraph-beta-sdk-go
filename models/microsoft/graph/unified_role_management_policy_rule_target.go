package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleManagementPolicyRuleTarget 
type UnifiedRoleManagementPolicyRuleTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The caller for the policy rule target. Allowed values are: None, Admin, EndUser.
    caller *string;
    // The list of settings which are enforced and cannot be overridden by child scopes. Use All for all settings.
    enforcedSettings []string;
    // The list of settings which can be inherited by child scopes. Use All for all settings.
    inheritableSettings []string;
    // The level for the policy rule target. Allowed values are: Eligibility, Assignment.
    level *string;
    // The operations for policy rule target. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
    operations []string;
    // 
    targetObjects []DirectoryObject;
}
// NewUnifiedRoleManagementPolicyRuleTarget instantiates a new unifiedRoleManagementPolicyRuleTarget and sets the default values.
func NewUnifiedRoleManagementPolicyRuleTarget()(*UnifiedRoleManagementPolicyRuleTarget) {
    m := &UnifiedRoleManagementPolicyRuleTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCaller gets the caller property value. The caller for the policy rule target. Allowed values are: None, Admin, EndUser.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetCaller()(*string) {
    if m == nil {
        return nil
    } else {
        return m.caller
    }
}
// GetEnforcedSettings gets the enforcedSettings property value. The list of settings which are enforced and cannot be overridden by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetEnforcedSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSettings
    }
}
// GetInheritableSettings gets the inheritableSettings property value. The list of settings which can be inherited by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetInheritableSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inheritableSettings
    }
}
// GetLevel gets the level property value. The level for the policy rule target. Allowed values are: Eligibility, Assignment.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.level
    }
}
// GetOperations gets the operations property value. The operations for policy rule target. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetOperations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetTargetObjects gets the targetObjects property value. 
func (m *UnifiedRoleManagementPolicyRuleTarget) GetTargetObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.targetObjects
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyRuleTarget) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["caller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaller(val)
        }
        return nil
    }
    res["enforcedSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnforcedSettings(res)
        }
        return nil
    }
    res["inheritableSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInheritableSettings(res)
        }
        return nil
    }
    res["level"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["targetObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTargetObjects(res)
        }
        return nil
    }
    return res
}
func (m *UnifiedRoleManagementPolicyRuleTarget) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyRuleTarget) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("caller", m.GetCaller())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcedSettings() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedSettings", m.GetEnforcedSettings())
        if err != nil {
            return err
        }
    }
    if m.GetInheritableSettings() != nil {
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
    if m.GetOperations() != nil {
        err := writer.WriteCollectionOfStringValues("operations", m.GetOperations())
        if err != nil {
            return err
        }
    }
    if m.GetTargetObjects() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCaller sets the caller property value. The caller for the policy rule target. Allowed values are: None, Admin, EndUser.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetCaller(value *string)() {
    if m != nil {
        m.caller = value
    }
}
// SetEnforcedSettings sets the enforcedSettings property value. The list of settings which are enforced and cannot be overridden by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetEnforcedSettings(value []string)() {
    if m != nil {
        m.enforcedSettings = value
    }
}
// SetInheritableSettings sets the inheritableSettings property value. The list of settings which can be inherited by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetInheritableSettings(value []string)() {
    if m != nil {
        m.inheritableSettings = value
    }
}
// SetLevel sets the level property value. The level for the policy rule target. Allowed values are: Eligibility, Assignment.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetLevel(value *string)() {
    if m != nil {
        m.level = value
    }
}
// SetOperations sets the operations property value. The operations for policy rule target. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetOperations(value []string)() {
    if m != nil {
        m.operations = value
    }
}
// SetTargetObjects sets the targetObjects property value. 
func (m *UnifiedRoleManagementPolicyRuleTarget) SetTargetObjects(value []DirectoryObject)() {
    if m != nil {
        m.targetObjects = value
    }
}
