package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PasswordSingleSignOnField struct {
    additionalData map[string]interface{};
    customizedLabel *string;
    defaultLabel *string;
    fieldId *string;
    type_escaped *string;
}
func NewPasswordSingleSignOnField()(*PasswordSingleSignOnField) {
    m := &PasswordSingleSignOnField{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PasswordSingleSignOnField) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PasswordSingleSignOnField) GetCustomizedLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customizedLabel
    }
}
func (m *PasswordSingleSignOnField) GetDefaultLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLabel
    }
}
func (m *PasswordSingleSignOnField) GetFieldId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fieldId
    }
}
func (m *PasswordSingleSignOnField) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *PasswordSingleSignOnField) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["customizedLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomizedLabel(val)
        return nil
    }
    res["defaultLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultLabel(val)
        return nil
    }
    res["fieldId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFieldId(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *PasswordSingleSignOnField) IsNil()(bool) {
    return m == nil
}
func (m *PasswordSingleSignOnField) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customizedLabel", m.GetCustomizedLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultLabel", m.GetDefaultLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fieldId", m.GetFieldId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
func (m *PasswordSingleSignOnField) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PasswordSingleSignOnField) SetCustomizedLabel(value *string)() {
    m.customizedLabel = value
}
func (m *PasswordSingleSignOnField) SetDefaultLabel(value *string)() {
    m.defaultLabel = value
}
func (m *PasswordSingleSignOnField) SetFieldId(value *string)() {
    m.fieldId = value
}
func (m *PasswordSingleSignOnField) SetType_escaped(value *string)() {
    m.type_escaped = value
}
