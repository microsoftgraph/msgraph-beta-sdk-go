package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScriptRunSchedule struct {
    additionalData map[string]interface{};
    interval *int32;
}
func NewDeviceHealthScriptRunSchedule()(*DeviceHealthScriptRunSchedule) {
    m := &DeviceHealthScriptRunSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceHealthScriptRunSchedule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceHealthScriptRunSchedule) GetInterval()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.interval
    }
}
func (m *DeviceHealthScriptRunSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["interval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInterval(val)
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRunSchedule) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthScriptRunSchedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("interval", m.GetInterval())
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
func (m *DeviceHealthScriptRunSchedule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceHealthScriptRunSchedule) SetInterval(value *int32)() {
    m.interval = value
}
