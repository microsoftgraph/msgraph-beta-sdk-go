package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MonitoringSettings struct {
    additionalData map[string]interface{};
    monitoringRules []MonitoringRule;
}
func NewMonitoringSettings()(*MonitoringSettings) {
    m := &MonitoringSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MonitoringSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MonitoringSettings) GetMonitoringRules()([]MonitoringRule) {
    if m == nil {
        return nil
    } else {
        return m.monitoringRules
    }
}
func (m *MonitoringSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["monitoringRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMonitoringRule() })
        if err != nil {
            return err
        }
        res := make([]MonitoringRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*MonitoringRule))
        }
        m.SetMonitoringRules(res)
        return nil
    }
    return res
}
func (m *MonitoringSettings) IsNil()(bool) {
    return m == nil
}
func (m *MonitoringSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonitoringRules()))
        for i, v := range m.GetMonitoringRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("monitoringRules", cast)
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
func (m *MonitoringSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MonitoringSettings) SetMonitoringRules(value []MonitoringRule)() {
    m.monitoringRules = value
}
