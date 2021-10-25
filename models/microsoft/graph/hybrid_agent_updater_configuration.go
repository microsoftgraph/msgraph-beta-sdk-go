package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type HybridAgentUpdaterConfiguration struct {
    additionalData map[string]interface{};
    allowUpdateConfigurationOverride *bool;
    deferUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    updateWindow *UpdateWindow;
}
func NewHybridAgentUpdaterConfiguration()(*HybridAgentUpdaterConfiguration) {
    m := &HybridAgentUpdaterConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *HybridAgentUpdaterConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *HybridAgentUpdaterConfiguration) GetAllowUpdateConfigurationOverride()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUpdateConfigurationOverride
    }
}
func (m *HybridAgentUpdaterConfiguration) GetDeferUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deferUpdateDateTime
    }
}
func (m *HybridAgentUpdaterConfiguration) GetUpdateWindow()(*UpdateWindow) {
    if m == nil {
        return nil
    } else {
        return m.updateWindow
    }
}
func (m *HybridAgentUpdaterConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowUpdateConfigurationOverride"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowUpdateConfigurationOverride(val)
        return nil
    }
    res["deferUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeferUpdateDateTime(val)
        return nil
    }
    res["updateWindow"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdateWindow() })
        if err != nil {
            return err
        }
        m.SetUpdateWindow(val.(*UpdateWindow))
        return nil
    }
    return res
}
func (m *HybridAgentUpdaterConfiguration) IsNil()(bool) {
    return m == nil
}
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
func (m *HybridAgentUpdaterConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *HybridAgentUpdaterConfiguration) SetAllowUpdateConfigurationOverride(value *bool)() {
    m.allowUpdateConfigurationOverride = value
}
func (m *HybridAgentUpdaterConfiguration) SetDeferUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deferUpdateDateTime = value
}
func (m *HybridAgentUpdaterConfiguration) SetUpdateWindow(value *UpdateWindow)() {
    m.updateWindow = value
}
