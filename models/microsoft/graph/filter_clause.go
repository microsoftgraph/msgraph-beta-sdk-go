package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new filterClause and sets the default values.
func NewFilterClause()(*FilterClause) {
    m := &FilterClause{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FilterClause) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the operatorName property value. Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
func (m *FilterClause) GetOperatorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatorName
    }
}
// Gets the sourceOperandName property value. Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
func (m *FilterClause) GetSourceOperandName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceOperandName
    }
}
// Gets the targetOperand property value. Values that the source operand will be tested against.
func (m *FilterClause) GetTargetOperand()(*FilterOperand) {
    if m == nil {
        return nil
    } else {
        return m.targetOperand
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *FilterClause) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the operatorName property value. Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
// Parameters:
//  - value : Value to set for the operatorName property.
func (m *FilterClause) SetOperatorName(value *string)() {
    m.operatorName = value
}
// Sets the sourceOperandName property value. Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
// Parameters:
//  - value : Value to set for the sourceOperandName property.
func (m *FilterClause) SetSourceOperandName(value *string)() {
    m.sourceOperandName = value
}
// Sets the targetOperand property value. Values that the source operand will be tested against.
// Parameters:
//  - value : Value to set for the targetOperand property.
func (m *FilterClause) SetTargetOperand(value *FilterOperand)() {
    m.targetOperand = value
}
