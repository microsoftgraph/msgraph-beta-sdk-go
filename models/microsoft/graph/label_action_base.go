package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LabelActionBase struct {
    additionalData map[string]interface{};
    name *string;
}
func NewLabelActionBase()(*LabelActionBase) {
    m := &LabelActionBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LabelActionBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LabelActionBase) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *LabelActionBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    return res
}
func (m *LabelActionBase) IsNil()(bool) {
    return m == nil
}
func (m *LabelActionBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *LabelActionBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LabelActionBase) SetName(value *string)() {
    m.name = value
}
