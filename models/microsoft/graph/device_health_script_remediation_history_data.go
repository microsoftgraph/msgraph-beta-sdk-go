package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptRemediationHistoryData 
type DeviceHealthScriptRemediationHistoryData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date on which devices were remediated by the device health script.
    date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The number of devices that were found to have no issue by the device health script.
    noIssueDeviceCount *int32;
    // The number of devices remediated by the device health script.
    remediatedDeviceCount *int32;
}
// NewDeviceHealthScriptRemediationHistoryData instantiates a new deviceHealthScriptRemediationHistoryData and sets the default values.
func NewDeviceHealthScriptRemediationHistoryData()(*DeviceHealthScriptRemediationHistoryData) {
    m := &DeviceHealthScriptRemediationHistoryData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationHistoryData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDate gets the date property value. The date on which devices were remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.date
    }
}
// GetNoIssueDeviceCount gets the noIssueDeviceCount property value. The number of devices that were found to have no issue by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetNoIssueDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noIssueDeviceCount
    }
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. The number of devices remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRemediationHistoryData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["date"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["noIssueDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoIssueDeviceCount(val)
        }
        return nil
    }
    res["remediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRemediationHistoryData) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRemediationHistoryData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("noIssueDeviceCount", m.GetNoIssueDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
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
func (m *DeviceHealthScriptRemediationHistoryData) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDate sets the date property value. The date on which devices were remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) SetDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.date = value
    }
}
// SetNoIssueDeviceCount sets the noIssueDeviceCount property value. The number of devices that were found to have no issue by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) SetNoIssueDeviceCount(value *int32)() {
    if m != nil {
        m.noIssueDeviceCount = value
    }
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. The number of devices remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) SetRemediatedDeviceCount(value *int32)() {
    if m != nil {
        m.remediatedDeviceCount = value
    }
}
