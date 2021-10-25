package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OutOfOfficeSettings struct {
    additionalData map[string]interface{};
    isOutOfOffice *bool;
    message *string;
}
func NewOutOfOfficeSettings()(*OutOfOfficeSettings) {
    m := &OutOfOfficeSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OutOfOfficeSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OutOfOfficeSettings) GetIsOutOfOffice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOutOfOffice
    }
}
func (m *OutOfOfficeSettings) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *OutOfOfficeSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isOutOfOffice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOutOfOffice(val)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    return res
}
func (m *OutOfOfficeSettings) IsNil()(bool) {
    return m == nil
}
func (m *OutOfOfficeSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isOutOfOffice", m.GetIsOutOfOffice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *OutOfOfficeSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OutOfOfficeSettings) SetIsOutOfOffice(value *bool)() {
    m.isOutOfOffice = value
}
func (m *OutOfOfficeSettings) SetMessage(value *string)() {
    m.message = value
}
