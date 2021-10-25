package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnPremisesPublishingSingleSignOn struct {
    additionalData map[string]interface{};
    kerberosSignOnSettings *KerberosSignOnSettings;
    singleSignOnMode *SingleSignOnMode;
}
func NewOnPremisesPublishingSingleSignOn()(*OnPremisesPublishingSingleSignOn) {
    m := &OnPremisesPublishingSingleSignOn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OnPremisesPublishingSingleSignOn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OnPremisesPublishingSingleSignOn) GetKerberosSignOnSettings()(*KerberosSignOnSettings) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnSettings
    }
}
func (m *OnPremisesPublishingSingleSignOn) GetSingleSignOnMode()(*SingleSignOnMode) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnMode
    }
}
func (m *OnPremisesPublishingSingleSignOn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["kerberosSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKerberosSignOnSettings() })
        if err != nil {
            return err
        }
        m.SetKerberosSignOnSettings(val.(*KerberosSignOnSettings))
        return nil
    }
    res["singleSignOnMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSingleSignOnMode)
        if err != nil {
            return err
        }
        cast := val.(SingleSignOnMode)
        m.SetSingleSignOnMode(&cast)
        return nil
    }
    return res
}
func (m *OnPremisesPublishingSingleSignOn) IsNil()(bool) {
    return m == nil
}
func (m *OnPremisesPublishingSingleSignOn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("kerberosSignOnSettings", m.GetKerberosSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSingleSignOnMode() != nil {
        cast := m.GetSingleSignOnMode().String()
        err := writer.WriteStringValue("singleSignOnMode", &cast)
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
func (m *OnPremisesPublishingSingleSignOn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OnPremisesPublishingSingleSignOn) SetKerberosSignOnSettings(value *KerberosSignOnSettings)() {
    m.kerberosSignOnSettings = value
}
func (m *OnPremisesPublishingSingleSignOn) SetSingleSignOnMode(value *SingleSignOnMode)() {
    m.singleSignOnMode = value
}
