package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsAutopilotDevicesSummary struct {
    additionalData map[string]interface{};
    devicesNotAutopilotRegistered *int32;
    devicesWithoutAutopilotProfileAssigned *int32;
}
func NewUserExperienceAnalyticsAutopilotDevicesSummary()(*UserExperienceAnalyticsAutopilotDevicesSummary) {
    m := &UserExperienceAnalyticsAutopilotDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetDevicesNotAutopilotRegistered()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.devicesNotAutopilotRegistered
    }
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetDevicesWithoutAutopilotProfileAssigned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.devicesWithoutAutopilotProfileAssigned
    }
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["devicesNotAutopilotRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDevicesNotAutopilotRegistered(val)
        return nil
    }
    res["devicesWithoutAutopilotProfileAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDevicesWithoutAutopilotProfileAssigned(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("devicesNotAutopilotRegistered", m.GetDevicesNotAutopilotRegistered())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("devicesWithoutAutopilotProfileAssigned", m.GetDevicesWithoutAutopilotProfileAssigned())
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
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetDevicesNotAutopilotRegistered(value *int32)() {
    m.devicesNotAutopilotRegistered = value
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetDevicesWithoutAutopilotProfileAssigned(value *int32)() {
    m.devicesWithoutAutopilotProfileAssigned = value
}
