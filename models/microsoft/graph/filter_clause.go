package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FilterClause struct {
    additionalData map[string]interface{};
    operatorName *string;
    sourceOperandName *string;
    targetOperand *FilterOperand;
}
func NewFilterClause()(*FilterClause) {
    m := &FilterClause{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FilterClause) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FilterClause) GetOperatorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatorName
    }
}
func (m *FilterClause) GetSourceOperandName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceOperandName
    }
}
func (m *FilterClause) GetTargetOperand()(*FilterOperand) {
    if m == nil {
        return nil
    } else {
        return m.targetOperand
    }
}
func (m *FilterClause) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["operatorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatorName(val)
        return nil
    }
    res["sourceOperandName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceOperandName(val)
        return nil
    }
    res["targetOperand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilterOperand() })
        if err != nil {
            return err
        }
        m.SetTargetOperand(val.(*FilterOperand))
        return nil
    }
    return res
}
func (m *FilterClause) IsNil()(bool) {
    return m == nil
}
func (m *FilterClause) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("operatorName", m.GetOperatorName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceOperandName", m.GetSourceOperandName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("targetOperand", m.GetTargetOperand())
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
func (m *FilterClause) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FilterClause) SetOperatorName(value *string)() {
    m.operatorName = value
}
func (m *FilterClause) SetSourceOperandName(value *string)() {
    m.sourceOperandName = value
}
func (m *FilterClause) SetTargetOperand(value *FilterOperand)() {
    m.targetOperand = value
}
