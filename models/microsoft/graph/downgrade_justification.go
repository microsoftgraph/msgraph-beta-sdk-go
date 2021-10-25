package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DowngradeJustification struct {
    additionalData map[string]interface{};
    isDowngradeJustified *bool;
    justificationMessage *string;
}
func NewDowngradeJustification()(*DowngradeJustification) {
    m := &DowngradeJustification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DowngradeJustification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DowngradeJustification) GetIsDowngradeJustified()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDowngradeJustified
    }
}
func (m *DowngradeJustification) GetJustificationMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justificationMessage
    }
}
func (m *DowngradeJustification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDowngradeJustified"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDowngradeJustified(val)
        return nil
    }
    res["justificationMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustificationMessage(val)
        return nil
    }
    return res
}
func (m *DowngradeJustification) IsNil()(bool) {
    return m == nil
}
func (m *DowngradeJustification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDowngradeJustified", m.GetIsDowngradeJustified())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationMessage", m.GetJustificationMessage())
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
func (m *DowngradeJustification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DowngradeJustification) SetIsDowngradeJustified(value *bool)() {
    m.isDowngradeJustified = value
}
func (m *DowngradeJustification) SetJustificationMessage(value *string)() {
    m.justificationMessage = value
}
