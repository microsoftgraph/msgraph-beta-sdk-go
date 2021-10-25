package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GovernanceNotificationTemplate struct {
    additionalData map[string]interface{};
    culture *string;
    id *string;
    source *string;
    type_escpaped *string;
    version *string;
}
func NewGovernanceNotificationTemplate()(*GovernanceNotificationTemplate) {
    m := &GovernanceNotificationTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GovernanceNotificationTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GovernanceNotificationTemplate) GetCulture()(*string) {
    if m == nil {
        return nil
    } else {
        return m.culture
    }
}
func (m *GovernanceNotificationTemplate) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *GovernanceNotificationTemplate) GetSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
func (m *GovernanceNotificationTemplate) GetType_escpaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *GovernanceNotificationTemplate) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *GovernanceNotificationTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["culture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCulture(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSource(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escpaped(val)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *GovernanceNotificationTemplate) IsNil()(bool) {
    return m == nil
}
func (m *GovernanceNotificationTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("culture", m.GetCulture())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escpaped", m.GetType_escpaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *GovernanceNotificationTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GovernanceNotificationTemplate) SetCulture(value *string)() {
    m.culture = value
}
func (m *GovernanceNotificationTemplate) SetId(value *string)() {
    m.id = value
}
func (m *GovernanceNotificationTemplate) SetSource(value *string)() {
    m.source = value
}
func (m *GovernanceNotificationTemplate) SetType_escpaped(value *string)() {
    m.type_escpaped = value
}
func (m *GovernanceNotificationTemplate) SetVersion(value *string)() {
    m.version = value
}
