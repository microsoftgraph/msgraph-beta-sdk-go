package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FilterOperand struct {
    additionalData map[string]interface{};
    values []string;
}
func NewFilterOperand()(*FilterOperand) {
    m := &FilterOperand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FilterOperand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FilterOperand) GetValues()([]string) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *FilterOperand) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetValues(res)
        return nil
    }
    return res
}
func (m *FilterOperand) IsNil()(bool) {
    return m == nil
}
func (m *FilterOperand) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
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
func (m *FilterOperand) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FilterOperand) SetValues(value []string)() {
    m.values = value
}
