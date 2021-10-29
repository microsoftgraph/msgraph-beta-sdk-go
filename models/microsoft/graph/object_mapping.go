package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ObjectMapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Attribute mappings define which attributes to map from the source object into the target object and how they should flow. A number of functions are available to support the transformation of the original source values.
    attributeMappings []AttributeMapping;
    // When true, this object mapping will be processed during synchronization. When false, this object mapping will be skipped.
    enabled *bool;
    // Which flow types are enabled for this object mapping. Add creates new objects in the target directory, Update modifies existing objects, and Delete deprovisions existing users. The default is Add, Update, Delete.
    flowTypes *ObjectFlowTypes;
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []MetadataEntry;
    // Human-friendly name of the object mapping.
    name *string;
    // Defines a filter to be used when deciding whether a given object should be provisioned. For example, you might want to only provision users that are located in the US.
    scope *Filter;
    // Name of the object in the source directory. Must match the object name from the source directory definition.
    sourceObjectName *string;
    // Name of the object in target directory. Must match the object name from the target directory definition.
    targetObjectName *string;
}
// Instantiates a new objectMapping and sets the default values.
func NewObjectMapping()(*ObjectMapping) {
    m := &ObjectMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectMapping) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attributeMappings property value. Attribute mappings define which attributes to map from the source object into the target object and how they should flow. A number of functions are available to support the transformation of the original source values.
func (m *ObjectMapping) GetAttributeMappings()([]AttributeMapping) {
    if m == nil {
        return nil
    } else {
        return m.attributeMappings
    }
}
// Gets the enabled property value. When true, this object mapping will be processed during synchronization. When false, this object mapping will be skipped.
func (m *ObjectMapping) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// Gets the flowTypes property value. Which flow types are enabled for this object mapping. Add creates new objects in the target directory, Update modifies existing objects, and Delete deprovisions existing users. The default is Add, Update, Delete.
func (m *ObjectMapping) GetFlowTypes()(*ObjectFlowTypes) {
    if m == nil {
        return nil
    } else {
        return m.flowTypes
    }
}
// Gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *ObjectMapping) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// Gets the name property value. Human-friendly name of the object mapping.
func (m *ObjectMapping) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the scope property value. Defines a filter to be used when deciding whether a given object should be provisioned. For example, you might want to only provision users that are located in the US.
func (m *ObjectMapping) GetScope()(*Filter) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Gets the sourceObjectName property value. Name of the object in the source directory. Must match the object name from the source directory definition.
func (m *ObjectMapping) GetSourceObjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceObjectName
    }
}
// Gets the targetObjectName property value. Name of the object in target directory. Must match the object name from the target directory definition.
func (m *ObjectMapping) GetTargetObjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetObjectName
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ObjectMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attributeMappings property value. Attribute mappings define which attributes to map from the source object into the target object and how they should flow. A number of functions are available to support the transformation of the original source values.
// Parameters:
//  - value : Value to set for the attributeMappings property.
func (m *ObjectMapping) SetAttributeMappings(value []AttributeMapping)() {
    m.attributeMappings = value
}
// Sets the enabled property value. When true, this object mapping will be processed during synchronization. When false, this object mapping will be skipped.
// Parameters:
//  - value : Value to set for the enabled property.
func (m *ObjectMapping) SetEnabled(value *bool)() {
    m.enabled = value
}
// Sets the flowTypes property value. Which flow types are enabled for this object mapping. Add creates new objects in the target directory, Update modifies existing objects, and Delete deprovisions existing users. The default is Add, Update, Delete.
// Parameters:
//  - value : Value to set for the flowTypes property.
func (m *ObjectMapping) SetFlowTypes(value *ObjectFlowTypes)() {
    m.flowTypes = value
}
// Sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
// Parameters:
//  - value : Value to set for the metadata property.
func (m *ObjectMapping) SetMetadata(value []MetadataEntry)() {
    m.metadata = value
}
// Sets the name property value. Human-friendly name of the object mapping.
// Parameters:
//  - value : Value to set for the name property.
func (m *ObjectMapping) SetName(value *string)() {
    m.name = value
}
// Sets the scope property value. Defines a filter to be used when deciding whether a given object should be provisioned. For example, you might want to only provision users that are located in the US.
// Parameters:
//  - value : Value to set for the scope property.
func (m *ObjectMapping) SetScope(value *Filter)() {
    m.scope = value
}
// Sets the sourceObjectName property value. Name of the object in the source directory. Must match the object name from the source directory definition.
// Parameters:
//  - value : Value to set for the sourceObjectName property.
func (m *ObjectMapping) SetSourceObjectName(value *string)() {
    m.sourceObjectName = value
}
// Sets the targetObjectName property value. Name of the object in target directory. Must match the object name from the target directory definition.
// Parameters:
//  - value : Value to set for the targetObjectName property.
func (m *ObjectMapping) SetTargetObjectName(value *string)() {
    m.targetObjectName = value
}
