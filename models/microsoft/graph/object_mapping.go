package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ObjectMapping struct {
    additionalData map[string]interface{};
    attributeMappings []AttributeMapping;
    enabled *bool;
    flowTypes *ObjectFlowTypes;
    metadata []MetadataEntry;
    name *string;
    scope *Filter;
    sourceObjectName *string;
    targetObjectName *string;
}
func NewObjectMapping()(*ObjectMapping) {
    m := &ObjectMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ObjectMapping) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ObjectMapping) GetAttributeMappings()([]AttributeMapping) {
    if m == nil {
        return nil
    } else {
        return m.attributeMappings
    }
}
func (m *ObjectMapping) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
func (m *ObjectMapping) GetFlowTypes()(*ObjectFlowTypes) {
    if m == nil {
        return nil
    } else {
        return m.flowTypes
    }
}
func (m *ObjectMapping) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
func (m *ObjectMapping) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ObjectMapping) GetScope()(*Filter) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
func (m *ObjectMapping) GetSourceObjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceObjectName
    }
}
func (m *ObjectMapping) GetTargetObjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetObjectName
    }
}
func (m *ObjectMapping) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attributeMappings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeMapping() })
        if err != nil {
            return err
        }
        res := make([]AttributeMapping, len(val))
        for i, v := range val {
            res[i] = *(v.(*AttributeMapping))
        }
        m.SetAttributeMappings(res)
        return nil
    }
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnabled(val)
        return nil
    }
    res["flowTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectFlowTypes)
        if err != nil {
            return err
        }
        cast := val.(ObjectFlowTypes)
        m.SetFlowTypes(&cast)
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMetadataEntry() })
        if err != nil {
            return err
        }
        res := make([]MetadataEntry, len(val))
        for i, v := range val {
            res[i] = *(v.(*MetadataEntry))
        }
        m.SetMetadata(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilter() })
        if err != nil {
            return err
        }
        m.SetScope(val.(*Filter))
        return nil
    }
    res["sourceObjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceObjectName(val)
        return nil
    }
    res["targetObjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetObjectName(val)
        return nil
    }
    return res
}
func (m *ObjectMapping) IsNil()(bool) {
    return m == nil
}
func (m *ObjectMapping) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttributeMappings()))
        for i, v := range m.GetAttributeMappings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("attributeMappings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetFlowTypes() != nil {
        cast := m.GetFlowTypes().String()
        err := writer.WriteStringValue("flowTypes", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceObjectName", m.GetSourceObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetObjectName", m.GetTargetObjectName())
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
func (m *ObjectMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ObjectMapping) SetAttributeMappings(value []AttributeMapping)() {
    m.attributeMappings = value
}
func (m *ObjectMapping) SetEnabled(value *bool)() {
    m.enabled = value
}
func (m *ObjectMapping) SetFlowTypes(value *ObjectFlowTypes)() {
    m.flowTypes = value
}
func (m *ObjectMapping) SetMetadata(value []MetadataEntry)() {
    m.metadata = value
}
func (m *ObjectMapping) SetName(value *string)() {
    m.name = value
}
func (m *ObjectMapping) SetScope(value *Filter)() {
    m.scope = value
}
func (m *ObjectMapping) SetSourceObjectName(value *string)() {
    m.sourceObjectName = value
}
func (m *ObjectMapping) SetTargetObjectName(value *string)() {
    m.targetObjectName = value
}
