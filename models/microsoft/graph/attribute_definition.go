package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeDefinition 
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
// NewAttributeDefinition instantiates a new attributeDefinition and sets the default values.
func NewAttributeDefinition()(*AttributeDefinition) {
    m := &AttributeDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAnchor gets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
func (m *AttributeDefinition) GetAnchor()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.anchor
    }
}
// GetApiExpressions gets the apiExpressions property value. 
func (m *AttributeDefinition) GetApiExpressions()([]StringKeyStringValuePair) {
    if m == nil {
        return nil
    } else {
        return m.apiExpressions
    }
}
// GetCaseExact gets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
func (m *AttributeDefinition) GetCaseExact()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.caseExact
    }
}
// GetDefaultValue gets the defaultValue property value. 
func (m *AttributeDefinition) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetFlowNullValues gets the flowNullValues property value. 'true' to allow null values for attributes.
func (m *AttributeDefinition) GetFlowNullValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flowNullValues
    }
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *AttributeDefinition) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetMultivalued gets the multivalued property value. true if an attribute can have multiple values. Default is false.
func (m *AttributeDefinition) GetMultivalued()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.multivalued
    }
}
// GetMutability gets the mutability property value. An attribute's mutability. Possible values are:  ReadWrite, ReadOnly, Immutable, WriteOnly. Default is ReadWrite.
func (m *AttributeDefinition) GetMutability()(*Mutability) {
    if m == nil {
        return nil
    } else {
        return m.mutability
    }
}
// GetName gets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
func (m *AttributeDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetReferencedObjects gets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
func (m *AttributeDefinition) GetReferencedObjects()([]ReferencedObject) {
    if m == nil {
        return nil
    } else {
        return m.referencedObjects
    }
}
// GetRequired gets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
func (m *AttributeDefinition) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// GetType gets the type property value. Attribute value type. Possible values are: String, Integer, Reference, Binary, Boolean,DateTime. Default is String.
func (m *AttributeDefinition) GetType()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["anchor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnchor(val)
        }
        return nil
    }
    res["apiExpressions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStringKeyStringValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyStringValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*StringKeyStringValuePair))
            }
            m.SetApiExpressions(res)
        }
        return nil
    }
    res["caseExact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaseExact(val)
        }
        return nil
    }
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["flowNullValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlowNullValues(val)
        }
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMetadataEntry() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetadataEntry, len(val))
            for i, v := range val {
                res[i] = *(v.(*MetadataEntry))
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["multivalued"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultivalued(val)
        }
        return nil
    }
    res["mutability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMutability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMutability(val.(*Mutability))
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["referencedObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReferencedObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReferencedObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*ReferencedObject))
            }
            m.SetReferencedObjects(res)
        }
        return nil
    }
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*AttributeType))
        }
        return nil
    }
    return res
}
func (m *AttributeDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttributeDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("anchor", m.GetAnchor())
        if err != nil {
            return err
        }
    }
    if m.GetApiExpressions() != nil {
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
    if m.GetMetadata() != nil {
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
        cast := (*m.GetMutability()).String()
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
    if m.GetReferencedObjects() != nil {
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AttributeDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAnchor sets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
func (m *AttributeDefinition) SetAnchor(value *bool)() {
    if m != nil {
        m.anchor = value
    }
}
// SetApiExpressions sets the apiExpressions property value. 
func (m *AttributeDefinition) SetApiExpressions(value []StringKeyStringValuePair)() {
    if m != nil {
        m.apiExpressions = value
    }
}
// SetCaseExact sets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
func (m *AttributeDefinition) SetCaseExact(value *bool)() {
    if m != nil {
        m.caseExact = value
    }
}
// SetDefaultValue sets the defaultValue property value. 
func (m *AttributeDefinition) SetDefaultValue(value *string)() {
    if m != nil {
        m.defaultValue = value
    }
}
// SetFlowNullValues sets the flowNullValues property value. 'true' to allow null values for attributes.
func (m *AttributeDefinition) SetFlowNullValues(value *bool)() {
    if m != nil {
        m.flowNullValues = value
    }
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *AttributeDefinition) SetMetadata(value []MetadataEntry)() {
    if m != nil {
        m.metadata = value
    }
}
// SetMultivalued sets the multivalued property value. true if an attribute can have multiple values. Default is false.
func (m *AttributeDefinition) SetMultivalued(value *bool)() {
    if m != nil {
        m.multivalued = value
    }
}
// SetMutability sets the mutability property value. An attribute's mutability. Possible values are:  ReadWrite, ReadOnly, Immutable, WriteOnly. Default is ReadWrite.
func (m *AttributeDefinition) SetMutability(value *Mutability)() {
    if m != nil {
        m.mutability = value
    }
}
// SetName sets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
func (m *AttributeDefinition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetReferencedObjects sets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
func (m *AttributeDefinition) SetReferencedObjects(value []ReferencedObject)() {
    if m != nil {
        m.referencedObjects = value
    }
}
// SetRequired sets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
func (m *AttributeDefinition) SetRequired(value *bool)() {
    if m != nil {
        m.required = value
    }
}
// SetType sets the type property value. Attribute value type. Possible values are: String, Integer, Reference, Binary, Boolean,DateTime. Default is String.
func (m *AttributeDefinition) SetType(value *AttributeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
