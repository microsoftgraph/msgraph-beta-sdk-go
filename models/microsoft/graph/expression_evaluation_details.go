package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExpressionEvaluationDetails 
type ExpressionEvaluationDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents expression which has been evaluated.
    expression *string;
    // Represents the details of the evaluation of the expression.
    expressionEvaluationDetails []ExpressionEvaluationDetails;
    // Represents the value of the result of the current expression.
    expressionResult *bool;
    // Defines the name of the property and the value of that property.
    propertyToEvaluate *PropertyToEvaluate;
}
// NewExpressionEvaluationDetails instantiates a new expressionEvaluationDetails and sets the default values.
func NewExpressionEvaluationDetails()(*ExpressionEvaluationDetails) {
    m := &ExpressionEvaluationDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpressionEvaluationDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpression gets the expression property value. Represents expression which has been evaluated.
func (m *ExpressionEvaluationDetails) GetExpression()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
// GetExpressionEvaluationDetails gets the expressionEvaluationDetails property value. Represents the details of the evaluation of the expression.
func (m *ExpressionEvaluationDetails) GetExpressionEvaluationDetails()([]ExpressionEvaluationDetails) {
    if m == nil {
        return nil
    } else {
        return m.expressionEvaluationDetails
    }
}
// GetExpressionResult gets the expressionResult property value. Represents the value of the result of the current expression.
func (m *ExpressionEvaluationDetails) GetExpressionResult()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.expressionResult
    }
}
// GetPropertyToEvaluate gets the propertyToEvaluate property value. Defines the name of the property and the value of that property.
func (m *ExpressionEvaluationDetails) GetPropertyToEvaluate()(*PropertyToEvaluate) {
    if m == nil {
        return nil
    } else {
        return m.propertyToEvaluate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpressionEvaluationDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["expressionEvaluationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExpressionEvaluationDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExpressionEvaluationDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExpressionEvaluationDetails))
            }
            m.SetExpressionEvaluationDetails(res)
        }
        return nil
    }
    res["expressionResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpressionResult(val)
        }
        return nil
    }
    res["propertyToEvaluate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPropertyToEvaluate() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyToEvaluate(val.(*PropertyToEvaluate))
        }
        return nil
    }
    return res
}
func (m *ExpressionEvaluationDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExpressionEvaluationDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    if m.GetExpressionEvaluationDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExpressionEvaluationDetails()))
        for i, v := range m.GetExpressionEvaluationDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("expressionEvaluationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("expressionResult", m.GetExpressionResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("propertyToEvaluate", m.GetPropertyToEvaluate())
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
func (m *ExpressionEvaluationDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpression sets the expression property value. Represents expression which has been evaluated.
func (m *ExpressionEvaluationDetails) SetExpression(value *string)() {
    if m != nil {
        m.expression = value
    }
}
// SetExpressionEvaluationDetails sets the expressionEvaluationDetails property value. Represents the details of the evaluation of the expression.
func (m *ExpressionEvaluationDetails) SetExpressionEvaluationDetails(value []ExpressionEvaluationDetails)() {
    if m != nil {
        m.expressionEvaluationDetails = value
    }
}
// SetExpressionResult sets the expressionResult property value. Represents the value of the result of the current expression.
func (m *ExpressionEvaluationDetails) SetExpressionResult(value *bool)() {
    if m != nil {
        m.expressionResult = value
    }
}
// SetPropertyToEvaluate sets the propertyToEvaluate property value. Defines the name of the property and the value of that property.
func (m *ExpressionEvaluationDetails) SetPropertyToEvaluate(value *PropertyToEvaluate)() {
    if m != nil {
        m.propertyToEvaluate = value
    }
}
