package parseexpression

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ParseExpressionRequestBody provides operations to call the parseExpression method.
type ParseExpressionRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    expression *string;
    // 
    targetAttributeDefinition i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinitionable;
    // 
    testInputObject i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObjectable;
}
// NewParseExpressionRequestBody instantiates a new parseExpressionRequestBody and sets the default values.
func NewParseExpressionRequestBody()(*ParseExpressionRequestBody) {
    m := &ParseExpressionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateParseExpressionRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParseExpressionRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewParseExpressionRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParseExpressionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpression gets the expression property value. 
func (m *ParseExpressionRequestBody) GetExpression()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParseExpressionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpression(val)
        }
        return nil
    }
    res["targetAttributeDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetAttributeDefinition(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinitionable))
        }
        return nil
    }
    res["testInputObject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateExpressionInputObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestInputObject(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObjectable))
        }
        return nil
    }
    return res
}
// GetTargetAttributeDefinition gets the targetAttributeDefinition property value. 
func (m *ParseExpressionRequestBody) GetTargetAttributeDefinition()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.targetAttributeDefinition
    }
}
// GetTestInputObject gets the testInputObject property value. 
func (m *ParseExpressionRequestBody) GetTestInputObject()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObjectable) {
    if m == nil {
        return nil
    } else {
        return m.testInputObject
    }
}
func (m *ParseExpressionRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParseExpressionRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpression sets the expression property value. 
func (m *ParseExpressionRequestBody) SetExpression(value *string)() {
    if m != nil {
        m.expression = value
    }
}
// SetTargetAttributeDefinition sets the targetAttributeDefinition property value. 
func (m *ParseExpressionRequestBody) SetTargetAttributeDefinition(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttributeDefinitionable)() {
    if m != nil {
        m.targetAttributeDefinition = value
    }
}
// SetTestInputObject sets the testInputObject property value. 
func (m *ParseExpressionRequestBody) SetTestInputObject(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExpressionInputObjectable)() {
    if m != nil {
        m.testInputObject = value
    }
}
