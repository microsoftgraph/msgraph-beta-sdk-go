package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDeviceAppImpact 
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
// NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact instantiates a new userExperienceAnalyticsBatteryHealthDeviceAppImpact and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpact) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceAppImpact{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppDisplayName gets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetAppName gets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// GetAppPublisher gets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// GetBatteryUsagePercentage gets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetBatteryUsagePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.batteryUsagePercentage
    }
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetIsForegroundApp gets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetIsForegroundApp()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isForegroundApp
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppName(val)
        }
        return nil
    }
    res["appPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPublisher(val)
        }
        return nil
    }
    res["batteryUsagePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryUsagePercentage(val)
        }
        return nil
    }
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
    res["isForegroundApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsForegroundApp(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAppDisplayName sets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetAppName sets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppName(value *string)() {
    if m != nil {
        m.appName = value
    }
}
// SetAppPublisher sets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppPublisher(value *string)() {
    if m != nil {
        m.appPublisher = value
    }
}
// SetBatteryUsagePercentage sets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetBatteryUsagePercentage(value *float64)() {
    if m != nil {
        m.batteryUsagePercentage = value
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetIsForegroundApp sets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetIsForegroundApp(value *bool)() {
    if m != nil {
        m.isForegroundApp = value
    }
}
