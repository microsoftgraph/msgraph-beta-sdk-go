package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBatteryHealthAppImpact struct {
    Entity
    // Number of active devices for using that app over a 14-day period. Valid values -2147483648 to 2147483647
    activeDevices *int32;
    // User friendly display name for the app. Eg: Outlook
    appDisplayName *string;
    // App name. Eg: oltk.exe
    appName *string;
    // App publisher. Eg: Microsoft Corporation
    appPublisher *string;
    // The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days computed across all devices in the tenant. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    batteryUsagePercentage *float64;
    // true if the user had active interaction with the app.
    isForegroundApp *bool;
}
// Instantiates a new userExperienceAnalyticsBatteryHealthAppImpact and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthAppImpact()(*UserExperienceAnalyticsBatteryHealthAppImpact) {
    m := &UserExperienceAnalyticsBatteryHealthAppImpact{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDevices property value. Number of active devices for using that app over a 14-day period. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetActiveDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDevices
    }
}
// Gets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// Gets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// Gets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days computed across all devices in the tenant. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetBatteryUsagePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.batteryUsagePercentage
    }
}
// Gets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetIsForegroundApp()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isForegroundApp
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDevices(val)
        }
        return nil
    }
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
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDevices", m.GetActiveDevices())
        if err != nil {
            return err
        }
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
        err = writer.WriteBoolValue("isForegroundApp", m.GetIsForegroundApp())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activeDevices property value. Number of active devices for using that app over a 14-day period. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDevices property.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetActiveDevices(value *int32)() {
    m.activeDevices = value
}
// Sets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appName property value. App name. Eg: oltk.exe
// Parameters:
//  - value : Value to set for the appName property.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetAppName(value *string)() {
    m.appName = value
}
// Sets the appPublisher property value. App publisher. Eg: Microsoft Corporation
// Parameters:
//  - value : Value to set for the appPublisher property.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// Sets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days computed across all devices in the tenant. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the batteryUsagePercentage property.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetBatteryUsagePercentage(value *float64)() {
    m.batteryUsagePercentage = value
}
// Sets the isForegroundApp property value. true if the user had active interaction with the app.
// Parameters:
//  - value : Value to set for the isForegroundApp property.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetIsForegroundApp(value *bool)() {
    m.isForegroundApp = value
}
