package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceHealthScriptRemediationHistoryData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date on which devices were remediated by the device health script.
    date *string;
    // The number of devices that were found to have no issue by the device health script.
    noIssueDeviceCount *int32;
    // The number of devices remediated by the device health script.
    remediatedDeviceCount *int32;
}
// Instantiates a new deviceHealthScriptRemediationHistoryData and sets the default values.
func NewDeviceHealthScriptRemediationHistoryData()(*DeviceHealthScriptRemediationHistoryData) {
    m := &DeviceHealthScriptRemediationHistoryData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationHistoryData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the date property value. The date on which devices were remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.date
    }
}
// Gets the noIssueDeviceCount property value. The number of devices that were found to have no issue by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetNoIssueDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noIssueDeviceCount
    }
}
// Gets the remediatedDeviceCount property value. The number of devices remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// The deserialization information for the current model
func (m *DeviceHealthScriptRemediationHistoryData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["date"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceHealthScriptRemediationHistoryData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("date", m.GetDate())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceHealthScriptRemediationHistoryData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the date property value. The date on which devices were remediated by the device health script.
// Parameters:
//  - value : Value to set for the date property.
func (m *DeviceHealthScriptRemediationHistoryData) SetDate(value *string)() {
    m.date = value
}
// Sets the noIssueDeviceCount property value. The number of devices that were found to have no issue by the device health script.
// Parameters:
//  - value : Value to set for the noIssueDeviceCount property.
func (m *DeviceHealthScriptRemediationHistoryData) SetNoIssueDeviceCount(value *int32)() {
    m.noIssueDeviceCount = value
}
// Sets the remediatedDeviceCount property value. The number of devices remediated by the device health script.
// Parameters:
//  - value : Value to set for the remediatedDeviceCount property.
func (m *DeviceHealthScriptRemediationHistoryData) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
