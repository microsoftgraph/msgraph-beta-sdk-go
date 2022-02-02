package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptRemediationHistory 
type DeviceHealthScriptRemediationHistory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of devices remediated by the device health script on the given date.
    historyData []DeviceHealthScriptRemediationHistoryData;
    // The date on which the results history is calculated for the healthscript.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewDeviceHealthScriptRemediationHistory instantiates a new deviceHealthScriptRemediationHistory and sets the default values.
func NewDeviceHealthScriptRemediationHistory()(*DeviceHealthScriptRemediationHistory) {
    m := &DeviceHealthScriptRemediationHistory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationHistory) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHistoryData gets the historyData property value. The number of devices remediated by the device health script on the given date.
func (m *DeviceHealthScriptRemediationHistory) GetHistoryData()([]DeviceHealthScriptRemediationHistoryData) {
    if m == nil {
        return nil
    } else {
        return m.historyData
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date on which the results history is calculated for the healthscript.
func (m *DeviceHealthScriptRemediationHistory) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRemediationHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["historyData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptRemediationHistoryData() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptRemediationHistoryData, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceHealthScriptRemediationHistoryData))
            }
            m.SetHistoryData(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRemediationHistory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRemediationHistory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetHistoryData() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistoryData()))
        for i, v := range m.GetHistoryData() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("historyData", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
func (m *DeviceHealthScriptRemediationHistory) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHistoryData sets the historyData property value. The number of devices remediated by the device health script on the given date.
func (m *DeviceHealthScriptRemediationHistory) SetHistoryData(value []DeviceHealthScriptRemediationHistoryData)() {
    if m != nil {
        m.historyData = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date on which the results history is calculated for the healthscript.
func (m *DeviceHealthScriptRemediationHistory) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
