package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deploymentSettings and sets the default values.
func NewDeploymentSettings()(*DeploymentSettings) {
    m := &DeploymentSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the monitoring property value. Settings governing conditions to monitor and automated actions to take.
func (m *DeploymentSettings) GetMonitoring()(*MonitoringSettings) {
    if m == nil {
        return nil
    } else {
        return m.monitoring
    }
}
// Gets the rollout property value. Settings governing how the content is rolled out.
func (m *DeploymentSettings) GetRollout()(*RolloutSettings) {
    if m == nil {
        return nil
    } else {
        return m.rollout
    }
}
// Gets the safeguard property value. Settings governing safeguard holds on offering content.
func (m *DeploymentSettings) GetSafeguard()(*SafeguardSettings) {
    if m == nil {
        return nil
    } else {
        return m.safeguard
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeploymentSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the monitoring property value. Settings governing conditions to monitor and automated actions to take.
// Parameters:
//  - value : Value to set for the monitoring property.
func (m *DeploymentSettings) SetMonitoring(value *MonitoringSettings)() {
    m.monitoring = value
}
// Sets the rollout property value. Settings governing how the content is rolled out.
// Parameters:
//  - value : Value to set for the rollout property.
func (m *DeploymentSettings) SetRollout(value *RolloutSettings)() {
    m.rollout = value
}
// Sets the safeguard property value. Settings governing safeguard holds on offering content.
// Parameters:
//  - value : Value to set for the safeguard property.
func (m *DeploymentSettings) SetSafeguard(value *SafeguardSettings)() {
    m.safeguard = value
}
