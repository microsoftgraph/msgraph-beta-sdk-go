package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeploymentSettings 
type DeploymentSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Settings governing conditions to monitor and automated actions to take.
    monitoring *MonitoringSettings;
    // Settings governing how the content is rolled out.
    rollout *RolloutSettings;
    // Settings governing safeguard holds on offering content.
    safeguard *SafeguardSettings;
}
// NewDeploymentSettings instantiates a new deploymentSettings and sets the default values.
func NewDeploymentSettings()(*DeploymentSettings) {
    m := &DeploymentSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMonitoring gets the monitoring property value. Settings governing conditions to monitor and automated actions to take.
func (m *DeploymentSettings) GetMonitoring()(*MonitoringSettings) {
    if m == nil {
        return nil
    } else {
        return m.monitoring
    }
}
// GetRollout gets the rollout property value. Settings governing how the content is rolled out.
func (m *DeploymentSettings) GetRollout()(*RolloutSettings) {
    if m == nil {
        return nil
    } else {
        return m.rollout
    }
}
// GetSafeguard gets the safeguard property value. Settings governing safeguard holds on offering content.
func (m *DeploymentSettings) GetSafeguard()(*SafeguardSettings) {
    if m == nil {
        return nil
    } else {
        return m.safeguard
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["monitoring"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMonitoringSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonitoring(val.(*MonitoringSettings))
        }
        return nil
    }
    res["rollout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRolloutSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRollout(val.(*RolloutSettings))
        }
        return nil
    }
    res["safeguard"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSafeguardSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafeguard(val.(*SafeguardSettings))
        }
        return nil
    }
    return res
}
func (m *DeploymentSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err := writer.WriteObjectValue("safeguard", m.GetSafeguard())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMonitoring sets the monitoring property value. Settings governing conditions to monitor and automated actions to take.
func (m *DeploymentSettings) SetMonitoring(value *MonitoringSettings)() {
    if m != nil {
        m.monitoring = value
    }
}
// SetRollout sets the rollout property value. Settings governing how the content is rolled out.
func (m *DeploymentSettings) SetRollout(value *RolloutSettings)() {
    if m != nil {
        m.rollout = value
    }
}
// SetSafeguard sets the safeguard property value. Settings governing safeguard holds on offering content.
func (m *DeploymentSettings) SetSafeguard(value *SafeguardSettings)() {
    if m != nil {
        m.safeguard = value
    }
}
