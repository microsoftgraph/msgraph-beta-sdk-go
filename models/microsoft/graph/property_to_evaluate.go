package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PropertyToEvaluate struct {
    additionalData map[string]interface{};
    propertyName *string;
    propertyValue *string;
}
func NewPropertyToEvaluate()(*PropertyToEvaluate) {
    m := &PropertyToEvaluate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PropertyToEvaluate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PropertyToEvaluate) GetPropertyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyName
    }
}
func (m *PropertyToEvaluate) GetPropertyValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyValue
    }
}
func (m *PropertyToEvaluate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["propertyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPropertyName(val)
        return nil
    }
    res["propertyValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPropertyValue(val)
        return nil
    }
    return res
}
func (m *PropertyToEvaluate) IsNil()(bool) {
    return m == nil
}
func (m *PropertyToEvaluate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("propertyName", m.GetPropertyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("propertyValue", m.GetPropertyValue())
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
func (m *PropertyToEvaluate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PropertyToEvaluate) SetPropertyName(value *string)() {
    m.propertyName = value
}
func (m *PropertyToEvaluate) SetPropertyValue(value *string)() {
    m.propertyValue = value
}
