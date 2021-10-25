package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type HyperlinkOrPictureColumn struct {
    additionalData map[string]interface{};
    isPicture *bool;
}
func NewHyperlinkOrPictureColumn()(*HyperlinkOrPictureColumn) {
    m := &HyperlinkOrPictureColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *HyperlinkOrPictureColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *HyperlinkOrPictureColumn) GetIsPicture()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPicture
    }
}
func (m *HyperlinkOrPictureColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isPicture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPicture(val)
        return nil
    }
    return res
}
func (m *HyperlinkOrPictureColumn) IsNil()(bool) {
    return m == nil
}
func (m *HyperlinkOrPictureColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isPicture", m.GetIsPicture())
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
func (m *HyperlinkOrPictureColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *HyperlinkOrPictureColumn) SetIsPicture(value *bool)() {
    m.isPicture = value
}
