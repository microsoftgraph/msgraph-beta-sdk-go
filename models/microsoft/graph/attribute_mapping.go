package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttributeMapping struct {
    additionalData map[string]interface{};
    defaultValue *string;
    exportMissingReferences *bool;
    flowBehavior *AttributeFlowBehavior;
    flowType *AttributeFlowType;
    matchingPriority *int32;
    source *AttributeMappingSource;
    targetAttributeName *string;
}
func NewAttributeMapping()(*AttributeMapping) {
    m := &AttributeMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttributeMapping) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttributeMapping) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
func (m *AttributeMapping) GetExportMissingReferences()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.exportMissingReferences
    }
}
func (m *AttributeMapping) GetFlowBehavior()(*AttributeFlowBehavior) {
    if m == nil {
        return nil
    } else {
        return m.flowBehavior
    }
}
func (m *AttributeMapping) GetFlowType()(*AttributeFlowType) {
    if m == nil {
        return nil
    } else {
        return m.flowType
    }
}
func (m *AttributeMapping) GetMatchingPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.matchingPriority
    }
}
func (m *AttributeMapping) GetSource()(*AttributeMappingSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
func (m *AttributeMapping) GetTargetAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetAttributeName
    }
}
func (m *AttributeMapping) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultValue(val)
        return nil
    }
    res["exportMissingReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetExportMissingReferences(val)
        return nil
    }
    res["flowBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeFlowBehavior)
        if err != nil {
            return err
        }
        cast := val.(AttributeFlowBehavior)
        m.SetFlowBehavior(&cast)
        return nil
    }
    res["flowType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeFlowType)
        if err != nil {
            return err
        }
        cast := val.(AttributeFlowType)
        m.SetFlowType(&cast)
        return nil
    }
    res["matchingPriority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMatchingPriority(val)
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeMappingSource() })
        if err != nil {
            return err
        }
        m.SetSource(val.(*AttributeMappingSource))
        return nil
    }
    res["targetAttributeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetAttributeName(val)
        return nil
    }
    return res
}
func (m *AttributeMapping) IsNil()(bool) {
    return m == nil
}
func (m *AttributeMapping) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exportMissingReferences", m.GetExportMissingReferences())
        if err != nil {
            return err
        }
    }
    if m.GetFlowBehavior() != nil {
        cast := m.GetFlowBehavior().String()
        err := writer.WriteStringValue("flowBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFlowType() != nil {
        cast := m.GetFlowType().String()
        err := writer.WriteStringValue("flowType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("matchingPriority", m.GetMatchingPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetAttributeName", m.GetTargetAttributeName())
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
func (m *AttributeMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttributeMapping) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
func (m *AttributeMapping) SetExportMissingReferences(value *bool)() {
    m.exportMissingReferences = value
}
func (m *AttributeMapping) SetFlowBehavior(value *AttributeFlowBehavior)() {
    m.flowBehavior = value
}
func (m *AttributeMapping) SetFlowType(value *AttributeFlowType)() {
    m.flowType = value
}
func (m *AttributeMapping) SetMatchingPriority(value *int32)() {
    m.matchingPriority = value
}
func (m *AttributeMapping) SetSource(value *AttributeMappingSource)() {
    m.source = value
}
func (m *AttributeMapping) SetTargetAttributeName(value *string)() {
    m.targetAttributeName = value
}
