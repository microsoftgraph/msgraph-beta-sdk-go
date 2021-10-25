package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PayloadTypes struct {
    additionalData map[string]interface{};
    rawContent *string;
    visualContent *VisualProperties;
}
func NewPayloadTypes()(*PayloadTypes) {
    m := &PayloadTypes{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PayloadTypes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PayloadTypes) GetRawContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rawContent
    }
}
func (m *PayloadTypes) GetVisualContent()(*VisualProperties) {
    if m == nil {
        return nil
    } else {
        return m.visualContent
    }
}
func (m *PayloadTypes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["rawContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRawContent(val)
        return nil
    }
    res["visualContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVisualProperties() })
        if err != nil {
            return err
        }
        m.SetVisualContent(val.(*VisualProperties))
        return nil
    }
    return res
}
func (m *PayloadTypes) IsNil()(bool) {
    return m == nil
}
func (m *PayloadTypes) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("rawContent", m.GetRawContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("visualContent", m.GetVisualContent())
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
func (m *PayloadTypes) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PayloadTypes) SetRawContent(value *string)() {
    m.rawContent = value
}
func (m *PayloadTypes) SetVisualContent(value *VisualProperties)() {
    m.visualContent = value
}
