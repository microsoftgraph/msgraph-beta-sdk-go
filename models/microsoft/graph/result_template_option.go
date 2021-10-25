package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ResultTemplateOption struct {
    additionalData map[string]interface{};
    enableResultTemplate *bool;
}
func NewResultTemplateOption()(*ResultTemplateOption) {
    m := &ResultTemplateOption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ResultTemplateOption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ResultTemplateOption) GetEnableResultTemplate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableResultTemplate
    }
}
func (m *ResultTemplateOption) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enableResultTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableResultTemplate(val)
        return nil
    }
    return res
}
func (m *ResultTemplateOption) IsNil()(bool) {
    return m == nil
}
func (m *ResultTemplateOption) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableResultTemplate", m.GetEnableResultTemplate())
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
func (m *ResultTemplateOption) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ResultTemplateOption) SetEnableResultTemplate(value *bool)() {
    m.enableResultTemplate = value
}
