package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceHealthScriptRemediationSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of devices remediated by device health scripts.
    remediatedDeviceCount *int32;
    // The number of device health scripts deployed.
    scriptCount *int32;
}
// Instantiates a new deviceHealthScriptRemediationSummary and sets the default values.
func NewDeviceHealthScriptRemediationSummary()(*DeviceHealthScriptRemediationSummary) {
    m := &DeviceHealthScriptRemediationSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the remediatedDeviceCount property value. The number of devices remediated by device health scripts.
func (m *DeviceHealthScriptRemediationSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// Gets the scriptCount property value. The number of device health scripts deployed.
func (m *DeviceHealthScriptRemediationSummary) GetScriptCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.scriptCount
    }
}
// The deserialization information for the current model
func (m *DeviceHealthScriptRemediationSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["scriptCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRemediationSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceHealthScriptRemediationSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("scriptCount", m.GetScriptCount())
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
func (m *DeviceHealthScriptRemediationSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the remediatedDeviceCount property value. The number of devices remediated by device health scripts.
// Parameters:
//  - value : Value to set for the remediatedDeviceCount property.
func (m *DeviceHealthScriptRemediationSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// Sets the scriptCount property value. The number of device health scripts deployed.
// Parameters:
//  - value : Value to set for the scriptCount property.
func (m *DeviceHealthScriptRemediationSummary) SetScriptCount(value *int32)() {
    m.scriptCount = value
}
