package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new unifiedRoleManagementPolicyRuleTarget and sets the default values.
func NewUnifiedRoleManagementPolicyRuleTarget()(*UnifiedRoleManagementPolicyRuleTarget) {
    m := &UnifiedRoleManagementPolicyRuleTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the caller property value. The caller for the policy rule target. Allowed values are: None, Admin, EndUser.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetCaller()(*string) {
    if m == nil {
        return nil
    } else {
        return m.caller
    }
}
// Gets the enforcedSettings property value. The list of settings which are enforced and cannot be overridden by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetEnforcedSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSettings
    }
}
// Gets the inheritableSettings property value. The list of settings which can be inherited by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetInheritableSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inheritableSettings
    }
}
// Gets the level property value. The level for the policy rule target. Allowed values are: Eligibility, Assignment.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.level
    }
}
// Gets the operations property value. The operations for policy rule target. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetOperations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// Gets the targetObjects property value. 
func (m *UnifiedRoleManagementPolicyRuleTarget) GetTargetObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.targetObjects
    }
}
// The deserialization information for the current model
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
            res[i] = *(v.(*string))
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
            res[i] = *(v.(*string))
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
            res[i] = *(v.(*string))
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the caller property value. The caller for the policy rule target. Allowed values are: None, Admin, EndUser.
// Parameters:
//  - value : Value to set for the caller property.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetCaller(value *string)() {
    m.caller = value
}
// Sets the enforcedSettings property value. The list of settings which are enforced and cannot be overridden by child scopes. Use All for all settings.
// Parameters:
//  - value : Value to set for the enforcedSettings property.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetEnforcedSettings(value []string)() {
    m.enforcedSettings = value
}
// Sets the inheritableSettings property value. The list of settings which can be inherited by child scopes. Use All for all settings.
// Parameters:
//  - value : Value to set for the inheritableSettings property.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetInheritableSettings(value []string)() {
    m.inheritableSettings = value
}
// Sets the level property value. The level for the policy rule target. Allowed values are: Eligibility, Assignment.
// Parameters:
//  - value : Value to set for the level property.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetLevel(value *string)() {
    m.level = value
}
// Sets the operations property value. The operations for policy rule target. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
// Parameters:
//  - value : Value to set for the operations property.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetOperations(value []string)() {
    m.operations = value
}
// Sets the targetObjects property value. 
// Parameters:
//  - value : Value to set for the targetObjects property.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetTargetObjects(value []DirectoryObject)() {
    m.targetObjects = value
}
