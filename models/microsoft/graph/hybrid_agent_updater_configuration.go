package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HybridAgentUpdaterConfiguration 
type HybridAgentUpdaterConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates if updater configuration will be skipped and the agent will receive an update when the next version of the agent is available.
    allowUpdateConfigurationOverride *bool;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    deferUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    updateWindow *UpdateWindow;
}
// NewHybridAgentUpdaterConfiguration instantiates a new hybridAgentUpdaterConfiguration and sets the default values.
func NewHybridAgentUpdaterConfiguration()(*HybridAgentUpdaterConfiguration) {
    m := &HybridAgentUpdaterConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HybridAgentUpdaterConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowUpdateConfigurationOverride gets the allowUpdateConfigurationOverride property value. Indicates if updater configuration will be skipped and the agent will receive an update when the next version of the agent is available.
func (m *HybridAgentUpdaterConfiguration) GetAllowUpdateConfigurationOverride()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUpdateConfigurationOverride
    }
}
// GetDeferUpdateDateTime gets the deferUpdateDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *HybridAgentUpdaterConfiguration) GetDeferUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deferUpdateDateTime
    }
}
// GetUpdateWindow gets the updateWindow property value. 
func (m *HybridAgentUpdaterConfiguration) GetUpdateWindow()(*UpdateWindow) {
    if m == nil {
        return nil
    } else {
        return m.updateWindow
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HybridAgentUpdaterConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowUpdateConfigurationOverride"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUpdateConfigurationOverride(val)
        }
        return nil
    }
    res["deferUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeferUpdateDateTime(val)
        }
        return nil
    }
    res["updateWindow"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdateWindow() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindow(val.(*UpdateWindow))
        }
        return nil
    }
    return res
}
func (m *HybridAgentUpdaterConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HybridAgentUpdaterConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowUpdateConfigurationOverride", m.GetAllowUpdateConfigurationOverride())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("deferUpdateDateTime", m.GetDeferUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("updateWindow", m.GetUpdateWindow())
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
func (m *HybridAgentUpdaterConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowUpdateConfigurationOverride sets the allowUpdateConfigurationOverride property value. Indicates if updater configuration will be skipped and the agent will receive an update when the next version of the agent is available.
func (m *HybridAgentUpdaterConfiguration) SetAllowUpdateConfigurationOverride(value *bool)() {
    if m != nil {
        m.allowUpdateConfigurationOverride = value
    }
}
// SetDeferUpdateDateTime sets the deferUpdateDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *HybridAgentUpdaterConfiguration) SetDeferUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deferUpdateDateTime = value
    }
}
// SetUpdateWindow sets the updateWindow property value. 
func (m *HybridAgentUpdaterConfiguration) SetUpdateWindow(value *UpdateWindow)() {
    if m != nil {
        m.updateWindow = value
    }
}
