package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FilterClause 
type FilterClause struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
    operatorName *string;
    // Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
    sourceOperandName *string;
    // Values that the source operand will be tested against.
    targetOperand *FilterOperand;
}
// NewFilterClause instantiates a new filterClause and sets the default values.
func NewFilterClause()(*FilterClause) {
    m := &FilterClause{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FilterClause) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOperatorName gets the operatorName property value. Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
func (m *FilterClause) GetOperatorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatorName
    }
}
// GetSourceOperandName gets the sourceOperandName property value. Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
func (m *FilterClause) GetSourceOperandName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceOperandName
    }
}
// GetTargetOperand gets the targetOperand property value. Values that the source operand will be tested against.
func (m *FilterClause) GetTargetOperand()(*FilterOperand) {
    if m == nil {
        return nil
    } else {
        return m.targetOperand
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FilterClause) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["operatorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatorName(val)
        }
        return nil
    }
    res["sourceOperandName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceOperandName(val)
        }
        return nil
    }
    res["targetOperand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilterOperand() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetOperand(val.(*FilterOperand))
        }
        return nil
    }
    return res
}
func (m *FilterClause) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FilterClause) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOperatorName sets the operatorName property value. Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
func (m *FilterClause) SetOperatorName(value *string)() {
    if m != nil {
        m.operatorName = value
    }
}
// SetSourceOperandName sets the sourceOperandName property value. Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
func (m *FilterClause) SetSourceOperandName(value *string)() {
    if m != nil {
        m.sourceOperandName = value
    }
}
// SetTargetOperand sets the targetOperand property value. Values that the source operand will be tested against.
func (m *FilterClause) SetTargetOperand(value *FilterOperand)() {
    if m != nil {
        m.targetOperand = value
    }
}
