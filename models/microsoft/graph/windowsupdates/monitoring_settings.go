package windowsupdates

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MonitoringSettings 
type MonitoringSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the rules through which monitoring signals can trigger actions on the deployment. Rules are combined using 'or'.
    monitoringRules []MonitoringRuleable;
}
// NewMonitoringSettings instantiates a new monitoringSettings and sets the default values.
func NewMonitoringSettings()(*MonitoringSettings) {
    m := &MonitoringSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMonitoringSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMonitoringSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMonitoringSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MonitoringSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MonitoringSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["monitoringRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMonitoringRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MonitoringRuleable, len(val))
            for i, v := range val {
                res[i] = v.(MonitoringRuleable)
            }
            m.SetMonitoringRules(res)
        }
        return nil
    }
    return res
}
// GetMonitoringRules gets the monitoringRules property value. Specifies the rules through which monitoring signals can trigger actions on the deployment. Rules are combined using 'or'.
func (m *MonitoringSettings) GetMonitoringRules()([]MonitoringRuleable) {
    if m == nil {
        return nil
    } else {
        return m.monitoringRules
    }
}
// Serialize serializes information the current object
func (m *MonitoringSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMonitoringRules() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonitoringRules()))
        for i, v := range m.GetMonitoringRules() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MonitoringSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMonitoringRules sets the monitoringRules property value. Specifies the rules through which monitoring signals can trigger actions on the deployment. Rules are combined using 'or'.
func (m *MonitoringSettings) SetMonitoringRules(value []MonitoringRuleable)() {
    if m != nil {
        m.monitoringRules = value
    }
}
