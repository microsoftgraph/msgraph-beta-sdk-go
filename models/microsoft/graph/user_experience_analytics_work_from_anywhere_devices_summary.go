package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsWorkFromAnywhereDevicesSummary struct {
    additionalData map[string]interface{};
    autopilotDevicesSummary *UserExperienceAnalyticsAutopilotDevicesSummary;
    cloudManagementDevicesSummary *UserExperienceAnalyticsCloudManagementDevicesSummary;
    windows10DevicesSummary *UserExperienceAnalyticsWindows10DevicesSummary;
}
func NewUserExperienceAnalyticsWorkFromAnywhereDevicesSummary()(*UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetAutopilotDevicesSummary()(*UserExperienceAnalyticsAutopilotDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.autopilotDevicesSummary
    }
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetCloudManagementDevicesSummary()(*UserExperienceAnalyticsCloudManagementDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.cloudManagementDevicesSummary
    }
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetWindows10DevicesSummary()(*UserExperienceAnalyticsWindows10DevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.windows10DevicesSummary
    }
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["autopilotDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAutopilotDevicesSummary() })
        if err != nil {
            return err
        }
        m.SetAutopilotDevicesSummary(val.(*UserExperienceAnalyticsAutopilotDevicesSummary))
        return nil
    }
    res["cloudManagementDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCloudManagementDevicesSummary() })
        if err != nil {
            return err
        }
        m.SetCloudManagementDevicesSummary(val.(*UserExperienceAnalyticsCloudManagementDevicesSummary))
        return nil
    }
    res["windows10DevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsWindows10DevicesSummary() })
        if err != nil {
            return err
        }
        m.SetWindows10DevicesSummary(val.(*UserExperienceAnalyticsWindows10DevicesSummary))
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("autopilotDevicesSummary", m.GetAutopilotDevicesSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cloudManagementDevicesSummary", m.GetCloudManagementDevicesSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("windows10DevicesSummary", m.GetWindows10DevicesSummary())
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetAutopilotDevicesSummary(value *UserExperienceAnalyticsAutopilotDevicesSummary)() {
    m.autopilotDevicesSummary = value
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetCloudManagementDevicesSummary(value *UserExperienceAnalyticsCloudManagementDevicesSummary)() {
    m.cloudManagementDevicesSummary = value
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetWindows10DevicesSummary(value *UserExperienceAnalyticsWindows10DevicesSummary)() {
    m.windows10DevicesSummary = value
}
