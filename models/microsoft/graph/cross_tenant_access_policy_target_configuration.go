package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CrossTenantAccessPolicyTargetConfiguration 
type CrossTenantAccessPolicyTargetConfiguration struct {
    // Defines whether access is allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
    accessType *CrossTenantAccessPolicyTargetConfigurationAccessType;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether to target users, groups, or applications with this rule.
    targets []CrossTenantAccessPolicyTarget;
}
// NewCrossTenantAccessPolicyTargetConfiguration instantiates a new crossTenantAccessPolicyTargetConfiguration and sets the default values.
func NewCrossTenantAccessPolicyTargetConfiguration()(*CrossTenantAccessPolicyTargetConfiguration) {
    m := &CrossTenantAccessPolicyTargetConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessType gets the accessType property value. Defines whether access is allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
func (m *CrossTenantAccessPolicyTargetConfiguration) GetAccessType()(*CrossTenantAccessPolicyTargetConfigurationAccessType) {
    if m == nil {
        return nil
    } else {
        return m.accessType
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyTargetConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTargets gets the targets property value. Specifies whether to target users, groups, or applications with this rule.
func (m *CrossTenantAccessPolicyTargetConfiguration) GetTargets()([]CrossTenantAccessPolicyTarget) {
    if m == nil {
        return nil
    } else {
        return m.targets
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyTargetConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCrossTenantAccessPolicyTargetConfigurationAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessType(val.(*CrossTenantAccessPolicyTargetConfigurationAccessType))
        }
        return nil
    }
    res["targets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CrossTenantAccessPolicyTarget, len(val))
            for i, v := range val {
                res[i] = *(v.(*CrossTenantAccessPolicyTarget))
            }
            m.SetTargets(res)
        }
        return nil
    }
    return res
}
func (m *CrossTenantAccessPolicyTargetConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyTargetConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := (*m.GetAccessType()).String()
        err := writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("targets", cast)
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
// SetAccessType sets the accessType property value. Defines whether access is allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
func (m *CrossTenantAccessPolicyTargetConfiguration) SetAccessType(value *CrossTenantAccessPolicyTargetConfigurationAccessType)() {
    if m != nil {
        m.accessType = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyTargetConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTargets sets the targets property value. Specifies whether to target users, groups, or applications with this rule.
func (m *CrossTenantAccessPolicyTargetConfiguration) SetTargets(value []CrossTenantAccessPolicyTarget)() {
    if m != nil {
        m.targets = value
    }
}
