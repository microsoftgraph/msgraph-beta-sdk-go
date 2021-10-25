package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttributeDefinition struct {
    additionalData map[string]interface{};
    anchor *bool;
    apiExpressions []StringKeyStringValuePair;
    caseExact *bool;
    defaultValue *string;
    flowNullValues *bool;
    metadata []MetadataEntry;
    multivalued *bool;
    mutability *Mutability;
    name *string;
    referencedObjects []ReferencedObject;
    required *bool;
    type_escpaped *AttributeType;
}
func NewAttributeDefinition()(*AttributeDefinition) {
    m := &AttributeDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttributeDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttributeDefinition) GetAnchor()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.anchor
    }
}
func (m *AttributeDefinition) GetApiExpressions()([]StringKeyStringValuePair) {
    if m == nil {
        return nil
    } else {
        return m.apiExpressions
    }
}
func (m *AttributeDefinition) GetCaseExact()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.caseExact
    }
}
func (m *AttributeDefinition) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
func (m *AttributeDefinition) GetFlowNullValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flowNullValues
    }
}
func (m *AttributeDefinition) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
func (m *AttributeDefinition) GetMultivalued()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.multivalued
    }
}
func (m *AttributeDefinition) GetMutability()(*Mutability) {
    if m == nil {
        return nil
    } else {
        return m.mutability
    }
}
func (m *AttributeDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *AttributeDefinition) GetReferencedObjects()([]ReferencedObject) {
    if m == nil {
        return nil
    } else {
        return m.referencedObjects
    }
}
func (m *AttributeDefinition) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
func (m *AttributeDefinition) GetType_escpaped()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *AttributeDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["anchor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAnchor(val)
        return nil
    }
    res["apiExpressions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStringKeyStringValuePair() })
        if err != nil {
            return err
        }
        res := make([]StringKeyStringValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*StringKeyStringValuePair))
        }
        m.SetApiExpressions(res)
        return nil
    }
    res["caseExact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCaseExact(val)
        return nil
    }
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultValue(val)
        return nil
    }
    res["flowNullValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFlowNullValues(val)
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
    res["multivalued"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMultivalued(val)
        return nil
    }
    res["mutability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMutability)
        if err != nil {
            return err
        }
        cast := val.(Mutability)
        m.SetMutability(&cast)
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
    res["referencedObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReferencedObject() })
        if err != nil {
            return err
        }
        res := make([]ReferencedObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*ReferencedObject))
        }
        m.SetReferencedObjects(res)
        return nil
    }
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequired(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        cast := val.(AttributeType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *AttributeDefinition) IsNil()(bool) {
    return m == nil
}
func (m *AttributeDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("anchor", m.GetAnchor())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApiExpressions()))
        for i, v := range m.GetApiExpressions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("apiExpressions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("caseExact", m.GetCaseExact())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("flowNullValues", m.GetFlowNullValues())
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
        err := writer.WriteBoolValue("multivalued", m.GetMultivalued())
        if err != nil {
            return err
        }
    }
    if m.GetMutability() != nil {
        cast := m.GetMutability().String()
        err := writer.WriteStringValue("mutability", &cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReferencedObjects()))
        for i, v := range m.GetReferencedObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("referencedObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err := writer.WriteStringValue("type_escpaped", &cast)
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
func (m *AttributeDefinition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttributeDefinition) SetAnchor(value *bool)() {
    m.anchor = value
}
func (m *AttributeDefinition) SetApiExpressions(value []StringKeyStringValuePair)() {
    m.apiExpressions = value
}
func (m *AttributeDefinition) SetCaseExact(value *bool)() {
    m.caseExact = value
}
func (m *AttributeDefinition) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
func (m *AttributeDefinition) SetFlowNullValues(value *bool)() {
    m.flowNullValues = value
}
func (m *AttributeDefinition) SetMetadata(value []MetadataEntry)() {
    m.metadata = value
}
func (m *AttributeDefinition) SetMultivalued(value *bool)() {
    m.multivalued = value
}
func (m *AttributeDefinition) SetMutability(value *Mutability)() {
    m.mutability = value
}
func (m *AttributeDefinition) SetName(value *string)() {
    m.name = value
}
func (m *AttributeDefinition) SetReferencedObjects(value []ReferencedObject)() {
    m.referencedObjects = value
}
func (m *AttributeDefinition) SetRequired(value *bool)() {
    m.required = value
}
func (m *AttributeDefinition) SetType_escpaped(value *AttributeType)() {
    m.type_escpaped = value
}
