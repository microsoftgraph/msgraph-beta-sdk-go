package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeploymentSettings struct {
    additionalData map[string]interface{};
    monitoring *MonitoringSettings;
    rollout *RolloutSettings;
}
func NewDeploymentSettings()(*DeploymentSettings) {
    m := &DeploymentSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeploymentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeploymentSettings) GetMonitoring()(*MonitoringSettings) {
    if m == nil {
        return nil
    } else {
        return m.monitoring
    }
}
func (m *DeploymentSettings) GetRollout()(*RolloutSettings) {
    if m == nil {
        return nil
    } else {
        return m.rollout
    }
}
func (m *DeploymentSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["monitoring"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMonitoringSettings() })
        if err != nil {
            return err
        }
        m.SetMonitoring(val.(*MonitoringSettings))
        return nil
    }
    res["rollout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRolloutSettings() })
        if err != nil {
            return err
        }
        m.SetRollout(val.(*RolloutSettings))
        return nil
    }
    return res
}
func (m *DeploymentSettings) IsNil()(bool) {
    return m == nil
}
func (m *DeploymentSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("monitoring", m.GetMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rollout", m.GetRollout())
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
func (m *DeploymentSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeploymentSettings) SetMonitoring(value *MonitoringSettings)() {
    m.monitoring = value
}
func (m *DeploymentSettings) SetRollout(value *RolloutSettings)() {
    m.rollout = value
}
