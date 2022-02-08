package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsAutopilotSettings 
type WindowsAutopilotSettings struct {
    Entity
    // Last data sync date time with DDS service.
    lastManualSyncTriggerDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last data sync date time with DDS service.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates the status of sync with Device data sync (DDS) service. Possible values are: unknown, inProgress, completed, failed.
    syncStatus *WindowsAutopilotSyncStatus;
}
// NewWindowsAutopilotSettings instantiates a new windowsAutopilotSettings and sets the default values.
func NewWindowsAutopilotSettings()(*WindowsAutopilotSettings) {
    m := &WindowsAutopilotSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetLastManualSyncTriggerDateTime gets the lastManualSyncTriggerDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) GetLastManualSyncTriggerDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastManualSyncTriggerDateTime
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetSyncStatus gets the syncStatus property value. Indicates the status of sync with Device data sync (DDS) service. Possible values are: unknown, inProgress, completed, failed.
func (m *WindowsAutopilotSettings) GetSyncStatus()(*WindowsAutopilotSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.syncStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastManualSyncTriggerDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastManualSyncTriggerDateTime(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["syncStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncStatus(val.(*WindowsAutopilotSyncStatus))
        }
        return nil
    }
    return res
}
func (m *WindowsAutopilotSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetSyncStatus()).String()
        err = writer.WriteStringValue("syncStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastManualSyncTriggerDateTime sets the lastManualSyncTriggerDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) SetLastManualSyncTriggerDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastManualSyncTriggerDateTime = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetSyncStatus sets the syncStatus property value. Indicates the status of sync with Device data sync (DDS) service. Possible values are: unknown, inProgress, completed, failed.
func (m *WindowsAutopilotSettings) SetSyncStatus(value *WindowsAutopilotSyncStatus)() {
    if m != nil {
        m.syncStatus = value
    }
}
