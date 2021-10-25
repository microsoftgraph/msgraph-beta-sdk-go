package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type VersionAction struct {
    additionalData map[string]interface{};
    newVersion *string;
}
func NewVersionAction()(*VersionAction) {
    m := &VersionAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *VersionAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *VersionAction) GetNewVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newVersion
    }
}
func (m *VersionAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNewVersion(val)
        return nil
    }
    return res
}
func (m *VersionAction) IsNil()(bool) {
    return m == nil
}
func (m *VersionAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newVersion", m.GetNewVersion())
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
func (m *VersionAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *VersionAction) SetNewVersion(value *string)() {
    m.newVersion = value
}
