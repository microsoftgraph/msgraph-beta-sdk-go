package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ParseExpressionResponse struct {
    additionalData map[string]interface{};
    error *PublicError;
    evaluationResult []string;
    evaluationSucceeded *bool;
    parsedExpression *AttributeMappingSource;
    parsingSucceeded *bool;
}
func NewParseExpressionResponse()(*ParseExpressionResponse) {
    m := &ParseExpressionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ParseExpressionResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ParseExpressionResponse) GetError()(*PublicError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *ParseExpressionResponse) GetEvaluationResult()([]string) {
    if m == nil {
        return nil
    } else {
        return m.evaluationResult
    }
}
func (m *ParseExpressionResponse) GetEvaluationSucceeded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.evaluationSucceeded
    }
}
func (m *ParseExpressionResponse) GetParsedExpression()(*AttributeMappingSource) {
    if m == nil {
        return nil
    } else {
        return m.parsedExpression
    }
}
func (m *ParseExpressionResponse) GetParsingSucceeded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.parsingSucceeded
    }
}
func (m *ParseExpressionResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*PublicError))
        return nil
    }
    res["evaluationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEvaluationResult(res)
        return nil
    }
    res["evaluationSucceeded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEvaluationSucceeded(val)
        return nil
    }
    res["parsedExpression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeMappingSource() })
        if err != nil {
            return err
        }
        m.SetParsedExpression(val.(*AttributeMappingSource))
        return nil
    }
    res["parsingSucceeded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetParsingSucceeded(val)
        return nil
    }
    return res
}
func (m *ParseExpressionResponse) IsNil()(bool) {
    return m == nil
}
func (m *ParseExpressionResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("evaluationResult", m.GetEvaluationResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("evaluationSucceeded", m.GetEvaluationSucceeded())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parsedExpression", m.GetParsedExpression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("parsingSucceeded", m.GetParsingSucceeded())
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
func (m *ParseExpressionResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ParseExpressionResponse) SetError(value *PublicError)() {
    m.error = value
}
func (m *ParseExpressionResponse) SetEvaluationResult(value []string)() {
    m.evaluationResult = value
}
func (m *ParseExpressionResponse) SetEvaluationSucceeded(value *bool)() {
    m.evaluationSucceeded = value
}
func (m *ParseExpressionResponse) SetParsedExpression(value *AttributeMappingSource)() {
    m.parsedExpression = value
}
func (m *ParseExpressionResponse) SetParsingSucceeded(value *bool)() {
    m.parsingSucceeded = value
}
