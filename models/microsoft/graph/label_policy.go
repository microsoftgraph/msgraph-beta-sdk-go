package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LabelPolicy struct {
    additionalData map[string]interface{};
    id *string;
    name *string;
}
func NewLabelPolicy()(*LabelPolicy) {
    m := &LabelPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LabelPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LabelPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *LabelPolicy) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *LabelPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
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
func (m *LabelPolicy) IsNil()(bool) {
    return m == nil
}
func (m *LabelPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
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
func (m *LabelPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LabelPolicy) SetId(value *string)() {
    m.id = value
}
func (m *LabelPolicy) SetName(value *string)() {
    m.name = value
}
