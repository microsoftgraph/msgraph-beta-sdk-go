package evaluateremoval

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EvaluateRemoval struct {
    additionalData map[string]interface{};
}
func NewEvaluateRemoval()(*EvaluateRemoval) {
    m := &EvaluateRemoval{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EvaluateRemoval) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EvaluateRemoval) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *EvaluateRemoval) IsNil()(bool) {
    return m == nil
}
func (m *EvaluateRemoval) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EvaluateRemoval) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
