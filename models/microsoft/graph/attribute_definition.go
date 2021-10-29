package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttributeDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
    anchor *bool;
    // 
    apiExpressions []StringKeyStringValuePair;
    // true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
    caseExact *bool;
    // 
    defaultValue *string;
    // 'true' to allow null values for attributes.
    flowNullValues *bool;
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []MetadataEntry;
    // true if an attribute can have multiple values. Default is false.
    multivalued *bool;
    // An attribute's mutability. Possible values are:  ReadWrite, ReadOnly, Immutable, WriteOnly. Default is ReadWrite.
    mutability *Mutability;
    // Name of the attribute. Must be unique within the object definition. Not nullable.
    name *string;
    // For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
    referencedObjects []ReferencedObject;
    // true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
    required *bool;
    // Attribute value type. Possible values are: String, Integer, Reference, Binary, Boolean,DateTime. Default is String.
    type_escaped *AttributeType;
}
// Instantiates a new attributeDefinition and sets the default values.
func NewAttributeDefinition()(*AttributeDefinition) {
    m := &AttributeDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
func (m *AttributeDefinition) GetAnchor()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.anchor
    }
}
// Gets the apiExpressions property value. 
func (m *AttributeDefinition) GetApiExpressions()([]StringKeyStringValuePair) {
    if m == nil {
        return nil
    } else {
        return m.apiExpressions
    }
}
// Gets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
func (m *AttributeDefinition) GetCaseExact()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.caseExact
    }
}
// Gets the defaultValue property value. 
func (m *AttributeDefinition) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// Gets the flowNullValues property value. 'true' to allow null values for attributes.
func (m *AttributeDefinition) GetFlowNullValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flowNullValues
    }
}
// Gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *AttributeDefinition) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// Gets the multivalued property value. true if an attribute can have multiple values. Default is false.
func (m *AttributeDefinition) GetMultivalued()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.multivalued
    }
}
// Gets the mutability property value. An attribute's mutability. Possible values are:  ReadWrite, ReadOnly, Immutable, WriteOnly. Default is ReadWrite.
func (m *AttributeDefinition) GetMutability()(*Mutability) {
    if m == nil {
        return nil
    } else {
        return m.mutability
    }
}
// Gets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
func (m *AttributeDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
func (m *AttributeDefinition) GetReferencedObjects()([]ReferencedObject) {
    if m == nil {
        return nil
    } else {
        return m.referencedObjects
    }
}
// Gets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
func (m *AttributeDefinition) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// Gets the type_escaped property value. Attribute value type. Possible values are: String, Integer, Reference, Binary, Boolean,DateTime. Default is String.
func (m *AttributeDefinition) GetType_escaped()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        cast := val.(AttributeType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *AttributeDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *AttributeDefinition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
// Parameters:
//  - value : Value to set for the anchor property.
func (m *AttributeDefinition) SetAnchor(value *bool)() {
    m.anchor = value
}
// Sets the apiExpressions property value. 
// Parameters:
//  - value : Value to set for the apiExpressions property.
func (m *AttributeDefinition) SetApiExpressions(value []StringKeyStringValuePair)() {
    m.apiExpressions = value
}
// Sets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
// Parameters:
//  - value : Value to set for the caseExact property.
func (m *AttributeDefinition) SetCaseExact(value *bool)() {
    m.caseExact = value
}
// Sets the defaultValue property value. 
// Parameters:
//  - value : Value to set for the defaultValue property.
func (m *AttributeDefinition) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// Sets the flowNullValues property value. 'true' to allow null values for attributes.
// Parameters:
//  - value : Value to set for the flowNullValues property.
func (m *AttributeDefinition) SetFlowNullValues(value *bool)() {
    m.flowNullValues = value
}
// Sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
// Parameters:
//  - value : Value to set for the metadata property.
func (m *AttributeDefinition) SetMetadata(value []MetadataEntry)() {
    m.metadata = value
}
// Sets the multivalued property value. true if an attribute can have multiple values. Default is false.
// Parameters:
//  - value : Value to set for the multivalued property.
func (m *AttributeDefinition) SetMultivalued(value *bool)() {
    m.multivalued = value
}
// Sets the mutability property value. An attribute's mutability. Possible values are:  ReadWrite, ReadOnly, Immutable, WriteOnly. Default is ReadWrite.
// Parameters:
//  - value : Value to set for the mutability property.
func (m *AttributeDefinition) SetMutability(value *Mutability)() {
    m.mutability = value
}
// Sets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
// Parameters:
//  - value : Value to set for the name property.
func (m *AttributeDefinition) SetName(value *string)() {
    m.name = value
}
// Sets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
// Parameters:
//  - value : Value to set for the referencedObjects property.
func (m *AttributeDefinition) SetReferencedObjects(value []ReferencedObject)() {
    m.referencedObjects = value
}
// Sets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
// Parameters:
//  - value : Value to set for the required property.
func (m *AttributeDefinition) SetRequired(value *bool)() {
    m.required = value
}
// Sets the type_escaped property value. Attribute value type. Possible values are: String, Integer, Reference, Binary, Boolean,DateTime. Default is String.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *AttributeDefinition) SetType_escaped(value *AttributeType)() {
    m.type_escaped = value
}
