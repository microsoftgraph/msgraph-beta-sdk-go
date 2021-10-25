package importresourceactions

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ImportResourceActionsRequestBody struct {
    additionalData map[string]interface{};
    format *string;
    overwriteResourceNamespace *bool;
    value *string;
}
func NewImportResourceActionsRequestBody()(*ImportResourceActionsRequestBody) {
    m := &ImportResourceActionsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ImportResourceActionsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ImportResourceActionsRequestBody) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *ImportResourceActionsRequestBody) GetOverwriteResourceNamespace()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overwriteResourceNamespace
    }
}
func (m *ImportResourceActionsRequestBody) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *ImportResourceActionsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFormat(val)
        return nil
    }
    res["overwriteResourceNamespace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOverwriteResourceNamespace(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *ImportResourceActionsRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ImportResourceActionsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteResourceNamespace", m.GetOverwriteResourceNamespace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ImportResourceActionsRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ImportResourceActionsRequestBody) SetFormat(value *string)() {
    m.format = value
}
func (m *ImportResourceActionsRequestBody) SetOverwriteResourceNamespace(value *bool)() {
    m.overwriteResourceNamespace = value
}
func (m *ImportResourceActionsRequestBody) SetValue(value *string)() {
    m.value = value
}
