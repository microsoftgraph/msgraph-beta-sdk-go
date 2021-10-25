package migratetotemplate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MigrateToTemplateRequestBody struct {
    additionalData map[string]interface{};
    newTemplateId *string;
    preserveCustomValues *bool;
}
func NewMigrateToTemplateRequestBody()(*MigrateToTemplateRequestBody) {
    m := &MigrateToTemplateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MigrateToTemplateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MigrateToTemplateRequestBody) GetNewTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newTemplateId
    }
}
func (m *MigrateToTemplateRequestBody) GetPreserveCustomValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preserveCustomValues
    }
}
func (m *MigrateToTemplateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNewTemplateId(val)
        return nil
    }
    res["preserveCustomValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPreserveCustomValues(val)
        return nil
    }
    return res
}
func (m *MigrateToTemplateRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *MigrateToTemplateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newTemplateId", m.GetNewTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preserveCustomValues", m.GetPreserveCustomValues())
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
func (m *MigrateToTemplateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MigrateToTemplateRequestBody) SetNewTemplateId(value *string)() {
    m.newTemplateId = value
}
func (m *MigrateToTemplateRequestBody) SetPreserveCustomValues(value *bool)() {
    m.preserveCustomValues = value
}
