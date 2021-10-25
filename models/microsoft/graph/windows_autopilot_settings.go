package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsAutopilotSettings struct {
    Entity
    lastManualSyncTriggerDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    syncStatus *WindowsAutopilotSyncStatus;
}
func NewWindowsAutopilotSettings()(*WindowsAutopilotSettings) {
    m := &WindowsAutopilotSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsAutopilotSettings) GetLastManualSyncTriggerDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastManualSyncTriggerDateTime
    }
}
func (m *WindowsAutopilotSettings) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
func (m *WindowsAutopilotSettings) GetSyncStatus()(*WindowsAutopilotSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.syncStatus
    }
}
func (m *WindowsAutopilotSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastManualSyncTriggerDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastManualSyncTriggerDateTime(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["syncStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotSyncStatus)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotSyncStatus)
        m.SetSyncStatus(&cast)
        return nil
    }
    return res
}
func (m *WindowsAutopilotSettings) IsNil()(bool) {
    return m == nil
}
func (m *WindowsAutopilotSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastManualSyncTriggerDateTime", m.GetLastManualSyncTriggerDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSyncStatus() != nil {
        cast := m.GetSyncStatus().String()
        err = writer.WriteStringValue("syncStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsAutopilotSettings) SetLastManualSyncTriggerDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastManualSyncTriggerDateTime = value
}
func (m *WindowsAutopilotSettings) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
func (m *WindowsAutopilotSettings) SetSyncStatus(value *WindowsAutopilotSyncStatus)() {
    m.syncStatus = value
}
