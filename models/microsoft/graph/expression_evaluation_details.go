package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExpressionEvaluationDetails struct {
    additionalData map[string]interface{};
    expression *string;
    expressionEvaluationDetails []ExpressionEvaluationDetails;
    expressionResult *bool;
    propertyToEvaluate *PropertyToEvaluate;
}
func NewExpressionEvaluationDetails()(*ExpressionEvaluationDetails) {
    m := &ExpressionEvaluationDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExpressionEvaluationDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ExpressionEvaluationDetails) GetExpression()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
func (m *ExpressionEvaluationDetails) GetExpressionEvaluationDetails()([]ExpressionEvaluationDetails) {
    if m == nil {
        return nil
    } else {
        return m.expressionEvaluationDetails
    }
}
func (m *ExpressionEvaluationDetails) GetExpressionResult()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.expressionResult
    }
}
func (m *ExpressionEvaluationDetails) GetPropertyToEvaluate()(*PropertyToEvaluate) {
    if m == nil {
        return nil
    } else {
        return m.propertyToEvaluate
    }
}
func (m *ExpressionEvaluationDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExpression(val)
        return nil
    }
    res["expressionEvaluationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExpressionEvaluationDetails() })
        if err != nil {
            return err
        }
        res := make([]ExpressionEvaluationDetails, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExpressionEvaluationDetails))
        }
        m.SetExpressionEvaluationDetails(res)
        return nil
    }
    res["expressionResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetExpressionResult(val)
        return nil
    }
    res["propertyToEvaluate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPropertyToEvaluate() })
        if err != nil {
            return err
        }
        m.SetPropertyToEvaluate(val.(*PropertyToEvaluate))
        return nil
    }
    return res
}
func (m *ExpressionEvaluationDetails) IsNil()(bool) {
    return m == nil
}
func (m *ExpressionEvaluationDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    {
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
func (m *ExpressionEvaluationDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExpressionEvaluationDetails) SetExpression(value *string)() {
    m.expression = value
}
func (m *ExpressionEvaluationDetails) SetExpressionEvaluationDetails(value []ExpressionEvaluationDetails)() {
    m.expressionEvaluationDetails = value
}
func (m *ExpressionEvaluationDetails) SetExpressionResult(value *bool)() {
    m.expressionResult = value
}
func (m *ExpressionEvaluationDetails) SetPropertyToEvaluate(value *PropertyToEvaluate)() {
    m.propertyToEvaluate = value
}
