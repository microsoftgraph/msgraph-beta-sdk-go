package parseexpression

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type ParseExpressionRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    expression *string;
    // 
    targetAttributeDefinition *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinition;
    // 
    testInputObject *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObject;
}
// Instantiates a new parseExpressionRequestBody and sets the default values.
func NewParseExpressionRequestBody()(*ParseExpressionRequestBody) {
    m := &ParseExpressionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParseExpressionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the expression property value. 
func (m *ParseExpressionRequestBody) GetExpression()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
// Gets the targetAttributeDefinition property value. 
func (m *ParseExpressionRequestBody) GetTargetAttributeDefinition()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinition) {
    if m == nil {
        return nil
    } else {
        return m.targetAttributeDefinition
    }
}
// Gets the testInputObject property value. 
func (m *ParseExpressionRequestBody) GetTestInputObject()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObject) {
    if m == nil {
        return nil
    } else {
        return m.testInputObject
    }
}
// The deserialization information for the current model
func (m *ParseExpressionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExpression(val)
        return nil
    }
    res["targetAttributeDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAttributeDefinition() })
        if err != nil {
            return err
        }
        m.SetTargetAttributeDefinition(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinition))
        return nil
    }
    res["testInputObject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExpressionInputObject() })
        if err != nil {
            return err
        }
        m.SetTestInputObject(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObject))
        return nil
    }
    return res
}
func (m *ParseExpressionRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ParseExpressionRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("targetAttributeDefinition", m.GetTargetAttributeDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("testInputObject", m.GetTestInputObject())
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
func (m *ParseExpressionRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the expression property value. 
// Parameters:
//  - value : Value to set for the expression property.
func (m *ParseExpressionRequestBody) SetExpression(value *string)() {
    m.expression = value
}
// Sets the targetAttributeDefinition property value. 
// Parameters:
//  - value : Value to set for the targetAttributeDefinition property.
func (m *ParseExpressionRequestBody) SetTargetAttributeDefinition(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinition)() {
    m.targetAttributeDefinition = value
}
// Sets the testInputObject property value. 
// Parameters:
//  - value : Value to set for the testInputObject property.
func (m *ParseExpressionRequestBody) SetTestInputObject(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObject)() {
    m.testInputObject = value
}
