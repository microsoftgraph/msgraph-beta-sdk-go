package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBatteryHealthDeviceAppImpact struct {
    Entity
    // User friendly display name for the app. Eg: Outlook
    appDisplayName *string;
    // App name. Eg: oltk.exe
    appName *string;
    // App publisher. Eg: Microsoft Corporation
    appPublisher *string;
    // The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    batteryUsagePercentage *float64;
    // The unique identifier of the device, Intune DeviceID or SCCM device id.
    deviceId *string;
    // true if the user had active interaction with the app.
    isForegroundApp *bool;
}
// Instantiates a new userExperienceAnalyticsBatteryHealthDeviceAppImpact and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpact) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceAppImpact{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// Gets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// Gets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetBatteryUsagePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.batteryUsagePercentage
    }
}
// Gets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetIsForegroundApp()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isForegroundApp
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["appName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppName(val)
        return nil
    }
    res["appPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppPublisher(val)
        return nil
    }
    res["batteryUsagePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetBatteryUsagePercentage(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["isForegroundApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsForegroundApp(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appName", m.GetAppName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPublisher", m.GetAppPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("batteryUsagePercentage", m.GetBatteryUsagePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isForegroundApp", m.GetIsForegroundApp())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appName property value. App name. Eg: oltk.exe
// Parameters:
//  - value : Value to set for the appName property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppName(value *string)() {
    m.appName = value
}
// Sets the appPublisher property value. App publisher. Eg: Microsoft Corporation
// Parameters:
//  - value : Value to set for the appPublisher property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// Sets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the batteryUsagePercentage property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetBatteryUsagePercentage(value *float64)() {
    m.batteryUsagePercentage = value
}
// Sets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the isForegroundApp property value. true if the user had active interaction with the app.
// Parameters:
//  - value : Value to set for the isForegroundApp property.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetIsForegroundApp(value *bool)() {
    m.isForegroundApp = value
}
