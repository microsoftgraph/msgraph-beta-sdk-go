package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExpressionInputObject 
type ExpressionInputObject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Definition of the test object.
    definition ObjectDefinitionable;
    // Property values of the test object.
    properties []StringKeyObjectValuePairable;
}
// NewExpressionInputObject instantiates a new expressionInputObject and sets the default values.
func NewExpressionInputObject()(*ExpressionInputObject) {
    m := &ExpressionInputObject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExpressionInputObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExpressionInputObjectFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExpressionInputObject(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpressionInputObject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefinition gets the definition property value. Definition of the test object.
func (m *ExpressionInputObject) GetDefinition()(ObjectDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpressionInputObject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateObjectDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(ObjectDefinitionable))
        }
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStringKeyObjectValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyObjectValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(StringKeyObjectValuePairable)
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. Property values of the test object.
func (m *ExpressionInputObject) GetProperties()([]StringKeyObjectValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// Serialize serializes information the current object
func (m *ExpressionInputObject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("properties", cast)
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
func (m *ExpressionInputObject) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefinition sets the definition property value. Definition of the test object.
func (m *ExpressionInputObject) SetDefinition(value ObjectDefinitionable)() {
    if m != nil {
        m.definition = value
    }
}
// SetProperties sets the properties property value. Property values of the test object.
func (m *ExpressionInputObject) SetProperties(value []StringKeyObjectValuePairable)() {
    if m != nil {
        m.properties = value
    }
}
