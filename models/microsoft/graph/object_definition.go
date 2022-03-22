package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ObjectDefinition 
type ObjectDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    attributes []AttributeDefinitionable;
    // 
    metadata []MetadataEntryable;
    // 
    name *string;
    // 
    supportedApis []string;
}
// NewObjectDefinition instantiates a new objectDefinition and sets the default values.
func NewObjectDefinition()(*ObjectDefinition) {
    m := &ObjectDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateObjectDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectDefinitionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewObjectDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttributes gets the attributes property value. 
func (m *ObjectDefinition) GetAttributes()([]AttributeDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.attributes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(AttributeDefinitionable)
            }
            m.SetAttributes(res)
        }
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetadataEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetadataEntryable, len(val))
            for i, v := range val {
                res[i] = v.(MetadataEntryable)
            }
            m.SetMetadata(res)
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
    res["supportedApis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedApis(res)
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. 
func (m *ObjectDefinition) GetMetadata()([]MetadataEntryable) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetName gets the name property value. 
func (m *ObjectDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSupportedApis gets the supportedApis property value. 
func (m *ObjectDefinition) GetSupportedApis()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedApis
    }
}
// Serialize serializes information the current object
func (m *ObjectDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAttributes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttributes()))
        for i, v := range m.GetAttributes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetSupportedApis() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttributes sets the attributes property value. 
func (m *ObjectDefinition) SetAttributes(value []AttributeDefinitionable)() {
    if m != nil {
        m.attributes = value
    }
}
// SetMetadata sets the metadata property value. 
func (m *ObjectDefinition) SetMetadata(value []MetadataEntryable)() {
    if m != nil {
        m.metadata = value
    }
}
// SetName sets the name property value. 
func (m *ObjectDefinition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSupportedApis sets the supportedApis property value. 
func (m *ObjectDefinition) SetSupportedApis(value []string)() {
    if m != nil {
        m.supportedApis = value
    }
}
