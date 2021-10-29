package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ObjectDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    attributes []AttributeDefinition;
    // 
    metadata []MetadataEntry;
    // 
    name *string;
    // 
    supportedApis []string;
}
// Instantiates a new objectDefinition and sets the default values.
func NewObjectDefinition()(*ObjectDefinition) {
    m := &ObjectDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attributes property value. 
func (m *ObjectDefinition) GetAttributes()([]AttributeDefinition) {
    if m == nil {
        return nil
    } else {
        return m.attributes
    }
}
// Gets the metadata property value. 
func (m *ObjectDefinition) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// Gets the name property value. 
func (m *ObjectDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the supportedApis property value. 
func (m *ObjectDefinition) GetSupportedApis()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedApis
    }
}
// The deserialization information for the current model
func (m *ObjectDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeDefinition() })
        if err != nil {
            return err
        }
        res := make([]AttributeDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*AttributeDefinition))
        }
        m.SetAttributes(res)
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
    res["supportedApis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSupportedApis(res)
        return nil
    }
    return res
}
func (m *ObjectDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ObjectDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttributes()))
        for i, v := range m.GetAttributes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("attributes", cast)
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
        err := writer.WriteCollectionOfStringValues("supportedApis", m.GetSupportedApis())
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
func (m *ObjectDefinition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attributes property value. 
// Parameters:
//  - value : Value to set for the attributes property.
func (m *ObjectDefinition) SetAttributes(value []AttributeDefinition)() {
    m.attributes = value
}
// Sets the metadata property value. 
// Parameters:
//  - value : Value to set for the metadata property.
func (m *ObjectDefinition) SetMetadata(value []MetadataEntry)() {
    m.metadata = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *ObjectDefinition) SetName(value *string)() {
    m.name = value
}
// Sets the supportedApis property value. 
// Parameters:
//  - value : Value to set for the supportedApis property.
func (m *ObjectDefinition) SetSupportedApis(value []string)() {
    m.supportedApis = value
}
