package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory struct {
    Entity
    // The unique identifier of the device, Intune DeviceID or SCCM device id.
    deviceId *string;
    // The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
    estimatedRuntimeInMinutes *int32;
    // The datetime for the instance of runtime history.
    runtimeDateTime *string;
}
// Instantiates a new userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetEstimatedRuntimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.estimatedRuntimeInMinutes
    }
}
// Gets the runtimeDateTime property value. The datetime for the instance of runtime history.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetRuntimeDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.runtimeDateTime
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["estimatedRuntimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEstimatedRuntimeInMinutes(val)
        }
        return nil
    }
    res["runtimeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuntimeDateTime(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("estimatedRuntimeInMinutes", m.GetEstimatedRuntimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("runtimeDateTime", m.GetRuntimeDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the estimatedRuntimeInMinutes property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) SetEstimatedRuntimeInMinutes(value *int32)() {
    m.estimatedRuntimeInMinutes = value
}
// Sets the runtimeDateTime property value. The datetime for the instance of runtime history.
// Parameters:
//  - value : Value to set for the runtimeDateTime property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) SetRuntimeDateTime(value *string)() {
    m.runtimeDateTime = value
}
